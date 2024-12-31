package models

type Family struct {
    FamilyID int `json:"id"`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Phone string `json:"phone_number"`
    Email string `json:"email"`
    Patient int `json:"patient_id"`
}
