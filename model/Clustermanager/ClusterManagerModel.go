package Clustermanager

import "database/sql"

type ClusterManagerModel struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	CreateTime  sql.NullString `db:"createtime"`
	File        sql.NullString `db:"file"`
	UpdateTime  sql.NullString `db:"updatetime"`
	Describtion sql.NullString `db:"describtion"`
}

type ClusterManagerModelView struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	CreateTime  string `db:"createtime"`
	File        string `db:"file"`
	UpdateTime  string `db:"updatetime"`
	Describtion string `db:"describtion"`
}
