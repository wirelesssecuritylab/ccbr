package controllers

import (
	"ccbr/model/Clustermanager"
	"ccbr/model/ResponseStruct"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func CisBenchMarkControoler_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cisbenchmark.html", nil)
}

func CisBenchMarkControoler_List(context *gin.Context) {
	db := utils.InitDB()
	var clusterManagerModels []Clustermanager.ClusterManagerModel
	var clusterManagerModelView []Clustermanager.ClusterManagerModelView
	sql := "select * from cluster_manager"
	err := db.Select(&clusterManagerModels, sql)
	if err != nil {
		utils.Response(context, 0, "Get Cluster manager failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	for i := 0; i < len(clusterManagerModels); i++ {
		var temp Clustermanager.ClusterManagerModelView
		temp.Id = clusterManagerModels[i].Id
		temp.Name = clusterManagerModels[i].Name.String
		temp.Describtion = clusterManagerModels[i].Describtion.String
		temp.CreateTime = clusterManagerModels[i].CreateTime.String
		temp.UpdateTime = clusterManagerModels[i].UpdateTime.String
		temp.File = clusterManagerModels[i].File.String
		clusterManagerModelView = append(clusterManagerModelView, temp)
	}
	var result []ResponseStruct.CisBenchmarkInfo
	for i, _ := range clusterManagerModelView {
		clientSet, err := utils.InitClientSet(clusterManagerModelView[i].File)
		if err != nil {
			utils.Response(context, 0, "get clientSet error", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		pods, err := utils.ListPods(&clientSet, "kube-system")
		if err != nil {
			utils.Response(context, 0, "get namespace kube-system pods error", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		files, _ := ioutil.ReadDir("./rego/kubernetes/")
		for _, f := range files {
			cisbenchmarkResult, err := utils.RegoQuery(f.Name(), pods)
			if err != nil {
				utils.Response(context, 0, "Rego Query  error", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
			for j, _ := range cisbenchmarkResult {
				var temp ResponseStruct.CisBenchmarkInfo
				temp.ClusterManager = clusterManagerModelView[i].Name
				temp.PodName = cisbenchmarkResult[j].PodName
				temp.Result = cisbenchmarkResult[j].Result
				temp.CisBenchmarkItem = cisbenchmarkResult[j].CisBenchmarkItem
				temp.Message = cisbenchmarkResult[j].Message
				result = append(result, temp)
			}

		}

	}
	/*responseResult := ResponseStruct.ResponseCisBenchmarkCheckResult{
		Code: 0,
		Msg:  "",
		Data: result,
	}*/
	//context.JSON(http.StatusOK, responseResult)
	
	context.HTML(http.StatusOK, "cisbenchmark.html", gin.H{"result": result})
}
