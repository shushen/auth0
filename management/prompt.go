package management

type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`
}

type PromptManager struct {
	m *Management
}

func NewPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

func (pm *PromptManager) Read() (*Prompt, error) {
	p := new(Prompt)
	err := pm.m.get(pm.m.uri("prompts"), p)
	return p, err
}

func (pm *PromptManager) Update(p *Prompt) error {
	return pm.m.patch(pm.m.uri("prompts"), p)
}
