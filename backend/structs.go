package main

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	CreatedAt time.Time `json:"createdAt"`
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"userId"`
}

type Student struct {
	PRN           int64     `json:"prn"`
	Name          string    `json:"name"`
	AadhaarNo     string    `json:"aadhaarNo"`
	BloodGroup    string    `json:"bloodGroup"`
	DOB           time.Time `json:"dob"`
	Gender        string    `json:"gender"`
	AdmissionDate time.Time `json:"admissionDate"`
	Semester      int       `json:"semester"`
	Email         string    `json:"email"`
	ProgramCode   string    `json:"programCode"`
	PhoneNo       string    `json:"phoneNo"`
	Address       string    `json:"address"`
	Picture       []byte    `json:"picture"`
}
