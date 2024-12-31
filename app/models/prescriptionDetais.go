package models

type PrescriptionDetail struct {
    PrescriptionDetailsID int `json:"id"`
    Prescription int `json:"prescription"`
    Medicine int `json:"medicine"`
    TimesPerDay int `json:"times_per_day"`
}
