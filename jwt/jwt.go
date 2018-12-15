package jwt

import (
	"time"
    "fmt"

    "github.com/dgrijalva/jwt-go"
)

/*
 * Generate Token with name
 */
func GenerateToken(name string) string {
    // create header
    token := jwt.New(jwt.SigningMethodHS256)
    claims := make(jwt.MapClaims)

    // add payload
    claims["iss"] = "CloudCompute711"
    claims["sub"] = "SWAPI"
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
    claims["iat"] = time.Now().Unix()
    token.Claims = claims

    // create signature
    tokenString, _ := token.SignedString([]byte(name))
    return tokenString
}

/*
 * Generate Token with name
 */
func ValidToken(tokenString string, name string) bool {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(name), nil
    })
    if token.Valid {
        return true
    } else {
        fmt.Println(err)
        return false
    }
}