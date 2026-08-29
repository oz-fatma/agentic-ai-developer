package capstone

import "time"

// Phase names the major learning blocks in the curriculum.
type Phase string

const (
	PhaseCaching  Phase = "caching-messaging"
	PhasePerf     Phase = "performance"
	PhaseTeam     Phase = "team-practices"
	PhaseCapstone Phase = "capstone"
)

// Milestone tracks a deliverable in the capstone plan.
type Milestone struct {
	Name        string
	Phase       Phase
	Done        bool
	Due         time.Time
	Description string
}

// Plan is the structured roadmap for the final project.
type Plan struct {
	Title       string
	Owner       string
	Milestones  []Milestone
	Hardening   []string
	DeployNotes string
}

func DefaultPlan() Plan {
	return Plan{
		Title: "Go Service Capstone",
		Owner: "learn-go-track",
		Milestones: []Milestone{
			{Name: "Cache layer", Phase: PhaseCaching, Description: "TTL + LRU caches"},
			{Name: "Async workers", Phase: PhaseCaching, Description: "Queue + pub/sub"},
			{Name: "Profile hot paths", Phase: PhasePerf, Description: "pprof and benches"},
			{Name: "Release checklist", Phase: PhaseTeam, Description: "Makefile, semver, reviews"},
			{Name: "Deploy", Phase: PhaseCapstone, Description: "K8s manifest + CI"},
		},
		Hardening: []string{
			"Rate limit public endpoints",
			"Graceful shutdown for workers",
			"Structured logging with request IDs",
		},
		DeployNotes: "Container image + K8s Deployment with probes",
	}
}

func (p Plan) CompletedCount() int {
	n := 0
	for _, m := range p.Milestones {
		if m.Done {
			n++
		}
	}
	return n
}
