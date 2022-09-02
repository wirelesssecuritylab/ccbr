package services

import (
	"ccbr/model/ResponseStruct"
	"ccbr/utils"
)

func ClusterManagerService_Test_Kubernetes(kubeconfig string, clusterName string) ([]ResponseStruct.ResponseKubernetesTestStruct, error) {

	var responseKubernetesTestResults []ResponseStruct.ResponseKubernetesTestStruct
	clentSet, err := utils.InitClient(kubeconfig)
	if err != nil {
		return nil, err
	}
	version := utils.K8sVersion(clentSet)
	namespaceList, err := utils.GetNamespaces(clentSet)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(namespaceList); i++ {
		var temp ResponseStruct.ResponseKubernetesTestStruct
		result, err := utils.GetPods(clentSet, namespaceList[i].Name)
		if err != nil {
			return nil, err
		}
		for j := 0; j < len(result); j++ {
			temp.Name = clusterName
			temp.Version = version
			temp.Namespace = namespaceList[i].Name
			temp.Pod = result[j].Name
			temp.Status = result[j].Status
			responseKubernetesTestResults = append(responseKubernetesTestResults, temp)
		}
	}
	return responseKubernetesTestResults, nil

}

func ClusterManagerService_Test_Gatekeeper(kubeconfig string, clusterName string) ([]ResponseStruct.ResponseKubernetesTestStruct, error) {
	var responseKubernetesTestResults []ResponseStruct.ResponseKubernetesTestStruct
	clentSet, err := utils.InitClient(kubeconfig)
	if err != nil {
		return nil, err
	}
	result, err := utils.GetPods(clentSet, "gatekeeper-system")
	for j := 0; j < len(result); j++ {
		var temp ResponseStruct.ResponseKubernetesTestStruct
		temp.Name = clusterName
		temp.Version = ""
		temp.Namespace = "gatekeeper-system"
		temp.Pod = result[j].Name
		temp.Status = result[j].Status
		responseKubernetesTestResults = append(responseKubernetesTestResults, temp)
	}
	return responseKubernetesTestResults, nil
}
