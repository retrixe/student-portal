package main

import (
	"time"
)

type Token struct {
	CreatedAt time.Time `json:"createdAt"`
	Token     string    `json:"token"`
	PRN       int64     `json:"prn"`
}

type Student struct {
	PRN           int64     `json:"prn"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	ProgramCode   string    `json:"program_code"`
	PhoneNo       string    `json:"phone_no"`
	AadhaarNo     string    `json:"aadhaar_no"`
	BloodGroup    string    `json:"blood_group"`
	DOB           time.Time `json:"dob"`
	Gender        string    `json:"gender"`
	AdmissionDate time.Time `json:"admission_date"`
	Semester      int       `json:"semester"`
	Address       string    `json:"address"`
	Picture       []byte    `json:"picture"`
}
