package models

import (
    "time"
)

type Prescription struct {
    PrescriptionID int `json:"id"`
    Doctor int `json:"doctor_id"`
    Patient int `json:"patient_id"`
    Date time.Time `json:"date"`
}
