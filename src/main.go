package main

import (
	"btcq/api"
	"btcq/config"
	"btcq/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	allDataFlag := flag.Bool("all-data", false, "Get all available Bitcoin data")
	priceFlag := flag.Bool("price", false, "Get the current Bitcoin price")
	volumeFlag := flag.Bool("volume", false, "Get the 24-hour Bitcoin trading volume")
	marketCapFlag := flag.Bool("market-cap", false, "Get the Bitcoin market capitalization")
	athPriceFlag := flag.Bool("ath-price", false, "Get the all-time high (ATH) price of Bitcoin")
	percentChangeFlags := map[string]*bool{
		"15-minute": flag.Bool("percent-change-15m", false, "Get the 15-minute percent change in Bitcoin price"),
		"30-minute": flag.Bool("percent-change-30m", false, "Get the 30-minute percent change in Bitcoin price"),
		"1-hour":    flag.Bool("percent-change-1h", false, "Get the 1-hour percent change in Bitcoin price"),
		"6-hour":    flag.Bool("percent-change-6h", false, "Get the 6-hour percent change in Bitcoin price"),
		"12-hour":   flag.Bool("percent-change-12h", false, "Get the 12-hour percent change in Bitcoin price"),
		"24-hour":   flag.Bool("percent-change-24h", false, "Get the 24-hour percent change in Bitcoin price"),
		"7-day":     flag.Bool("percent-change-7d", false, "Get the 7-day percent change in Bitcoin price"),
		"30-day":    flag.Bool("percent-change-30d", false, "Get the 30-day percent change in Bitcoin price"),
		"1-year":    flag.Bool("percent-change-1y", false, "Get the 1-year percent change in Bitcoin price"),
	}
	
	flag.Parse()

	// Start loading animation
	stopChan := make(chan bool)
	go utils.LoadingAnimation(stopChan)

	// Fetch data from the API
	data, err := api.FetchBitcoinData(config.BitcoinAPIURL)
	if err != nil {
		stopChan <- true
		fmt.Printf("Error fetching Bitcoin data: %v\n", err)
		os.Exit(1)
	}
	stopChan <- true

	// Handle the flags
	if len(os.Args) == 1 || *priceFlag {
		handlePrice(data)
	}
	if *allDataFlag {
		handleAllData(data)
	}
	if *volumeFlag {
		handleVolume(data)
	}
	if *marketCapFlag {
		handleMarketCap(data)
	}
	if *athPriceFlag {
		handleATHPrice(data)
	}
	for period, flag := range percentChangeFlags {
		if *flag {
			handlePercentChange(period, data)
		}
	}
}

func handlePrice(data *api.BitcoinData) {
	fmt.Printf("Current Bitcoin Price: %.2f USD\n", data.Price)
}

func handleAllData(data *api.BitcoinData) {
	fmt.Printf("Price: %.2f USD\n", data.Price)
	fmt.Printf("24-hour Trading Volume: %.2f USD\n", data.Volume24h)
	fmt.Printf("Market Capitalization: %.2f USD\n", data.MarketCap)
	fmt.Printf("All-Time High Price: %.2f USD\n", data.ATHPrice)
	fmt.Printf("15-minute Percent Change: %.2f%%\n", data.PercentChange15m)
	fmt.Printf("30-minute Percent Change: %.2f%%\n", data.PercentChange30m)
	fmt.Printf("1-hour Percent Change: %.2f%%\n", data.PercentChange1h)
	fmt.Printf("6-hour Percent Change: %.2f%%\n", data.PercentChange6h)
	fmt.Printf("12-hour Percent Change: %.2f%%\n", data.PercentChange12h)
	fmt.Printf("24-hour Percent Change: %.2f%%\n", data.PercentChange24h)
	fmt.Printf("7-day Percent Change: %.2f%%\n", data.PercentChange7d)
	fmt.Printf("30-day Percent Change: %.2f%%\n", data.PercentChange30d)
	fmt.Printf("1-year Percent Change: %.2f%%\n", data.PercentChange1y)
}

func handleVolume(data *api.BitcoinData) {
	fmt.Printf("24-hour Trading Volume: %.2f USD\n", data.Volume24h)
}

func handleMarketCap(data *api.BitcoinData) {
	fmt.Printf("Market Capitalization: %.2f USD\n", data.MarketCap)
}

func handleATHPrice(data *api.BitcoinData) {
	fmt.Printf("All-Time High Price: %.2f USD\n", data.ATHPrice)
}

func handlePercentChange(period string, data *api.BitcoinData) {
	var percentChange float64
	switch period {
	case "15-minute":
		percentChange = data.PercentChange15m
	case "30-minute":
		percentChange = data.PercentChange30m
	case "1-hour":
		percentChange = data.PercentChange1h
	case "6-hour":
		percentChange = data.PercentChange6h
	case "12-hour":
		percentChange = data.PercentChange12h
	case "24-hour":
		percentChange = data.PercentChange24h
	case "7-day":
		percentChange = data.PercentChange7d
	case "30-day":
		percentChange = data.PercentChange30d
	case "1-year":
		percentChange = data.PercentChange1y
	}
	fmt.Printf("%s Percent Change: %.2f%%\n", period, percentChange)
}
