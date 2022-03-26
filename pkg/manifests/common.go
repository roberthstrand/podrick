package manifests

type Metadata struct {
	Name        string `yaml:"name" json:"name"`
	Namespace   string `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Labels      string `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations string `yaml:"annotations,omitempty" json:"annotations,omitempty"`
}
