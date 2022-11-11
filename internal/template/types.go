package template

type TemplateEntry struct {
	Values   string `json:"values" yaml:"values"`
	Template string `json:"template" yaml:"template"`
	Output   string `json:"output" yaml:"output"`
}
