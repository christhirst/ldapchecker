package ldap

import (
	"testing"
)

func TestScimArray(t *testing.T) {
	path := ""
	fileName := "ldapconfig"
	scimPort := 8380

	var ldapServer = AllLdap{AllConns: map[string]*ConnsLdap{}}
	_, err := ldapServer.InitConn(path, fileName)
	if err != nil {
		t.Error()
	}
	ldapServer.ScimArray(scimPort)
	//time.Sleep(20 * time.Second)

	//t.Error()
}
