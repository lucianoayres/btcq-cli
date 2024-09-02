package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type BitcoinData struct {
    Price            float64 `json:"price"`
    Volume24h        float64 `json:"volume_24h"`
    MarketCap        float64 `json:"market_cap"`
    ATHPrice         float64 `json:"ath_price"`
    PercentChange15m float64 `json:"percent_change_15m"`
    PercentChange30m float64 `json:"percent_change_30m"`
    PercentChange1h  float64 `json:"percent_change_1h"`
    PercentChange6h  float64 `json:"percent_change_6h"`
    PercentChange12h float64 `json:"percent_change_12h"`
    PercentChange24h float64 `json:"percent_change_24h"`
    PercentChange7d  float64 `json:"percent_change_7d"`
    PercentChange30d float64 `json:"percent_change_30d"`
    PercentChange1y  float64 `json:"percent_change_1y"`
}

func FetchBitcoinData(apiURL string) (*BitcoinData, error) {
    response, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, errors.New("failed to fetch Bitcoin data")
    }

    var jsonResponse struct {
        Quotes struct {
            USD BitcoinData `json:"USD"`
        } `json:"quotes"`
    }

    err = json.NewDecoder(response.Body).Decode(&jsonResponse)
    if err != nil {
        return nil, err
    }

    return &jsonResponse.Quotes.USD, nil
}
