package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	ue "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateToken Generate Token JWT
func GenerateToken(user ue.User) (string, error) {
	// Create Error Variable
	var err error

	// Creating New JWT
	key := os.Getenv("API_JWT_SECRET")

	// Create New Claims Object
	claims := jwt.MapClaims{}

	// Assign data to Claims
	claims["authorized"] = true
	claims["user_id"] = user.Id
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["expired"] = time.Now().Add(time.Hour * 720).Unix() // Valid for 30 days

	// Create JWT
	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create Token Variable ( String )
	token, err := unsignedJWT.SignedString([]byte(key))

	// Check if there is error when signing JWT
	if err != nil {
		return "", err
	}

	// Return Token
	return token, nil
}

// ExtractToken Extract token
func ExtractToken(context *gin.Context) string {
	// Get Token
	authorizationToken := context.Request.Header.Get("Authorization")

	// Check if token not provided
	if authorizationToken != "" {
		//remove 'Bearer '
		authorizationToken = strings.Replace(authorizationToken, "Bearer ", "", 1)
		return authorizationToken
	}

	// Return empty string when token not provided
	return ""
}

func GetUserCredentialFromToken(context *gin.Context) (ad.UserClaims, error) {
	// Extract Token Data
	tokenString := ExtractToken(context)

	// Get Secret Key from ENV
	key := os.Getenv("API_JWT_SECRET")

	// User Claims Object
	var userClaims ad.UserClaims

	// Parse JWT and validate
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check sign in method token
		if jwt.GetSigningMethod("HS256") != token.Method {
			// When sign in method not same
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return key
		return []byte(key), nil
	})

	// Check if user exist in database & Token Expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Catch User Claims from token
		authorized := claims["authorized"].(bool)
		userId := uint(claims["user_id"].(float64))
		email := fmt.Sprintf("%v", claims["email"])
		roleId := uint(claims["role_id"].(float64))
		roleName := fmt.Sprintf("%v", claims["role_name"])
		expired := int64(math.Round(claims["expired"].(float64)))

		// Assign to User Claims Object
		userClaims.Authorized = authorized
		userClaims.UserId = userId
		userClaims.Email = email
		userClaims.RoleId = roleId
		userClaims.RoleName = roleName
		userClaims.Expired = expired

		// Return Value
		return userClaims, nil
	} else {
		// Return Error
		return userClaims, errors.New("error parsing jwt")
	}

}
