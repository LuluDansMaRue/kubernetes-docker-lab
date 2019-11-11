package database

import (
	"github.com/LuluDansMaRue/kubernetes-docker-lab/bobba"
)

func (db *BobbaCon) OrderBobba(order bobba.Order) (int64, error) {
	stmt, err := db.connection.Prepare("INSERT INTO order SET bobba_id = ?, client_name = ?")

	if err != nil {
		return 0, err
	}

	row, errInsert := stmt.Exec(
		order.Bobba_id,
		order.Client,
	)

	if errInsert != nil {
		return 0, errInsert
	}

	insertedId, _ := row.LastInsertId()

	return insertedId, nil
}
