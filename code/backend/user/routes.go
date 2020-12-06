package user

import (
	"fmt"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserService interface {
	CreateEmployee(*cognito.AdminCreateUserInput) (*cognito.AdminCreateUserOutput, error)
	DeleteUser(*cognito.AdminDeleteUserInput) (*cognito.AdminDeleteUserOutput, error)
	AddUserToGroup(input *cognito.AdminAddUserToGroupInput) (*cognito.AdminAddUserToGroupOutput, error)
	RemoveUserFromGroup(input *cognito.AdminRemoveUserFromGroupInput) (*cognito.AdminRemoveUserFromGroupOutput, error)
	GetUser(input *cognito.AdminGetUserInput) (*cognito.AdminGetUserOutput, error)
	ListUsersInGroup(input *cognito.ListUsersInGroupInput) ([]*User, error)
	Login(*cognito.InitiateAuthInput) (*cognito.InitiateAuthOutput, error)
	// UpdatePreferredStore(username string, preferredStore int) error
	CreateProfile(Username string, StoreID int, FirstName string, LastName string, Email string) error
	GetProfile(username string) (*User, error)
	UpdateProfile(user *User) (*User, error)
	// DeleteUser(username string) (bool, error)
}

//cognito = CognitoIdentityProvider

// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminCreateUser
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

/* https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.AdminGetUser
Request Syntax
{
   "Username": "string",
   "UserPoolId": "string"
}
*/
func (s *userService) GetUser(input *cognito.AdminGetUserInput) (*cognito.AdminGetUserOutput, error) {
	output, err := s.cognito.AdminGetUser(input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

/* https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#CognitoIdentityProvider.ListUsersInGroup
Request Syntax
{
   "GroupName": "string",
   "Limit": number,
   "NextToken": "string",
   "UserPoolId": "string"
}
*/

func (s *userService) ListUsersInGroup(input *cognito.ListUsersInGroupInput) ([]*User, error) {
	output, err := s.cognito.ListUsersInGroup(input)

	// this grabs the username. they're stored in an array
	//fmt.Println(*output.Users[1].Username)
	length := len(output.Users)
	//fmt.Println(length)
	var userArray []*User
	var user *User

	for i := 0; i < length; i++ {
		fmt.Println(*output.Users[i].Username)
		user, err = s.db.getProfile(*output.Users[i].Username)
		userArray = append(userArray, user)
		//fmt.Println(*user)
	}

	//fmt.Println(userArray)
	if err != nil {
		return nil, err
	}

	return userArray, nil
}

func (s *userService) Login(input *cognito.InitiateAuthInput) (*cognito.InitiateAuthOutput, error) {
	output, err := s.cognito.InitiateAuth(input)
	if err != nil {
		return nil, err
	}
	return output, err
}

func (s *userService) CreateProfile(Username string, StoreID int, FirstName string, LastName string, Email string) error {
	err := s.db.createProfile(Username, StoreID, FirstName, LastName, Email)
	if err != nil {
		// log.Printf("%v", err)
		return err
	}

	return nil
}

func (s *userService) GetProfile(username string) (*User, error) {

	user, err := s.db.getProfile(username)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateProfile(user *User) (*User, error) {
	user, err := s.db.updateProfile(user)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return user, nil
}
