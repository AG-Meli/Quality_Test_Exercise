/*
	Author: AG-Meli - Andrés Ghione
*/

package main

import (
	"github.com/AG-Meli/Quality_Test_Exercise/cmd/server/routes"
	"github.com/AG-Meli/Quality_Test_Exercise/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	db := store.New(store.FileType, "internal/transactions/transactions.json")
	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
