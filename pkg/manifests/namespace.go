package manifests

type Namespace struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata" json:"metadata"`
}

func CreateNamespace(name string) Namespace {
	return Namespace{
		ApiVersion: "v1",
		Kind:       "Namespace",
		Metadata: Metadata{
			Name: name,
		},
	}
}
