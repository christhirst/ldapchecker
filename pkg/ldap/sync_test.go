package ldap

import (
	"fmt"
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestSyncSingle(t *testing.T) {
	path := ""
	fileName := "ldapconfig"
	var hostBundle string = "THIRD"
	var allConnections = &AllLdap{map[string]*ConnsLdap{}}
	if err := allConnections.InitConn(path, fileName); err != nil {
		t.Error()
	}

	con := allConnections.AllConns[hostBundle]
	respCodes, err := SyncSingle(hostBundle, con)
	if err != nil {
		t.Error()
	}
	if respCodes.Code68 == 0 {
		t.Error(respCodes.Code68)
	}
}

func TestSync(t *testing.T) {
	path := ""
	fileName := "ldapconfig"

	var ii = &AllLdap{map[string]*ConnsLdap{}}

	err := ii.InitConn(path, fileName)
	if err != nil {
		t.Errorf("got %s", err)
	}
	Code68, err := ii.Sync()
	if err != nil {
		t.Errorf("got %s", err)
	}
	if Code68.Code68 < 0 {
		t.Errorf("got %+v", Code68)
	}

}

func TestCountString(t *testing.T) {
	var source string = "Count every instance and when you do, you got an E"
	var search string = "instance"
	var count int = 0

	strExists := CountString(source, search, count)
	if strExists != 1 {
		t.Error()
	}
}

func TestEntrysDiff(t *testing.T) {
	path := ""
	fileName := "ldapconfig"
	filterCheckOut := "(&(objectclass=inetOrgPerson))"
	baseDnOut := "dc=example,dc=org"
	baseDnIn := "dc=example,dc=org"

	var ii = &AllLdap{map[string]*ConnsLdap{}}

	err := ii.InitConn(path, fileName)
	if err != nil {
		t.Errorf("got %+v", err)
	}
	outData, err := ii.AllConns["FIRST"].Conns["OUT"].SearchLDAP(baseDnOut, 0, filterCheckOut, []string{"*"})
	if err != nil {
		t.Errorf("got %s", err)
	}
	inData, err := ii.AllConns["FIRST"].Conns["IN"].SearchLDAP(baseDnIn, 0, filterCheckOut, []string{"*"})
	if err != nil {
		t.Errorf("got %s", err)
	}
	dataIN := make(map[string][]*ldap.EntryAttribute)
	for _, v := range inData.Entries {
		dataIN[dnPathClean(v.DN)] = v.Attributes
	}

	dataOUT := make(map[string][]*ldap.EntryAttribute)
	for _, v := range outData.Entries {
		dataOUT[dnPathClean(v.DN)] = v.Attributes
	}

	add, delete, err := EntrysDiff(dataOUT, dataIN, baseDnOut, baseDnIn)
	fmt.Println(len(add))
	fmt.Println(len(delete))
	if err != nil {
		t.Errorf("got %+v %+v", err, add)
	}

	t.Error()
}
func TestStopSync(t *testing.T) {
	path := ""
	fileName := "ldapconfig"
	var ii = &AllLdap{map[string]*ConnsLdap{}}

	err := ii.InitConn(path, fileName)
	if err != nil {
		t.Errorf("got %+v", err)
	}
	//err = ii.SyncStop()

	if err != nil {
		t.Errorf("got %+v", err)
	}
}
