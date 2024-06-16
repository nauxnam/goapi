package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"hasan": {
		AuthToken: "123ABC",
		Username:  "hasan",
	},
	"iso": {
		AuthToken: "456DEF",
		Username:  "iso",
	},
	"beyret": {
		AuthToken: "789GHI",
		Username:  "beyret",
	},
	"batu": {
		AuthToken: "nigga",
		Username:  "batu",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"hasan": {
		Coins:    100,
		Username: "hasan",
	},
	"iso": {
		Coins:    10000,
		Username: "iso",
	},
	"beyret": {
		Coins:    31,
		Username: "beyret",
	},
	"batu": {
		Coins:    0,
		Username: "batu",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
