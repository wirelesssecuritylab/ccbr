package utils

import (
	"ccbr/model/Constraint"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

func K8sVersion(clientSet kubernetes.Clientset) string {
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		return err.Error()
	}
	return serverVersion.String()

}

func GetNamespaces(clientSet kubernetes.Clientset) ([]K8sResultStruct, error) {
	var result []K8sResultStruct
	namespacesList, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, v := range namespacesList.Items {
		var temp K8sResultStruct
		temp.Name = v.Name
		temp.Status = fmt.Sprintf("%v", v.Status.Phase)
		result = append(result, temp)
	}
	return result, nil

}

type K8sResultStruct struct {
	Name   string
	Status string
}

func GetPods(clientSet kubernetes.Clientset, namespace string) ([]K8sResultStruct, error) {
	var result []K8sResultStruct
	podsList, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, v := range podsList.Items {
		var temp K8sResultStruct
		temp.Name = v.Name
		temp.Status = fmt.Sprintf("%v", v.Status.Phase)
		result = append(result, temp)
	}
	return result, err
}
func InitClient(filename string) (kubernetes.Clientset, error) {

	restConf, err := GetRestConf(filename)
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	// 生成clientset配置
	clientset, err := kubernetes.NewForConfig(&restConf)
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	return *clientset, nil
}

type ConstraintTemplateSpec struct {
	CRD     CRD      `json:"crd,omitempty"`
	Targets []Target `json:"targets,omitempty"`
}

type CRD struct {
	Spec CRDSpec `json:"spec,omitempty"`
}

type CRDSpec struct {
	Names Names `json:"names,omitempty"`
	// +kubebuilder:default={legacySchema: false}
	Validation *Validation `json:"validation,omitempty"`
}

type Names struct {
	Kind       string   `json:"kind,omitempty"`
	ShortNames []string `json:"shortNames,omitempty"`
}

type Validation struct {
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	OpenAPIV3Schema *apiextensionsv1.JSONSchemaProps `json:"openAPIV3Schema,omitempty"`
	// +kubebuilder:default=false
	LegacySchema *bool `json:"legacySchema,omitempty"` // *bool allows for "unset" state which we need to apply appropriate defaults
}

type Target struct {
	Target string   `json:"target,omitempty"`
	Rego   string   `json:"rego,omitempty"`
	Libs   []string `json:"libs,omitempty"`
}

// CreateCRDError represents a single error caught during parsing, compiling, etc.
type CreateCRDError struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Location string `json:"location,omitempty"`
}

// ByPodStatus defines the observed state of ConstraintTemplate as seen by
// an individual controller
// +kubebuilder:pruning:PreserveUnknownFields
type ByPodStatus struct {
	// a unique identifier for the pod that wrote the status
	ID                 string           `json:"id,omitempty"`
	ObservedGeneration int64            `json:"observedGeneration,omitempty"`
	Errors             []CreateCRDError `json:"errors,omitempty"`
}

// ConstraintTemplateStatus defines the observed state of ConstraintTemplate.
type ConstraintTemplateStatus struct {
	Created bool          `json:"created,omitempty"`
	ByPod   []ByPodStatus `json:"byPod,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

// ConstraintTemplate is the Schema for the constrainttemplates API
// +k8s:openapi-gen=true
// +k8s:conversion-gen-external-types=github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates
type ConstraintTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConstraintTemplateSpec   `json:"spec,omitempty"`
	Status ConstraintTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConstraintTemplateList contains a list of ConstraintTemplate.
type ConstraintTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConstraintTemplate `json:"items"`
}

var gvrTemplates = schema.GroupVersionResource{
	Group:    "templates.gatekeeper.sh",
	Version:  "v1beta1",
	Resource: "constrainttemplates",
}

func GetDynamicClient(configk8s string) (dynamic.Interface, error) {
	config, err := GetRestConf(configk8s)
	if err != nil {
		return nil, err
	}

	dynamicClient, err := dynamic.NewForConfig(&config)
	if err != nil {
		return nil, err
	} else {
		return dynamicClient, nil
	}

}

func GetDiscoverityClient(configk8s string) (*discovery.DiscoveryClient, error) {
	config, err := GetRestConf(configk8s)
	if err != nil {
		return nil, err
	}
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(&config)
	return discoveryClient, nil
}

func ListConstraintTemplate(client dynamic.Interface) (*ConstraintTemplateList, error) {

	list, err := client.Resource(gvrTemplates).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	cts := &ConstraintTemplateList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(list.UnstructuredContent(), &cts)
	if err != nil {
		return nil, err
	}
	return cts, nil
}

func ConstraintTemplateExisted(client dynamic.Interface, ctname string) (bool, error) {
	flag := false
	list, err := client.Resource(gvrTemplates).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return flag, err
	}
	cts := &ConstraintTemplateList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(list.UnstructuredContent(), &cts)
	if err != nil {
		return flag, err
	}
	for i, _ := range cts.Items {
		if cts.Items[i].GetName() == ctname {
			flag = true
		}
	}
	return flag, nil
}

func GetConstraintTemplate(client dynamic.Interface, name string) (*ConstraintTemplate, error) {
	result, err := client.Resource(gvrTemplates).Get(context.TODO(), name, metav1.GetOptions{})
	ct := &ConstraintTemplate{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.UnstructuredContent(), &ct)
	if err != nil {
		return nil, err
	}
	return ct, err
}

func DeleteConstraintTemplate(client dynamic.Interface, name string) (bool, error) {
	err := client.Resource(gvrTemplates).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}

func CreateConstraintTemplate(yamlData string, k8sconfig string) (*ConstraintTemplate, error) {
	obj := &unstructured.Unstructured{}
	_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(yamlData), nil, obj)
	if err != nil {
		return nil, err
	}
	dr, err := GetGVRdyClient(gvk, k8sconfig)
	if err != nil {
		return nil, err
	}
	result, err := dr.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	ct := &ConstraintTemplate{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.UnstructuredContent(), &ct)
	if err != nil {
		return nil, err
	}
	return ct, nil

}

func UpdateConstraintTemplate(yamlData string, k8sconfig string) (*ConstraintTemplate, error) {
	obj := &unstructured.Unstructured{}
	_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(yamlData), nil, obj)
	if err != nil {
		return nil, err
	}
	dr, err := GetGVRdyClient(gvk, k8sconfig)
	if err != nil {
		return nil, err
	}
	utd, err := dr.Get(context.TODO(), obj.GetName(), metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	obj.SetResourceVersion(utd.GetResourceVersion())

	result, err := dr.Update(context.TODO(), obj, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	ct := &ConstraintTemplate{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.UnstructuredContent(), &ct)
	if err != nil {
		return nil, err
	}
	return ct, nil

}

type ClusterManagerModel struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	CreateTime  sql.NullString `db:"createtime"`
	File        sql.NullString `db:"file"`
	UpdateTime  sql.NullString `db:"updatetime"`
	Describtion sql.NullString `db:"describtion"`
}

func InitClientSet(filename string) (kubernetes.Clientset, error) {

	restConf, err := GetRestConf(filename)
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	// 生成clientset配置
	clientset, err := kubernetes.NewForConfig(&restConf)
	if err != nil {
		return kubernetes.Clientset{}, err
	}

	return *clientset, nil
}

func InitDynamicClient(filename string) {

}

// 获取k8s restful client配置
func GetRestConf(kubeconfigcode string) (rest.Config, error) {

	// 读kubeconfig文件
	kubeconfig := []byte(kubeconfigcode)
	// 生成rest client配置
	restConf, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return rest.Config{}, err
	}

	return *restConf, nil
}

type ConstraintRecordStruct struct {
	Kind string
	Name string
}

/**
result, err := client.Resource(gvrTemplates).Get(context.TODO(), name, metav1.GetOptions{})
	ct := &ConstraintTemplate{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.UnstructuredContent(), &ct)
	if err != nil {
		return nil, err
	}
*/

func GetConstraintes2(client dynamic.Interface, constrainttemplate string, constraintname string) (*Constraint.Constraint, error) {
	GvrConstraint.Resource = constrainttemplate
	result, err := client.Resource(GvrConstraint).Get(context.TODO(), constraintname, metav1.GetOptions{})

	constraint := &Constraint.Constraint{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.UnstructuredContent(), &constraint)
	if err != nil {
		return nil, err
	}

	return constraint, nil

}

func GetConstraintes(client dynamic.Interface, constrainttemplate string, constraintname string) (*unstructured.Unstructured, error) {
	GvrConstraint.Resource = constrainttemplate
	result, err := client.Resource(GvrConstraint).Get(context.TODO(), constraintname, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil

}
func DeleteConstraintes(client dynamic.Interface, constrainttemplate string, constraintname string) (bool, error) {
	GvrConstraint.Resource = constrainttemplate
	err := client.Resource(GvrConstraint).Delete(context.TODO(), constraintname, metav1.DeleteOptions{})
	if err != nil {
		return false, err

	} else {
		return true, nil
	}

}
func CreateConstraints(yamlData string, k8sconfig string) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(yamlData), nil, obj)
	if err != nil {
		return nil, err
	}
	dr, err := GetGVRdyClient(gvk, k8sconfig)
	if err != nil {
		return nil, err
	}
	result, err := dr.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateConstraintes(yamlData string, k8sconfig string) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(yamlData), nil, obj)
	if err != nil {
		return nil, err
	}
	dr, err := GetGVRdyClient(gvk, k8sconfig)
	if err != nil {
		return nil, err
	}
	utd, err := dr.Get(context.TODO(), obj.GetName(), metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	obj.SetResourceVersion(utd.GetResourceVersion())

	result, err := dr.Update(context.TODO(), obj, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil

}

func ListConstraintes(client dynamic.Interface) ([]ConstraintRecordStruct, error) {
	constraintTemplateList, err := ListConstraintTemplate(client)
	if err != nil {
		return nil, err
	}
	var constraintRecords []ConstraintRecordStruct
	for i, _ := range constraintTemplateList.Items {
		//constraintList, err := listConstraintByKind(constraintTemplateList.Items[i].Name, k8sconfig)
		GvrConstraint.Resource = constraintTemplateList.Items[i].Name
		constraintList, err := client.Resource(GvrConstraint).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for j, _ := range constraintList.Items {
			var temp ConstraintRecordStruct
			temp.Name = constraintList.Items[j].GetName()
			temp.Kind = constraintTemplateList.Items[i].Name
			constraintRecords = append(constraintRecords, temp)
		}
	}
	return constraintRecords, nil
}

var GvrConstraint = schema.GroupVersionResource{
	Group:    "constraints.gatekeeper.sh",
	Version:  "v1beta1",
	Resource: "",
}

func ListConstraintes2(client dynamic.Interface, constrainttemplatename string) ([]string, error) {
	GvrConstraint.Resource = constrainttemplatename
	constraintList, err := client.Resource(GvrConstraint).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var result []string
	for i, _ := range constraintList.Items {
		result = append(result, constraintList.Items[i].GetName())
	}
	return result, nil
}

func GetDynamicClientByGVK(gvk *schema.GroupVersionKind, discoveryClient *discovery.DiscoveryClient, dynamicClient dynamic.Interface) (dynamic.ResourceInterface, error) {
	mapperGVRGVK := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))
	// 根据资源GVK 获取资源的GVR GVK映射
	resourceMapper, err := mapperGVRGVK.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}
	return dynamicClient.Resource(resourceMapper.Resource), nil
}

func GetGVRdyClient(gvk *schema.GroupVersionKind, configk8s string) (dynamic.ResourceInterface, error) {
	config, err := GetRestConf(configk8s)
	if err != nil {
		return nil, err
	}
	// 创建discovery客户端
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(&config)
	if err != nil {
		return nil, err
	}
	// 获取GVK GVR 映射
	mapperGVRGVK := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

	resourceMapper, err := mapperGVRGVK.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}
	// 创建动态客户端
	dynamicClient, err := dynamic.NewForConfig(&config)
	if err != nil {
		return nil, err
	}
	return dynamicClient.Resource(resourceMapper.Resource), nil

}
