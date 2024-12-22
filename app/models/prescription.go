package main

import "time"

type prescription struct {
    ID int `json:"id"`
    DoctorID int `json:"doctor_id"`
    PatientId int `json:patient_id`
    Date time.Time `json:"date"`
}
