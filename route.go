package defillama

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// proBase is the DefiLlama Pro API origin. The API key is appended as a single
// URL path segment — never as a header or query parameter.
const proBase = "https://pro-api.llama.fi"

// routeParam describes one declared path/query parameter of a route.
type routeParam struct {
	Name     string
	In       string // "path" | "query"
	Required bool
	Type     string // string | integer | number | boolean
	Enum     []string
}

// route is one entry of the closed route registry generated from
// spec/defillama-api.json.
type route struct {
	Method   string
	Path     string
	Tier     string // "free" | "pro"
	Server   string // per-operation Free origin, or proBase
	ProPath  string // official Pro override for Free routes ("" for Pro routes)
	Service  string
	GoMethod string
	DocsURL  string
	Summary  string
	Params   []routeParam
}

// lookupRoute fetches a registry entry by id ("GET /path").
func lookupRoute(id string) (route, error) {
	r, ok := routeRegistry[id]
	if !ok {
		return route{}, fmt.Errorf("defillama: unknown route %q", id)
	}
	return r, nil
}

// resolve builds the request URL and its redacted twin for diagnostics.
//
//	path params are escaped exactly once; a []string value is escaped
//	element-wise and joined with ',' (the coins segment).
func (c *config) resolve(r route, pathParams map[string]any) (string, string, error) {
	path := r.Path
	base := r.Server
	keySegment := ""

	if r.Tier == "pro" {
		if c.apiKey == "" {
			return "", "", &ProAPIKeyRequiredError{Route: r.Method + " " + r.Path}
		}
		keySegment = "/" + url.PathEscape(c.apiKey)
		base = proBase
	} else if c.preferProForFree && c.apiKey != "" && r.ProPath != "" {
		keySegment = "/" + url.PathEscape(c.apiKey)
		base = proBase
		path = r.ProPath
	}

	for _, p := range r.Params {
		if p.In != "path" {
			continue
		}
		v, ok := pathParams[p.Name]
		if !ok {
			return "", "", fmt.Errorf("defillama: missing path parameter %q for %s", p.Name, r.Path)
		}
		if len(p.Enum) > 0 && !inStrings(fmt.Sprint(v), p.Enum) {
			return "", "", fmt.Errorf("defillama: invalid value %q for path parameter %q", v, p.Name)
		}
		var encoded string
		switch t := v.(type) {
		case []string:
			if len(t) == 0 {
				return "", "", fmt.Errorf("defillama: path parameter %q must not be empty", p.Name)
			}
			parts := make([]string, len(t))
			for i, s := range t {
				if s == "" {
					return "", "", fmt.Errorf("defillama: path parameter %q contains an empty element", p.Name)
				}
				parts[i] = url.PathEscape(s)
			}
			encoded = strings.Join(parts, ",")
		case string:
			if t == "" {
				return "", "", fmt.Errorf("defillama: path parameter %q must not be empty", p.Name)
			}
			encoded = url.PathEscape(t)
		case int64:
			if t < 0 {
				return "", "", fmt.Errorf("defillama: path parameter %q must be >= 0", p.Name)
			}
			encoded = strconv.FormatInt(t, 10)
		case float64:
			if t < 0 {
				return "", "", fmt.Errorf("defillama: path parameter %q must be >= 0", p.Name)
			}
			encoded = strconv.FormatFloat(t, 'f', -1, 64)
		case int:
			if t < 0 {
				return "", "", fmt.Errorf("defillama: path parameter %q must be >= 0", p.Name)
			}
			encoded = strconv.Itoa(t)
		default:
			encoded = url.PathEscape(fmt.Sprint(v))
		}
		path = strings.ReplaceAll(path, "{"+p.Name+"}", encoded)
	}

	if override, ok := c.baseURLs[base]; ok {
		base = override
	}

	u := base + keySegment + path
	redacted := u
	if keySegment != "" {
		redacted = base + "/{API_KEY}" + path
	}
	return u, redacted, nil
}

// validateQuery checks required flags and enum values declared by the spec.
func validateQuery(r route, q url.Values) error {
	for _, p := range r.Params {
		if p.In != "query" {
			continue
		}
		if p.Required && (!q.Has(p.Name) || strings.TrimSpace(q.Get(p.Name)) == "") {
			return fmt.Errorf("defillama: missing required query parameter %q for %s", p.Name, r.Path)
		}
		if len(p.Enum) > 0 && q.Has(p.Name) && !inStrings(q.Get(p.Name), p.Enum) {
			return fmt.Errorf("defillama: invalid value %q for query parameter %q", q.Get(p.Name), p.Name)
		}
	}
	return nil
}

// inStrings reports whether s is in the list.
func inStrings(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// boolStr serializes a boolean query flag as lowercase "true"/"false".
func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
