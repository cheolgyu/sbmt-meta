package info

import (
	"github.com/cheolgyu/base/logging"
	"github.com/jmoiron/sqlx"
)

const NAME_UPDATE_GRAPH_BOUND = "update_graph_bound"
const NAME_UPDATE_STAT_PRICE = "update_stat_price"

func UpdateNow(conn *sqlx.DB, name string) error {
	query := `INSERT INTO public.info( name, updated) VALUES ('`
	query += name
	query += `', now()) ON CONFLICT ("name") DO UPDATE SET  updated= now()  `

	_, err := conn.Exec(query)
	if err != nil {
		logging.Log.Fatalln(err, query)
		panic(err)
	}
	return err
}
