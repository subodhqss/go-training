//package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// var jwtKey = []byte("secret_key")

// var employee = map[string]string{
// 	"dmurphy@classicmodelcars.com": "pass3",
// 	//"user2": "password2",

// }

// type Credentials struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	var credentials Credentials
// 	err := json.NewDecoder(r.Body).Decode(&credentials)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	expectedPassword, ok := employee[credentials.Email]

// 	if !ok || expectedPassword != credentials.Password {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	expirationTime := time.Now().Add(time.Minute * 5)
// 	claims := &Claims{
// 		Email: credentials.Email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})

// }

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tokenStr := cookie.Value

// 	claims := &Claims{}

// 	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
// 		func(t *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil

// 		})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Email)))

// }

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{

	//"user1": "password1",
	//"user2": "password2",
	"abcxyz@gmail.com_update":        "pass2",
	"shubi@gmail.com_update":         "pass1",
	"dmurphy@classicmodelcars.com":   "pass3",
	"mpatterso@classicmodelcars.com": "pass4",
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//Firstname string `json:"firstname"`
}

type Claims struct {
	Email string `json:"email"`
	//Firstname string `json:"firstname"`

	jwt.StandardClaims
}

type ResponseData struct {
	Expires time.Time `json:"expire"`
	Token   string    `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Print("Err1 : ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[credentials.Email]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		log.Print("Err2 : ", err)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Err3 : ", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// res := &ResponseData{
	// 	Token:   tokenString,
	// 	Expires: expirationTime,
	// }

	jsonData, _ := json.Marshal(claims)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Email)))

}
