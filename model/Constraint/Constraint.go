package Constraint

import "time"

type Constraint struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Annotations struct {
			KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
		} `json:"annotations"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Generation        int       `json:"generation"`
		ManagedFields     []struct {
			ApiVersion string `json:"apiVersion"`
			FieldsType string `json:"fieldsType"`
			FieldsV1   struct {
				FMetadata struct {
					FAnnotations struct {
						Field1 struct {
						} `json:"."`
						FKubectlKubernetesIoLastAppliedConfiguration struct {
						} `json:"f:kubectl.kubernetes.io/last-applied-configuration"`
					} `json:"f:annotations"`
				} `json:"f:metadata,omitempty"`
				FSpec struct {
					Field1 struct {
					} `json:"."`
					FMatch struct {
						Field1 struct {
						} `json:"."`
						FKinds struct {
						} `json:"f:kinds"`
					} `json:"f:match"`
				} `json:"f:spec,omitempty"`
				FStatus struct {
				} `json:"f:status,omitempty"`
			} `json:"fieldsV1"`
			Manager   string    `json:"manager"`
			Operation string    `json:"operation"`
			Time      time.Time `json:"time"`
		} `json:"managedFields"`
		Name            string `json:"name"`
		ResourceVersion string `json:"resourceVersion"`
		Uid             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Match struct {
			Kinds []struct {
				ApiGroups []string `json:"apiGroups"`
				Kinds     []string `json:"kinds"`
			} `json:"kinds"`
		} `json:"match"`
	} `json:"spec"`
	Status struct {
		AuditTimestamp time.Time `json:"auditTimestamp"`
		ByPod          []struct {
			ConstraintUID      string   `json:"constraintUID"`
			Enforced           bool     `json:"enforced"`
			Id                 string   `json:"id"`
			ObservedGeneration int      `json:"observedGeneration"`
			Operations         []string `json:"operations"`
		} `json:"byPod"`
		TotalViolations int `json:"totalViolations"`
		Violations      []struct {
			EnforcementAction string `json:"enforcementAction"`
			Kind              string `json:"kind"`
			Message           string `json:"message"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
		} `json:"violations"`
	} `json:"status"`
}
