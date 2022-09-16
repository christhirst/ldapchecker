package ldap

import (
	"fmt"
	"testing"
)

var ldapServer = &AllLdap{AllConns: map[string]*ConnsLdap{}}
var path = ""
var fileName = "ldapconfig"

func TestAvailabiltiy(t *testing.T) {
	var ldapServer = &AllLdap{AllConns: map[string]*ConnsLdap{}}
	ee, err := ldapServer.InitConn(path, fileName)
	fmt.Println(err)
	Availabiltiy(ldapServer, ee)
	fmt.Println("v")
	t.Error()
}
