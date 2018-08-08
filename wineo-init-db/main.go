package main

import (
	"github.com/pkg/errors"

	rtdb "gopkg.in/gorethink/gorethink.v4"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	session, err := rtdb.Connect(rtdb.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to rethinkdb"))
	}

	dbsCurs, err := rtdb.DBList().Run(session)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list of databases"))
	}
	var dbs []string
	err = dbsCurs.All(&dbs)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list of databases"))
	}

	if !stringInSlice("wineo", dbs) {
		_, err = rtdb.DBCreate("wineo").RunWrite(session)
		if err != nil {
			panic(errors.Wrap(err, "failed to create wineo database"))
		}
	}

	tabsCurs, err := rtdb.DB("wineo").TableList().Run(session)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list of tables"))
	}
	var tabs []string
	err = tabsCurs.All(&tabs)
	if err != nil {
		panic(errors.Wrap(err, "failed to get list of databases"))
	}

	if !stringInSlice("accounts", tabs) {
		_, err = rtdb.DB("wineo").TableCreate("accounts").RunWrite(session)
		if err != nil {
			panic(errors.Wrap(err, "failed to create accounts table"))
		}
	}
}
