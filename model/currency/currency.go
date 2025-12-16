package currency

type CBRResponse struct {
	Date   string            `json:"Date"`
	Valute map[Code]Currency `json:"Valute"`
}

type Currency struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}
