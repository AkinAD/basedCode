package user

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type userService struct {
	cognito *cognito.CognitoIdentityProvider
	db      UserRepo
}

func NewService(awsRegion, awsID, awsSecret string, conn string) UserService {
	mySession, err := awsSession(awsRegion, awsID, awsSecret)
	if err != nil {
		panic(err)
	}

	svc := cognito.New(mySession)

	return &userService{
		cognito: svc,
		db:      NewDatabase(conn),
	}
}

func awsSession(awsRegion, awsID, awsSecret string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsID, awsSecret, ""),
	})

	if err != nil {
		return nil, err
	}

	return session, nil
}
