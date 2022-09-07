package ldap

import (
	"io"
	"os"

	_ "github.com/christhirst/ldapchecker/pkg/testing_init"
	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func OpenFile(path string, filename string) ([]byte, error) {
	//rootPath := UniversalConfigPath()
	/* err := os.Chdir(path)
	if err != nil {
		log.Error().Err(err).Msg("Path does not exist")
		panic(err)
	} */

	file, err := os.Open(path + filename + ".json")
	if err != nil {
		log.Error().Err(err).Msg("Path does not exist")
		panic(err)
	}

	defer file.Close()
	log.Info().Msg("Successfully Opened: " + filename)
	byteValue, _ := io.ReadAll(file)
	return byteValue, err
}

func MapToAddRequest(FieldData map[string][]string) []ldap.Attribute {
	var attr ldap.Attribute
	var aslice []ldap.Attribute
	for key, entry := range FieldData {
		attr.Type = key
		attr.Vals = entry
		aslice = append(aslice, attr)
	}
	return aslice
}
