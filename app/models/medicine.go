package models

import "gorm.io/gorm"

type Medicine struct {
    gorm.Model
    MedicineID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Name string `json:"name"`
    ActiveIngredient string `json:"active_ingredient"`
}
