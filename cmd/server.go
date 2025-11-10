package cmd

import (
	"GrpcMessangerAccServer/internal"
	"GrpcMessangerAccServer/internal/migration/database"
	"GrpcMessangerAccServer/pkg"
	consts "GrpcMessangerAccServer/pkg/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strconv"
)

type Server struct {
	GinServ  *gin.Engine
	Config   *pkg.Config
	DbStruct *database.LoginTracker
}

func NewAccServer(db *database.LoginTracker) *Server {
	var emailServ, emailServPass string
	// Если нет почты и пароля от почты, то и сервис нам нахуй не нужен
	if emailServPass = os.Getenv("EMAIL_GENERATOR_PASSWORD"); emailServPass == "" {
		panic(consts.EMAIL_GENERATOR_ADDRESS_NIL)
		return nil
	}
	if emailServ = os.Getenv("EMAIL_GENERATOR"); emailServ == "" {
		panic(consts.EMAIL_GENERATOR_NIL)
		return nil
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/auth/code/:email", func(c *gin.Context) {
		var code string
		email := c.Param("email")
		clientIP := c.RemoteIP()

		dbIP, err := db.GetIPbyEmail(email)
		if err != nil {
			fmt.Println(err)
		}
		if clientIP == dbIP {
			code = "CACHING"
		} else {
			code = strconv.Itoa(rand.Intn(90000) + 9999)
			go internal.SendCodeToTheEmail(code, email, emailServ, emailServPass)

			if errStore := db.StoreLogin(email, clientIP); errStore != nil {
				fmt.Println(errStore)
			}
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"code":   code,
		})
	})

	return &Server{
		GinServ:  router,
		Config:   pkg.ReadConfig(),
		DbStruct: db,
	}
}
