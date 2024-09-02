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
	// Define command-line flags with short versions
	allDataFlag := flag.Bool("all-data", false, "Get all available Bitcoin data")
	priceFlag := flag.Bool("price", false, "Get the current Bitcoin price")
	volumeFlag := flag.Bool("volume", false, "Get the 24-hour Bitcoin trading volume")
	marketCapFlag := flag.Bool("market-cap", false, "Get the Bitcoin market capitalization")
	athPriceFlag := flag.Bool("ath-price", false, "Get the all-time high (ATH) price of Bitcoin")
	percentChangeFlags := map[string]*bool{
		"15m": flag.Bool("pc-15m", false, "Get the 15-minute percent change in Bitcoin price"),
		"30m": flag.Bool("pc-30m", false, "Get the 30-minute percent change in Bitcoin price"),
		"1h":  flag.Bool("pc-1h", false, "Get the 1-hour percent change in Bitcoin price"),
		"6h":  flag.Bool("pc-6h", false, "Get the 6-hour percent change in Bitcoin price"),
		"12h": flag.Bool("pc-12h", false, "Get the 12-hour percent change in Bitcoin price"),
		"24h": flag.Bool("pc-24h", false, "Get the 24-hour percent change in Bitcoin price"),
		"7d":  flag.Bool("pc-7d", false, "Get the 7-day percent change in Bitcoin price"),
		"30d": flag.Bool("pc-30d", false, "Get the 30-day percent change in Bitcoin price"),
		"1y":  flag.Bool("pc-1y", false, "Get the 1-year percent change in Bitcoin price"),
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
	fmt.Printf("Percent Change in 15m: %.2f%%\n", data.PercentChange15m)
	fmt.Printf("Percent Change in 30m: %.2f%%\n", data.PercentChange30m)
	fmt.Printf("Percent Change in 1h: %.2f%%\n", data.PercentChange1h)
	fmt.Printf("Percent Change in 6h: %.2f%%\n", data.PercentChange6h)
	fmt.Printf("Percent Change in 12h: %.2f%%\n", data.PercentChange12h)
	fmt.Printf("Percent Change in 24h: %.2f%%\n", data.PercentChange24h)
	fmt.Printf("Percent Change in 7d: %.2f%%\n", data.PercentChange7d)
	fmt.Printf("Percent Change in 30d: %.2f%%\n", data.PercentChange30d)
	fmt.Printf("Percent Change in 1y: %.2f%%\n", data.PercentChange1y)
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
	var timePeriod string

	switch period {
	case "15m":
		percentChange = data.PercentChange15m
		timePeriod = "15m"
	case "30m":
		percentChange = data.PercentChange30m
		timePeriod = "30m"
	case "1h":
		percentChange = data.PercentChange1h
		timePeriod = "1h"
	case "6h":
		percentChange = data.PercentChange6h
		timePeriod = "6h"
	case "12h":
		percentChange = data.PercentChange12h
		timePeriod = "12h"
	case "24h":
		percentChange = data.PercentChange24h
		timePeriod = "24h"
	case "7d":
		percentChange = data.PercentChange7d
		timePeriod = "7d"
	case "30d":
		percentChange = data.PercentChange30d
		timePeriod = "30d"
	case "1y":
		percentChange = data.PercentChange1y
		timePeriod = "1y"
	}

	fmt.Printf("Percent Change in %s: %.2f%%\n", timePeriod, percentChange)
}
