package fixer

type Currencies struct {
	Base    string `json:"base"`
	Date    string `json:"date"`
	Rates   Rates  `json:"rates"`
	Success bool   `json:"success"`
}

type Rates map[string]float32
