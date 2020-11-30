package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	auth "github.com/AkinAD/basedCode/auth"
	item "github.com/AkinAD/basedCode/item"
	user "github.com/AkinAD/basedCode/user"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws"
)

var (
	userSrv user.UserService
	itemSrv item.ItemService
	//storeSrv db.DbService
	// authSrv auth.AuthService

	port               string
	connString         string
	userPoolID         string
	awsRegion          string
	awsID              string
	awsSecret          string
	cognitoAppClientID string
)

func main() {
	router := gin.Default()

	router.Use(corsMiddleware)

	userSrv = user.NewService(awsRegion, awsID, awsSecret, connString)
	itemSrv = item.NewService(connString)
	// authSrv = auth.NewService()

	// heartbeat
	router.GET("/", homeHandler)
	router.GET("/heartbeat", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), heartbeat)

	//login
	router.POST("/login", login)

	//all account types
	router.GET("/account/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), getAccount)
	router.PUT("/account/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"user", "employee", "manager", "admin"}), updateAccount)

	//users
	router.GET("/user", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), getGroupUser)

	//employees
	router.GET("/employee", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), getGroupEmployee)
	router.POST("/employee", auth.AuthMiddleware(awsRegion, userPoolID, []string{"manager", "admin"}), createEmployee)
	// router.PUT("/employee", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{"employee", "manager", "admin"}), updateEmployee)
	// router.DELETE("/employee", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{manager", "admin"}), deleteEmployee)

	//managers
	router.GET("/manager", auth.AuthMiddleware(awsRegion, userPoolID, []string{"manager", "admin"}), getGroupManager)
	router.POST("/manager", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), promoteToManager)
	// router.DELETE("manager/:id", deleteManager)

	//admin
	router.GET("/admin", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), getGroupAdmin)
	router.POST("/admin)", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), promoteToAdmin)
	// router.DELTE("admin/:id", auth.AuthMiddleware(cognitoRegion, userPoolID, []string{"admin"}) ,deleteFromAdmin)

	//items
	router.GET("/item", getItems) //?storeID= to get the items/stock for a specific store
	router.GET("/item/:id", getItem)
	router.POST("/item", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), addItem)
	router.PUT("/item/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), updateItem)
	router.DELETE("/item/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager", "admin"}), deleteItem)

	//store
	router.GET("/store", getStores)
	router.POST("/store", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), createStore)
	router.PUT("/store", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), updateStore)
	router.DELETE("/store", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), deleteStore)

	//stock
	router.POST("/stock/mystore/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager"}), addItemToStock)
	router.PUT("/stock/mystore/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager"}), editItemInStock)
	router.DELETE("/stock/mystore/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"employee", "manager"}), deleteItemInStock)
	router.POST("/stock/admin/:store/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), adimnAddItemToStock)
	router.PUT("/stock/admin/:store/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), adminEditItemInStock)
	router.DELETE("/stock/admin/:store/:id", auth.AuthMiddleware(awsRegion, userPoolID, []string{"admin"}), adminDeleteItemInStock)

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
	cognitoAppClientID = defaulter("COGNITO_APP_CLIENT_ID", "")
}

func initPostgres() string {
	// PGHost := defaulter("PG_HOST", "localhost")
	PGHost := defaulter("PG_HOST", "localhost")
	PGPort := defaulter("PG_PORT", "5432")
	PGUser := defaulter("PG_USER", "postgres")
	PGPass := defaulter("PG_PASS", "")
	PGName := defaulter("PG_NAME", "postgres")
	// PGSSL := defaulter("PG_SSLMODE", "disable")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		PGUser, PGPass, PGName, PGHost, PGPort)
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
		gin.H{"message": "whats poppin"},
	)
}

func heartbeat(c *gin.Context) {
	groups, _ := c.Get("groups")

	c.JSON(
		200,
		gin.H{"groups": groups},
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
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(username),
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

func getGroupUser(c *gin.Context) {
	getGroup(c, "user")
}

func getGroupEmployee(c *gin.Context) {
	getGroup(c, "employee")
}

func getGroupManager(c *gin.Context) {
	getGroup(c, "manager")
}

func getGroupAdmin(c *gin.Context) {
	getGroup(c, "admin")
}

func getGroup(c *gin.Context, group string) {
	// employee UserPoolId = 3
	// var groupName string
	// err := c.ShouldBind(&groupName)
	// if err != nil {
	// 	c.JSON(500, err)
	// }

	input := &cognito.ListUsersInGroupInput{
		GroupName:  aws.String(group),
		NextToken:  aws.String("1"),
		UserPoolId: aws.String(userPoolID),
	}

	resp, err := userSrv.ListUsersInGroup(input)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func createEmployee(c *gin.Context) {
	//binded variables change depending on what is being sent from front-end
	type CreateEmployeeInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var input CreateEmployeeInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(500, err)
	}

	payload := &cognito.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{aws.String("email")},
		// ForceAliasCreation:     aws.Bool(true),
		UserAttributes: []*cognito.AttributeType{&cognito.AttributeType{Name: aws.String(cognito.UsernameAttributeTypeEmail), Value: aws.String(input.Email)}},
		UserPoolId:     aws.String(userPoolID),
		Username:       aws.String(input.Username),
	}

	resp, err := userSrv.CreateEmployee(payload)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func promoteToManager(c *gin.Context) {
	// need user to be promoted, changing their UserPoolId to 2 (manager)
	promoteTo(c, "manager")

}

func promoteToAdmin(c *gin.Context) {
	// need user to be promoted, changing their UserPoolId to 1 (admin)
	promoteTo(c, "admin")
}

func promoteTo(c *gin.Context, group string) {
	var username string
	err := c.ShouldBind(&username)
	if err != nil {
		c.JSON(500, err)
	}
	input := &cognito.AdminAddUserToGroupInput{
		GroupName:  aws.String(group),
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(username),
	}

	resp, err := userSrv.AddUserToGroup(input)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func getItems(c *gin.Context) {
	c.JSON(200, gin.H{"message": "yo"})
}

func getItem(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(500, err)
	}

	item, err := itemSrv.GetItem(id)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{"message": "yo yo", "item ID": c.Param("id"), "item": item})
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

func getStores(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func createStore(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func updateStore(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func deleteStore(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func addItemToStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func editItemInStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func deleteItemInStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func adimnAddItemToStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func adminEditItemInStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func adminDeleteItemInStock(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
