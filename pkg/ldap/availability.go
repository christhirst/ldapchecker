package ldap

import (
	"fmt"
)

// verbindung herstellbar
// login möglich
// suche möglich
// todo connection: methode that creates a connection instead of a pointer to a connection
func Availabiltiy(von *AllLdap, s map[string][]string) {

	for i, v := range von.AllConns {
		fmt.Println(i)
		fmt.Println("---------------------")

		fmt.Println(v.Conns["OUT"].ConnData)
		username := v.Conns["OUT"].ConnData.Bindusername
		password := v.Conns["OUT"].ConnData.Bindpassword
		err := v.Conns["OUT"].Conn.Bind(username, password)
		if err != nil {
			fmt.Println("conn does not work")
		}

		v.Conns["OUT"].Conn.Close()
	}
}
