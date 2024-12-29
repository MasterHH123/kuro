package models

import (
    "time"

    "gorm.io/gorm"
)

type Prescription struct {
    gorm.Model
    PrescriptionID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Doctor int `json:"doctor_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    Patient int `json:"patient_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    Date time.Time `json:"date"`
}
