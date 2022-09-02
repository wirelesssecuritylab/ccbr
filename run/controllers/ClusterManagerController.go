package controllers

import (
	"ccbr/model/Clustermanager"
	"ccbr/model/ResponseStruct"
	"ccbr/run/services"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func ClusterManager_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "clustermanager.html", nil)
}
func ClusterManager_Update(context *gin.Context) {
	Id := context.PostForm("Id")
	Name := context.PostForm("Name")
	File := context.PostForm("File")
	Describtion := context.PostForm("Describtion")
	Updatetime := time.Now().Format("2006-01-02 15:04:05")
	db := utils.InitDB()
	if (Id != "") && (Name != "") && (File != "") {
		sql := "update cluster_manager set name =?,describtion=?,file=?,updatetime=? where id= ?"
		res, err := db.Exec(sql, Name, Describtion, File, Updatetime, Id)

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

func ClusterManager_Save(context *gin.Context) {
	db := utils.InitDB()
	Name := context.PostForm("Name")
	CreateTime := context.PostForm("CreateTime")
	File := context.PostForm("File")
	Updatetime := context.PostForm("CreateTime")
	Describtion := context.PostForm("Describtion")
	if (Name != "") && (File != "") && (CreateTime != "") {
		sql := "insert into cluster_manager(name,describtion,createtime,file,updatetime)values (?,?,?,?,?)"
		//执行SQL语句
		r, err := db.Exec(sql, Name, Describtion, CreateTime, File, Updatetime)
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

func ClusterManager_Delete(context *gin.Context) {
	Id := context.PostForm("Id")
	db := utils.InitDB()
	sql := "delete from cluster_manager where id=?"

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

func ClusterManager_List(context *gin.Context) {
	db := utils.InitDB()
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))

	Name := context.Query("Name")
	var clusterManagerModels []Clustermanager.ClusterManagerModel
	var clusterManagerModelView []Clustermanager.ClusterManagerModelView
	var sql string
	var err error
	var count int
	if Name != "" {
		sql = "select * from cluster_manager where name like ? limit ?,?"
		err = db.Select(&clusterManagerModels, sql, "%"+Name+"%", (page-1)*limit, limit)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())

		}
		err = db.QueryRow("select count(*) from cluster_manager where name like ?", "%"+Name+"%").Scan(&count)

	} else {
		sql = "select * from cluster_manager limit ?,?"
		err = db.Select(&clusterManagerModels, sql, (page-1)*limit, limit)

		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from cluster_manager").Scan(&count)

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
	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	result := ResponseStruct.ResponseClusterManagerStruct{
		Code:  0,
		Msg:   "",
		Count: count,
		Data:  clusterManagerModelView,
	}
	context.JSON(http.StatusOK, result)

}
func ClusterManager_Test(context *gin.Context) {
	Name := context.PostForm("Name")
	Testitem := context.PostForm("Testitem")
	File := context.PostForm("File")

	if Testitem == "Kubernetes" {
		result, err := services.ClusterManagerService_Test_Kubernetes(File, Name)
		if err != nil {
			messgae_map := map[string]interface{}{
				"code":  400,
				"msg":   "Kubernetes 环境测试失败",
				"count": 100,
				"data":  nil,
			}
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			context.JSON(http.StatusOK, messgae_map)
		} else {
			messgae_map := map[string]interface{}{
				"code":  200,
				"msg":   "Kubernetes 环境测试成功",
				"count": 100,
				"data":  result,
			}
			context.JSON(http.StatusOK, messgae_map)
		}

	} else if Testitem == "Gatekeeper" {
		result, err := services.ClusterManagerService_Test_Gatekeeper(File, Name)
		if err != nil {
			messgae_map := map[string]interface{}{
				"code":  400,
				"msg":   "Gatekeeper 环境测试失败",
				"count": 100,
				"data":  nil,
			}
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			context.JSON(http.StatusOK, messgae_map)
		} else {
			messgae_map := map[string]interface{}{
				"code":  200,
				"msg":   "Gatekeeper 环境测试成功",
				"count": 100,
				"data":  result,
			}
			context.JSON(http.StatusOK, messgae_map)
		}
	}

}
