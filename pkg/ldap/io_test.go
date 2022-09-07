package ldap

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	path := ""
	filename := "ldapconfig"
	_, err := OpenFile(path, filename)
	if err != nil {
		t.Error(err)
	}
}

func TestMapToAddRequest(t *testing.T) {
	attribute := "mail"
	FieldData := map[string][]string{attribute: {"test@mail.com", "test2@mail.com"}}

	ldapAttribute := MapToAddRequest(FieldData)
	for _, v := range ldapAttribute {
		if v.Type != attribute {
			t.Error("String map is not able to convert to ldapAttributes")
		}
	}
}
