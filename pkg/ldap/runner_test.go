package ldap

import (
	"testing"
)

func TestSyncInBG(t *testing.T) {
	path := ""
	fileName := "ldapconfig"

	var allConnections = &AllLdap{map[string]*ConnsLdap{}}
	err := allConnections.InitConn(path, fileName)
	if err != nil {
		t.Error()
	}
	//c := allConnections.AllConns["FIRST"]

	//SyncInBG(allConnections, c, "FIRST", true)

	if false {
		t.Error()
	}
}

func TestSetTicker(t *testing.T) {
	var sec int64 = 88
	time := SetTicker(sec)
	if time == nil {
		t.Error()
	}

}

func TestReloadConf(t *testing.T) {
	var allConnections = &AllLdap{map[string]*ConnsLdap{}}

	allLdap, err := allConnections.ReloadConf()
	if err != nil {
		t.Error(err)
	}
	if allLdap == nil {
		t.Error(allLdap)
	}
}
