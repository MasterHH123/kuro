package models

type Hospital struct {
    HospitalID int `json:"id"`
    Name string `form:"name" json:"name"`
    Address string `form:"address" json:"address"`
    City string `form:"city" json:"city"`
}
