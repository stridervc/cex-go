package cex

/*
cex-go is a library to get currency exchange information from
https://currency-api.appspot.com
*/

import (
	"encoding/json"
	"errors"
	"net/http"
)

const baseURL = "https://currency-api.appspot.com/api/"

type CurrencyData struct {
	Success bool    `json:"success"`
	Source  string  `json:"source"`
	Target  string  `json:"target"`
	Rate    float64 `json:"rate"`
	Amount  float64 `json:"amount"`
	Message string  `json:"message"`
}

func ExchangeRate(source, target string) (CurrencyData, error) {
	url := baseURL + source + "/" + target + ".json"
	r, err := http.Get(url)
	if err != nil {
		return CurrencyData{}, err
	}
	defer r.Body.Close()

	var data CurrencyData
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&data)
	if err != nil {
		return data, err
	}

	if data.Success == false {
		return data, errors.New(data.Message)
	}

	return data, nil
}
