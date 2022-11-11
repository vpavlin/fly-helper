package template

type Template struct {
	Values string          `json:"values" yaml:"values"`
	Items  []TemplateEntry `json:"items" yaml:"items"`
}

type TemplateEntry struct {
	Template string `json:"template" yaml:"template"`
	Output   string `json:"output" yaml:"output"`
}
