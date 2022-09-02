package ConstraintTemplate

type ConstraintTemplate struct {
	Object Object `json:"Object"`
}
type Annotations struct {
	KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
}
type NAMING_FAILED struct {
}
type FByPod struct {
}
type FCreated struct {
}
type FStatus struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FByPod        FByPod        `json:"f:byPod"`
	FCreated      FCreated      `json:"f:created"`
}
type FieldsV2 struct {
	FStatus FStatus `json:"f:status"`
}
type FKubectlKubernetesIoLastAppliedConfiguration struct {
}
type FAnnotations struct {
	NAMING_FAILED                                NAMING_FAILED                                `json:"."`
	FKubectlKubernetesIoLastAppliedConfiguration FKubectlKubernetesIoLastAppliedConfiguration `json:"f:kubectl.kubernetes.io/last-applied-configuration"`
}
type FMetadata struct {
	FAnnotations FAnnotations `json:"f:annotations"`
}
type FKind struct {
}
type FNames struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FKind         FKind         `json:"f:kind"`
}
type FLegacySchema struct {
}
type FProperties struct {
}
type FOpenAPIV3Schema struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FProperties   FProperties   `json:"f:properties"`
}
type FValidation struct {
	NAMING_FAILED    NAMING_FAILED    `json:"."`
	FLegacySchema    FLegacySchema    `json:"f:legacySchema"`
	FOpenAPIV3Schema FOpenAPIV3Schema `json:"f:openAPIV3Schema"`
}
type FSpec_1 struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FNames        FNames        `json:"f:names"`
	FValidation   FValidation   `json:"f:validation"`
}
type FCrd struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FSpec         FSpec_1       `json:"f:spec"`
}
type FTargets struct {
}
type FSpec struct {
	NAMING_FAILED NAMING_FAILED `json:"."`
	FCrd          FCrd          `json:"f:crd"`
	FTargets      FTargets      `json:"f:targets"`
}
type FieldsV1 struct {
	FMetadata FMetadata `json:"f:metadata"`
	FSpec     FSpec     `json:"f:spec"`
}
type ManagedFields struct {
	APIVersion string   `json:"apiVersion"`
	FieldsType string   `json:"fieldsType"`
	FieldsV1   FieldsV1 `json:"fieldsV1,omitempty"`
	Manager    string   `json:"manager"`
	Operation  string   `json:"operation"`
	Time2      string   `json:"time"`
}
type Metadata struct {
	Annotations       Annotations     `json:"annotations"`
	CreationTimestamp string          `json:"creationTimestamp"`
	Generation        int             `json:"generation"`
	ManagedFields     []ManagedFields `json:"managedFields"`
	Name              string          `json:"name"`
	ResourceVersion   string          `json:"resourceVersion"`
	UID               string          `json:"uid"`
}
type Names struct {
	Kind string `json:"kind"`
}
type Prefix struct {
	Type string `json:"type"`
}
type Properties struct {

}
type OpenAPIV3Schema struct {
	Description string `json:"description"`
	Properties interface{} `json:"properties"`
}
type Validation struct {
	LegacySchema    bool            `json:"legacySchema"`
	OpenAPIV3Schema OpenAPIV3Schema `json:"openAPIV3Schema"`
}
type Spec_1 struct {
	Names      Names      `json:"names"`
	Validation Validation `json:"validation"`
}
type Crd struct {
	Spec Spec_1 `json:"spec"`
}
type Targets struct {
	Rego string `json:"rego"`
	Target string `json:"target"`
	Libs []string `json:"libs"`
}
type Spec struct {
	Crd     Crd       `json:"crd"`
	Targets []Targets `json:"targets"`
}
type ByPod struct {
	ID string `json:"id"`
	ObservedGeneration int `json:"observedGeneration"`
	Operations []string `json:"operations"`
	TemplateUID string `json:"templateUID"`
}
type Status struct {
	ByPod []ByPod `json:"byPod"`
	Created bool  `json:"created"`
}
type Object struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}
