module github.com/AkinAD/basedCode

go 1.15

replace github.com/AkinAD/basedCode/user => ./user

replace github.com/AkinAD/basedCode/auth => ./auth

replace github.com/AkinAD/basedCode/shop => ./shop

require (
	github.com/AkinAD/basedCode/auth v1.0.0
	github.com/AkinAD/basedCode/shop v1.0.0
	github.com/AkinAD/basedCode/user v1.0.0
	github.com/aws/aws-sdk-go v1.35.35
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jackc/pgx/v4 v4.9.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect

)
