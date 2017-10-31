package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/iBatStat/extractor/db"
	"github.com/iBatStat/extractor/model"
	"golang.org/x/crypto/bcrypt"
)

// need the following functions handlers

func LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// extract user and password
	// bcrypt the password
	// look up the user on mongo
	// verify encoded password
	// if valid, return a token else unauthorised
	if r.Method == "POST" {
		var loginUser model.User

		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error decoding body")
			return
		}

		// validate if useremail and password not nil

		if loginUser.Email == "" || loginUser.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("UserEmail and Password mandatory"))
			return
		}

		// validate if the user exists in the db
		existingUser := db.DBAccess.GetUser(loginUser.Email)
		if existingUser == nil {
			writeError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid user name and or password", loginUser.Email)), w, "")
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUser.Password))
		if err != nil {
			writeError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid user name and or password", loginUser.Email)), w, "")
			return

		}

		token, err := generateNew(loginUser.Email, loginUser.PhoneModel)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error logging in user")
			return
		}
		w.WriteHeader(http.StatusCreated)
		out, _ := json.Marshal(model.Token{token})
		w.Write(out)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func AuthenticateHandlerFunc(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if authenticate(r) {
		next(w, r)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func NewUserHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// encode the password and save it in db
	// generate a token and return
	if r.Method == "POST" {
		var loginUser model.User

		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error decoding body")
			return
		}

		// validate if useremail and password not nil

		if loginUser.Email == "" || loginUser.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("UserEmail and Password mandatory"))
			return
		}

		// validate if the user exists in the db
		existingUser := db.DBAccess.GetUser(loginUser.Email)
		if existingUser != nil {
			writeError(http.StatusBadRequest, errors.New(fmt.Sprintf("user %s already exists", loginUser.Email)), w, "")
			return
		}

		dbpass, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), 21)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error creating new  user")
			return
		}
		loginUser.Password = string(dbpass)
		err = db.DBAccess.SaveUser(loginUser)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error creating new user")
			return
		}

		token, err := generateNew(loginUser.Email, loginUser.PhoneModel)
		if err != nil {
			writeError(http.StatusInternalServerError, err, w, "error logging in user")
			return
		}
		w.WriteHeader(http.StatusCreated)
		out, _ := json.Marshal(model.Token{token})
		w.Write(out)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

type errormsg struct {
	Error   string `json:"error"`
	Message string `json:"msg"`
}

func writeError(httpCode int, err error, w http.ResponseWriter, msg string) {
	outerror := errormsg{err.Error(), msg}
	out, _ := json.Marshal(outerror)
	w.WriteHeader(httpCode)
	w.Header().Add("Content-type", "application/json")
	w.Write(out)
}

func UploadImageHandlerFunc(w http.ResponseWriter, r *http.Request) {}
