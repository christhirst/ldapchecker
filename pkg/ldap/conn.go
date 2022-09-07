package ldap

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func HardCon() (*ConnLdap, error) {
	hostname := "165.22.85.167"
	port := 389
	bindusername := "cn=admin,dc=example,dc=org"
	bindpassword := "admin"
	starttls := false
	var ls = &ConnLdap{ConnData: &ConnDataLdap{
		Hostname: hostname, Port: port, Bindusername: bindusername, Bindpassword: bindpassword,
		Starttls: starttls, Filter: "(&(objectclass=inetOrgPerson))", Basedn: "dc=example,dc=org", Uid: "uid"}}
	_, err := ls.SingleCon(hostname, port, starttls, bindusername, bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Connection can't not be established")
	}
	return ls, err
}

func (a *AllLdap) InitConn(path string, fileName string) (map[string][]string, error) {
	conerror := map[string][]string{}
	configsFromFile, err := ParseConfig(path, fileName)
	fmt.Println(configsFromFile)
	if err != nil {
		log.Error().Err(err).Msg("Failed reading connection data from file")
	}
	for k, configFromFile := range configsFromFile {
		ess, err := AttributeMappingMem(k)
		cons := &ConnsLdap{map[string]*ConnLdap{}, ess, make(chan bool)}

		conerrors := []string{}
		for i, v := range configFromFile {
			savedPointer := v
			cons.Conns[i] = &ConnLdap{ConnData: &savedPointer}
			a.AllConns[k] = cons
			//create Connection
			log.Info().Msgf("Trying connect to %s", i)
			cc, err := a.AllConns[k].Conns[i].SingleCon(v.Hostname, v.Port, v.Starttls, v.Bindusername, v.Bindpassword)
			if cc == nil {
				conerrors = append(conerrors, i)
			}
			if err != nil {
				log.Error().Err(err).Msg("Can' read connection data from file")
			}
		}
		conerror[k] = conerrors

		if err != nil {
			log.Error().Err(err).Msg("Can' read connection data from file")
		}
	}
	return conerror, err
}

func (c *ConnLdap) SingleCon(hostname string, port int, starttls bool, bindusername string, bindpassword string) (*ldap.Conn, error) {
	var err error
	c.Conn, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		log.Error().Err(err).Str(hostname, strconv.Itoa(port)).Msg("Connection Failed")
		return nil, err
	}
	// Reconnect with TLS
	if starttls == true {
		err = c.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			log.Error().Err(err).Msg("TLS failed")
			return nil, err
		}
	}
	err = c.Conn.Bind(bindusername, bindpassword)
	if err != nil {
		log.Error().Err(err).Msg("Authentication failed")
		return nil, err
	}
	return c.Conn, err
}

func AddCon() {

}

func WriteCon() {

}

/*
func GetCon(target map[string]string) (map[string]*ConnLdap, error) {
	configsFromFile, _ := ReadConData("")
	var ii = make(map[string]*ConnLdap)
	for i, v := range target {
		ii[v].ConnData.Bindusername = (*configsFromFile)[i][v].Bindusername
		ii[v].ConnData.Bindpassword = (*configsFromFile)[i][v].Bindpassword
		ii[v].ConnData.Hostname = (*configsFromFile)[i][v].Hostname
		ii[v].ConnData.Port = (*configsFromFile)[i][v].Port
		ii[v].ConnData.Starttls = (*configsFromFile)[i][v].Starttls
	}
	return ii, nil
} */

/* func PopulateConData(d *connDataLdap) ConnLdap {
	lconf := ConnLdap{Conn: &ldap.Conn{}, ConnData: &connDataLdap{}}
	lconf.ConnData.Hostname = d.Hostname
	lconf.ConnData.Bindpassword = d.Bindpassword
	lconf.ConnData.Port = d.Port
	lconf.ConnData.Bindusername = d.Bindusername
	lconf.ConnData.Starttls = d.Starttls
	return lconf
}
*/
