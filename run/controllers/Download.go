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

func DOMNLOAD_Index(context *gin.Context) {
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

	var result []ResponseStruct.Report
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

				result_constraint, err := utils.GetConstraintes2(dynamicClient, constrainttemplatelist.Items[j].GetName(), resconstraint[k])
				if err != nil {
					utils.Response(context, 0, "Get Constraint failure", nil)
					logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
					return
				}

				for z, _ := range result_constraint.Status.Violations {
					var temp ResponseStruct.Report
					temp.ClusterManager = clusterManagerModelView[i].Name
					temp.ConstraintTemplate = constrainttemplatelist.Items[j].GetName()
					temp.Constraint = resconstraint[k]
					temp.Action = result_constraint.Status.Violations[z].EnforcementAction
					temp.Kind = result_constraint.Status.Violations[z].Kind
					temp.Namespace = result_constraint.Status.Violations[z].Namespace
					temp.Name = result_constraint.Status.Violations[z].Name
					temp.Message = result_constraint.Status.Violations[z].Message
					result = append(result, temp)
				}
			}

		}

	}

	context.HTML(http.StatusOK, "downloadreport.html", gin.H{"result": result, "length": len(result)})
}
