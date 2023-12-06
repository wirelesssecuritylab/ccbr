package cmd

import (
	"ccbr/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"strings"
)

var check string
var kubeconfig string

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	kubernetesCmd.Flags().StringVarP(&check, "check", "c", "all", "cis benchmark item")
	kubernetesCmd.Flags().StringVarP(&kubeconfig, "kubernetesconfig", "k", "~/.kube/config", "kubernetes configure file")

}

func parsedkubernetesArgs(k8sconfig string) (*kubernetes.Clientset, error) {

	config, err := utils.ParseConfig(k8sconfig)
	if err != nil {
		return nil, err
	}
	return config, nil

}

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "kubernetes cis benchmark check used rego",
	Long:  "kubernetes cis benchmark check used rego",
	Run: func(cmd *cobra.Command, args []string) {
		clientSet, err := parsedkubernetesArgs(kubeconfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		pods, err := utils.ListPods(clientSet, "kube-system")
		if err != nil {
			log.Fatal(err)
			return
		}
		if check != "all" {
			Checkitem(check, pods)
		} else {
			files, _ := ioutil.ReadDir("./rego/kubernetes/")
			for _, f := range files {
				Checkitem(f.Name(), pods)
			}

		}
	},
}

func inputArgsCheck(inputArgs string) bool {
	//k8s配置文件解析
	fileInfoList, err := ioutil.ReadDir("./rego/kubernetes/")
	if err != nil {
		log.Fatal(err)
	}
	var flag bool = false
	for _, file := range fileInfoList {
		if inputArgs == file.Name() {
			flag = true
			break
		}
	}
	return flag
}

func Checkitem(itemName string, pods *v1.PodList) {
	falg := strings.Contains(itemName, ",")
	if falg {
		item_array := strings.Split(itemName, ",")
		for _, str := range item_array {
			if true == inputArgsCheck(str) {
				utils.RegoQuery(str, pods)
			} else {
				log.Printf("input args has error: %v", str)
			}
		}
	} else {
		if true == inputArgsCheck(itemName) {
			utils.RegoQuery(itemName, pods)
		} else {
			log.Printf("input args has error: %v", itemName)
		}
	}
}
