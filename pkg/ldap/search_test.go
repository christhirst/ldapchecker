package ldap

import (
	"fmt"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}

	baseDN := "dc=example,dc=org"
	filterLDAP := "(&(objectclass=*)(cn=Dwight))"
	attributeList := []string{"memberOf"}

	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	fmt.Println("##########")
	for _, v := range result.Entries {
		fmt.Println(v.GetAttributeValues("memberOf"))
		fmt.Println("##########")
		if strings.Contains(baseDN, v.DN) && !(baseDN == v.DN) {
			fmt.Println(v)
			t.Errorf("got %s want %+v", baseDN, v.DN)
		}
	}

	if err != nil {
		t.Errorf("got %s want %+v", err, ls)
	}
	t.Error()
}

func TestLdapSearchAttr(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	filterMap := map[string]string{"cn": "*", "objectclass": "inetOrgPerson"}

	baseDN := "dc=example,dc=org"
	attributeList := []string{"*"}
	attributteValue := "cn"
	result, err := ls.LdapSearchByAttr(baseDN, 0, LdapFilter(filterMap), attributeList, attributteValue)

	if result == nil {
		t.Error(err)
	} else {
		for _, v := range result.Attributes {
			fmt.Println(v)
		}
	}
}

func TestLdapGetGroups(t *testing.T) {
	baseDN := "dc=example,dc=org"
	filterLDAP := "(&(objectclass=inetOrgPerson))"
	attributeList := []string{"dn"}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	//baseDN := "ou=Myrl,dc=example,dc=org"

	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	for _, v := range result.Entries[:1] {
		result, err := ls.LdapGetGroups(baseDN)
		if err != nil {
			t.Errorf("got %s want %+v", err, v.DN)
		}
		for _, v := range result {
			if v == "" {
				t.Error()
			}
		}
	}

	if err != nil {
		t.Error(err)
	}

}

func TestLdapGetUsersOfGroup(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	baseDN := "ou=Myrl,dc=example,dc=org"
	result, err := ls.LdapGetUsersOfGroup(baseDN)
	fmt.Println(result)

	if err != nil {
		t.Error(err)
	}
	for _, v := range result {
		if v == "" {
			t.Error()
		}
	}
}

func TestLdapUserGroupsMap(t *testing.T) {
	baseDN := "dc=example,dc=org"
	filterLDAP := "(&(objectclass=inetOrgPerson))"
	attributeList := []string{"dn"}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}

	user := "uid=Sylvester,dc=example,dc=org"

	if err != nil {
		t.Error(err)
	}

	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	if err != nil {
		t.Errorf("got %s", err)
	}
	for _, v := range result.Entries[:1] {
		result, err := ls.LdapUserGroupMap(baseDN, user)
		if err != nil {
			t.Errorf("got %s want %+v", err, v.DN)
		}
		for _, v := range result {
			for _, vv := range v {
				if vv == "" {
					t.Error()
				}
			}
		}
	}

}

func TestDNbuilder(t *testing.T) {}
