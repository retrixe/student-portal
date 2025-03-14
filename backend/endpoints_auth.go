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

	"github.com/google/uuid"
)

var ErrNotAuthenticated = errors.New("request not authenticated")

func IsAuthenticatedHTTP(w http.ResponseWriter, r *http.Request) (*User, *Token) {
	user, token, err := IsAuthenticated(GetTokenFromHTTP(r))
	if errors.Is(err, ErrNotAuthenticated) {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
	} else if err != nil {
		handleInternalServerError(w, err)
	}
	return user, token
}

func IsAuthenticated(token string) (*User, *Token, error) {
	if token == "" {
		return nil, nil, ErrNotAuthenticated
	}

	user := User{}
	var tokenCreatedAt time.Time
	err := findUserByTokenStmt.QueryRow(token).Scan(
		&user.Username,
		&user.Password,
		&user.Email,
		&user.ID,
		&user.CreatedAt,
		&user.Verified,
		&token,
		&tokenCreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotAuthenticated
	} else if err != nil {
		return nil, nil, err
	} else {
		return &user, &Token{CreatedAt: tokenCreatedAt, Token: token, UserID: user.ID}, nil
	}
}

func StatusEndpoint(w http.ResponseWriter, r *http.Request) {
	user, _, err := IsAuthenticated(GetTokenFromHTTP(r))
	if errors.Is(err, ErrNotAuthenticated) {
		w.Write([]byte("{\"online\":true,\"authenticated\":false}"))
	} else if err != nil {
		handleInternalServerError(w, err)
	} else {
		usernameJson, _ := json.Marshal(user.Username)
		userIdJson, _ := json.Marshal(user.ID)
		w.Write([]byte("{\"online\":true,\"authenticated\":true," +
			"\"username\":" + string(usernameJson) + ",\"userId\":" + string(userIdJson) + "}"))
	}
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	// Check the body for JSON containing username and password and return a token.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, errorJson("Unable to read body!"), http.StatusBadRequest)
		return
	}
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, errorJson("Unable to read body!"), http.StatusBadRequest)
		return
	} else if data.Username == "" || data.Password == "" {
		http.Error(w, errorJson("No username or password provided!"), http.StatusBadRequest)
		return
	}
	var user User
	err = findUserByNameOrEmailStmt.QueryRow(data.Username, data.Username).Scan(
		&user.Username, &user.Password, &user.Email, &user.ID, &user.CreatedAt, &user.Verified)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		http.Error(w, errorJson("No account with this username/email exists!"), http.StatusUnauthorized)
		return
	} else if err != nil {
		handleInternalServerError(w, err)
		return
	} else if !user.Verified {
		http.Error(w, errorJson("Your account is not verified yet!"), http.StatusForbidden)
		return
	} else if !ComparePassword(data.Password, user.Password) {
		http.Error(w, errorJson("Incorrect password!"), http.StatusUnauthorized)
		return
	}
	tokenBytes := make([]byte, 64)
	_, _ = rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)
	result, err := insertTokenStmt.Exec(token, time.Now().UTC(), user.ID)
	if err != nil {
		handleInternalServerError(w, err)
		return
	} else if rows, err := result.RowsAffected(); err != nil || rows != 1 {
		handleInternalServerError(w, err) // nil err solved by Ostrich algorithm
		return
	}
	// Add cookie to browser.
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
		Token    string `json:"token"`
		Username string `json:"username"`
	}{Token: token, Username: user.Username})
}

func LogoutEndpoint(w http.ResponseWriter, r *http.Request) {
	token := GetTokenFromHTTP(r)
	if token == "" {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
		return
	}
	var userID uuid.UUID
	err := deleteTokenStmt.QueryRow(token).Scan(&userID)
	if err == sql.ErrNoRows {
		http.Error(w, errorJson("You are not authenticated to access this resource!"),
			http.StatusUnauthorized)
		return
	} else if err != nil {
		handleInternalServerError(w, err)
		return
	}
	// Delete cookie on browser.
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
