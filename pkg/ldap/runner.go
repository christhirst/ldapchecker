package ldap

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func (c *AllLdap) BgTask(hostBundle string, v *ConnsLdap) {
	v.Control = make(chan bool)
	/* quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				// Do other stuff
			}
		}
	}() */

	// Do stuff
	//quit := make(chan bool)
	// Quit goroutine
	go SyncInBG(c, v, hostBundle, false)
}

func SyncInBG(c *AllLdap, v *ConnsLdap, hostBundle string, reload bool) (respCodes, error) {
	var err error
	var resCo respCodes
	var frequence int64
	frequence = 100

	if v.Conns["IN"].ConnData.Frequence != 0 {
		frequence = v.Conns["IN"].ConnData.Frequence
	}

	t := SetTicker(frequence)
	for range t.C {
		/* if reload == true {
			c, err = c.ReloadConf()
			if err != nil {
				log.Error().Err(err).Msg("Resetting up connection failed")
			}
		} */
		log.Info().Msg("Ticker")
		select {
		case <-v.Control:
			log.Info().Msg("Stopped sync")
			return resCo, err
		default:
			resCo, err = SyncSingle(hostBundle, v)
		}
		//v.Control <- true
	}

	return resCo, err
}

func SetTicker(sec int64) *time.Ticker {
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	return ticker
}

func (c *AllLdap) ReloadConf() (*AllLdap, error) {
	path := ""
	fileName := "ldapconfig"
	err := c.InitConn(path, fileName)

	return c, err
}

func (c *AllLdap) StartSync() {
	path := ""
	fileName := "ldapconfig"
	fmt.Println(c.AllConns)
	err := c.InitConn(path, fileName)
	if err != nil {
		log.Error().Err(err)
	}
	fmt.Println(c.AllConns)
	fmt.Println(path)
	fmt.Println(c.AllConns["FIRST"].Conns["IN"].ConnData)
	//go ldap.Setup()
	_, err = c.Sync()
	if err != nil {
		log.Error().Err(err)
	}

}
