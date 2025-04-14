package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

var ErrNotAuthenticated = errors.New("request not authenticated")

func IsAuthenticatedHTTP(w http.ResponseWriter, r *http.Request) (*Student, *Token) {
	student, token, err := IsAuthenticated(GetTokenFromHTTP(r))
	if errors.Is(err, ErrNotAuthenticated) {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
	} else if err != nil {
		handleInternalServerError(w, err)
	}
	return student, token
}

func IsAuthenticated(token string) (*Student, *Token, error) {
	if token == "" {
		return nil, nil, ErrNotAuthenticated
	}

	student := Student{}
	var tokenCreatedAt time.Time
	err := findStudentByTokenStmt.QueryRow(token).Scan(
		&student.PRN,
		&student.Name,
		&student.Email,
		&student.ProgramCode,
		&student.PhoneNo,
		&tokenCreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotAuthenticated
	} else if err != nil {
		return nil, nil, err
	} else {
		return &student, &Token{CreatedAt: tokenCreatedAt, Token: token}, nil
	}
}

func StatusEndpoint(w http.ResponseWriter, r *http.Request) {
	student, _, err := IsAuthenticated(GetTokenFromHTTP(r))
	if errors.Is(err, ErrNotAuthenticated) {
		w.Write([]byte("{\"online\":true,\"authenticated\":false}"))
	} else if err != nil {
		handleInternalServerError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Online        bool    `json:"online"`
			Authenticated bool    `json:"authenticated"`
			Student       Student `json:"student"`
		}{
			Online:        true,
			Authenticated: true,
			Student:       *student,
		})
	}
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, errorJson("Unable to read body!"), http.StatusBadRequest)
		return
	}
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil || data.Email == "" || data.Password == "" {
		http.Error(w, errorJson("Invalid email or password!"), http.StatusBadRequest)
		return
	}

	var student Student
	var hashedPassword string

	err = findStudentByEmailStmt.QueryRow(data.Email).Scan(
		&student.PRN,
		&hashedPassword,
	)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, errorJson("No student account with this email exists!"), http.StatusUnauthorized)
		return
	} else if err != nil {
		handleInternalServerError(w, err)
		return
	} else if !ComparePassword(data.Password, hashedPassword) {
		http.Error(w, errorJson("Incorrect password!"), http.StatusUnauthorized)
		return
	}

	tokenBytes := make([]byte, 64)
	_, _ = rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	result, err := insertTokenStmt.Exec(token, time.Now().UTC(), student.PRN)

	if err != nil {
		handleInternalServerError(w, err)
		return
	} else if rows, err := result.RowsAffected(); err != nil || rows != 1 {
		handleInternalServerError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   config.SecureCookies,
		MaxAge:   3600 * 24 * 31,
		SameSite: http.SameSiteStrictMode,
		Path:     config.BasePath,
	})

	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
		PRN   int64  `json:"prn"`
	}{Token: token, PRN: student.PRN})
}

func LogoutEndpoint(w http.ResponseWriter, r *http.Request) {
	token := GetTokenFromHTTP(r)
	if token == "" {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
		return
	}
	var prn int64
	err := deleteTokenStmt.QueryRow(token).Scan(&prn)

	if err == sql.ErrNoRows {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
		return
	} else if err != nil {
		handleInternalServerError(w, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "null",
		HttpOnly: true,
		Secure:   config.SecureCookies,
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	})
	w.Write([]byte("{\"success\":true}"))
}
