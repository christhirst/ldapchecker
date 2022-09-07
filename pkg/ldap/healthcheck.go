package ldap

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func (a *AllLdap) HealthCheck() []string {
	//v.Hostname, v.Port, v.Starttls, v.Bindusername, v.Bindpassword
	var returnstring = []string{}
	for i, v := range a.AllConns {
		fmt.Println(i, v)
		condata := v.Conns[i].ConnData

		_, err := v.Conns[i].SingleCon(condata.Hostname, condata.Port, condata.Starttls, condata.Bindusername, condata.Bindpassword)

		if err != nil {
			returnstring = append(returnstring, err.Error())
			log.Error().Err(err).Str("Connection Health Test:", "Failed").Msg("")
		}
	}
	return returnstring
}
