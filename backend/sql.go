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

CREATE TABLE IF NOT EXISTS tokens (
	token VARCHAR(128) NOT NULL PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	user_id UUID NOT NULL REFERENCES users(id));

CREATE TABLE students (
    prn BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    aadhaarNo VARCHAR(12) NOT NULL,
    bloodGroup VARCHAR(4) NOT NULL,
    dob DATE NOT NULL,
    gender VARCHAR(16) NOT NULL,
    admissionDate DATE NOT NULL,
    semester INT NOT NULL,
    email VARCHAR(100) NOT NULL,
    programCode VARCHAR(16) NOT NULL,
    phoneNo VARCHAR(100) NOT NULL,
    address VARCHAR(1000) NOT NULL
);

CREATE TABLE guardians (
    id BIGINT PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    profession VARCHAR(30) NOT NULL,
    phoneNo VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL,
    name VARCHAR(30) NOT NULL,
    prn BIGINT NOT NULL,
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE serviceRequest (
    id INT PRIMARY KEY,
    creationTime DATETIME NOT NULL,
    prn BIGINT NOT NULL,
    title VARCHAR(30) NOT NULL,
    status VARCHAR(30) NOT NULL,
    description VARCHAR(5000) NOT NULL,
    department VARCHAR(30) NOT NULL,
    category VARCHAR(30) NOT NULL,
    assignedTo VARCHAR(30) NOT NULL,
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE serviceRequestReplies (
    id INT PRIMARY KEY,
    requestid INT NOT NULL,
    time DATETIME NOT NULL,
    prn BIGINT NOT NULL,
    content VARCHAR(5000) NOT NULL,
    FOREIGN KEY (requestid) REFERENCES serviceRequest(id),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE programs (
    code VARCHAR(30) PRIMARY KEY,
    degree VARCHAR(30) NOT NULL,
    major VARCHAR(30) NOT NULL,
    fees BIGINT NOT NULL
);

CREATE TABLE courses (
    code VARCHAR(16) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    credits INT NOT NULL,
    prerequisites VARCHAR(100) NOT NULL,
    semester INT NOT NULL
);

CREATE TABLE enrollments (
    courseCode VARCHAR(16),
    prn BIGINT,
    attempts INT NOT NULL,
    grade VARCHAR(10) NOT NULL,
    semester VARCHAR(20) NOT NULL,
    externalMarks INT NOT NULL,
    internalMarks INT NOT NULL,
    lectureId INT NOT NULL,
    PRIMARY KEY (courseCode, prn),
    FOREIGN KEY (courseCode) REFERENCES courses(code),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE receipts (
    id INT PRIMARY KEY,
    prn BIGINT NOT NULL,
    ipAddress VARBINARY(16) NOT NULL,
    purchaseTime TIMESTAMP NOT NULL,
    total DECIMAL(10,3) NOT NULL,
    paymentSystemId VARCHAR(255),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE receiptItems (
    receiptId INT,
    item VARCHAR(64),
    quantity INT NOT NULL,
    subtotal DECIMAL(10,3) NOT NULL,
    PRIMARY KEY (receiptId, item),
    FOREIGN KEY (receiptId) REFERENCES receipts(id)
);

CREATE TABLE internalAssessments (
    courseCode VARCHAR(16),
    prn BIGINT,
    name VARCHAR(30) NOT NULL,
    marks INT NOT NULL,
    total BIGINT NOT NULL,
    PRIMARY KEY (courseCode, prn),
    FOREIGN KEY (courseCode) REFERENCES courses(code),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE circulars (
    id INT PRIMARY KEY,
    date DATE NOT NULL,
    subject VARCHAR(100) NOT NULL,
    description VARCHAR(5000) NOT NULL
);

CREATE TABLE holidays (
    name VARCHAR(30) PRIMARY KEY,
    start DATE NOT NULL,
    end DATE NOT NULL,
    type VARCHAR(30) NOT NULL
);

CREATE TABLE lectures (
    id INT PRIMARY KEY,
    courseCode VARCHAR(16) NOT NULL,
    room VARCHAR(20) NOT NULL,
    timings VARCHAR(500) NOT NULL,
    year INT NOT NULL,
    professor VARCHAR(30) NOT NULL,
    FOREIGN KEY (courseCode) REFERENCES courses(code)
);

CREATE TABLE attendance (
    prn BIGINT,
    lectureId INT,
    timing VARCHAR(32),
    present BOOL NOT NULL,
    professor VARCHAR(30) NOT NULL,
    PRIMARY KEY (prn, lectureId, timing),
    FOREIGN KEY (prn) REFERENCES students(prn),
    FOREIGN KEY (lectureId) REFERENCES lectures(id)
);


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
