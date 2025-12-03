package dto

type FXRate struct {
	Pair string  `json:"pair"`
	Rate float64 `json:"rate"`
}

type FXConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type FXConvertResponse struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Rate      float64 `json:"rate"`
	Converted float64 `json:"converted"`
}
