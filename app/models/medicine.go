package models

type Medicine struct {
    MedicineID int `json:"id"`
    Name string `json:"name"`
    ActiveIngredient string `json:"active_ingredient"`
}
