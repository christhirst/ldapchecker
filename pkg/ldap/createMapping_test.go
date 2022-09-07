package ldap

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCreateFolder(t *testing.T) {
	path := "./testFolder"
	err := CreateFolder(path)
	if err == nil {
		t.Error()
	}

	err = CreateFolder(path)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove(path)
	if err != nil {
		t.Error()
	}
}
func TestLdifToJson(t *testing.T) {
	path := ""
	fileName := "ldapconfig"
	syncMode := "clone"
	var allConnections = &AllLdap{map[string]*ConnsLdap{}}

	err := allConnections.InitConn(path, fileName)
	if err != nil {
		log.Error().Err(err).Msg("TLS failed")
	}

	baseDNin := "dc=example,dc=example,dc=org"
	filterOut := "(&(objectclass=*))"
	err = allConnections.LdifToJson(filterOut, baseDNin, syncMode)
	if err != nil {
		log.Error().Err(err).Msg("TLS failed")
	}
}
