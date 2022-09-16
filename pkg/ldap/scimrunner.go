package ldap

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func (c AllLdap) ScimArray(scimPort int) {
	fmt.Println(c.AllConns)

	serverCount := 0
	for i, v := range c.AllConns {
		log.Info().Msg("Starting webserver for: " + i)
		for _, vv := range v.Conns {
			go ScimServed(*vv, scimPort+serverCount)
			serverCount += 1
		}
	}
}

func (c AllLdap) ScimArrays(scimPort int) {
	fmt.Println(c.AllConns)

	serverCount := 0
	for i, v := range c.AllConns {
		log.Info().Msg("Starting webserver for: " + i)
		for _, vv := range v.Conns {
			go ScimServed(*vv, scimPort+serverCount)
			serverCount += 1
		}
	}
}
