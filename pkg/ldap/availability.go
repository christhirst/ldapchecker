package ldap

import "fmt"

//verbindung herstellbar
//login möglich
//suche möglich

func Availabiltiy(s map[string][]string) {
	for i, v := range s {
		fmt.Println(i)
		fmt.Println(v)
	}
}
