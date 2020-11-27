package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//https://github.com/yoskeoka/gognito/blob/master/main.go
func AuthMiddleware(region, userPoolID string, authedGroups []string) gin.HandlerFunc {

	// 1. Download and store the JSON Web Key (JWK) for your user pool.
	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", region, userPoolID)
	// fmt.Println(jwkURL)
	jwk := getJWK(jwkURL)

	for key, val := range getJWK("https://www.googleapis.com/oauth2/v3/certs") {
		jwk[key] = val
	}

	// fmt.Println(jwk)

	return func(c *gin.Context) {
		tokenString, ok := getBearer(c.Request.Header["Authorization"])

		if !ok {
			c.AbortWithStatusJSON(401, res{Text: "Authorization Bearer Header is missing"})
			return
		}

		token, username, groups, err := validateToken(tokenString, region, userPoolID, jwk, authedGroups)
		if !token.Valid {
			fmt.Printf("token is not valid\n")
			c.AbortWithStatusJSON(401, res{Text: fmt.Sprintf("token is not valid")})
		} else if err != nil {
			fmt.Printf("token is not valid\n%v\n", err)
			c.AbortWithStatusJSON(401, res{Text: fmt.Sprintf("token is not valid%v", err)})
		} else {
			c.Set("token", token)
			c.Set("username", username)
			c.Set("groups", groups)
			c.Next()
		}
	}
}

func validateAWSJwtClaims(claims jwt.MapClaims, region, userPoolID string, authedGroups []string) (string, []string, error) {
	var err error
	// 3. Check the iss claim. It should match your user pool.
	issShoudBe := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", region, userPoolID)
	err = validateClaimItem("iss", []string{issShoudBe}, claims)
	if err != nil {
		return "", nil, err
	}

	// 4. Check the token_use claim.
	validateTokenUse := func() error {
		if tokenUse, ok := claims["token_use"]; ok {
			if tokenUseStr, ok := tokenUse.(string); ok {
				if tokenUseStr == "id" || tokenUseStr == "access" {
					return nil
				}
			}
		}
		return errors.New("token_use should be id or access")
	}

	err = validateTokenUse()
	if err != nil {
		return "", nil, err
	}

	getUsername := func() (string, error) {
		if username, ok := claims["username"]; ok {
			if usernameStr, ok := username.(string); ok {
				return usernameStr, nil
			}
		}
		return "", errors.New("could not retrieve username")
	}

	username, err := getUsername()
	if err != nil {
		return "", nil, err
	}

	validateValidGroup := func() ([]string, error) {
		if groupsFromClaims, ok := claims["cognito:groups"]; ok {
			// log.Printf("Groups1 | %t | %v\n", groupsFromClaims, groupsFromClaims)
			groups := groupsToStringArray(groupsFromClaims)
			// log.Printf("Groups2 | %v\n", groupsFromClaimsStr)
			for _, group := range authedGroups {
				for _, grp := range groups {
					if strings.Index(grp, group) > -1 {
						return groups, nil
					}

				}
			}
		}
		return nil, errors.New("user unauthorized to perform this action")
	}

	groups, err := validateValidGroup()
	if err != nil {
		return username, nil, err
	}

	// 7. Check the exp claim and make sure the token is not expired.
	err = validateExpired(claims)
	if err != nil {
		return username, groups, err
	}

	return username, groups, nil
}

func groupsToStringArray(t ...interface{}) []string {
	s := make([]string, len(t))
	for i, v := range t {
		s[i] = fmt.Sprint(v)
	}
	return s
}

// // validateAWSJwtClaims validates Google JWT
// func validateGoogleJwtClaims(claims jwt.MapClaims) error {
// 	var err error
// 	issShoudBe := []string{"accounts.google.com", "https://accounts.google.com"}
// 	err = validateClaimItem("iss", issShoudBe, claims)
// 	if err != nil {
// 		return err
// 	}
// 	aud := []string{os.Getenv("GOOGLE_CLIENT_ID")}
// 	err = validateClaimItem("aud", aud, claims)
// 	if err != nil {
// 		return err
// 	}
// 	err = validateExpired(claims)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func validateClaimItem(key string, keyShouldBe []string, claims jwt.MapClaims) error {
	if val, ok := claims[key]; ok {
		if valStr, ok := val.(string); ok {
			for _, shouldbe := range keyShouldBe {
				if valStr == shouldbe {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("%v does not match any of valid values: %v", key, keyShouldBe)
}

func validateExpired(claims jwt.MapClaims) error {
	if tokenExp, ok := claims["exp"]; ok {
		if exp, ok := tokenExp.(float64); ok {
			now := time.Now().Unix()
			// fmt.Printf("current unixtime : %v\n", now)
			// fmt.Printf("expire unixtime  : %v\n", int64(exp))
			if int64(exp) > now {
				return nil
			}
		}
		return errors.New("cannot parse token exp")
	}
	return errors.New("token is expired")
}

func validateToken(tokenStr, region, userPoolID string, jwk map[string]JWKKey, authedGroups []string) (*jwt.Token, string, []string, error) {

	// 2. Decode the token string into JWT format.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// cognito user pool, googleは : RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// 5. Get the kid from the JWT token header and retrieve the corresponding JSON Web Key that was stored
		if kid, ok := token.Header["kid"]; ok {
			if kidStr, ok := kid.(string); ok {
				key := jwk[kidStr]
				// 6. Verify the signature of the decoded JWT token.
				rsaPublicKey := convertKey(key.E, key.N)
				return rsaPublicKey, nil
			}
		}

		// rsa public key取得できず
		return "", nil
	})

	if err != nil {
		return token, "", nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	// log.Printf("CLAIMS | %v\n", claims)
	iss, ok := claims["iss"]
	if !ok {
		return token, "", nil, fmt.Errorf("token does not contain issuer")
	}
	issStr := iss.(string)
	if strings.Contains(issStr, "cognito-idp") {
		// 3. 4. 7.のチェックをまとめて
		username, groups, err := validateAWSJwtClaims(claims, region, userPoolID, authedGroups)
		if err != nil {
			return token, username, groups, err
		}

		if token.Valid {
			return token, username, groups, nil
		}
	}

	return token, "", nil, err
}

func getBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) > 1 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJWK(jwkURL string) map[string]JWKKey {

	jwk := &JWK{}

	getJSON(jwkURL, jwk)

	jwkMap := make(map[string]JWKKey, 0)
	for _, jwk := range jwk.Keys {
		jwkMap[jwk.Kid] = jwk
	}
	return jwkMap
}

// https://gist.github.com/MathieuMailhos/361f24316d2de29e8d41e808e0071b13
func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	// fmt.Println(decodedN)
	// fmt.Println(decodedE)
	// fmt.Printf("%#v\n", *pubKey)
	return pubKey
}
