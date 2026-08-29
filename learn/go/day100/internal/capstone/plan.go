package capstone

type Phase string

const (
	PhaseCaching  Phase = "caching-messaging"
	PhasePerf     Phase = "performance"
	PhaseTeam     Phase = "team-practices"
	PhaseCapstone Phase = "capstone"
)

type Milestone struct {
	Name  string
	Phase Phase
}

type Plan struct {
	Title      string
	Milestones []Milestone
}

func DefaultPlan() Plan {
	return Plan{
		Title: "Go Service Capstone",
		Milestones: []Milestone{
			{Name: "Cache layer", Phase: PhaseCaching},
			{Name: "Profile hot paths", Phase: PhasePerf},
			{Name: "Release checklist", Phase: PhaseTeam},
			{Name: "Deploy", Phase: PhaseCapstone},
		},
	}
}
