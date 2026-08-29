package hardening

var Checklist = []struct {
	Category string
	Item     string
}{
	{Category: "Security", Item: "Validate input"},
	{Category: "Reliability", Item: "Health endpoints"},
	{Category: "Observability", Item: "Structured logs"},
}

func ByCategory() map[string][]string {
	out := make(map[string][]string)
	for _, c := range Checklist {
		out[c.Category] = append(out[c.Category], c.Item)
	}
	return out
}
