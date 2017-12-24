package util

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Log logs information to console
func Log(message string) {
	fmt.Println("[ticker] " + message)
}

// LogErr logs information including time
func LogErr(message string) {
	log.Println("[ticker] " + getCurrentTime() + message)
}

// getCurrentTime returns current time in RFC1123 format as string
func getCurrentTime() string {
	return " [" + time.Now().Format(time.RFC1123) + "] "
}

// FormatPrice returns a formatted version of the price grabbed from the crypto APIs.
func FormatPrice(rawPrice string, both bool) string {
	currency := "$"

	minusCents := strings.Split(rawPrice, ".")[0]

	// $15.4k
	if both {
		// contains comma
		if strings.Contains(minusCents, ",") {
			parts := strings.Split(minusCents, ",")
			// 15,433
			// ret -> $15.4k
			return currency + parts[0] + "." + strings.Split(parts[1], "")[0] + "k"
		}

		//format for thousands with no comma
		if len(minusCents) > 3 {
			parts := strings.Split(minusCents, "")
			if len(minusCents) == 4 {
				// 1655
				// ret -> $1.6k
				return currency + parts[0] + "." + parts[1] + "k"
			} else if len(minusCents) == 5 {
				// 15632
				// ret -> $15.6k
				return currency + parts[0] + parts[1] + "." + parts[2] + "k"
			}
		}
		// $15,402
	} else {
		// contains comma
		if strings.Contains(minusCents, ",") {
			// 15,433
			// ret -> $15,443
			return currency + minusCents
		}

		//format for thousands with no comma
		if len(minusCents) > 3 {
			parts := strings.Split(minusCents, "")
			if len(minusCents) == 4 {
				// 1655
				// ret -> $1,655
				return currency + parts[0] + "," + parts[1] + parts[2] + parts[3]
			} else if len(minusCents) == 5 {
				// 15632
				// ret -> $15,632
				return currency + parts[0] + parts[1] + "," + parts[2] + parts[3] + parts[4]
			}
		}
	}

	// otherwise return hundreds or tens value
	// 744
	// ret -> $744
	return currency + minusCents
}

func trendToArrow(trend string) string {
	if strings.Contains(trend, "-") {
		return "↓"
	}
	return "↑"
}

func fmtTrend(trend string) string {
	if strings.Contains(trend, "-") {
		return trend + "%"
	}
	return "+" + trend + "%"
}

// GenSSID takes prices and trends and generates the SSID of the ticker network
func GenSSID(btcp string, btct string, ethp string, etht string, both bool) string {
	if Conf.BTCEnabled && Conf.ETHEnabled {
		// return both
		return "[BTC " + trendToArrow(btct) + " " + FormatPrice(btcp, both) + "] [ETH " + trendToArrow(etht) + " " + FormatPrice(ethp, both) + "]"
	} else if Conf.BTCEnabled && !Conf.ETHEnabled {
		// return btc only
		return "(BTC) " + trendToArrow(btct) + " " + FormatPrice(btcp, both) + " [" + fmtTrend(btct) + "]"
	}

	// return eth only
	return "(ETH) " + trendToArrow(etht) + " " + FormatPrice(ethp, both) + " [" + fmtTrend(etht) + "]"
}
