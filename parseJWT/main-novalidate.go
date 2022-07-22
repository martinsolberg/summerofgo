package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var rsaSampleSecret []byte

func main() {
	// sample token string taken from the New example
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldC8iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC85ZDJhYzAxOC1lODQzLTRlMTQtOWUyYi00ZTBkZGFjNzU0NTAvIiwiaWF0IjoxNjU4MzA4NDg3LCJuYmYiOjE2NTgzMDg0ODcsImV4cCI6MTY1ODMxMzI2OSwiYWNyIjoiMSIsImFpbyI6IkFWUUFxLzhUQUFBQTlTWjh0SGt6UTc3LzluajJEMzNOQTdKUG1jSDRIYjNuZVFadE9PMi9yWlhrY3Z6ZWx0OWFMb2F4RzdEdjdQbGpQWFozU0ZsYlNjR0ZuNkpZajJWL0FOckF1clhLc3haWTB4dVpLRk0rdHdRPSIsImFtciI6WyJwd2QiLCJtZmEiXSwiYXBwaWQiOiIwNGIwNzc5NS04ZGRiLTQ2MWEtYmJlZS0wMmY5ZTFiZjdiNDYiLCJhcHBpZGFjciI6IjAiLCJmYW1pbHlfbmFtZSI6IlNvbGJlcmcgVG1wIiwiZ2l2ZW5fbmFtZSI6Ik1hcnRpbiBCcnVzZXQiLCJncm91cHMiOlsiYzQyZTU4M2UtYjM4Yy00YTk4LWEyMTItNjEyMTJkY2VmMjM5IiwiMmExODc4ZTEtYTY0Yy00NTIwLTgyOTMtYzU3YzdlYjM4Y2U5IiwiM2EyNjllZmItMDU2ZC00OTAwLTkyOGUtZDQ3YjliNjE3OTQ0IiwiNzlmMzQxODktYTNmNS00MDVmLWE2YWUtNmE2MmY3ZjMyZGI4IiwiNTAxYjkyM2EtZTQxZS00OTRhLWI0NDItZmYxZTg5MDNjOWM3IiwiOGViNTM0Y2YtNDJjMy00OGNjLTkwMWEtODI4N2UyNTZlMTVlIiwiZjA1ZTcxMmUtMTQzOS00OWI2LTg5NmEtYzE5ZjVjZDU5MDQwIiwiZTMxN2ExZGMtNWIzMC00ZDYyLWI0NTctYzI3NGY2YmY4OTlhIiwiMzVjOTQ5ODEtZGFhZS00OTdmLThiNjgtMWU2NDMyYjg3ODZkIiwiYmEyNTE2OTctMmMyOS00ZGE3LTk5ZDgtM2MxNjc0MWFjNTFlIiwiNmI2MzBmMTQtYzQ5ZS00NWY2LTgwYjMtODBjNWRiYmMxNmJhIiwiN2JiZjdhNmYtYjI3Ny00OTVmLWI1Y2EtODA0MTZkMjk2NGJhIiwiMTBlY2Q5ZWItODA5Ni00ZjRiLThlYTgtYTk1NmU3OTAxYjk2IiwiMTdjMGNlZGYtN2Q5Yy00NzFlLWJkYTktNTVjYTNjOTU4ODQzIiwiMzE0Njc2NWUtYWY5Mi00ZGYxLWIyYzMtYTMxZjBkYmQwZTI3IiwiODNhYTkzYmUtYTU0ZC00MTdkLTg5MDAtODFiNDExOTAwODNjIiwiNmYwMjVhNmEtZmJjOC00N2I1LTg4N2YtMzQyNGY3NTE3MmI3IiwiZWE3ZGUwMjAtY2Q5YS00NzgzLTk4MDQtMTc0NzcxZjA5YTUxIiwiOWUxNjA4NjQtZjI5ZS00NmJjLTk4ZTktNzEzNzE4ZWRjYTA0IiwiMmExMzhiYzEtY2ZiNy00NzNmLTljMjEtODhiZjM0MWNmMmU2IiwiZjE0MmQ4NDAtNmNmYi00OTNlLTgxNGItMjA1ODNkNDI1NTYyIiwiYjI3NjZmMmQtOGI1OS00ZTZkLWI0NWEtYWM4YTYwZDg5OTZhIiwiNGJmNGIxOGYtYmQyZi00ZDNiLWE2ZTktZWI4Yjc2NjQ5MDA3Il0sImlwYWRkciI6IjE2MC42Ny4xMTUuNDciLCJuYW1lIjoiTWFydGluIEJydXNldCBTb2xiZXJnIFRtcCIsIm9pZCI6IjM3OTA2MzVhLTM4Y2UtNDA4Yy1hODg0LTQwNTYxOGE0MmM2NyIsIm9ucHJlbV9zaWQiOiJTLTEtNS0yMS0zMTI3MjU1My0xMjcwMzEwMDQ0LTE4MDQzNjQ5ODUtMzEyNTQxIiwicHVpZCI6IjEwMDMyMDAxRTI1M0Y2NzYiLCJyaCI6IjAuQVR3QUdNQXFuVVBvRkU2ZUswNE4yc2RVVUVaSWYza0F1dGRQdWtQYXdmajJNQk04QUpNLiIsInNjcCI6InVzZXJfaW1wZXJzb25hdGlvbiIsInN1YiI6Ik5KR1lQOFlCeTZDS3VucXhBOEZURXlLZTE3S1dWTkdBcXY2dXpVQVFFOTQiLCJ0aWQiOiI5ZDJhYzAxOC1lODQzLTRlMTQtOWUyYi00ZTBkZGFjNzU0NTAiLCJ1bmlxdWVfbmFtZSI6InRtcDY0OTcwNkBucmsubm8iLCJ1cG4iOiJ0bXA2NDk3MDZAbnJrLm5vIiwidXRpIjoiLU9NdGc3X18wMFdpS0hTM0lZUWVBQSIsInZlciI6IjEuMCIsIndpZHMiOlsiYjc5ZmJmNGQtM2VmOS00Njg5LTgxNDMtNzZiMTk0ZTg1NTA5Il0sInhtc190Y2R0IjoxMzY2ODkyOTQwfQ.fvHJKExg1ERDY32y53AdaJcXKBQFTWnG590mWT5K7Z3PEJrkUDjineGWa0bmyb7pHD0TpZHEAVB7agJvgVNoII2At1CnpysSk-wlSV-9sZkN2x3Z7UYuYu1WnzGf3YrRzt2xXFnB5szq8zbu3nJY9Q5t4aN0Vh2t0IM2meTLsp0ySuSiCoQBJobetxPIOz05GD84NNhKZRb9Vgj9hQYyRQ8rC2XqKlfd6Nt7OZ3QjMVeM8GvT4adzyspUKLICqDysD3qJRez27mdPhJ0ddBugi_-LPSMssm047uy6lRaPqrisIp2g8TeHqM5yLcxyboV9rz-PuJz2RmsynhHYkbvMA"
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return rsaSampleSecret, nil
	// })

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("%T\n", token)

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Println(claims["foo"], claims["nbf"])
	// } else {
	// 	fmt.Println(err)
	// }
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("Name: ", claims["name"])
		fmt.Println("Groups: ", claims["groups"])
	} else {
		fmt.Println("Error: ", err)
	}

}
