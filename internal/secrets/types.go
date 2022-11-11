package secrets

const PREFIX = "FLY_SECRET"

type Entry struct {
	Name string
	Path string
}

type Secrets struct {
	Input  []Entry
	Output []Entry
}
