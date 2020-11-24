package user

import cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type UserService interface {
	CreateEmployee(*cognito.AdminCreateUserInput) (*cognito.AdminCreateUserOutput, error)
	DeleteUser(*cognito.AdminDeleteUserInput) (*cognito.AdminDeleteUserOutput, error)
	AddUserToGroup(input *cognito.AdminAddUserToGroupInput) (*cognito.AdminAddUserToGroupOutput, error)
	RemoveUserFromGroup(input *cognito.AdminRemoveUserFromGroupInput) (*cognito.AdminRemoveUserFromGroupOutput, error)
	GetUser(input *cognito.AdminGetUserInput) (*cognito.AdminGetUserOutput, error)
}

//cognito = CognitoIdentityProvider

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminCreateUser
func (s *userService) CreateEmployee(input *cognito.AdminCreateUserInput) (*cognito.AdminCreateUserOutput, error) {
	output, err := s.cognito.AdminCreateUser(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminDeleteUser
func (s *userService) DeleteUser(input *cognito.AdminDeleteUserInput) (*cognito.AdminDeleteUserOutput, error) {
	output, err := s.cognito.AdminDeleteUser(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminAddUserToGroup
//Add to Employee, Add to Manager add to Admin
/*	Request Syntax
	{
	"GroupName": "string",
	"Username": "string",
	"UserPoolId": "string"
	}
*/
func (s *userService) AddUserToGroup(input *cognito.AdminAddUserToGroupInput) (*cognito.AdminAddUserToGroupOutput, error) {
	output, err := s.cognito.AdminAddUserToGroup(input)
	if err != nil {
		return nil, err
	}
	return output, nil

}

// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminRemoveUserFromGroup
func (s *userService) RemoveUserFromGroup(input *cognito.AdminRemoveUserFromGroupInput) (*cognito.AdminRemoveUserFromGroupOutput, error) {
	output, err := s.cognito.AdminRemoveUserFromGroup(input)
	if err != nil {
		return nil, err
	}
	return output, nil

}

// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminGetUser
func (s *userService) GetUser(input *cognito.AdminGetUserInput) (*cognito.AdminGetUserOutput, error) {
	output, err := s.cognito.AdminGetUser(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
