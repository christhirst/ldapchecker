package ldap

import "github.com/rs/zerolog/log"

//verbindung herstellbar
//login möglich
//suche möglich

func Availabiltiy() {
	var ldapServer = &AllLdap{AllConns: map[string]*ConnsLdap{}}
	path := ""
	fileName := "ldapconfig"
	err := ldapServer.InitConn(path, fileName)
	if err != nil {
		log.Error().Err(err).Msg("Can't open badgerDB")
	}
}
