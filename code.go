package main

import (
	"github.com/cheolgyu/stock-write-common/logging"
	"github.com/cheolgyu/stock-write-model/model"
	"github.com/jmoiron/sqlx"
)

/*
meta.code 조회
*/
func GetCodeList(conn *sqlx.DB) ([]model.Code, error) {
	var res []model.Code
	rows, err := conn.Query("select id, code, code_type from meta.code   order by id  ")
	if err != nil {
		logging.Log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Code{}
		if err := rows.Scan(&c.Id, &c.Code, &c.Code_type); err != nil {
			logging.Log.Fatal(err)
			panic(err)
		}
		res = append(res, c)
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		logging.Log.Fatal(err)
		panic(err)
	}
	return res, err
}
