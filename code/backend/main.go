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
	//storeSrv db.DbService
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

	userSrv = user.NewService(awsRegion, awsID, awsSecret, connString)
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

func heartbeatUser(c *gin.Context) {
	heartbeat(c, "user")
}
func heartbeatEmployee(c *gin.Context) {
	heartbeat(c, "employee")
}
func heartbeatManager(c *gin.Context) {
	heartbeat(c, "manager")
}
func heartbeatAdmin(c *gin.Context) {
	heartbeat(c, "admin")
}

func heartbeat(c *gin.Context, group string) {
	c.JSON(
		200,
		gin.H{"group": group},
	)
}

func login(c *gin.Context) {
	type LoginRequest struct {
		Username string
		Password string
	}
	var login LoginRequest
	err := c.ShouldBind(&login)

	if err != nil {
		c.JSON(401, err)
	}

	authParams := make(map[string]*string)
	authParams["USERNAME"] = aws.String(login.Username)
	authParams["PASSWORD"] = aws.String(login.Password)

	input := &cognito.InitiateAuthInput{
		AuthFlow:       aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: authParams,
		ClientId:       aws.String(cognitoAppClientID),
	}

	res, err := userSrv.Login(input)
	if err != nil {
		c.JSON(401, err)
	}

	c.JSON(200, res)
}

func getAccount(c *gin.Context) {
	var username string

	err := c.ShouldBind(&username)
	if err != nil {
		c.JSON(500, err)
	}

	input := &cognito.AdminGetUserInput{
		Username:   aws.String(username),
		UserPoolId: aws.String(userPoolID),
	}

	resp, err := userSrv.GetUser(input)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func updateAccount(c *gin.Context) {
	//get the user's account id and their preferred store id
	type updateReqest struct {
		username       string
		preferredStore int
	}
	var update updateReqest
	err := c.ShouldBind(&update)
	if err != nil {
		c.JSON(401, err)
	}
	// grab the username and connect to the userDB to find them
	// then update the db with the preferred location
	//err = itemSrv.db.Table("account").Where("username = ?", update.username).Update("storeID", update.preferredStore)
	err = userSrv.updatePreferredStore(update.username, update.preferredStore)
	if err != nil {
		c.JSON(401, err)
	}

}

func getEmployees(c *gin.Context) {
	// employee UserPoolId = 3
	var groupName string
	bindedVar := c.ShouldBind(&groupName)
	resp, err := userSrv.ListUsersInGroup(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func createEmployee(c *gin.Context) {
	//binded variables change depending on what is being sent from front-end
	var name string
	bindedVar := c.ShouldBind(&name)
	resp, err := userSrv.CreateEmployee(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func getManagers(c *gin.Context) {
	//binded variables change depending on what is being sent from front-end
	// managers UserPoolId = 2
	var groupName string
	bindedVar := c.ShouldBind(&groupName)
	resp, err := userSrv.ListUsersInGroup(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func promoteToManager(c *gin.Context) {
	// need user to be promoted, changing their UserPoolId to 2 (manager)
	//
	var accountID string
	bindedVar := c.ShouldBind(&accountID)
	resp, err := userSrv.AddUserToGroup(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)

}

func getAdmins(c *gin.Context) {
	// admin - UserPoolId = 1
	var groupName string
	bindedVar := c.ShouldBind(&groupName)
	resp, err := userSrv.ListUsersInGroup(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func promoteToAdmin(c *gin.Context) {
	// need user to be promoted, changing their UserPoolId to 1 (admin)
	var accountID string
	bindedVar := c.ShouldBind(&accountID)
	resp, err := userSrv.AddUserToGroup(bindedVar)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
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
