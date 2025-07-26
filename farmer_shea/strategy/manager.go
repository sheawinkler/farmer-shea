package strategy

// Manager manages all the strategies.
type Manager struct {
	Strategies []Strategy
}

// NewManager creates a new strategy manager.
func NewManager() *Manager {
	return &Manager{Strategies: []Strategy{}}
}

// Add adds a new strategy to the manager.
func (m *Manager) Add(s Strategy) {
	m.Strategies = append(m.Strategies, s)
}