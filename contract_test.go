package defillama

import (
	"encoding/json"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// Contract test: the route registry, public service methods and
// docs/api-index.md must exactly cover the pinned OpenAPI spec.

type specDoc struct {
	Servers []struct {
		URL string `json:"url"`
	} `json:"servers"`
	Paths map[string]specPath `json:"paths"`
}

type specPath struct {
	Servers []struct {
		URL string `json:"url"`
	} `json:"servers"`
	Parameters []specParam `json:"parameters"`
	Get        *specOp     `json:"get"`
}

type specOp struct {
	Servers []struct {
		URL string `json:"url"`
	} `json:"servers"`
	Parameters []specParam `json:"parameters"`
}

type specParam struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Schema   struct {
		Type string `json:"type"`
	} `json:"schema"`
}

func loadSpec(t *testing.T) (map[string]struct {
	op     specOp
	server string
}, *specDoc,
) {
	t.Helper()
	raw, err := os.ReadFile("spec/defillama-api.json")
	if err != nil {
		t.Fatalf("read spec: %v", err)
	}
	var doc specDoc
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parse spec: %v", err)
	}
	ops := map[string]struct {
		op     specOp
		server string
	}{}
	for path, p := range doc.Paths {
		if p.Get == nil {
			continue
		}
		server := ""
		if len(p.Get.Servers) > 0 {
			server = p.Get.Servers[0].URL
		} else if len(p.Servers) > 0 {
			server = p.Servers[0].URL
		} else if len(doc.Servers) > 0 {
			server = doc.Servers[0].URL
		}
		params := append(append([]specParam{}, p.Parameters...), p.Get.Parameters...)
		op := *p.Get
		op.Parameters = params
		ops["GET "+path] = struct {
			op     specOp
			server string
		}{op: op, server: server}
	}
	return ops, &doc
}

func TestSpecDefines132Operations(t *testing.T) {
	ops, _ := loadSpec(t)
	if len(ops) != 132 {
		t.Fatalf("spec defines %d GET ops, want 132", len(ops))
	}
}

func TestRegistryCoversSpecExactly(t *testing.T) {
	ops, _ := loadSpec(t)
	if len(routeRegistry) != len(ops) {
		t.Fatalf("registry has %d routes, spec has %d", len(routeRegistry), len(ops))
	}
	for id := range ops {
		if _, ok := routeRegistry[id]; !ok {
			t.Errorf("missing registry route: %s", id)
		}
	}
	for id := range routeRegistry {
		if _, ok := ops[id]; !ok {
			t.Errorf("registry route not in spec: %s", id)
		}
	}
}

func TestRegistryMatchesSpecMetadata(t *testing.T) {
	ops, _ := loadSpec(t)
	for id, info := range ops {
		r := routeRegistry[id]
		if r.Server != info.server {
			t.Errorf("%s: server %q, spec %q", id, r.Server, info.server)
		}
		tier := "free"
		if strings.Contains(info.server, "pro-api") {
			tier = "pro"
		}
		if r.Tier != tier {
			t.Errorf("%s: tier %q, want %q", id, r.Tier, tier)
		}
		want := map[string][2]any{}
		for _, p := range info.op.Parameters {
			want[p.In+":"+p.Name] = [2]any{p.Required, p.Schema.Type}
		}
		got := map[string][2]any{}
		for _, p := range r.Params {
			got[p.In+":"+p.Name] = [2]any{p.Required, p.Type}
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("%s: params mismatch\n spec: %v\n reg:  %v", id, want, got)
		}
	}
}

func TestEveryRouteHasPublicServiceMethod(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	services := map[string]any{
		"TVL": c.TVL(), "Prices": c.Prices(), "Stablecoins": c.Stablecoins(),
		"Yields": c.Yields(), "Volumes": c.Volumes(), "Fees": c.Fees(),
		"Emissions": c.Emissions(), "Ecosystem": c.Ecosystem(), "Bridges": c.Bridges(),
		"ETFs": c.ETFs(), "Narratives": c.Narratives(), "Account": c.Account(),
		"DAT": c.DAT(), "Treasury": c.Treasury(), "Oracles": c.Oracles(),
		"Forks": c.Forks(), "Dimensions": c.Dimensions(),
		"FinancialStatements": c.FinancialStatements(), "Equities": c.Equities(),
		"PreIPO": c.PreIPO(), "RWA": c.RWA(),
	}
	for id, r := range routeRegistry {
		svc, ok := services[r.Service]
		if !ok {
			t.Errorf("%s: unknown service %q", id, r.Service)
			continue
		}
		m := reflect.ValueOf(svc).MethodByName(r.GoMethod)
		if !m.IsValid() {
			t.Errorf("%s: %sService.%s not found", id, r.Service, r.GoMethod)
		}
	}
}

func TestAPIIndexHas132UniqueDocURLs(t *testing.T) {
	raw, err := os.ReadFile("docs/api-index.md")
	if err != nil {
		t.Fatalf("read api-index: %v", err)
	}
	re := regexp.MustCompile(`\((https://api-docs\.defillama\.com/[^)]+)\)`)
	urls := re.FindAllStringSubmatch(string(raw), -1)
	if len(urls) != 132 {
		t.Fatalf("api-index.md has %d doc links, want 132", len(urls))
	}
	seen := map[string]bool{}
	for _, u := range urls {
		if seen[u[1]] {
			t.Errorf("duplicate docs URL: %s", u[1])
		}
		seen[u[1]] = true
	}
	for id, r := range routeRegistry {
		if !strings.Contains(string(raw), r.DocsURL) {
			t.Errorf("%s: docs URL %s missing from api-index.md", id, r.DocsURL)
		}
	}
}

func TestAllFreeRoutesHaveProOverride(t *testing.T) {
	n := 0
	for id, r := range routeRegistry {
		if r.Tier == "free" {
			n++
			if r.ProPath == "" {
				t.Errorf("%s: free route missing pro override", id)
			}
		}
	}
	if n != 31 {
		t.Fatalf("found %d free routes, want 31", n)
	}
}
