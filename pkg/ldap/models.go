package ldap

import (
	"time"

	"github.com/go-ldap/ldap/v3"
)

type ConnDataLdap struct {
	Hostname     string
	Port         int
	Bindusername string
	Bindpassword string
	Starttls     bool
	Filter       string
	Basedn       string
	Uid          string
	SyncMode     string
	Mapping      bool
	Frequence    int64
	SCIM         bool
}

type ConnLdap struct {
	ConnData *ConnDataLdap
	Conn     *ldap.Conn
}
type ConnsLdap struct {
	Conns   map[string]*ConnLdap
	Mapping map[string]interface{}
	Control chan bool
}

type AllLdap struct {
	AllConns map[string]*ConnsLdap
}

type respCodes struct {
	Code68 int
}

type FileConfig struct {
	Path          string
	MappingFolder string
	ConfigFile    string
}

type Fakedata struct {
	Int     int
	Name    string    `fake:"{firstname}"`   // Any available function all lowercase
	Number  string    `fake:"{number:1,10}"` // Comma separated for multiple values
	Created time.Time // Can take in a fake tag as well as a format tag
	Mail    string
}
