package defillama

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
)

// Exercise every public endpoint through the real transport. Besides catching
// missing methods, this checks the path, query serialization, response shape,
// and propagation of an HTTP error for each generated service wrapper.
func TestEveryServiceRoute(t *testing.T) {
	ids := make([]string, 0, len(routeRegistry))
	for id := range routeRegistry {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	for _, id := range ids {
		r := routeRegistry[id]
		t.Run(id, func(t *testing.T) {
			var failing atomic.Bool
			payload := ""
			_, c, log := testServer(t, func(w http.ResponseWriter, _ *http.Request) {
				if failing.Load() {
					w.WriteHeader(http.StatusBadRequest)
					_, _ = fmt.Fprint(w, `{"error":"invalid request"}`)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = fmt.Fprint(w, payload)
			})

			getter := reflect.ValueOf(c).MethodByName(r.Service)
			if !getter.IsValid() {
				t.Fatalf("missing %s service", r.Service)
			}
			service := getter.Call(nil)[0]
			method := service.MethodByName(r.GoMethod)
			if !method.IsValid() {
				t.Fatalf("missing %s.%s", r.Service, r.GoMethod)
			}
			payload = fixtureForRoute(r, method.Type().Out(0))
			args := routeArguments(t, r, method.Type())

			result := method.Call(args)
			if !result[1].IsNil() {
				t.Fatalf("success response: %v", result[1].Interface())
			}
			assertPopulatedResult(t, result[0])
			requests := log.all()
			if len(requests) != 1 {
				t.Fatalf("success response sent %d requests, want 1", len(requests))
			}
			assertRouteRequest(t, r, requests[0])

			failing.Store(true)
			result = method.Call(args)
			var apiErr *APIError
			if !errors.As(valueError(result[1]), &apiErr) || apiErr.StatusCode != http.StatusBadRequest {
				t.Fatalf("error response: got %v, want HTTP 400 APIError", valueError(result[1]))
			}
			if len(log.all()) != 2 {
				t.Fatalf("error response sent %d total requests, want 2", len(log.all()))
			}
		})
	}
}

func valueError(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}

func fixtureForRoute(r route, resultType reflect.Type) string {
	switch r.GoMethod {
	case "GetCurrentPrices", "GetHistoricalPrices", "GetFirstPrices":
		return `{"coins":{"coingecko:bitcoin":{"price":42.5}}}`
	case "GetPools":
		return `{"data":[{"pool":"pool-1"}]}`
	case "GetStablecoins":
		return `{"peggedAssets":[{"name":"Example"}]}`
	}
	if r.Service == "PreIPO" && r.GoMethod == "GetCompaniesList" {
		return `{"data":[{"name":"Example"}]}`
	}
	switch resultType.Kind() {
	case reflect.Slice:
		if resultType.Elem().Kind() == reflect.String {
			return `["sample"]`
		}
		return `[{"name":"Example","pool":"pool-1"}]`
	case reflect.Map:
		return `{"marker":true}`
	case reflect.Pointer:
		return `{"name":"Example","height":42,"credits_left":42}`
	case reflect.Float64:
		return `42.5`
	default:
		return `null`
	}
}

func routeArguments(t *testing.T, r route, typ reflect.Type) []reflect.Value {
	t.Helper()
	args := []reflect.Value{reflect.ValueOf(context.Background())}
	pathString := "sample"
	for _, p := range r.Params {
		if p.In == "path" && len(p.Enum) > 0 {
			pathString = p.Enum[0]
		}
	}
	for i := 1; i < typ.NumIn(); i++ {
		paramType := typ.In(i)
		switch paramType.Kind() {
		case reflect.String:
			args = append(args, reflect.ValueOf(pathString))
		case reflect.Int64:
			args = append(args, reflect.ValueOf(int64(42)))
		case reflect.Float64:
			args = append(args, reflect.ValueOf(42.5))
		case reflect.Slice:
			if paramType.Elem().Kind() != reflect.String {
				t.Fatalf("unexpected slice argument %s", paramType)
			}
			args = append(args, reflect.ValueOf([]string{"coingecko:bitcoin"}))
		case reflect.Map:
			args = append(args, reflect.ValueOf(map[string][]int64{"coingecko:bitcoin": {42}}))
		case reflect.Struct:
			v := reflect.New(paramType).Elem()
			populateRouteOptions(v, r)
			args = append(args, v)
		case reflect.Pointer:
			v := reflect.New(paramType.Elem())
			populateRouteOptions(v.Elem(), r)
			args = append(args, v)
		default:
			t.Fatalf("unexpected argument type %s", paramType)
		}
	}
	return args
}

func populateRouteOptions(v reflect.Value, r route) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := normalizeParamName(v.Type().Field(i).Name)
		str := "sample"
		declared := false
		for _, p := range r.Params {
			if normalizeParamName(p.Name) == name {
				declared = true
				if len(p.Enum) > 0 {
					str = p.Enum[0]
				}
			}
		}
		if field.Kind() == reflect.Pointer {
			if !declared {
				continue
			}
			field.Set(reflect.New(field.Type().Elem()))
			field = field.Elem()
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(str)
		case reflect.Bool:
			field.SetBool(true)
		case reflect.Int64:
			field.SetInt(42)
		case reflect.Float64:
			field.SetFloat(42.5)
		}
	}
}

func normalizeParamName(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "_", "")
}

func assertRouteRequest(t *testing.T, r route, req *http.Request) {
	t.Helper()
	if req.Method != http.MethodGet {
		t.Errorf("method = %s, want GET", req.Method)
	}
	wantPath := r.Path
	pathString := "sample"
	for _, p := range r.Params {
		if p.In == "path" && len(p.Enum) > 0 {
			pathString = p.Enum[0]
		}
	}
	for _, p := range r.Params {
		if p.In != "path" {
			continue
		}
		value := pathString
		if p.Type == "integer" || p.Name == "timestamp" {
			value = "42"
		} else if p.Type == "number" {
			value = "42.5"
		} else if p.Name == "coins" {
			value = "coingecko:bitcoin"
		}
		wantPath = strings.ReplaceAll(wantPath, "{"+p.Name+"}", value)
	}
	if r.Tier == "pro" {
		wantPath = "/SECRET_TEST_KEY_1" + wantPath
	}
	if got := req.URL.Path; got != wantPath {
		t.Errorf("path = %q, want %q", got, wantPath)
	}
	for _, p := range r.Params {
		if p.In != "query" {
			continue
		}
		q := req.URL.Query()
		if !q.Has(p.Name) {
			t.Errorf("missing query parameter %q in %s", p.Name, req.URL.RawQuery)
			continue
		}
		want := "sample"
		switch {
		case p.Name == "coins":
			want = `{"coingecko:bitcoin":[42]}`
		case len(p.Enum) > 0:
			want = p.Enum[0]
		case p.Type == "boolean":
			want = "true"
		case p.Type == "integer":
			want = "42"
		case p.Type == "number":
			want = "42.5"
		}
		if got := q.Get(p.Name); got != want {
			t.Errorf("query parameter %q = %q, want %q", p.Name, got, want)
		}
	}
}

func assertPopulatedResult(t *testing.T, value reflect.Value) {
	t.Helper()
	switch value.Kind() {
	case reflect.Slice, reflect.Map:
		if value.Len() != 1 {
			t.Errorf("decoded result length = %d, want 1", value.Len())
		}
	case reflect.Pointer:
		if value.IsNil() {
			t.Error("decoded pointer is nil")
		}
	case reflect.Float64:
		if value.Float() != 42.5 {
			t.Errorf("decoded value = %v, want 42.5", value.Float())
		}
	default:
		t.Errorf("unhandled result type %s", value.Type())
	}
}
