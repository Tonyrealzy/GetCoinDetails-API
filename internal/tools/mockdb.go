package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alexis": {
		AuthToken: "123XE4",
		Username:  "alexis",
	},
	"jason": {
		AuthToken: "879UZE",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "AZ3LE4",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alexis",
	},
	"jason": {
		Coins:    200,
		Username: "statham",
	},
	"marie": {
		Coins:    300,
		Username: "mariacaire",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}
