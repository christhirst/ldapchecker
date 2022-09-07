package ldap

import (
	"testing"
)

func TestInitConn(t *testing.T) {
	path := ""
	fileName := "ldapconfig"

	var allConnections = &AllLdap{map[string]*ConnsLdap{}}
	err := allConnections.InitConn(path, fileName)
	if err != nil {
		t.Errorf(err.Error(), "Initialize failed")
	}
	for _, v := range allConnections.AllConns {
		for _, vv := range v.Conns {
			if vv == nil {
				t.Errorf(vv.ConnData.Bindpassword)

				t.Errorf("Read connection data from file: %+v", v.Conns)
			}
			defer vv.Conn.Close()
		}
	}

}

func TestSingleCon(t *testing.T) {
	hostname := "165.22.85.167"
	port := 388
	bindusername := "cn=admin,dc=example,dc=org"
	bindpassword := "admin"
	starttls := false
	var jj = &ConnLdap{}
	ls, err := jj.SingleCon(hostname, port, starttls, bindusername, bindpassword)
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}
	defer ls.Close()
}

func TestAddCon(t *testing.T) {
	/* var ll = []string{"FIRST", "SECOND"}

	oo := map[string]*connLdap{
		"first": &connLdap{connData: &connDataLdap{Bindusername: "cn=admin,dc=example,dc=org",
			Bindpassword: "admins",
			Hostname:     "165.22.85.167",
			Port:         389,
			Starttls:     false}},
	} */
}

func TestDeleteCon(t *testing.T) {
	/* var ll = []string{"FIRST", "SECOND"}

	oo := map[string]*connLdap{
		"first": &connLdap{connData: &connDataLdap{Bindusername: "cn=admin,dc=example,dc=org",
			Bindpassword: "admins",
			Hostname:     "165.22.85.167",
			Port:         389,
			Starttls:     false}},
	} */
}
