package user

import cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type UserService interface {
	CreateEmployee(*cognito.AdminCreateUserInput) (*cognito.AdminCreateUserOutput, error)
	//PromoteToManager
	//PromoteToAdmin
	//DeleteUser
}

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminCreateUser
func (s *userService) CreateEmployee(input *cognito.AdminCreateUserInput) (*cognito.AdminCreateUserOutput, error) {
	output, err := s.cognito.AdminCreateUser(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminDeleteUser
//Delete User

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminAddUserToGroup
//Add to Employee, Add to Manager add to Admin
