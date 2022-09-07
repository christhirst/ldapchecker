package ldap

func LdapFilter(sm map[string]string) string {
	stringBuild := "(&"

	for i, v := range sm {
		stringBuild += "(" + i + "=" + v + ")"

	}
	stringBuild += ")"
	return stringBuild
	//filterLDAP := "(&(" + filterEntry + ")(objectclass=inetOrgPerson))"

}
