package models

type Hospital struct {
    HospitalID int `json:"id"`
    Name string `json:"name"`
    Address string `json:"address"`
    City string `json:"city"`
}
