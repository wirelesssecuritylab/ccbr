package controllers

import (
	"ccbr/model/Gatekeeper"
	"ccbr/model/ResponseStruct"
	"ccbr/run/services"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func OpaPolicies_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "opapolicies.html", nil)
}
func OpaPolicies_Update(context *gin.Context) {
	Id := context.PostForm("Id")
	Name := context.PostForm("Name")
	Type := context.PostForm("Type")
	Version := context.PostForm("Version")
	Constraintlist := context.PostForm("Constraintlist")
	Process := context.PostForm("Process")
	Describtion := context.PostForm("Describtion")
	Clustername := context.PostForm("Clustername")
	Updatetime := time.Now().Format("2006-01-02 15:04:05")

	db := utils.InitDB()
	if (Id != "") && (Name != "") && (Type != "") && (Version != "") && (Process != "") && (Constraintlist != "") {
		sql := "update opa_gatekeeper_policies set name =?,version=?,constraintlist=?,process=?,describtion=?,updatetime=?,type=?,Clustername=? where id= ?"
		res, err := db.Exec(sql, Name, Version, Constraintlist, Process, Describtion, Updatetime, Type, Clustername, Id)

		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			utils.Response_Error(context)
		}

		//查询影响的行数，判断修改插入成功
		row, err := res.RowsAffected()
		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			utils.Response_Error(context)
		}
		utils.Response_Ajax(row, context, "更新")
	} else {
		utils.Response_Error(context)
	}

}

func OpaPolicies_Save(context *gin.Context) {
	db := utils.InitDB()
	Name := context.PostForm("Name")
	Version := context.PostForm("Version")
	Constraintlist := context.PostForm("Constraintlist")
	Process := context.PostForm("Process")
	Describtion := context.PostForm("Describtion")
	Clustername := context.PostForm("Clustername")
	CreateTime := context.PostForm("CreateTime")
	Updatetime := context.PostForm("CreateTime")
	Type := context.PostForm("Type")
	if (Name != "") && (Type != "") && (Version != "") && (CreateTime != "") && (Process != "") && (Constraintlist != "") {
		sql := "insert into opa_gatekeeper_policies(name,version,constraintlist, process,describtion,createtime,updatetime,type,status,clustername)values (?,?,?,?,?,?,?,?,?,?)"
		//执行SQL语句
		r, err := db.Exec(sql, Name, Version, Constraintlist, Process, Describtion, CreateTime, Updatetime, Type, "新建策略", Clustername)
		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			utils.Response_Error(context)
		}

		//判断是否插入成功
		id, err := r.LastInsertId()
		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			utils.Response_Error(context)
		}
		utils.Response_Ajax(id, context, "保存")
	} else {
		utils.Response_Error(context)
	}
}

func OpaPolicies_Delete(context *gin.Context) {
	Id := context.PostForm("Id")
	db := utils.InitDB()
	sql := "delete from opa_gatekeeper_policies where id=?"

	res, err := db.Exec(sql, Id)
	if err != nil {
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		utils.Response_Error(context)
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}
	utils.Response_Ajax(row, context, "删除")
}

func OpaPolicies_List(context *gin.Context) {
	db := utils.InitDB()
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))

	Name := context.Query("Name")
	var opaGatekeeperPolicies []Gatekeeper.OpaGatekeeperPolicies
	var opaGatekeeperPolicies2 []Gatekeeper.OpaGatekeeperPolicies2
	var sql string
	var err error
	var count int
	if Name != "" {
		sql = "select * from opa_gatekeeper_policies where name like ? limit ?,?"
		err = db.Select(&opaGatekeeperPolicies, sql, "%"+Name+"%", (page-1)*limit, limit)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_policies where name like ?", "%"+Name+"%").Scan(&count)

	} else {
		sql = "select * from opa_gatekeeper_policies limit ?,?"
		err = db.Select(&opaGatekeeperPolicies, sql, (page-1)*limit, limit)

		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_policies").Scan(&count)

	}

	for i := 0; i < len(opaGatekeeperPolicies); i++ {
		var opaGatekeeperPolicie Gatekeeper.OpaGatekeeperPolicies2
		opaGatekeeperPolicie.Id = opaGatekeeperPolicies[i].Id
		opaGatekeeperPolicie.Name = opaGatekeeperPolicies[i].Name.String
		opaGatekeeperPolicie.Version = opaGatekeeperPolicies[i].Version.String
		opaGatekeeperPolicie.Constraintlist = opaGatekeeperPolicies[i].Constraintlist.String
		opaGatekeeperPolicie.Process = opaGatekeeperPolicies[i].Process.String
		opaGatekeeperPolicie.Describtion = opaGatekeeperPolicies[i].Describtion.String
		opaGatekeeperPolicie.CreateTime = opaGatekeeperPolicies[i].CreateTime.String
		opaGatekeeperPolicie.UpdateTime = opaGatekeeperPolicies[i].UpdateTime.String
		opaGatekeeperPolicie.Type = opaGatekeeperPolicies[i].Type.String
		opaGatekeeperPolicie.Status = opaGatekeeperPolicies[i].Status.String
		opaGatekeeperPolicie.Clustername = opaGatekeeperPolicies[i].Clustername.String

		opaGatekeeperPolicies2 = append(opaGatekeeperPolicies2, opaGatekeeperPolicie)
	}
	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	result := ResponseStruct.ResponseOPAPoliciesStruct{
		Code:  0,
		Msg:   "",
		Count: count,
		Data:  opaGatekeeperPolicies2,
	}
	context.JSON(http.StatusOK, result)

}

func OpaPolicies_ListConstraint(context *gin.Context) {
	db := utils.InitDB()
	var Constraint_Name []string
	sql := "select name from opa_gatekeeper_constraint"
	err := db.Select(&Constraint_Name, sql)

	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}
	var resultData []ResponseStruct.ResonseOPAConstraintMultiSelect
	for i := 0; i < len(Constraint_Name); i++ {
		var temp ResponseStruct.ResonseOPAConstraintMultiSelect
		temp.Name = Constraint_Name[i]
		temp.Value = i + 1
		temp.Selected = false
		temp.Disabled = false
		resultData = append(resultData, temp)
	}

	messgae_map := map[string]interface{}{
		"code": 200,
		"msg":  "操作成功",
		"data": resultData,
	}
	context.JSON(http.StatusOK, messgae_map)

}

func OpaPolicies_Deploy(context *gin.Context) {
	db := utils.InitDB()
	Constraintlist := context.PostForm("Constraintlist")
	Clustername := context.PostForm("Clustername")
	Id := context.PostForm("Id")
	var cmct []Gatekeeper.ConstraintMapperConstraintTemplate
	k8sconfig, err := services.GetK8sconfigByClusterName(db, Clustername)
	if err != nil {
		utils.Response(context, 0, "Get k8sconfig error", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	// operation etcd databased
	dynamicClient, err := utils.GetDynamicClient(k8sconfig)
	if err != nil {
		utils.Response(context, 0, "get dynamic client error", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	constrainttemplates, err := utils.ListConstraintTemplate(dynamicClient)
	if err != nil {
		utils.Response(context, 0, "list constraint template error", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	for i := 0; i < len(strings.Split(Constraintlist, ",")); i++ {
		//Constraint Name
		var temp Gatekeeper.ConstraintMapperConstraintTemplate
		ConstraintName := strings.Split(Constraintlist, ",")[i]
		constraintTemplateName, err := services.GetConstraintTemplateNameByConstraintName(db, ConstraintName)
		if err != nil {
			utils.Response(context, 0, "get constraint template by constraint name error", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		temp.Constraint = ConstraintName
		temp.ConstraintTemplate = constraintTemplateName
		cmct = append(cmct, temp)
	}

	for i, _ := range cmct {
		//utils.

		//创建约束模板
		flag := false
		for j, _ := range constrainttemplates.Items {
			if constrainttemplates.Items[j].Name == cmct[i].ConstraintTemplate {
				flag = true
				break
			}
		}
		if !flag {
			//create constraint template
			yamlFile, err := services.GetConstrsintTemplateFileByName(db, cmct[i].ConstraintTemplate)
			if err != nil {
				utils.Response(context, 0, "Get constraint template yaml file error", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
			_, err = utils.CreateConstraintTemplate(yamlFile, k8sconfig)
			if err != nil {
				utils.Response(context, 0, "create constraint template error", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
		}

		constrainttemplate, err := utils.GetConstraintTemplate(dynamicClient, cmct[i].ConstraintTemplate)
		if err != nil {
			utils.Response(context, 0, "Get constraint template error", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		if constrainttemplate != nil {
			constraintlist, err := utils.ListConstraintes2(dynamicClient, cmct[i].ConstraintTemplate)
			if err != nil {
				utils.Response(context, 0, "list constraint error", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}

			flag2 := false
			for k, _ := range constraintlist {
				if constraintlist[k] == cmct[i].Constraint {
					flag2 = true
					break
				}
			}
			if !flag2 {
				yamlFile, err := services.GetConstrsintFileByName(db, cmct[i].Constraint)
				if err != nil {
					utils.Response(context, 0, "Get constraint yaml file error", nil)
					logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
					return
				}
				_, err = utils.CreateConstraints(yamlFile, k8sconfig)
				if err != nil {
					utils.Response(context, 0, "create constraint error", nil)
					logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
					return
				}
			}
		}

		//创建约束

	}

	flag3, err := services.UpdateOPAPoliciesStatus(db, "策略部署成功", Id)
	if err != nil {
		utils.Response(context, 0, "status update error", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	if flag3 {
		utils.Response(context, 0, "部署成功", nil)
	} else {
		utils.Response(context, 0, "部署失败", nil)
	}

}
func OpaPolicies_Stop(context *gin.Context) {
	db := utils.InitDB()
	Constraintlist := context.PostForm("Constraintlist")
	Clustername := context.PostForm("Clustername")
	Id := context.PostForm("Id")
	k8sconfig, err := services.GetK8sconfigByClusterName(db, Clustername)
	var cmct []Gatekeeper.ConstraintMapperConstraintTemplate
	if err != nil {
		utils.Response(context, 0, "obtain k8sconfigure failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}

	for i := 0; i < len(strings.Split(Constraintlist, ",")); i++ {
		//Constraint Name
		var temp Gatekeeper.ConstraintMapperConstraintTemplate
		ConstraintName := strings.Split(Constraintlist, ",")[i]
		constraintTemplateName, err := services.GetConstraintTemplateNameByConstraintName(db, ConstraintName)
		if err != nil {
			utils.Response(context, 0, "obtain constraint template failure", nil)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		temp.Constraint = ConstraintName
		temp.ConstraintTemplate = constraintTemplateName
		cmct = append(cmct, temp)
	}
	dynamicClient, err := utils.GetDynamicClient(k8sconfig)
	if err != nil {
		utils.Response(context, 0, "create dynamicClient failure", nil)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		return
	}
	for i, _ := range cmct {
		//delete constraint
		result, err := utils.GetConstraintes(dynamicClient, cmct[i].ConstraintTemplate, cmct[i].Constraint)

		if result != nil {
			flag, err := utils.DeleteConstraintes(dynamicClient, cmct[i].ConstraintTemplate, cmct[i].Constraint)
			if err != nil || !flag {
				utils.Response(context, 0, "delete Constrainte failure", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
		}
		//delete constraint template

		res, err := utils.ListConstraintes2(dynamicClient, cmct[i].ConstraintTemplate)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			return
		}
		if len(res) == 0 {
			flag, err := utils.DeleteConstraintTemplate(dynamicClient, cmct[i].ConstraintTemplate)
			if err != nil || !flag {
				utils.Response(context, 0, "Delete Constrainte template failure", nil)
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
				return
			}
		}

	}
	flag2, err := services.UpdateOPAPoliciesStatus(db, "策略停止成功", Id)
	if flag2 && err == nil {
		utils.Response(context, 0, "停止策略成功", nil)
	} else {
		utils.Response(context, 0, "停止策略失败", nil)

	}

}

func OpaPolicies_ListClusterManager(context *gin.Context) {
	db := utils.InitDB()
	var Clustermanager_Name []string
	sql := "select name from cluster_manager"
	err := db.Select(&Clustermanager_Name, sql)

	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	messgae_map := map[string]interface{}{
		"code": 200,
		"msg":  "操作成功",
		"data": Clustermanager_Name,
	}
	context.JSON(http.StatusOK, messgae_map)
}
