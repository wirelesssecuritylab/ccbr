package controllers

import (
	"ccbr/model/Gatekeeper"
	"ccbr/model/ResponseStruct"
	"ccbr/run/response"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func OpaCT_Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "opaconstrainttemplate.html", nil)
}
func OpaCT_Update(context *gin.Context) {
	Id := context.PostForm("Id")
	Name := context.PostForm("Name")
	Type := context.PostForm("Type")
	PackageType := context.PostForm("PackageType")
	File := context.PostForm("File")
	Describtion := context.PostForm("Describtion")
	Updatetime := time.Now().Format("2006-01-02 15:04:05")

	if strings.Replace(Name, " ", "", -1) != utils.FindName(File) {
		utils.Response(context, 0, "Constraint Template Name Different", nil)
	}

	db := utils.InitDB()
	if (Id != "") && (Name != "") && (Type != "") && (PackageType != "") && (File != "") {
		sql := "update opa_gatekeeper_constrainttemplate set name =?,type=?,describtion=?,packagetype=?,file=?,updatetime=? where id= ?"
		res, err := db.Exec(sql, Name, Type, Describtion, PackageType, File, Updatetime, Id)

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

func OpaCT_Save(context *gin.Context) {
	db := utils.InitDB()
	Name := context.PostForm("Name")
	Type := context.PostForm("Type")
	PackageType := context.PostForm("PackageType")
	CreateTime := context.PostForm("CreateTime")
	File := context.PostForm("File")
	Updatetime := context.PostForm("CreateTime")
	Describtion := context.PostForm("Describtion")
	if strings.Replace(Name, " ", "", -1) != utils.FindName(File) {
		utils.Response(context, 0, "Constraint Template Name Different", nil)
	}
	if (Name != "") && (Type != "") && (PackageType != "") && (File != "") && (CreateTime != "") {
		sql := "insert into opa_gatekeeper_constrainttemplate(name,type,describtion, packagetype,createtime,file,updatetime)values (?,?,?,?,?,?,?)"
		//执行SQL语句
		r, err := db.Exec(sql, Name, Type, Describtion, PackageType, CreateTime, File, Updatetime)
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

func OpaCT_Delete(context *gin.Context) {
	Id := context.PostForm("Id")
	Name := context.PostForm("Name")
	db := utils.InitDB()
	sql := "delete from opa_gatekeeper_constrainttemplate where id=?"

	res, err := db.Exec(sql, Id)
	if err != nil {
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		utils.Response_Error(context)
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	//delete constraint
	var opaGatekeeperConstraints []Gatekeeper.OpaGatekeeperConstraint
	constraintSelectSql := "select * from opa_gatekeeper_constraint where ctname=?"
	err = db.Select(&opaGatekeeperConstraints, constraintSelectSql, Name)
	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}
	if len(opaGatekeeperConstraints) > 0 {
		constraintDelSql := "delete from opa_gatekeeper_constraint where ctname=?"

		res, err := db.Exec(constraintDelSql, Name)
		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			utils.Response_Error(context)
		}

		row_1, err := res.RowsAffected()
		if err != nil {
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		logrus.Println(time.Now().Format("2006-01-02 15:04:05") + " " + strconv.FormatInt(row_1, 10))

	}

	utils.Response_Ajax(row, context, "删除")
}

func OpaCT_List(context *gin.Context) {
	db := utils.InitDB()
	page, _ := strconv.Atoi(context.Query("page"))
	limit, _ := strconv.Atoi(context.Query("limit"))

	Name := context.Query("Name")
	var opaGatekeeperConstraintTemplaters []Gatekeeper.OpaGatekeeperConstraintTemplater
	var opaGatekeeperConstraintTemplaters2 []Gatekeeper.OpaGatekeeperConstraintTemplater2
	var sql string
	var err error
	var count int
	if Name != "" {
		sql = "select * from opa_gatekeeper_constrainttemplate where name like ? limit ?,?"
		err = db.Select(&opaGatekeeperConstraintTemplaters, sql, "%"+Name+"%", (page-1)*limit, limit)
		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_constrainttemplate where name like ?", "%"+Name+"%").Scan(&count)

	} else {
		sql = "select * from opa_gatekeeper_constrainttemplate limit ?,?"
		err = db.Select(&opaGatekeeperConstraintTemplaters, sql, (page-1)*limit, limit)

		if err != nil {
			utils.Response_Error(context)
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		err = db.QueryRow("select count(*) from opa_gatekeeper_constrainttemplate").Scan(&count)

	}

	for i := 0; i < len(opaGatekeeperConstraintTemplaters); i++ {
		var opaGatekeeperConstraintTemplater Gatekeeper.OpaGatekeeperConstraintTemplater2
		opaGatekeeperConstraintTemplater.Id = opaGatekeeperConstraintTemplaters[i].Id
		opaGatekeeperConstraintTemplater.Name = opaGatekeeperConstraintTemplaters[i].Name.String
		opaGatekeeperConstraintTemplater.Type = opaGatekeeperConstraintTemplaters[i].Type.String
		opaGatekeeperConstraintTemplater.PackageType = opaGatekeeperConstraintTemplaters[i].PackageType.String
		opaGatekeeperConstraintTemplater.CreateTime = opaGatekeeperConstraintTemplaters[i].CreateTime.String
		opaGatekeeperConstraintTemplater.File = opaGatekeeperConstraintTemplaters[i].File.String
		opaGatekeeperConstraintTemplater.UpdateTime = opaGatekeeperConstraintTemplaters[i].UpdateTime.String
		opaGatekeeperConstraintTemplater.Describtion = opaGatekeeperConstraintTemplaters[i].Describtion.String

		opaGatekeeperConstraintTemplaters2 = append(opaGatekeeperConstraintTemplaters2, opaGatekeeperConstraintTemplater)
	}
	if err != nil {
		utils.Response_Error(context)
		logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
	}

	result := ResponseStruct.ResponseOPAGCTStruct{
		Code:  0,
		Msg:   "",
		Count: count,
		Data:  opaGatekeeperConstraintTemplaters2,
	}
	context.JSON(http.StatusOK, result)

}

func OpaConstraintTemplate(ctx *gin.Context) {
	db := utils.InitDB()
	id := ctx.Query("id")
	operator := ctx.Query("operator")
	if operator == "" && id == "" {
		var opaGatekeeperConstraintTemplaters []Gatekeeper.OpaGatekeeperConstraintTemplater
		sql := "select * from opa_gatekeeper_constrainttemplate"
		err := db.Select(&opaGatekeeperConstraintTemplaters, sql)
		if err != nil {
			response.Fail(ctx, "Select OPAConstraintTemplate Fail", gin.H{"result": "-1"})
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		} else {
			//ctx.HTML(http.StatusOK, "opaconstrainttemplate.html", gin.H{"result": opaGatekeeperConstraintTemplaters})
			type user struct {
				Id         int    `json:"id"`
				Username   string `json:"username"`
				Sex        string `json:"sex"`
				City       string `json:"city"`
				Sign       string `json:"sign"`
				Experience int    `json:"experience"`
				Wealth     int    `json:"wealth"`
				Classify   string `json:"classify"`
				Score      int    `json:"score"`
			}
			type resu struct {
				Code  int    `json:"code"`
				Msg   string `json:"msg"`
				Count int    `json:"count"`
				Data  []user `json:"data"`
			}
			data := []user{{Id: 1000, Username: "USER", Sex: "男", City: "城市", Sign: "sign-1", Experience: 884, Wealth: 5444, Classify: "hh", Score: 27}}
			resu123 := resu{
				Code:  0,
				Msg:   "",
				Count: 1000,
				Data:  data,
			}
			ctx.JSON(http.StatusOK, resu123)
		}
	} else if operator == "edit" && id != "" {

		sql := "select * from opa_gatekeeper_constrainttemplate where id=?"
		var opaGatekeeperConstraintTemplater Gatekeeper.OpaGatekeeperConstraintTemplater
		if err := db.Get(&opaGatekeeperConstraintTemplater, sql, id); err != nil {
			response.Fail(ctx, "Select OPAConstraintTemplate Fail", gin.H{"result": "-1"})
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		} else {
			ctx.HTML(http.StatusOK, "opaconstrainttemplateedit.html", gin.H{"id": opaGatekeeperConstraintTemplater.Id, "name": opaGatekeeperConstraintTemplater.Name,
				"type": opaGatekeeperConstraintTemplater.Type, "packagetype": opaGatekeeperConstraintTemplater.PackageType, "createtime": opaGatekeeperConstraintTemplater.CreateTime, "file": opaGatekeeperConstraintTemplater.File})
		}

	} else if operator == "delete" && id != "" {
		sql := "delete from opa_gatekeeper_constrainttemplate where id=?"
		res, err := db.Exec(sql, id)
		if err != nil {
			response.Fail(ctx, "Delete OPAConstraintTemplate Fail", gin.H{"result": "-1"})
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}

		row, err := res.RowsAffected()
		if err != nil {
			response.Fail(ctx, "Delete OPAConstraintTemplate Fail", gin.H{"result": "-1"})
			logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
		}
		if row > 0 {
			var opaGatekeeperConstraintTemplaters []Gatekeeper.OpaGatekeeperConstraintTemplater
			sql = "select * from opa_gatekeeper_constrainttemplate"
			err = db.Select(&opaGatekeeperConstraintTemplaters, sql)
			if err != nil {
				response.Fail(ctx, "Delete OPAConstraintTemplate Fail", gin.H{"result": "-1"})
				logrus.Error(time.Now().Format("2006-01-02 15:04:05") + " " + err.Error())
			} else {
				ctx.HTML(http.StatusOK, "opaconstrainttemplate.html", gin.H{"result": opaGatekeeperConstraintTemplaters})
			}
		} else {
			response.Fail(ctx, "Delete OPAConstraintTemplate Fail", gin.H{"result": "-1"})
		}

	} else {
		response.Fail(ctx, "Select OPAConstraintTemplate Fail", gin.H{"result": "-1"})
	}
}
