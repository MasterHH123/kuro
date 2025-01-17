package models

type Doctor struct {
    DoctorID int `json:"id"`
    Name string `form:"name" json:"name"`
    LastName string `form:"last_name" json:"last_name"`
    Hospital string `form:"hospital" json:"hospital"`
}
