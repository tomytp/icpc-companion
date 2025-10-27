package platform

import "strings"

type Manager struct {
    resolvers []Resolver
}

func NewManager() *Manager {
    return &Manager{resolvers: []Resolver{
        CodeforcesResolver{},
        VJudgeResolver{},
        GenericResolver{},
    }}
}

func (m *Manager) Resolve(basePath, url string) (ProblemInfo, error) {
    lower := strings.ToLower(url)
    for _, r := range m.resolvers {
        if r.Matches(lower) {
            if pi, err := r.Resolve(basePath, lower); err == nil {
                return pi, nil
            }
        }
    }
    // fallback to generic as last resort
    return GenericResolver{}.Resolve(basePath, lower)
}

