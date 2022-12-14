package controllers

// http headers must be set before sending any response through http.ResponseWriter else they will not be sent.

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sauravrana646/golang-learning/Golang_Auth/jwthttpserver/utils"
)

// Implement login function
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Function ran..")
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userClaim, err := utils.CreateClaim(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, expire, err := utils.CreateToken(&userClaim)
	if err != nil {
		http.Error(w, "Login Failed", http.StatusUnauthorized)
	}
	cookie := &http.Cookie{
		Name:    "jwttoken",
		Value:   token,
		Expires: time.Unix(expire, 0),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(`Login Success`))
	return
}

// Implement Home function
func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := fetchCookie(w, r)
	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println("Error Calling fetchCookie")
		http.Error(w, "Unauthorized. Error in Cookie", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	// wnew := RefreshToken(w,r)
	// if wnew != nil {
	// 	wnew.Write([]byte(`Token refreshed`))
	// 	wnew.Write([]byte("Hello " + claims.Username + " "))
	// 	return
	// }
	w.Write([]byte("Hello " + claims.Username + " "))
	return
}

// Implement RefreshToken function
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwttoken")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	}

	// fmt.Println("")
	// fmt.Println("Cookie:")
	// fmt.Println(cookie)

	claims, err := utils.ParseToken(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	// fmt.Println("")
	// fmt.Println("Claim before")
	// fmt.Println(claims)
	// Calculate time remaining and covert in human readable form
	tmr := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())

	if tmr > 20*time.Second {
		w.Write([]byte(`Still long time...`))
		w.Write([]byte("Time Remaing : " + tmr.String()))
		return
	}

	newexpire := time.Now().Add(60 * time.Second).Unix()
	claims.ExpiresAt = newexpire

	// fmt.Println("Claim after")
	// fmt.Println("")
	// fmt.Println(claims)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	newtoken, clmExpire, err := utils.CreateToken(claims)
	if err != nil {
		http.Error(w, "Refresh Failed", http.StatusUnauthorized)
	}

	newcookie := &http.Cookie{
		Name:    "jwttoken",
		Value:   newtoken,
		Expires: time.Unix(clmExpire, 0),
	}

	// fmt.Println("")
	// fmt.Println("New Cookie")
	// fmt.Println(newcookie)

	// w.Header().Add("Set-Cookie", newcookie.String())
	http.SetCookie(w, newcookie)

	// fmt.Println("")
	// fmt.Println("Header")
	// fmt.Println(w.Header())

	w.Write([]byte("Time Remaing : " + tmr.String()))
	return
}

func Base(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte("<center>Welcome to the JWT Server. Please login..!!!</center>"))
	// http.Redirect(w,r, "/login", http.StatusTemporaryRedirect)
}

func fetchCookie(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("jwttoken")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return nil, err
	}
	return cookie, nil
}
