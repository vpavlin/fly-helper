package secrets

const PREFIX = "FLY_SECRET"

type Entry struct {
	Name string `json:"name" yaml:"name"`
	Path string `json:"path" yaml:"path"`
}

type Secrets struct {
	Input  []Entry `json:"input" yaml:"input"`
	Output []Entry `json:"output" yaml:"output"`
}
