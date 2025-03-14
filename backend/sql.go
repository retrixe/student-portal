package main

import (
	"database/sql"
	"log"
)

func CreateSqlTables() {
	if _, err := db.Exec(`BEGIN;

CREATE TABLE IF NOT EXISTS users (
	username VARCHAR(16) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(319) NOT NULL UNIQUE,
	id UUID NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	verified BOOLEAN NOT NULL DEFAULT FALSE);

COMMIT;`); err != nil {
		log.Fatalln("Failed to create tables and indexes!", err)
	}
}

var (
	findUserByTokenStmt       *sql.Stmt
	findUserByNameOrEmailStmt *sql.Stmt

	insertTokenStmt *sql.Stmt
	deleteTokenStmt *sql.Stmt
)

func PrepareSqlStatements() {
	findUserByTokenStmt = prepareQuery("SELECT username, password, email, id, users.created_at " +
		"AS user_created_at, verified, token, tokens.created_at AS token_created_at FROM tokens " +
		"JOIN users ON tokens.user_id = users.id WHERE token = ?;")
	findUserByNameOrEmailStmt = prepareQuery("SELECT username, password, email, id, created_at, verified FROM users " +
		"WHERE username = ? OR email = ? LIMIT 1;")

	insertTokenStmt = prepareQuery("INSERT INTO tokens (token, created_at, user_id) VALUES (?, ?, ?);")
	deleteTokenStmt = prepareQuery("DELETE FROM tokens WHERE token = ? RETURNING user_id;")
}

func prepareQuery(query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("failed to build SQL query:", query, err)
	}
	return stmt
}
