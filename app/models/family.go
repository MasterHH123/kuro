package models

import "gorm.io/gorm"

type Family struct {
    gorm.Model
    FamilyID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Phone string `json:"phone_number"`
    Email string `json:"email"`
    Patient int `json:"patient_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}
