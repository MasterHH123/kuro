package models

import "gorm.io/gorm"

type Patient struct {
    gorm.Model
    PatientID int `json:"id gorm:"unique;primaryKey;autoIncrement""`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Address string `json:"address"`
    Phone string `json:"phone_number"`
    Age int `json:"age"`
    Doctor int `json:"doctor_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    Prescription int `json:"prescription_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    Hospital int `json:"hospital_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}
