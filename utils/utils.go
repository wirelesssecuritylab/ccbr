package utils

import (
	"bytes"
	"ccbr/model/Constraint"
	"ccbr/model/ConstraintList"
	"ccbr/model/ResponseStruct"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"os/user"
	"reflect"
	"strings"
)

var gvr = schema.GroupVersionResource{
	Group:    "templates.gatekeeper.sh",
	Version:  "v1beta1",
	Resource: "constrainttemplates",
}
var gvr1 = schema.GroupVersionResource{
	Group:   "constraints.gatekeeper.sh",
	Version: "v1beta1",
}

var gvr2 = schema.GroupVersionResource{
	Group:    "constraints.gatekeeper.sh",
	Version:  "v1beta1",
	Resource: "",
}

func ListPods(c *kubernetes.Clientset, ns string) (*v1.PodList, error) {
	pods, err := c.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {

		return nil, err
	}
	/*for _, v := range pods.Items {
		fmt.Printf("namespace: %v podname: %v podstatus: %v \n", v.Namespace, v.Name, v.Status.Phase)
	}*/
	return pods, nil
}

func ListNodes(c *kubernetes.Clientset) error {
	nodeList, err := c.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, node := range nodeList.Items {
		fmt.Printf("nodeName: %v, status: %v", node.GetName(), node.GetCreationTimestamp())
	}
	return nil
}
func ListDeployment(c *kubernetes.Clientset, ns string) error {
	deployments, err := c.AppsV1().Deployments(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, v := range deployments.Items {
		fmt.Printf("deploymentname: %v, available: %v, ready: %v", v.GetName(), v.Status.AvailableReplicas, v.Status.ReadyReplicas)
	}
	return nil
}
func DescribePod(c *kubernetes.Clientset, ns string, podName string) (string, error) {
	podInfo, err := c.CoreV1().Pods(ns).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	bs, _ := json.Marshal(podInfo)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	return out.String(), nil
	//fmt.Printf("podInfo=%v\n", out.String())

}

func ParseConfig(kubeconfig string) (*kubernetes.Clientset, error) {
	var config string
	if kubeconfig == "~/.kube/config" {
		u, err := user.Current()
		if err != nil {
			fmt.Println("error")
		}
		config = u.HomeDir + "/.kube/config"

	} else {
		config = kubeconfig
	}

	config_, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return nil, err
	}
	// 生成clientSet
	clientSet, err := kubernetes.NewForConfig(config_)
	if err != nil {
		return clientSet, err
	}
	return clientSet, nil
}

/**
 */

func ParserArgs(kubeconfig string) string {
	var config string
	if kubeconfig == "~/.kube/config" {
		u, err := user.Current()
		if err != nil {
			fmt.Println("error")
		}
		config = u.HomeDir + "/.kube/config"

	} else {
		config = kubeconfig
	}
	return config
}
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
func ListConstraint(dynamicClient dynamic.Interface) ConstraintList.ConstraintList {
	list, err := dynamicClient.Resource(gvr1).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := json.Marshal(list)
	var data ConstraintList.ConstraintList
	err_1 := json.Unmarshal(bs, &data)
	if err_1 != nil {
		log.Fatal(err_1)
	}
	return data
}

func GetConstraint(dynamicClient dynamic.Interface, constraintName string) Constraint.Constraint {
	gvr2.Resource = constraintName
	list, err := dynamicClient.Resource(gvr2).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return Constraint.Constraint{}
	} else {
		bs, _ := json.Marshal(list)
		var data Constraint.Constraint
		err_1 := json.Unmarshal(bs, &data)
		if err_1 != nil {
			return Constraint.Constraint{}
		}
		return data
	}
	return Constraint.Constraint{}
}

func Response_Ajax(row int64, ctx *gin.Context, msg string) {

	if row > 0 {

		messgae_map := map[string]interface{}{
			"code": 200,
			"msg":  msg + "成功",
		}
		ctx.JSON(http.StatusOK, messgae_map)
	} else {

		messgae_map := map[string]interface{}{
			"code": 400,
			"msg":  msg + "失败",
		}
		ctx.JSON(http.StatusBadRequest, messgae_map)
	}
}

func Response_Error(ctx *gin.Context) {
	messgae_map := map[string]interface{}{
		"code": 400,
		"msg":  "操作失败",
	}
	ctx.JSON(http.StatusBadRequest, messgae_map)
}

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	messgae_map := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	ctx.JSON(http.StatusOK, messgae_map)
}
func FindName(file string) string {
	firstSub := string([]rune(file)[strings.Index(file, "metadata:"):strings.Index(file, "spec:")])
	secondSub := firstSub[strings.Index(firstSub, "name:"):len(firstSub)]
	threeSub := secondSub[strings.Index(secondSub, "name:"):strings.Index(secondSub, "\n")]
	forSub := threeSub[strings.Index(threeSub, ":")+1 : len(threeSub)]
	return strings.Replace(forSub, " ", "", -1)
}
func RegoQueryExecute(pods *v1.PodList, podNameStr string, queryitem string, query rego.PreparedEvalQuery, context context.Context) ([]ResponseStruct.CisBenchmarkInfo, error) {
	var result []ResponseStruct.CisBenchmarkInfo
	for _, v := range pods.Items {
		if strings.Contains(v.Name, podNameStr) {
			var temp ResponseStruct.CisBenchmarkInfo
			cisbenchmarkitem, err := Query2(v, queryitem, query, context)
			if err != nil {
				return nil, err
			}
			temp.PodName = v.Name
			temp.Result = cisbenchmarkitem.Result
			temp.CisBenchmarkItem = cisbenchmarkitem.CisBenchmarkItem
			temp.Message = cisbenchmarkitem.Message
			result = append(result, temp)

		}

	}
	return result, nil
}

func Query2(inputdata interface{}, queryitem string, query rego.PreparedEvalQuery, ctx context.Context) (*ResponseStruct.CisBenchmarkInfo, error) {
	var resultRes ResponseStruct.CisBenchmarkInfo
	resultRes.CisBenchmarkItem = queryitem
	rs, err := query.Eval(ctx, rego.EvalInput(inputdata))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if len(rs) > 0 {
		result := rs[0].Bindings["msg"].(string)
		resultRes.Message = result
		start := strings.IndexByte(result, '(')
		flag := result[start+1 : len(result)-1]
		if flag == "Manual" {
			yellow := color.New(color.FgYellow).PrintfFunc()
			yellow("[WARN]")
			fmt.Println(" " + queryitem + ": " + result)
			resultRes.Result = "[WARN]"

		} else if flag == "Automated" {
			red := color.New(color.FgRed).PrintfFunc()
			red("[FAIL]")
			fmt.Println(" " + queryitem + ": " + result)
			resultRes.Result = "[FAIL]"
		}
		return &resultRes, nil
	} else {
		green := color.New(color.FgGreen).PrintfFunc()
		green("[PASS]")
		fmt.Println(" " + queryitem + ": CHECK PASS")
		resultRes.Result = "[PASS]"
		resultRes.Message = "CHECK PASS"
		return &resultRes, nil
	}
}

func RegoQuery(queryitem string, pods *v1.PodList) ([]ResponseStruct.CisBenchmarkInfo, error) {

	ctx := context.Background()
	var queryCondition string = "msg =data." + queryitem + ".violation[msg]"
	var regoPath string = "./rego/kubernetes/" + queryitem + "/" + queryitem + ".rego"
	//var regoInput string = "./rego/kubernetes/" + queryitem +"/" + queryitem + "_input.json"
	r := rego.New(
		rego.Query(queryCondition),
		rego.Load([]string{regoPath}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var result []ResponseStruct.CisBenchmarkInfo
	if strings.Contains(queryitem, "cis_1_2_") {
		//kube-apiserver
		result, err = RegoQueryExecute(pods, "kube-apiserver", queryitem, query, ctx)

	} else if strings.Contains(queryitem, "cis_1_3_") {
		//kube-controller-manager
		result, err = RegoQueryExecute(pods, "kube-controller-manager", queryitem, query, ctx)
	} else if strings.Contains(queryitem, "cis_1_4_") {
		//kube-scheduler
		result, err = RegoQueryExecute(pods, "kube-scheduler", queryitem, query, ctx)
	} else if strings.Contains(queryitem, "cis_2_") {
		//etcd
		result, err = RegoQueryExecute(pods, "etcd", queryitem, query, ctx)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil

}
