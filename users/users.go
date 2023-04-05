package users

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"questions-BE/DB"
)

func getUsers() (error, pgconn.CommandTag) {
	conn, err := DB.ConnectToDB("postgres://ftcjjrjb:TFy5c90a0aYRafWE2p_bd69h-MXUYStv@snuffleupagus.db.elephantsql.com/ftcjjrjb")
	if err != nil {
		return err, pgconn.CommandTag{}
	}
	users, _ := conn.Exec(context.Background(), "SELECT * FROM users")
	return nil, users
}
