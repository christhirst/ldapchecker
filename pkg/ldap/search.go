package ldap

import (
	"errors"

	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

func (c *ConnLdap) SearchLDAP(baseDN string, sizeLimit int, filterLDAP string, attributeList []string) (*ldap.SearchResult, error) {
	err := c.Conn.Bind(c.ConnData.Bindusername, c.ConnData.Bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Getting data with ldapsearch failed")
	}

	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		sizeLimit,
		0,
		false,
		filterLDAP,    // The filter to apply
		attributeList, // A list attributes to retrieve
		nil,
	)

	result, err := c.Conn.Search(searchRequest)
	if err != nil {
		log.Error().Err(err).Msg("Getting data with ldapsearch failed")
	}

	return result, err
}

func (c *ConnLdap) LdapSearchByAttr(baseDN string, sizeLimit int, filterLDAP string, attributeList []string, attributeValue string) (*ldap.Entry, error) {
	searchRequest, err := c.SearchLDAP(baseDN, 0, filterLDAP, attributeList)

	for _, v := range searchRequest.Entries {
		attr := v.GetAttributeValue(attributeValue)
		if attr != "" {
			return v, err
		}

	}

	return nil, errors.New("No match found")

}

func (c *ConnLdap) LdapGetGroups(baseDN string) ([]string, error) {
	filterMap := map[string]string{"ou": "*"}
	filter := LdapFilter(filterMap)
	attributeList := []string{"groupOfNames", "groupOfUniqueNames"}
	searchRequest, err := c.SearchLDAP(baseDN, 0, filter, attributeList)
	var groupsNames []string
	for _, v := range searchRequest.Entries {
		groupsNames = append(groupsNames, v.DN)
	}

	return groupsNames, err

}

func (c *ConnLdap) LdapGetUsersOfGroup(groupDN string) ([]string, error) {
	filterMap := map[string]string{"ou": "*"}
	filter := LdapFilter(filterMap)
	attributeList := []string{"groupOfNames", "groupOfUniqueNames"}
	searchRequest, err := c.SearchLDAP(groupDN, 0, filter, attributeList)

	memberList := searchRequest.Entries[0].GetAttributeValues("member")
	if len(memberList) == 0 {
		memberList = searchRequest.Entries[0].GetAttributeValues("uniqueMember")
	}

	return memberList, err
}
func (c *ConnLdap) LdapUserGroupMap(baseDN string, user string) (map[string][]string, error) {
	groupMap := map[string][]string{}

	groups, err := c.LdapGetGroups(baseDN)

	if err != nil {
		log.Error().Err(err).Msg("Getting groups failed")
	}
	//fmt.Println(groups)
	for _, group := range groups {

		memberList, err := c.LdapGetUsersOfGroup(group)
		if err != nil {
			log.Error().Err(err).Msg("Getting groups failed")
		}
		isPresent := slices.Contains(memberList, user)
		if isPresent {
			groupMap[user] = append(groupMap[user], group)
		}

	}
	return groupMap, err
}
