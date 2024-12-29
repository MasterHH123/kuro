package models

import "gorm.io/gorm"

type Doctor struct {
    gorm.Model
    DoctorID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Hospital string `json:"hospital_id"`
}
