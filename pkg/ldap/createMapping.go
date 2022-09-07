package ldap

import (
	"os"
	"path/filepath"

	_ "github.com/christhirst/ldapchecker/pkg/testing_init"
	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func CreateFolder(path string) error {
	folderInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal().Err(err).Msg("Path does not exist")
		}
	} else {
		log.Debug().Str(folderInfo.Name(), "0").Msg("")
	}

	return err
}

func (c *AllLdap) LdifToJson(filterOut string, baseDnIn string, syncMode string) error {
	jsonMap := make(map[string]interface{})
	err := CreateFolder("mappings/")
	if err != nil {
		log.Error().Err(err).Msg("Reading file failed")
	}

	path := "mappings"
	rootPath, err := os.Getwd()

	concatPath := filepath.Join(rootPath, path)
	var neu []ldap.Attribute

	//interate over hostBundle
	for hostBundle, ldapConns := range c.AllConns {
		if hostBundle == "SCIM" {
			continue
		}

		//read mapping file
		fileData, err := ReadFileToJson(concatPath+"/", hostBundle)
		if err != nil {
			log.Error().Err(err).Msg("Reading file failed")
		}
		//interate over Connections
		for _, ldapCon := range ldapConns.Conns {
			if ldapCon.ConnData.Mapping {
				//load data from server
				outData, err := ldapCon.SearchLDAP(ldapCon.ConnData.Basedn, 0, filterOut, []string{"*"})
				if err != nil {
					log.Error().Err(err).Interface("ConnData", ldapCon.ConnData).Msg("")
				}
				//interate over Connections Entries
				for _, entry := range outData.Entries {
					neu = ExtractEntryAttribute(entry.Attributes)
					for _, r := range neu {
						if r.Type != "objectClass" {
							if fileData[r.Type] != nil {
								jsonMap[r.Type] = fileData[r.Type]
							} else {
								jsonMap[r.Type] = r.Type
							}
						}
					}
				}
				err = MapToJson(hostBundle, &jsonMap, path)
				if err != nil {
					log.Error().Err(err).Msg("TLS failed")
				}
			}
		}
	}

	return err
}
