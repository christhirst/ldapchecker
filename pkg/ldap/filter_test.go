package ldap

import (
	"testing"
)

func TestLdapFilter(t *testing.T) {
	filterMap := map[string]string{"cn": "Aaliyah", "objectclass": "inetOrgPerson"}
	filter := LdapFilter(filterMap)
	println(filter)

	if filter == "" {
		t.Error()
	}
}
