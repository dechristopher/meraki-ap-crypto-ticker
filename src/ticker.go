package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	u "github.com/dechristopher/meraki-ap-crypto-ticker/src/util"
	cron "github.com/robfig/cron"
)

var (
	// BTCPrice is current BTC price
	BTCPrice = "0"
	// BTCTrendPCT percent change 24h
	BTCTrendPCT = ""

	// ETHPrice is current ETH price
	ETHPrice = "0"
	// ETHTrendPCT percent change 24h
	ETHTrendPCT = ""

	// BothCoins is set if both coins are enabled
	BothCoins = false
)

func main() {
	// Leggo
	fmt.Println(">> Init ticker...")

	// Read global configuration into memory from config.json
	u.Conf = u.ReadConfig()

	if !u.Conf.BTCEnabled && !u.Conf.ETHEnabled {
		u.LogErr("You must have at least one cryptocurrency enabled!")
		os.Exit(1)
		return
	}

	if u.Conf.BTCEnabled && u.Conf.ETHEnabled {
		BothCoins = true
	}

	// Okay now lets really go
	u.Log("Starting...")

	// Check for cron mode and execute then exit
	if len(os.Args) > 1 && os.Args[1] == "-cron" {
		setTicker()
		os.Exit(0)
		return
	}

	// Otherwise, we start the service to run on the specified interval
	c := cron.New()
	c.AddFunc("0 */"+strconv.Itoa(u.Conf.UpdateInterval)+" * * * *", func() { setTicker() })
	c.Start()
}

// Hits the Meraki API to set the ticker SSID
func setTicker() {
	errPull := pullPrice()

	if errPull != nil {
		return
	}

	url := "https://n6.meraki.com/api/v0/networks/" + u.Conf.NetworkID + "/ssids/" + u.Conf.SSIDNum
	u.Log("Setting ticker...")

	ssidName := u.GenSSID(BTCPrice, BTCTrendPCT, ETHPrice, ETHTrendPCT, BothCoins)
	fmt.Println("SSID: " + ssidName)

	var jsonStr = []byte(`{"name":"` + ssidName + `", "enabled": true}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Cisco-Meraki-API-Key", u.Conf.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	/*body, _ :=*/ ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	u.Log("Ticker set!")
	return
}

// Pulls the current BTC price and calculates the trend based on historical data
func pullPrice() error {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	//r, err := myClient.Get("https://blockchain.info/ticker")
	r, err := myClient.Get("https://api.coinmarketcap.com/v1/ticker/?convert=" + u.Conf.Currency + "&limit=2")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	/*body, err := ioutil.ReadAll(r.Body)
	var f interface{}
	json.Unmarshal(body, &f)
	m := f.(map[string]interface{})
	listing := m[u.Conf.Currency].(map[string]interface{})

	price := fmt.Sprintf("%v", listing["last"])
	symbol := fmt.Sprintf("%v", listing["symbol"])*/

	body, err := ioutil.ReadAll(r.Body)
	var tickresp u.TickerResponse
	errUnm := json.Unmarshal(body, &tickresp)

	if errUnm != nil {
		u.LogErr(errUnm.Error())
		return errUnm
	}

	BTCPrice = tickresp[0].PriceUSD
	BTCTrendPCT = tickresp[0].PctChange24H
	ETHPrice = tickresp[1].PriceUSD
	ETHTrendPCT = tickresp[1].PctChange24H
	return nil
}
