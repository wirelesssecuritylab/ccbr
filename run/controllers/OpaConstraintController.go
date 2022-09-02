package controllers

import (
	"ccbr/model/Gatekeeper"
	"ccbr/model/ResponseStruct"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func OpaC_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "opaconstraint.html", nil)
}
func OpaC_Update(context *gin.Context) {
	Id := context.PostForm("Id")
	Name := context.PostForm("Name")
	Ctname := context.PostForm("Ctname")
	Type := context.PostForm("Type")
	PackageType := context.PostForm("PackageType")
	File := context.PostForm("File")
	Describtion := context.PostForm("Describtion")
	Updatetime := time.Now().Format("2006-01-02 15:04:05")
	if strings.Replace(Name, " ", "", -1) != utils.FindName(File) {
		utils.Response(context, 0, "Constraint Name Different", nil)
	}
	db := utils.InitDB()
	if (Id != "") && (Name != "") && (Type != "") && (PackageType != "") && (File != "") && (Ctname != "") {
		sql := "update opa_gatekeeper_constraint set name =?,type=?,ctname=?,describtion=?,packagetype=?,file=?,updatetime=? where id= ?"
		res, err := db.Exec(sql, Name, Type, Ctname, Describtion, PackageType, File, Updatetime, Id)

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

func OpaC_Save(context *gin.Context) {
	db := utils.InitDB()
	Name := context.PostForm("Name")
	Ctname := context.PostForm("Ctname")
	Type := context.PostForm("Type")
	PackageType := context.PostForm("PackageType")
	CreateTime := context.PostForm("CreateTime")
	File := context.PostForm("File")
	Updatetime := context.PostForm("CreateTime")
	Describtion := context.PostForm("Describtion")

	if strings.Replace(Name, " ", "", -1) != utils.FindName(File) {
		utils.Response(context, 0, "Constraint Name Different", nil)
	}
	if (Name != "") && (Type != "") && (PackageType != "") && (File != "") && (Ctname != "") && (CreateTime != "") {
		sql := "insert into opa_gatekeeper_constraint(name,ctname,type, describtion,packagetype,createtime,file,updatetime)values (?,?,?,?,?,?,?,?)"
		//执行SQL语句
		r, err := db.Exec(sql, Name, Ctname, Type, Describtion, PackageType, CreateTime, File, Updatetime)
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

func OpaC_Delete(context *gin.Context) {
	Id := context.PostForm("Id")
	db := utils.InitDB()
	sql := "delete from opa_gatekeeper_constraint where id=?"

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

func OpaC_List(context *gin.Context) {
	db := utils.InitDB()
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))
	Name := context.Query("Name")
	var opaGatekeeperConstraints []Gatekeeper.OpaGatekeeperConstraint
	var opaGatekeeperConstraints2 []Gatekeeper.OpaGatekeeperConstraint2
	var sql string
	var err error
	var count int
	if Name != "" {
		sql = "select * from opa_gatekeeper_constraint where name like ? limit ?,?"
		err = db.Select(&opaGatekeeperConstraints, sql, "%"+Name+"%", (page-1)*limit, limit)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_constraint where name like ?", "%"+Name+"%").Scan(&count)

	} else {
		sql = "select * from opa_gatekeeper_constraint limit ?,?"
		err = db.Select(&opaGatekeeperConstraints, sql, (page-1)*limit, limit)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_constraint").Scan(&count)

	}

	for i := 0; i < len(opaGatekeeperConstraints); i++ {
		var opaGatekeeperConstraint Gatekeeper.OpaGatekeeperConstraint2
		opaGatekeeperConstraint.Id = opaGatekeeperConstraints[i].Id
		opaGatekeeperConstraint.Name = opaGatekeeperConstraints[i].Name.String
		opaGatekeeperConstraint.Ctname = opaGatekeeperConstraints[i].Ctname.String
		opaGatekeeperConstraint.Type = opaGatekeeperConstraints[i].Type.String
		opaGatekeeperConstraint.PackageType = opaGatekeeperConstraints[i].PackageType.String
		opaGatekeeperConstraint.CreateTime = opaGatekeeperConstraints[i].CreateTime.String
		opaGatekeeperConstraint.UpdateTime = opaGatekeeperConstraints[i].UpdateTime.String
		opaGatekeeperConstraint.File = opaGatekeeperConstraints[i].File.String
		opaGatekeeperConstraint.Describtion = opaGatekeeperConstraints[i].Describtion.String

		opaGatekeeperConstraints2 = append(opaGatekeeperConstraints2, opaGatekeeperConstraint)
	}
	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	result := ResponseStruct.ResponseOPAGConstraintStruct{
		Code:  0,
		Msg:   "",
		Count: count,
		Data:  opaGatekeeperConstraints2,
	}
	context.JSON(http.StatusOK, result)

}

func OpaC_List_CT(context *gin.Context) {
	db := utils.InitDB()
	var CT_Name []string
	sql := "select name from opa_gatekeeper_constrainttemplate"
	err := db.Select(&CT_Name, sql)

	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}
	messgae_map := map[string]interface{}{
		"code": 400,
		"msg":  "操作失败",
		"data": CT_Name,
	}
	context.JSON(http.StatusOK, messgae_map)

}
