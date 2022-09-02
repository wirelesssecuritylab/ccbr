package ResponseStruct

import (
	"ccbr/model/Clustermanager"
	"ccbr/model/Gatekeeper"
)

type ResponseOPAGCTStruct struct {
	Code  int                                            `json:"code"`
	Msg   string                                         `json:"msg"`
	Count int                                            `json:"count"`
	Data  []Gatekeeper.OpaGatekeeperConstraintTemplater2 `json:"data"`
}

type ResponseOPAGConstraintStruct struct {
	Code  int                                   `json:"code"`
	Msg   string                                `json:"msg"`
	Count int                                   `json:"count"`
	Data  []Gatekeeper.OpaGatekeeperConstraint2 `json:"data"`
}

type ResponseOPAPoliciesStruct struct {
	Code  int                                 `json:"code"`
	Msg   string                              `json:"msg"`
	Count int                                 `json:"count"`
	Data  []Gatekeeper.OpaGatekeeperPolicies2 `json:"data"`
}
type ResonseOPAConstraintMultiSelect struct {
	Name     string `json:"name"`
	Value    int    `json:"value"`
	Selected bool   `json:"selected"`
	Disabled bool   `json:"disabled"`
}
type ResponseClusterManagerStruct struct {
	Code  int                                      `json:"code"`
	Msg   string                                   `json:"msg"`
	Count int                                      `json:"count"`
	Data  []Clustermanager.ClusterManagerModelView `json:"data"`
}

type ResponseClusterManagerConstraintTemplateStruct struct {
	Code  int                                           `json:"code"`
	Msg   string                                        `json:"msg"`
	Count int                                           `json:"count"`
	Data  []ResponseClustermanagerAndConstraintTenplate `json:"data"`
}

type ResponseKubernetesTestStruct struct {
	Name      string
	Version   string
	Namespace string
	Pod       string
	Status    string
}

type ResponseClustermanagerAndConstraintTenplate struct {
	Clustername            string
	Constrainttemplatename string
	File                   string
	Constraint             string
}

type Report struct {
	ClusterManager     string
	ConstraintTemplate string
	Constraint         string
	Action             string
	Kind               string
	Namespace          string
	Name               string
	Message            string
}

type CisBenchmarkInfo struct {
	CisBenchmarkItem string
	Result           string
	Message          string
	PodName          string
	ClusterManager   string
}

type ResponseCisBenchmarkCheckResult struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data []CisBenchmarkInfo `json:"data"`
}
