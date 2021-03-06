package config

type ItemType string

const (
	ItemTypeInput       = "input"
	ItemTypeTextarea    = "textarea"
	ItemTypeSelect      = "select"
	ItemTypeMultiSelect = "multiSelect"
	ItemTypeTags        = "tags"
	ItemTypeImage       = "image"
	ItemTypeSwitch      = "switch"
)

// ThemeConfig is the struct for config.yaml under theme folder
type ThemeConfig struct {
	Tabs []Tab `yaml:"tabs" json:"tabs"`
}

type Tab struct {
	Name  string `yaml:"name" json:"name"`
	Key   string `yaml:"key" json:"key"`
	Items []Item `yaml:"items" json:"items"`
}

type Item struct {
	Name        string      `yaml:"name" json:"name"`
	Key         string      `yaml:"key" json:"key"`
	Type        string      `yaml:"type" json:"type"`
	Default     string      `yaml:"default" json:"default"`
	Description string      `yaml:"description" json:"description"`
	Placeholder string      `yaml:"placeholder" json:"placeholder"`
	Value       interface{} `yaml:"value" json:"value"`
	Values      []string    `yaml:"values" json:"values"`
	Options     []string    `yaml:"options" json:"options"`
}
