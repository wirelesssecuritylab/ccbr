package controllers

import (
	"ccbr/model/Clustermanager"
	"ccbr/model/ResponseStruct"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func CMC_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cmconstraint.html", nil)
}
func CMC_List(context *gin.Context) {
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

	var result []ResponseStruct.ResponseClustermanagerAndConstraintTenplate
	for i, _ := range clusterManagerModelView {
		dynamicClient, err := utils.GetDynamicClient(clusterManagerModelView[i].File)
		if err != nil {
			utils.Response(context, 0, "Get Dynamic Client failure", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		constrainttemplatelist, err := utils.ListConstraintTemplate(dynamicClient)
		if err != nil {
			utils.Response(context, 0, "Get Dynamic Client failure", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		for j, _ := range constrainttemplatelist.Items {

			resconstraint, err := utils.ListConstraintes2(dynamicClient, constrainttemplatelist.Items[j].GetName())
			if err != nil {
				utils.Response(context, 0, "Get Constraint failure", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
			for k := 0; k < len(resconstraint); k++ {
				var temp ResponseStruct.ResponseClustermanagerAndConstraintTenplate
				temp.Clustername = clusterManagerModelView[i].Name
				temp.File = clusterManagerModelView[i].File
				temp.Constrainttemplatename = constrainttemplatelist.Items[j].GetName()
				temp.Constraint = resconstraint[k]
				result = append(result, temp)

			}

		}

	}
	response := ResponseStruct.ResponseClusterManagerConstraintTemplateStruct{
		Code:  0,
		Msg:   "",
		Count: -1,
		Data:  result,
	}
	context.JSON(http.StatusOK, response)
}

func CMC_Detail(context *gin.Context) {

	Constrainttemplatename := context.Query("Constrainttemplatename")
	Constraintname := context.Query("Constraintname")
	Clustername := context.Query("Clustername")
	db := utils.InitDB()
	var clusterManagerModels []Clustermanager.ClusterManagerModel

	sql := "select * from cluster_manager where name =?"
	err := db.Select(&clusterManagerModels, sql, Clustername)

	if err != nil {
		utils.Response(context, 0, "Get Cluster manager failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	dynamicClient, err := utils.GetDynamicClient(clusterManagerModels[0].File.String)

	if err != nil {
		utils.Response(context, 0, "Get Dynamic Client failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	result, err := utils.GetConstraintes2(dynamicClient, Constrainttemplatename, Constraintname)
	if err != nil {
		utils.Response(context, 0, "Get Constraint failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}

	context.HTML(http.StatusOK, "cmconstraintdetail.html", gin.H{"result": result})
}
