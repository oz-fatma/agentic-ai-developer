package hardening

// Checklist covers production readiness items for the capstone service.
var Checklist = []struct {
	Category string
	Item     string
}{
	{Category: "Security", Item: "Validate and sanitize all external input"},
	{Category: "Security", Item: "Use least-privilege credentials for dependencies"},
	{Category: "Reliability", Item: "Configure timeouts on outbound HTTP and DB calls"},
	{Category: "Reliability", Item: "Add health and readiness endpoints"},
	{Category: "Observability", Item: "Export metrics and structured logs"},
	{Category: "Operations", Item: "Document rollback steps for each release"},
}

func ByCategory() map[string][]string {
	out := make(map[string][]string)
	for _, c := range Checklist {
		out[c.Category] = append(out[c.Category], c.Item)
	}
	return out
}
