package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"kyle" : {
		AuthToken: "123ABC",
		Username: "kyle",
	},
	"shawn" : {
		AuthToken: "AAA444",
		Username: "shawn",
	},
	"whit" : {
		AuthToken: "020122",
		Username: "whit",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"kyle" : {
		Coins: 100,
		Username: "kyle",
	},
	"shawn" : {
		Coins: 200,
		Username: "shawn",
	},
	"whit" : {
		Coins: 500,
		Username: "whit",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// sim db call
	time.Sleep(time.Millisecond*500)

	var clientData = LoginDetails{}
	clientData, okay := mockLoginDetails[username]

	if !okay {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// sim db call
	time.Sleep(time.Millisecond*750)

	var clientData = CoinDetails{}
	clientData, okay := mockCoinDetails[username]

	if !okay {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}