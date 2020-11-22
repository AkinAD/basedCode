package user

import cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type userService struct {
	cognito *cognito.CognitoIdentityProvider
}

func NewService() UserService {
	return &userService{
		cognito: &cognito.CognitoIdentityProvider{},
	}
}
