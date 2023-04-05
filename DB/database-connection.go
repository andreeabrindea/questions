package DB

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func ConnectToDB(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return conn, nil

}
