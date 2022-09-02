package controllers

import (
	"ccbr/model/Constraint"
	"ccbr/run/response"
	"ccbr/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

func ListConstraint(ctx *gin.Context) {
	dynamicClient, err :=utils.GetDynamicClient("/root/.kube/config")
	if err!=nil {
		log.Fatal(err)
	}
	constraintlist := utils.ListConstraint(dynamicClient)
	//response.Success(ctx, gin.H{"result": result}, "obtain_successful")
	resultConstraints :=make(map[string]Constraint.Constraint)
	for i:=0;i<len(constraintlist.Resources);i++ {
		fmt.Println(constraintlist.Resources[i].Name)
		var constraint Constraint.Constraint
		constraint = utils.GetConstraint(dynamicClient,constraintlist.Resources[i].Name)
		if !reflect.DeepEqual(constraint,Constraint.Constraint{}) {
			resultConstraints[constraintlist.Resources[i].Name] = constraint
		}

	}
	ctx.HTML(http.StatusOK,"constraint.html",gin.H{"resultConstraints":resultConstraints})
}

func GetConstraint(ctx *gin.Context)  {
	constraintName := ctx.Param("constraintName")
	dynamicClient, err :=utils.GetDynamicClient("/root/.kube/config")
	if err!=nil {
		log.Fatal(err)
	}
	result := utils.GetConstraint(dynamicClient,constraintName)
	response.Success(ctx, gin.H{"result": result}, "obtain_successful")
}
