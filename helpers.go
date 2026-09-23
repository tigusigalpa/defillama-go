package defillama

// Ptr returns a pointer to v — a convenience for filling pointer-based
// optional fields on the option structs (e.g. &ChartOptions{Period: Ptr("1d")}).
func Ptr[T any](v T) *T { return &v }
