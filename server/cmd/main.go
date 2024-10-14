package main

import (
	"log"

	"github.com/VaheMuradyan/vv-chat/server/db"
	"github.com/VaheMuradyan/vv-chat/server/internal/user"
	"github.com/VaheMuradyan/vv-chat/server/internal/ws"
	"github.com/VaheMuradyan/vv-chat/server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal(" !!! could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
