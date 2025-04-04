package main

import (
	"database/sql"
	"fmt"
	"log"
)

func PostDataPostgres(person []Person, wallet []Wallet) {

	client := "user=postgres password=123123 dbname=bank sslmode=disable"
	db, err := sql.Open("postgres", client)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for _, p := range person {
		insert := fmt.Sprintf("insert into person(first_name, last_name, country, city, district) values ('%s', '%s', '%s', '%s', '%s')", p.FirstName, p.LastName, p.Country, p.City, p.District)
		_, err := db.Exec(insert)
		if err != nil {
			log.Fatal(err)
		}

	}

	for _, w := range wallet {
		insert := fmt.Sprintf("insert into wallet(wallet_address, wallet_id, currency, amount, amount_usd) values ('%s','%d', '%s', '%d', '%d')", w.WallerAddress, w.WalletID, w.WalletInfo.Currency, w.WalletInfo.Amount, w.WalletInfo.AmountUsd)
		_, err := db.Exec(insert)
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println("Success!")
}
