package ldap

import "github.com/go-ldap/ldap/v3"

func buildAddRequest(DN string, Attributes []ldap.Attribute, Controls []ldap.Control) *ldap.AddRequest {
	var allConnections = &ldap.AddRequest{DN: DN,
		Attributes: Attributes,
	}
	return allConnections
}
