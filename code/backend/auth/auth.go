package auth

// type authService struct {
// 	Repo AuthRepo
// }

// func NewService() AuthService {
// 	return &authService{
// 		Repo: NewAuthRepo(),
// 	}
// }

type Account struct {
	ID    int
	Name  string
	Roles []Role
}

type res struct {
	Text string `json:"text"`
}

// JWKKey is json data struct for cognito jwk key
type JWKKey struct {
	Alg string
	E   string
	Kid string
	Kty string
	N   string
	Use string
}

type Role string

const (
	user     Role = "user"
	employee Role = "employee"
	manager  Role = "manager"
	admin    Role = "admin"
)

// JWK is json data struct for JSON Web Key
type JWK struct {
	Keys []JWKKey
}
