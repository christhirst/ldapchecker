package ldap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

//result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)

func LdapAddHelper() ([]*ldap.Entry, error) {
	baseDN := "dc=example,dc=org"

	attributeList := []string{"*"}
	ls, err := HardCon()
	if err != nil {
		log.Error().Err(err)
	}
	defer ls.Conn.Close()

	var fakeData Fakedata

	err = gofakeit.Struct(&fakeData)
	if err != nil {
		log.Error().Err(err)
	}
	var basenameOpts = []ldap.Attribute{
		{Type: "uid",
			Vals: []string{fakeData.Name},
		},
		{Type: "cn",
			Vals: []string{fakeData.Name},
		},
		{Type: "sn",
			Vals: []string{"3"},
		},
		{Type: "Mail",
			Vals: []string{fakeData.Name},
		},
		{Type: "userPassword",
			Vals: []string{"testpw"},
		},
		{Type: "objectClass",
			Vals: []string{"top", "inetOrgPerson"},
		},
	}

	r := ldap.AddRequest{
		DN:         "cn=" + fakeData.Name + ",dc=example,dc=org",
		Attributes: basenameOpts,
	}
	fmt.Println(r)
	uid := fakeData.Name
	err = ls.LdapAdd(&r)
	if err != nil {
		log.Error().Err(err)
	}
	filterLDAP := "(&(objectclass=inetOrgPerson)(uid=" + uid + "))"
	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	fmt.Println(result.Entries[0].Attributes)

	return result.Entries, err
}

func TestLdapAddGroup(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	defer ls.Conn.Close()
	var fakeData Fakedata
	for i := 0; i < 1; i++ {
		err = gofakeit.Struct(&fakeData)
		if err != nil {
			t.Error(err)
		}

		var basenameOpts = []ldap.Attribute{
			{Type: "cn",
				Vals: []string{fakeData.Name},
			},
			{Type: "objectClass",
				Vals: []string{"top", "groupOfUniqueNames"},
			},
			{Type: "uniqueMember",
				Vals: []string{"cn=Joy,dc=example,dc=org"},
			},
		}
		/* 	dn: cn=dbagrp,ou=groups,dc=tgs,dc=com
		objectClass: top
		objectClass: posixGroup
		gidNumber: 678 */

		r := ldap.AddRequest{
			DN:         "ou=" + fakeData.Name + ",dc=example,dc=org",
			Attributes: basenameOpts,
		}
		err = ls.LdapAdd(&r)
		if (err != nil) && (!strings.Contains(err.Error(), "Result Code 68")) {
			t.Errorf("got %s want %+v", err, r)
		}
	}

}

func TestLdapAdd(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	defer ls.Conn.Close()
	var fakeData Fakedata
	for i := 0; i < 1; i++ {
		err = gofakeit.Struct(&fakeData)
		if err != nil {
			t.Error(err)
		}
		var basenameOpts = []ldap.Attribute{
			{Type: "uid",
				Vals: []string{fakeData.Name},
			},
			{Type: "cn",
				Vals: []string{fakeData.Name},
			},
			{Type: "sn",
				Vals: []string{"3"},
			},
			{Type: "Mail",
				Vals: []string{fakeData.Name},
			},
			{Type: "userPassword",
				Vals: []string{"testpw"},
			},
			{Type: "objectClass",
				Vals: []string{"top", "inetOrgPerson"},
			},
		}
		/* 	dn: cn=dbagrp,ou=groups,dc=tgs,dc=com
		objectClass: top
		objectClass: posixGroup
		gidNumber: 678 */

		r := ldap.AddRequest{
			DN:         "cn=" + fakeData.Name + ",dc=example,dc=org",
			Attributes: basenameOpts,
		}
		err = ls.LdapAdd(&r)
		if (err != nil) && (!strings.Contains(err.Error(), "Result Code 68")) {
			t.Errorf("got %s want %+v", err, r)
		}
	}

}

func TestLdapDelete(t *testing.T) {
	result, err := LdapAddHelper()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", err)
	}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v %+v", ls, err)
	}
	defer ls.Conn.Close()
	for _, v := range result[:1] {
		err = ls.LdapDelete(v)
		if err != nil {
			t.Errorf("got %s want %+v", err, v.DN)
		}
	}
}

func TestLdapModify(t *testing.T) {
	ldapType := "uid"
	attributeList := []string{"MODIFIED"}
	result, err := LdapAddHelper()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", err)
	}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v %+v", ls, err)
	}
	defer ls.Conn.Close()
	newChange := []ldap.Change{{
		Operation: 2,
		Modification: ldap.PartialAttribute{
			Type: ldapType,
			Vals: attributeList}}}
	testControl := []ldap.Control{}

	testChange := &ldap.ModifyRequest{DN: result[0].DN,
		Changes: newChange, Controls: testControl}

	err = ls.LdapModify(testChange)
	if err != nil && !strings.Contains(err.Error(), "Result Code 32") {
		t.Errorf("Modify failed: %+v", err)
	}
}
