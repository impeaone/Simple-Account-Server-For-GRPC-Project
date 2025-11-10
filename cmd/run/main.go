package main

import (
	"GrpcMessangerAccServer/cmd"
	"GrpcMessangerAccServer/internal/migration/database"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	db, err := database.NewLoginTracker()
	if err != nil {
		panic(err)
	}
	server := cmd.NewAccServer(db)
	if server == nil {
		fmt.Println("Server Error, NIL")
		return
	}

	if errServ := server.GinServ.Run(fmt.Sprintf("%s:%s", server.Config.IPAddress, server.Config.Port)); errServ != nil {
		panic(errServ)
	}
}
