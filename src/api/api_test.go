package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchBitcoinData_Success(t *testing.T) {
    // Create a mock server with a successful response
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulate a successful API response
        response := struct {
            Quotes struct {
                USD BitcoinData `json:"USD"`
            } `json:"quotes"`
        }{
            Quotes: struct {
                USD BitcoinData `json:"USD"`
            }{
                USD: BitcoinData{
                    Price:            50000.0,
                    Volume24h:        1000000.0,
                    MarketCap:        1000000000.0,
                    ATHPrice:         60000.0,
                    PercentChange15m: 1.5,
                    PercentChange30m: 2.5,
                    PercentChange1h:  3.5,
                    PercentChange6h:  4.5,
                    PercentChange12h: 5.5,
                    PercentChange24h: 6.5,
                    PercentChange7d:  7.5,
                    PercentChange30d: 8.5,
                    PercentChange1y:  9.5,
                },
            },
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }))
    defer mockServer.Close()

    data, err := FetchBitcoinData(mockServer.URL)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    expected := &BitcoinData{
        Price:            50000.0,
        Volume24h:        1000000.0,
        MarketCap:        1000000000.0,
        ATHPrice:         60000.0,
        PercentChange15m: 1.5,
        PercentChange30m: 2.5,
        PercentChange1h:  3.5,
        PercentChange6h:  4.5,
        PercentChange12h: 5.5,
        PercentChange24h: 6.5,
        PercentChange7d:  7.5,
        PercentChange30d: 8.5,
        PercentChange1y:  9.5,
    }

    if *data != *expected {
        t.Errorf("expected %+v, got %+v", expected, data)
    }
}

func TestFetchBitcoinData_Failure(t *testing.T) {
    // Create a mock server that returns an error
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))
    defer mockServer.Close()

    data, err := FetchBitcoinData(mockServer.URL)
    if err == nil {
        t.Fatal("expected an error, got nil")
    }
    if data != nil {
        t.Fatalf("expected nil data, got %+v", data)
    }
}

func TestFetchBitcoinData_InvalidJSON(t *testing.T) {
    // Create a mock server with invalid JSON response
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"quotes":{"USD":{"price":50000}}`)) // Invalid JSON
    }))
    defer mockServer.Close()

    data, err := FetchBitcoinData(mockServer.URL)
    if err == nil {
        t.Fatal("expected an error, got nil")
    }
    if data != nil {
        t.Fatalf("expected nil data, got %+v", data)
    }
}
