package ldap

import "strings"

func ExtractGroupName(ldapGroup string) string {
	tmpGroup := strings.Split(ldapGroup, ",")[0]
	group := strings.Split(tmpGroup, "=")[1]
	return group
}
