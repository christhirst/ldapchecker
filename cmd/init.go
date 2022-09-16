package cmd

import (
	"flag"
	"fmt"

	"github.com/christhirst/ldapchecker/pkg/ldap"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	Users     []string
	PortFront string
	PortLdap  string
}

func Init(logLevel string) {
	debug := flag.Bool(logLevel, false, "sets log level to debug")
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	switch os := logLevel; os {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		log.Error().Str("Path:", os).Msg("Path does not exist")
	}
}

func Run() {

	var ldapServer = &ldap.AllLdap{AllConns: map[string]*ldap.ConnsLdap{}}
	path := ""
	fileName := "ldapconfig"
	ee, err := ldapServer.InitConn(path, fileName)

	fmt.Println(err)
	ldap.Availabiltiy(ldapServer, ee)
}
