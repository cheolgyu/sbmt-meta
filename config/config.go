package config

import (
	"github.com/cheolgyu/sbm-base/logging"
	"github.com/cheolgyu/sbm-struct/model"
	"github.com/jmoiron/sqlx"
)

const CONFIG_UPPER_CODE_PRICE_TYPE = "price_type"

/*
칼럼 upper_code 으로 meta.config 조회
*/
func GetConfigListByUpperCode(conn *sqlx.DB, upper_code string) ([]model.Config, error) {
	var res []model.Config
	rows, err := conn.Query("select id, upper_code, upper_name, code, name from meta.config where upper_code =$1   order by id  ", upper_code)
	if err != nil {
		logging.Log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Config{}
		if err := rows.Scan(&c.Id, &c.Upper_code, &c.Upper_name, &c.Code, &c.Name); err != nil {
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
