package manifests

import "testing"

func TestCreateNamespace(t *testing.T) {
	name := "test"
	expected := Namespace{
		ApiVersion: "v1",
		Kind:       "Namespace",
		Metadata: Metadata{
			Name: name,
		},
	}
	actual := CreateNamespace(name)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
