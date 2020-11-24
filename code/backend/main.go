package main

import (
	"os"
	"time"

	auth "github.com/AkinAD/basedCode/auth"
	item "github.com/AkinAD/basedCode/item"
	user "github.com/AkinAD/basedCode/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	userSrv user.UserService
	// authSrv auth.AuthService
	itemSrv            item.ItemService
	itemRepoConnString = ""
	region             = ""
	userPoolID         = ""
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://wheypal.com", "http://localhost:8080"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		// AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userSrv = user.NewService()
	// authSrv = auth.NewService()
	itemSrv = item.NewService(itemRepoConnString)

	router.GET("/", homeHandler)
	router.GET("/", auth.AuthMiddleware(region, userPoolID, []string{"user", "employee", "manager", "admin"}), homeHandler)

	if port == "443" {
		router.RunTLS(":"+port, "add cert here", "add key here")
	} else {
		router.Run(":" + port)
	}
}
