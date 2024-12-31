package models

type Patient struct {
    PatientID int `json:"id"`
    Name string `json:"name"`
    LastName string `json:"last_name"`
    Address string `json:"address"`
    Phone string `json:"phone_number"`
    Age int `json:"age"`
    Doctor int `json:"doctor_id"`
    Prescription int `json:"prescription_id"`
    Hospital int `json:"hospital_id"`
}
