package main

import (
	"database/sql"
	"log"
)

func CreateSqlTables() {
	if _, err := db.Exec(`BEGIN;

CREATE TABLE IF NOT EXISTS students (
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
    address VARCHAR(1000) NOT NULL,
    picture BLOB NOT NULL,
    password VARCHAR(255) NOT NULL  -- Added password field
);

CREATE TABLE IF NOT EXISTS tokens (
    token VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    prn BIGINT NOT NULL REFERENCES students(prn)
);

CREATE TABLE IF NOT EXISTS guardians (
    id BIGINT PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    profession VARCHAR(30) NOT NULL,
    phoneNo VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL,
    name VARCHAR(30) NOT NULL,
    prn BIGINT NOT NULL,
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE IF NOT EXISTS serviceRequest (
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

CREATE TABLE IF NOT EXISTS serviceRequestReplies (
    id INT PRIMARY KEY,
    requestid INT NOT NULL,
    time DATETIME NOT NULL,
    prn BIGINT NOT NULL,
    content VARCHAR(5000) NOT NULL,
    FOREIGN KEY (requestid) REFERENCES serviceRequest(id),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE IF NOT EXISTS programs (
    code VARCHAR(30) PRIMARY KEY,
    degree VARCHAR(30) NOT NULL,
    major VARCHAR(30) NOT NULL,
    fees BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    code VARCHAR(16) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    credits INT NOT NULL,
    prerequisites VARCHAR(100) NOT NULL,
    semester INT NOT NULL
);

CREATE TABLE IF NOT EXISTS enrollments (
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

CREATE TABLE IF NOT EXISTS receipts (
    id INT PRIMARY KEY,
    prn BIGINT NOT NULL,
    ipAddress VARBINARY(16) NOT NULL,
    purchaseTime DATETIME NOT NULL,
    total DECIMAL(10,3) NOT NULL,
    paymentSystemId VARCHAR(255),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE IF NOT EXISTS receiptItems (
    receiptId INT,
    item VARCHAR(64),
    quantity INT NOT NULL,
    subtotal DECIMAL(10,3) NOT NULL,
    PRIMARY KEY (receiptId, item),
    FOREIGN KEY (receiptId) REFERENCES receipts(id)
);

CREATE TABLE IF NOT EXISTS internalAssessments (
    courseCode VARCHAR(16),
    prn BIGINT,
    name VARCHAR(30) NOT NULL,
    marks INT NOT NULL,
    total BIGINT NOT NULL,
    PRIMARY KEY (courseCode, prn),
    FOREIGN KEY (courseCode) REFERENCES courses(code),
    FOREIGN KEY (prn) REFERENCES students(prn)
);

CREATE TABLE IF NOT EXISTS circulars (
    id INT PRIMARY KEY,
    date DATE NOT NULL,
    subject VARCHAR(100) NOT NULL,
    description VARCHAR(5000) NOT NULL
);

CREATE TABLE IF NOT EXISTS holidays (
    name VARCHAR(30) PRIMARY KEY,
    start DATE NOT NULL,
    end DATE NOT NULL,
    type VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS lectures (
    id INT PRIMARY KEY,
    courseCode VARCHAR(16) NOT NULL,
    room VARCHAR(20) NOT NULL,
    timings VARCHAR(500) NOT NULL,
    year INT NOT NULL,
    professor VARCHAR(30) NOT NULL,
    FOREIGN KEY (courseCode) REFERENCES courses(code)
);

CREATE TABLE IF NOT EXISTS attendance (
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
	findStudentByTokenStmt *sql.Stmt
	findStudentByEmailStmt *sql.Stmt

	insertTokenStmt *sql.Stmt
	deleteTokenStmt *sql.Stmt
)

func PrepareSqlStatements() {
	findStudentByEmailStmt = prepareQuery("SELECT prn, password FROM students WHERE email = ?;")
	insertTokenStmt = prepareQuery("INSERT INTO tokens (token, created_at, prn) VALUES (?, ?, ?);")
	deleteTokenStmt = prepareQuery("DELETE FROM tokens WHERE token = ? RETURNING prn;")
}

func prepareQuery(query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("failed to build SQL query:", query, err)
	}
	return stmt
}
