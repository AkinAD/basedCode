package main

import (
	"fmt"
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
	itemSrv item.ItemService
	// authSrv auth.AuthService

	port       string
	connString string
	userPoolID string
	awsRegion  string
	awsID      string
	awsSecret  string
)

func main() {
	router := gin.Default()

	router.Use(corsMiddleware)

	userSrv = user.NewService(awsRegion, awsID, awsSecret)
	itemSrv = item.NewService(connString)
	// authSrv = auth.NewService()

	// heartbeat
	router.GET("/", homeHandler)
	router.GET("/", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), homeHandler)

	//users
	router.GET("/account/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), getAccount)
	router.PUT("/account/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), updateAccount)

	//employees
	router.GET("/employee", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), getEmployees)
	router.POST("/employee", auth.AuthMiddleware(awsRegion, userPoolID, []string{"manager", "admin"}), createEmployee)
	// router.PUT("/employee", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{"employee", "manager", "admin"}), updateEmployee)
	// router.DELETE("/employee", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{manager", "admin"}), deleteEmployee)

	//managers
	router.GET("/manager", auth.AuthMiddleware(awsRegion, userPoolID, []string{"manager", "admin"}), getManagers)
	router.POST("/manager/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), promoteToManager)
	// router.DELETE("manager/:id", deleteManager)

	//admin
	router.GET("/admin", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), getAdmins)
	router.POST("/admin/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), promoteToAdmin)
	// router.DELTE("admin/:id", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{"admin"}) ,deleteFromAdmin)

	//items
	router.GET("/item", getItems)
	router.GET("/item/:id", getItem)
	router.POST("/item", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), addItem)
	router.PUT("/item/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), updateItem)
	router.DELETE("/item/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), deleteItem)

	if port == "443" {
		router.RunTLS(":"+port, "add cert here", "add key here")
	} else {
		router.Run(":" + port)
	}
}

func init() {
	connString = initPostgres()
	port = defaulter("PORT", "8081")
	awsRegion = defaulter("AWS_REGION", "us-east-2")
	awsID = defaulter("AWS_ID", "")
	awsSecret = defaulter("AWS_SECRET", "")
	userPoolID = defaulter("COGNITO_USER_POOL_ID", "")
}

func initPostgres() string {
	PGHost := defaulter("PG_HOST", "localhost")
	PGPort := defaulter("PG_PORT", "5432")
	PGUser := defaulter("PG_USER", "postgres")
	PGPass := defaulter("PG_PASS", "postgres")
	PGName := defaulter("PG_NAME", "smartshopper")
	PGSSL := defaulter("PG_SSLMODE", "disable")

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		PGUser, PGPass, PGHost, PGPort, PGName, PGSSL)
}

func defaulter(envName, defaultValue string) string {
	input := os.Getenv(envName)
	if len(input) == 0 {
		input = defaultValue
	}

	return input
}

var corsMiddleware = cors.New(cors.Config{
	// AllowOrigins:     []string{"https://wheypal.com", "http://localhost:8080"},
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	// AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
})

func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}

func getAccount(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func updateAccount(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func getEmployees(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func createEmployee(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func getManagers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func promoteToManager(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func getAdmins(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func promoteToAdmin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func getItems(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func getItem(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func addItem(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func updateItem(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func deleteItem(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
