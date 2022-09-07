package ldap

import (
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func (c *ConnLdap) LdapAdd(r *ldap.AddRequest) error {
	err := c.Conn.Bind(c.ConnData.Bindusername, c.ConnData.Bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Bind failed")
	}
	err = c.Conn.Add(r)
	if (err != nil) && !strings.Contains(err.Error(), "Result Code 68") {
		log.Error().Err(err).Msg("Add data failed")
	}
	return err
}

func (c *ConnLdap) LdapDelete(e *ldap.Entry) error {
	err := c.Conn.Bind(c.ConnData.Bindusername, c.ConnData.Bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Bind failed")
	}
	err = c.Conn.Del(EntryToDelRequestDN(e))
	if err != nil {
		log.Error().Err(err).Msg("Delete data failed")
	}
	return err
}

func (c *ConnLdap) LdapModify(modifyRequest *ldap.ModifyRequest) error {
	err := c.Conn.Bind(c.ConnData.Bindusername, c.ConnData.Bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Bind failed")
	}
	err = c.Conn.Modify(modifyRequest)
	if err != nil {
		log.Error().Err(err).Msg("Modifiy data failed")
	}
	return err

}
