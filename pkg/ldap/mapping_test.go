package ldap

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/elimity-com/scim"
	"github.com/go-ldap/ldap/v3"
)

func TestAtributeMapping(t *testing.T) {
	b := []*ldap.EntryAttribute{{Name: "sn", Values: []string{"Jim"}}}
	b = append(b, &ldap.EntryAttribute{Name: "mail", Values: []string{"Jim@ee.de"}})
	b = append(b, &ldap.EntryAttribute{Name: "dc", Values: []string{"dcBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "member", Values: []string{"memberBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "o", Values: []string{"oBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "ou", Values: []string{"ouBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "uid", Values: []string{"IDBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "cn", Values: []string{"CNBoo"}})

	var ii = &AllLdap{AllConns: map[string]*ConnsLdap{}}
	path := ""
	fileName := "ldapconfig"
	err := ii.InitConn(path, fileName)
	if err != nil {
		t.Error(err)
	}
	entry, err := AttributeMapping(b, ii.AllConns["SECOND"])

	if err != nil {
		t.Error(err)
	}
	if entry == nil {
		t.Error("")
	}
}

func TestExtractValue(t *testing.T) {
	var ii = &AllLdap{AllConns: map[string]*ConnsLdap{}}
	path := ""
	fileName := "ldapconfig"
	err := ii.InitConn(path, fileName)
	if err != nil {
		t.Error(err)
	}
	replace := "SNJOHN"
	b := []*ldap.EntryAttribute{{Name: "sn", Values: []string{replace}}}
	b = append(b, &ldap.EntryAttribute{Name: "mail", Values: []string{"Jim@ee.de"}})
	b = append(b, &ldap.EntryAttribute{Name: "dc", Values: []string{"dc=John,dc=member"}})
	b = append(b, &ldap.EntryAttribute{Name: "member", Values: []string{"memberBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "o", Values: []string{"oBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "ou", Values: []string{"ouBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "uid", Values: []string{"IDBoo"}})
	b = append(b, &ldap.EntryAttribute{Name: "cn", Values: []string{"CNBoo"}})

	strResult, _ := AttributeMapping(b, ii.AllConns["SECOND"])

	if !strings.Contains(strResult[0].Values[0], replace) && strings.Contains(strResult[0].Values[0], replace) {
		t.Error()
	}

}

func TestAttrStatValAdded(t *testing.T) {

}

func TestAttrReplaceVariable(t *testing.T) {

}

func TestAttributeMappingMem(t *testing.T) {
	filename := "FIRST"

	_, err := AttributeMappingMem(filename)

	if err != nil {
		t.Error()
	}
}

/*
func TestScimMapping(t *testing.T) {
	file := "SCIM"
	key := "mail"
	_, jsonF, err := ScimMapping(file, key)
	if err != nil {
		t.Error(err)
	}
	if jsonF[key] == nil {
		t.Error(jsonF)
	}
} */

func TestScimToLdap(t *testing.T) {
	js := `{"name":{"familyName":"Schrute","givenName":"Dwight"}}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(js), &result)
	if err != nil {
		t.Errorf("%+v", err)
	}

	file := "SCIM"
	DN := "uid=Mekhi,dc=example,dc=org"
	genScimAttr := make(scim.ResourceAttributes)
	genScimAttr["id"] = DN
	genScimAttr["userName"] = "Mekhi"
	genScimAttr["name"] = result["name"]

	/* basenameOpts := []ldap.Attribute{}
	basenameOpts = append(basenameOpts, ldap.Attribute{Type: "objectClass",
		Vals: []string{"top", "inetOrgPerson"},
	})
	basenameOpts = append(basenameOpts, ldap.Attribute{Type: "sn",
		Vals: []string{"top"},
	})
	basenameOpts = append(basenameOpts, ldap.Attribute{Type: "cn",
		Vals: []string{"top"},
	})
	*/
	ld, err := ScimToLdap(file, genScimAttr)
	if err != nil {
		t.Errorf("%v", err)
	}
	if ld == nil {
		t.Errorf("%v", ld)
	}
}

func TestMapToMap(t *testing.T) {
	file := "SCIM"
	back, err := MapToMap(file)
	if err != nil {
		t.Error(back, err)
	}
}

func TestLdapToScim(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	defer ls.Conn.Close()

	baseDN := "dc=example,dc=org"
	filterLDAP := "(&(objectclass=inetOrgPerson))"
	attributeList := []string{"*"}

	file := "SCIM"
	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	if err != nil {
		t.Error(err)
	}
	ldapE, err := LdapToScim(result.Entries, file)
	if err != nil {
		t.Error()
		t.Error(ldapE)
	}

}

func TestLdapGroupToScim(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	defer ls.Conn.Close()

	baseDN := "dc=example,dc=org"
	filterLDAP := "(&(objectclass=groupOfNames))"
	attributeList := []string{"*"}

	file := "SCIM"
	result, err := ls.SearchLDAP(baseDN, 0, filterLDAP, attributeList)
	if err != nil {
		t.Error()
	}
	/* for _, v := range result.Entries {
		fmt.Println(v)
		for _, vv := range v.Attributes {
			fmt.Println(vv)
		}

	} */

	ldapE, err := LdapGroupToScim(result.Entries, file)
	if err != nil {
		t.Error()
	}
	b, err := json.Marshal(ldapE)
	fmt.Println(string(b))
	if err != nil {
		t.Error()
	}
	fmt.Println(ldapE)

}
