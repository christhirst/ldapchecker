package ldap

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func TestExtractEntryAttribute(t *testing.T) {
	var fakeData Fakedata
	var attrName = "cn"
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		t.Errorf("got %s", err)
	}
	testAttr := []*ldap.EntryAttribute{
		{
			Name:   attrName,
			Values: []string{fakeData.Name}}}
	wantedAttr := []ldap.Attribute{{
		Type: attrName,
		Vals: []string{fakeData.Name}}}
	entryAttr := ExtractEntryAttribute(testAttr)

	for i, v := range entryAttr {
		if v.Type != wantedAttr[i].Type {
			t.Error()
		}
		if v.Vals[0] != wantedAttr[i].Vals[0] {
			t.Error()
		}
	}
}

func TestEntryAttrToMap(t *testing.T) {
	var fakeData Fakedata
	var attrName = "cn"
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		t.Errorf("got %s", err)
	}

	testAttr := []*ldap.EntryAttribute{
		{Name: attrName, Values: []string{fakeData.Name}}}
	fmt.Println(testAttr[0])
	mapedEntry := EntryAttrToMap(testAttr)

	for i, v := range mapedEntry {
		if i != attrName {
			t.Error(attrName)
		}
		if v[0] != fakeData.Name {
			t.Error(fakeData.Name)
		}
	}
}
func TestMapToEntryAttribute(t *testing.T) {
	var fakeData Fakedata
	var attrName = "cn"
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		log.Error().Err(err).Msg("Can't read json body")
	}
	testMap := map[string][]string{
		attrName: {fakeData.Name}}
	entryAttr := MapToEntryAttribute(testMap)

	for _, v := range entryAttr {
		if v.Type != attrName {
			t.Error()
		}
		if v.Vals[0] != fakeData.Name {
			t.Error()
		}

	}

}
func TestJSONToEntryAttribute(t *testing.T) {
	var fakeData Fakedata
	var attrName = "cn"
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		t.Errorf("got %s", err)
	}
	testMap := map[string][]string{
		attrName: {fakeData.Name}}
	entryAttr := JSONToEntryAttribute(testMap)

	for _, v := range entryAttr {
		if v.Type != attrName {
			t.Error()
		}
		if v.Vals[0] != fakeData.Name {
			t.Error()
		}
	}
}

func TestEntryToDelRequestDN(t *testing.T) {
	testDN := "dc=example,dc=org"
	var fakeData Fakedata
	var attrName = "cn"
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		t.Errorf("got %s", err)
	}
	entryAttr := []*ldap.EntryAttribute{{Name: attrName, Values: []string{gofakeit.Name()}}}
	testLdap := &ldap.Entry{DN: testDN, Attributes: entryAttr}

	delRequest := EntryToDelRequestDN(testLdap)

	if delRequest.DN != testDN {
		t.Error()
	}
	if delRequest.Controls != nil {
		t.Error()
	}
	delRequest = EntryToDelRequestDN(nil)
	if delRequest != nil {
		t.Error()
	}
}
