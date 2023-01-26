// fjwt v1.0.0
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin
package cracker

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/alpkeskin/fjwt/cmd/utils.go"
	"github.com/gammazero/workerpool"
	"github.com/golang-jwt/jwt/v4"
)

func Handler(tokenString string, wordlist string, threads int) {
	utils.Spinner.Start()
	var alg string = GetJWTAlgorithm(tokenString)    // get algorithm from jwt header
	sign := jwt.GetSigningMethod(alg)                // get signing method
	arr, err := utils.ReadFileToStringList(wordlist) // read wordlist to string array
	if err != nil {
		log.Fatal(err)
	}
	Crack(arr, tokenString, sign, threads) // crack jwt
	utils.Spinner.Stop()
}

func Crack(arr []string, jwt string, sign jwt.SigningMethod, threads int) {
	start := time.Now()
	wp := workerpool.New(threads)
	for _, s := range arr {
		s := s
		wp.Submit(func() {
			if IsCorrectSecret(jwt, s, sign) {
				finish := time.Now()
				utils.PrintResult(s, finish.Sub(start))
				os.Exit(0)
			}
		})
	}
	wp.StopWait()
	println("\x1b[31m[NOT FOUND]\x1b[0m")
}

func IsCorrectSecret(tokenString string, secret string, SignMethod jwt.SigningMethod) bool {
	defer func() {
		atomic.AddUint64(&utils.Counter, 1)
	}()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != SignMethod {
			return nil, nil
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}

func GetJWTAlgorithm(token string) string {
	parts := strings.Split(token, ".")
	headerDecoded, err := base64.RawURLEncoding.DecodeString(parts[0]) // header
	if err != nil {
		log.Fatal(err)
	}
	jsonParsed, err := gabs.ParseJSON(headerDecoded)
	if err != nil {
		log.Fatal(err)
	}
	return jsonParsed.Path("alg").Data().(string)
}
