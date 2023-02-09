package main

import (
	"example/db"
	"example/di"
	"example/services"
	"log"
)

func main() {
	dbCfg := &db.Config{}
	mailCfg := &services.MailConfig{}

	s, err := di.NewService(dbCfg, mailCfg)
	if err != nil {
		log.Fatalln(err)
	}
	s.UserSignUp()
}
