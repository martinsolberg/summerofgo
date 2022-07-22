package main

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

/*
- last inn en JWT access token
- hent ut KID fra token
- hent ut JWK nøkkel fra Azure basert på KID
- konvertere JWK nøkkel til rsa.PublicKey
- parse JWT access token
*/

func getJTWAccessToken() string {
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldC8iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC85ZDJhYzAxOC1lODQzLTRlMTQtOWUyYi00ZTBkZGFjNzU0NTAvIiwiaWF0IjoxNjU4MzA4NDg3LCJuYmYiOjE2NTgzMDg0ODcsImV4cCI6MTY1ODMxMzI2OSwiYWNyIjoiMSIsImFpbyI6IkFWUUFxLzhUQUFBQTlTWjh0SGt6UTc3LzluajJEMzNOQTdKUG1jSDRIYjNuZVFadE9PMi9yWlhrY3Z6ZWx0OWFMb2F4RzdEdjdQbGpQWFozU0ZsYlNjR0ZuNkpZajJWL0FOckF1clhLc3haWTB4dVpLRk0rdHdRPSIsImFtciI6WyJwd2QiLCJtZmEiXSwiYXBwaWQiOiIwNGIwNzc5NS04ZGRiLTQ2MWEtYmJlZS0wMmY5ZTFiZjdiNDYiLCJhcHBpZGFjciI6IjAiLCJmYW1pbHlfbmFtZSI6IlNvbGJlcmcgVG1wIiwiZ2l2ZW5fbmFtZSI6Ik1hcnRpbiBCcnVzZXQiLCJncm91cHMiOlsiYzQyZTU4M2UtYjM4Yy00YTk4LWEyMTItNjEyMTJkY2VmMjM5IiwiMmExODc4ZTEtYTY0Yy00NTIwLTgyOTMtYzU3YzdlYjM4Y2U5IiwiM2EyNjllZmItMDU2ZC00OTAwLTkyOGUtZDQ3YjliNjE3OTQ0IiwiNzlmMzQxODktYTNmNS00MDVmLWE2YWUtNmE2MmY3ZjMyZGI4IiwiNTAxYjkyM2EtZTQxZS00OTRhLWI0NDItZmYxZTg5MDNjOWM3IiwiOGViNTM0Y2YtNDJjMy00OGNjLTkwMWEtODI4N2UyNTZlMTVlIiwiZjA1ZTcxMmUtMTQzOS00OWI2LTg5NmEtYzE5ZjVjZDU5MDQwIiwiZTMxN2ExZGMtNWIzMC00ZDYyLWI0NTctYzI3NGY2YmY4OTlhIiwiMzVjOTQ5ODEtZGFhZS00OTdmLThiNjgtMWU2NDMyYjg3ODZkIiwiYmEyNTE2OTctMmMyOS00ZGE3LTk5ZDgtM2MxNjc0MWFjNTFlIiwiNmI2MzBmMTQtYzQ5ZS00NWY2LTgwYjMtODBjNWRiYmMxNmJhIiwiN2JiZjdhNmYtYjI3Ny00OTVmLWI1Y2EtODA0MTZkMjk2NGJhIiwiMTBlY2Q5ZWItODA5Ni00ZjRiLThlYTgtYTk1NmU3OTAxYjk2IiwiMTdjMGNlZGYtN2Q5Yy00NzFlLWJkYTktNTVjYTNjOTU4ODQzIiwiMzE0Njc2NWUtYWY5Mi00ZGYxLWIyYzMtYTMxZjBkYmQwZTI3IiwiODNhYTkzYmUtYTU0ZC00MTdkLTg5MDAtODFiNDExOTAwODNjIiwiNmYwMjVhNmEtZmJjOC00N2I1LTg4N2YtMzQyNGY3NTE3MmI3IiwiZWE3ZGUwMjAtY2Q5YS00NzgzLTk4MDQtMTc0NzcxZjA5YTUxIiwiOWUxNjA4NjQtZjI5ZS00NmJjLTk4ZTktNzEzNzE4ZWRjYTA0IiwiMmExMzhiYzEtY2ZiNy00NzNmLTljMjEtODhiZjM0MWNmMmU2IiwiZjE0MmQ4NDAtNmNmYi00OTNlLTgxNGItMjA1ODNkNDI1NTYyIiwiYjI3NjZmMmQtOGI1OS00ZTZkLWI0NWEtYWM4YTYwZDg5OTZhIiwiNGJmNGIxOGYtYmQyZi00ZDNiLWE2ZTktZWI4Yjc2NjQ5MDA3Il0sImlwYWRkciI6IjE2MC42Ny4xMTUuNDciLCJuYW1lIjoiTWFydGluIEJydXNldCBTb2xiZXJnIFRtcCIsIm9pZCI6IjM3OTA2MzVhLTM4Y2UtNDA4Yy1hODg0LTQwNTYxOGE0MmM2NyIsIm9ucHJlbV9zaWQiOiJTLTEtNS0yMS0zMTI3MjU1My0xMjcwMzEwMDQ0LTE4MDQzNjQ5ODUtMzEyNTQxIiwicHVpZCI6IjEwMDMyMDAxRTI1M0Y2NzYiLCJyaCI6IjAuQVR3QUdNQXFuVVBvRkU2ZUswNE4yc2RVVUVaSWYza0F1dGRQdWtQYXdmajJNQk04QUpNLiIsInNjcCI6InVzZXJfaW1wZXJzb25hdGlvbiIsInN1YiI6Ik5KR1lQOFlCeTZDS3VucXhBOEZURXlLZTE3S1dWTkdBcXY2dXpVQVFFOTQiLCJ0aWQiOiI5ZDJhYzAxOC1lODQzLTRlMTQtOWUyYi00ZTBkZGFjNzU0NTAiLCJ1bmlxdWVfbmFtZSI6InRtcDY0OTcwNkBucmsubm8iLCJ1cG4iOiJ0bXA2NDk3MDZAbnJrLm5vIiwidXRpIjoiLU9NdGc3X18wMFdpS0hTM0lZUWVBQSIsInZlciI6IjEuMCIsIndpZHMiOlsiYjc5ZmJmNGQtM2VmOS00Njg5LTgxNDMtNzZiMTk0ZTg1NTA5Il0sInhtc190Y2R0IjoxMzY2ODkyOTQwfQ.fvHJKExg1ERDY32y53AdaJcXKBQFTWnG590mWT5K7Z3PEJrkUDjineGWa0bmyb7pHD0TpZHEAVB7agJvgVNoII2At1CnpysSk-wlSV-9sZkN2x3Z7UYuYu1WnzGf3YrRzt2xXFnB5szq8zbu3nJY9Q5t4aN0Vh2t0IM2meTLsp0ySuSiCoQBJobetxPIOz05GD84NNhKZRb9Vgj9hQYyRQ8rC2XqKlfd6Nt7OZ3QjMVeM8GvT4adzyspUKLICqDysD3qJRez27mdPhJ0ddBugi_-LPSMssm047uy6lRaPqrisIp2g8TeHqM5yLcxyboV9rz-PuJz2RmsynhHYkbvMA"

	return tokenString
}

func getJWKKey() jwk.Key {

	myJwkString := []byte(`{
		"kty": "RSA",
		"use": "sig",
		"kid": "nOo3ZDrODXEK1jKWhXslHR_KXEg",
		"x5t": "nOo3ZDrODXEK1jKWhXslHR_KXEg",
		"n": "oaLLT9hkcSj2tGfZsjbu7Xz1Krs0qEicXPmEsJKOBQHauZ_kRM1HdEkgOJbUznUspE6xOuOSXjlzErqBxXAu4SCvcvVOCYG2v9G3-uIrLF5dstD0sYHBo1VomtKxzF90Vslrkn6rNQgUGIWgvuQTxm1uRklYFPEcTIRw0LnYknzJ06GC9ljKR617wABVrZNkBuDgQKj37qcyxoaxIGdxEcmVFZXJyrxDgdXh9owRmZn6LIJlGjZ9m59emfuwnBnsIQG7DirJwe9SXrLXnexRQWqyzCdkYaOqkpKrsjuxUj2-MHX31FqsdpJJsOAvYXGOYBKJRjhGrGdONVrZdUdTBQ",
		"e": "AQAB",
		"x5c": [
		  "MIIDBTCCAe2gAwIBAgIQN33ROaIJ6bJBWDCxtmJEbjANBgkqhkiG9w0BAQsFADAtMSswKQYDVQQDEyJhY2NvdW50cy5hY2Nlc3Njb250cm9sLndpbmRvd3MubmV0MB4XDTIwMTIyMTIwNTAxN1oXDTI1MTIyMDIwNTAxN1owLTErMCkGA1UEAxMiYWNjb3VudHMuYWNjZXNzY29udHJvbC53aW5kb3dzLm5ldDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKGiy0/YZHEo9rRn2bI27u189Sq7NKhInFz5hLCSjgUB2rmf5ETNR3RJIDiW1M51LKROsTrjkl45cxK6gcVwLuEgr3L1TgmBtr/Rt/riKyxeXbLQ9LGBwaNVaJrSscxfdFbJa5J+qzUIFBiFoL7kE8ZtbkZJWBTxHEyEcNC52JJ8ydOhgvZYykete8AAVa2TZAbg4ECo9+6nMsaGsSBncRHJlRWVycq8Q4HV4faMEZmZ+iyCZRo2fZufXpn7sJwZ7CEBuw4qycHvUl6y153sUUFqsswnZGGjqpKSq7I7sVI9vjB199RarHaSSbDgL2FxjmASiUY4RqxnTjVa2XVHUwUCAwEAAaMhMB8wHQYDVR0OBBYEFI5mN5ftHloEDVNoIa8sQs7kJAeTMA0GCSqGSIb3DQEBCwUAA4IBAQBnaGnojxNgnV4+TCPZ9br4ox1nRn9tzY8b5pwKTW2McJTe0yEvrHyaItK8KbmeKJOBvASf+QwHkp+F2BAXzRiTl4Z+gNFQULPzsQWpmKlz6fIWhc7ksgpTkMK6AaTbwWYTfmpKnQw/KJm/6rboLDWYyKFpQcStu67RZ+aRvQz68Ev2ga5JsXlcOJ3gP/lE5WC1S0rjfabzdMOGP8qZQhXk4wBOgtFBaisDnbjV5pcIrjRPlhoCxvKgC/290nZ9/DLBH3TbHk8xwHXeBAnAjyAqOZij92uksAv7ZLq4MODcnQshVINXwsYshG1pQqOLwMertNaY5WtrubMRku44Dw7R"
		],
		"issuer": "https://login.microsoftonline.com/9d2ac018-e843-4e14-9e2b-4e0ddac75450/v2.0"
	  }"
	  `)

	myJwkKey, err := jwk.ParseKey(myJwkString)
	if err != nil {
		panic(err)
	}

	return myJwkKey
}

func getJWTKid() int {
	myJWTKid := 0
	return myJWTKid
}

func convJWKtoRSA(myJWKKey jwk.Key) rsa.PublicKey {

	var rawkey rsa.PublicKey

	err := myJWKKey.Raw(&rawkey)
	if err != nil {
		panic(err)
	}

	return rawkey

}

func main() {

	myJWTAccessToken := getJTWAccessToken()

	myJWKKey := getJWKKey()

	myRSAKey := convJWKtoRSA(myJWKKey)

	token, err := jwt.Parse(myJWTAccessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return myRSAKey, nil
	})

	if err != nil {
		panic(err)
	}

	// token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("%T\n", token)

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Println(claims["foo"], claims["nbf"])
	// } else {
	// 	fmt.Println(err)
	// }
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Name: ", claims["name"])
		fmt.Println("Groups: ", claims["groups"])
	} else {
		fmt.Println("Validation error: ", err)
	}

}
