package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"net/http"
	"questions-BE/DB"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = DB.ConnectToDB("postgres://ftcjjrjb:TFy5c90a0aYRafWE2p_bd69h-MXUYStv@snuffleupagus.db.elephantsql.com/ftcjjrjb")
	if err != nil {
		fmt.Println(err)
		return
	}
}
