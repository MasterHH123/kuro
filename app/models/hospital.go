package kuro

import "gorm.io/gorm"

type Hospital struct {
    gorm.Model
    HospitalID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Name string `json:"name"`
    Address string `json:"address"`
    City string `json:"city"`
}
