package config

import (
	"errors"
	"os"
	"github.com/codepitaka/pitaka-web/src/config/constants"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
    databaseURL   string
}

func New() *Configuration{
	c := Configuration{}
	appmode := getENV("APPMODE")
	fmt.Println("APPMODE:",appmode)
	dbURL, e := initializeDatabaseURL(appmode)
	if e != nil {
		panic(e)
	}
	c.databaseURL = dbURL
    return &c
}

func getENV(envVariable string) (appmode string){
	appmode = os.Getenv(envVariable)
	if appmode == ""{
		log.WithField("APPMODE", appmode).Fatal("$APPMODE must be set. Forcing APPMODE to be set as `LOCAL`.")
		appmode = "LOCAL"
	}
	return
}

func initializeDatabaseURL(appmode string) (dbURL string, e error){
	switch {
    case appmode=="PRD":
        dbURL = constants.PrdServerURL
		e = nil
    case appmode=="DEV" || appmode=="RVW":
        dbURL = constants.DevServerURL
		e = nil
    case appmode=="LOCAL":
        dbURL = constants.LclServerURL
		e = nil
	default:
		dbURL = "INVALID"
		e = errors.New("Invalid AppMode.")
    }
	return
}

func (c Configuration) DatabaseURL() string{
	return c.databaseURL
}

