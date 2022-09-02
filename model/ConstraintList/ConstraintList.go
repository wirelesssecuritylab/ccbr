package ConstraintList

type ConstraintList struct {
	APIVersion string     `json:"apiVersion"`
	GroupVersion string   `json:"groupVersion"`
	Items []interface{}   `json:"items"`
	Kind string           `json:"kind"`
	Resources []Resources `json:"resources"`
}
type Resources struct {
	Categories []string `json:"categories,omitempty"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	Namespaced bool `json:"namespaced"`
	SingularName string `json:"singularName"`
	StorageVersionHash string `json:"storageVersionHash,omitempty"`
	Verbs []string `json:"verbs"`
}
