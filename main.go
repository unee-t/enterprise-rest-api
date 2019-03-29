// main.go

package main

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/unee-t/env"
)

func main() {
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
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), env.Protect(a.Router, a.APIAccessToken)); err != nil {
			log.WithError(err).Fatal("error listening")
		}
	}

}
