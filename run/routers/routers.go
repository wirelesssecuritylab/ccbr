package routers

import "github.com/gin-gonic/gin"
import CTC "ccbr/run/controllers"

func CollectRoute(r *gin.Engine) *gin.Engine {
	//r.GET("/opaconstrainttemplate", CTC.OpaConstraintTemplate)
	r.GET("/", CTC.Index)
	opaconstrainttemplate := r.Group("/opaconstrainttemplate")
	{
		opaconstrainttemplate.GET("/index", CTC.OpaCT_Index)
		opaconstrainttemplate.POST("/update", CTC.OpaCT_Update)
		opaconstrainttemplate.POST("/save", CTC.OpaCT_Save)
		opaconstrainttemplate.POST("/delete", CTC.OpaCT_Delete)
		opaconstrainttemplate.GET("/list", CTC.OpaCT_List)
	}

	clustermanager := r.Group("/clustermanager")
	{
		clustermanager.GET("/index", CTC.ClusterManager_Index)
		clustermanager.POST("/update", CTC.ClusterManager_Update)
		clustermanager.POST("/save", CTC.ClusterManager_Save)
		clustermanager.POST("/delete", CTC.ClusterManager_Delete)
		clustermanager.GET("/list", CTC.ClusterManager_List)
		clustermanager.POST("/test", CTC.ClusterManager_Test)
	}

	//
	opaconstraint := r.Group("/opaconstraint")
	{
		opaconstraint.GET("/index", CTC.OpaC_Index)
		opaconstraint.POST("/update", CTC.OpaC_Update)
		opaconstraint.POST("/save", CTC.OpaC_Save)
		opaconstraint.POST("/delete", CTC.OpaC_Delete)
		opaconstraint.GET("/list", CTC.OpaC_List)
		opaconstraint.POST("/list_CT", CTC.OpaC_List_CT)

	}
	opapolicies := r.Group("/opapolicies")
	{
		opapolicies.GET("/index", CTC.OpaPolicies_Index)
		opapolicies.POST("/update", CTC.OpaPolicies_Update)
		opapolicies.POST("/save", CTC.OpaPolicies_Save)
		opapolicies.POST("/delete", CTC.OpaPolicies_Delete)
		opapolicies.POST("/deploy", CTC.OpaPolicies_Deploy)
		opapolicies.POST("/stop", CTC.OpaPolicies_Stop)
		opapolicies.GET("/list", CTC.OpaPolicies_List)
		opapolicies.GET("/listconstraint", CTC.OpaPolicies_ListConstraint)
		opapolicies.GET("/listclustermanager", CTC.OpaPolicies_ListClusterManager)

	}
	cmconstrainttemplate := r.Group("/cmconstrainttemplate")
	{
		cmconstrainttemplate.GET("/index", CTC.CMCT_Index)
		cmconstrainttemplate.GET("/list", CTC.CMCT_List)
		cmconstrainttemplate.GET("/detail", CTC.CMCT_Detail)
	}

	cmconstraint := r.Group("/cmconstraint")
	{
		cmconstraint.GET("/index", CTC.CMC_Index)
		cmconstraint.GET("/list", CTC.CMC_List)
		cmconstraint.GET("/detail", CTC.CMC_Detail)
	}

	download := r.Group("download")
	{
		download.GET("/index", CTC.DOMNLOAD_Index)
	}
	cisbenchmark := r.Group("cisbenchmark")
	{
		cisbenchmark.GET("/index", CTC.CisBenchMarkControoler_Index)
		cisbenchmark.GET("/list", CTC.CisBenchMarkControoler_List)
	}

	return r
}
