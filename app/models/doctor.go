package models

type Doctor struct {
    DoctorID int `json:"id"`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Hospital string `json:"hospital_id"`
}
