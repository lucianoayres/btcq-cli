# BTCQ - Bitcoin CLI Query Tool

`btcq` is a command-line tool for querying Bitcoin data from a cryptocurrency API. This tool allows you to fetch various pieces of information about Bitcoin, including the current price, trading volume, market capitalization, all-time high price, and percentage changes over different time windows.

## Installation

To build the `btcq` CLI tool, run the following command in `src` directory of the project:

```bash
go build -o btcq
```

This will generate an executable named `btcq` in your directory. After building the project, you can run the executable with:

```sh
./btcq
```

Alternatively, you can move the binary to a directory that's included in your `PATH` to run it directly. Example:

```sh
mv btcq /usr/local/bin/
btcq
```

This will allow you to execute the tool from anywhere in your terminal.

## Usage

You can use `btcq` with various command-line flags to retrieve specific information about Bitcoin. The available flags are:

### Flags

-   `-all-data` : Get all available Bitcoin data.
-   `-price` : Get the current Bitcoin price.
-   `-volume` : Get the 24-hour Bitcoin trading volume.
-   `-market-cap` : Get the Bitcoin market capitalization.
-   `-ath-price` : Get the all-time high (ATH) price of Bitcoin.

### Percent Change Flags

You can use either the long or short flags to get percentage changes:

-   `-percent-change-15m` or `-pc-15m` : Get the 15-minute percent change in Bitcoin price.
-   `-percent-change-30m` or `-pc-30m` : Get the 30-minute percent change in Bitcoin price.
-   `-percent-change-1h` or `-pc-1h` : Get the 1-hour percent change in Bitcoin price.
-   `-percent-change-6h` or `-pc-6h` : Get the 6-hour percent change in Bitcoin price.
-   `-percent-change-12h` or `-pc-12h` : Get the 12-hour percent change in Bitcoin price.
-   `-percent-change-24h` or `-pc-24h` : Get the 24-hour percent change in Bitcoin price.
-   `-percent-change-7d` or `-pc-7d` : Get the 7-day percent change in Bitcoin price.
-   `-percent-change-30d` or `-pc-30d` : Get the 30-day percent change in Bitcoin price.
-   `-percent-change-1y` or `-pc-1y` : Get the 1-year percent change in Bitcoin price.

### Examples

-   Get the current Bitcoin price:

    ```bash
    ./btcq
    ```

-   Get all available Bitcoin data:

    ```bash
    ./btcq -all-data
    ```

-   Get the 1-hour percent change in Bitcoin price:

    ```bash
    ./btcq -percent-change-1h
    ```

-   Get the 7-day percent change in Bitcoin price using the short flag:

    ```bash
    ./btcq -pc-7d
    ```

## API Reference

The data is fetched from the [CoinPaprika API](https://api.coinpaprika.com/v1/tickers/btc-bitcoin). For more information about the API, refer to their [documentation](https://api.coinpaprika.com/).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
