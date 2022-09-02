package services

import (
	"github.com/jmoiron/sqlx"
)
import _ "github.com/go-sql-driver/mysql"

func GetK8sconfigByClusterName(db *sqlx.DB, Clustername string) (string, error) {
	var k8sconfig []string
	sql := "select file from cluster_manager where name = ?"
	err := db.Select(&k8sconfig, sql, Clustername)
	if err != nil {
		return "", err
	}
	return k8sconfig[0], nil
}

func GetConstraintTemplateNameByConstraintName(db *sqlx.DB, constraintname string) (string, error) {
	var Ctame []string
	sql := "select ctname from opa_gatekeeper_constraint where name = ?"
	err := db.Select(&Ctame, sql, constraintname)
	if err != nil {
		return "", err
	}
	return Ctame[0], nil
}

func GetConstrsintTemplateFileByName(db *sqlx.DB, constraintTemplateName string) (string, error) {
	var yamlFile []string
	sql := "select file from opa_gatekeeper_constrainttemplate where name = ?"
	err := db.Select(&yamlFile, sql, constraintTemplateName)
	if err != nil {
		return "", err
	}
	return yamlFile[0], nil
}

func GetConstrsintFileByName(db *sqlx.DB, constraintName string) (string, error) {
	var yamlFile []string
	sql := "select file from opa_gatekeeper_constraint where name = ?"
	err := db.Select(&yamlFile, sql, constraintName)
	if err != nil {
		return "", err
	}
	return yamlFile[0], nil
}

func UpdateOPAPoliciesStatus(db *sqlx.DB, status string, Id string) (bool, error) {
	sql := "update opa_gatekeeper_policies set status =? where id=?"
	result, err := db.Exec(sql, status, Id)
	if err != nil {
		return false, err
	}
	flag, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if flag > 0 {
		return true, nil
	} else {
		return false, nil
	}

}
