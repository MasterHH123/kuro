package main

type medicine struct {
    ID int `json:"id"`
    Name string `json:"name"`
    ActiveIngredient string `json:"active_ingredient"`
    DosageForm int `json:"dosage_form"`
}
