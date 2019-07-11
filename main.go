// main.go

package main

import (
	"net/http"
	"os"

	jsonlog "github.com/apex/log/handlers/json"
	textlog "github.com/apex/log/handlers/json"

	"github.com/apex/log"
)

func main() {
	if os.Getenv("UP_STAGE") != "" {
		log.SetHandler(jsonlog.Default)
	} else {
		log.SetHandler(textlog.Default)
	}
	a, err := New()
	if err != nil {
		log.WithError(err).Fatal("error setting configuration")
	}
	defer a.DB.Close()
	if os.Getenv("UP_STAGE") == "" {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), a.Router); err != nil {
			log.WithError(err).Fatal("error listening")
		}

	} else {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), a.Router); err != nil {
			log.WithError(err).Fatal("error listening")
		}
	}

}
