package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

// 法2：
func main() {
	maxAge := 60 * 60 * 24
	// 这样自定义claims
	// 原本的做法是，先声明一个结构体，然后定义claims
	// 因为声明的结构体包含StandardClaims，所以相当于实现了Claims接口，所以可以传参
	// 现在是直接用包里自带的MapClaims类型(而MapClaims已经实现了Cliams接口)
	claims := jwt.MapClaims{
		"id":11,
		"name":"mary",
		"exp":time.Now().Add(time.Duration(maxAge)*time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n",tokenString)
	ret, err := ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user: %v\n",ret)
}

// 法1：
//func main() {
//	// create claims
//	maxAge := 60 * 60 * 24
//	customClaims := &CustomClaims{
//		UserId: 11,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
//			Issuer:    "marry",
//		},
//	}
//
//	// new a token
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
//	fmt.Printf("1:%s\n", token)
//	// sign the token, tokenString = xxx.yyy.zzz
//	tokenString, err := token.SignedString([]byte(SECRETKEY))
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("token:", tokenString)
//
//	// parse the token
//	ret, err := ParseToken(tokenString)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("userInfo: %v", ret)
//}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	// 接口类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
