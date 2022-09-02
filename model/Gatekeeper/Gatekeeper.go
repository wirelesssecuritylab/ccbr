package Gatekeeper

import "database/sql"

type OpaGatekeeperConstraintTemplater struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Type        sql.NullString `db:"type"`
	PackageType sql.NullString `db:"packagetype"`
	CreateTime  sql.NullString `db:"createtime"`
	File        sql.NullString `db:"file"`
	UpdateTime  sql.NullString `db:"updatetime"`
	Describtion sql.NullString `db:"describtion"`
}

type OpaGatekeeperConstraintTemplater2 struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	PackageType string `db:"packagetype"`
	CreateTime  string `db:"createtime"`
	File        string `db:"file"`
	UpdateTime  string `db:"updatetime"`
	Describtion string `db:"describtion"`
}

type OpaGatekeeperConstraint struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Ctname      sql.NullString `db:"ctname"`
	Type        sql.NullString `db:"type"`
	PackageType sql.NullString `db:"packagetype"`
	CreateTime  sql.NullString `db:"createtime"`
	File        sql.NullString `db:"file"`
	UpdateTime  sql.NullString `db:"updatetime"`
	Describtion sql.NullString `db:"describtion"`
}
type OpaGatekeeperConstraint2 struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Ctname      string `db:"ctname"`
	Type        string `db:"type"`
	PackageType string `db:"packagetype"`
	CreateTime  string `db:"createtime"`
	File        string `db:"file"`
	UpdateTime  string `db:"updatetime"`
	Describtion string `db:"describtion"`
}

type OpaGatekeeperPolicies struct {
	Id             int            `db:"id"`
	Name           sql.NullString `db:"name"`
	Version        sql.NullString `db:"version"`
	Constraintlist sql.NullString `db:"constraintlist"`
	Process        sql.NullString `db:"process"`
	Describtion    sql.NullString `db:"describtion"`
	CreateTime     sql.NullString `db:"createtime"`
	UpdateTime     sql.NullString `db:"updatetime"`
	Type           sql.NullString `db:"type"`
	Status         sql.NullString `db:"status"`
	Clustername    sql.NullString `db:"clustername"`
}
type OpaGatekeeperPolicies2 struct {
	Id             int    `db:"id"`
	Name           string `db:"name"`
	Version        string `db:"version"`
	Constraintlist string `db:"constraintlist"`
	Process        string `db:"process"`
	Describtion    string `db:"describtion"`
	CreateTime     string `db:"createtime"`
	UpdateTime     string `db:"updatetime"`
	Type           string `db:"type"`
	Status         string `db:"status"`
	Clustername    string `db:"clustername"`
}

type ConstraintMapperConstraintTemplate struct {
	ConstraintTemplate string
	Constraint         string
}
