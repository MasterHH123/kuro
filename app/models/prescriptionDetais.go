package kuro

import "gorm.io/gorm"

type PrescriptionDetail struct {
    gorm.Model
    PrescriptionDetailsID int `json:"id" gorm:"unique;primaryKey;autoIncrement"`
    Prescription int `json:"prescription" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    Medicine int `json:"medicine" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
    TimesPerDay int `json:"times_per_day"`
}
