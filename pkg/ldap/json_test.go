package ldap

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

/* func TestJsonParse(t *testing.T) {
	err := JsonParse(2)
	if err != nil {
		t.Error("ss")
	}
} */

func TestReadFileToJson(t *testing.T) {
	path := ""
	file := "ldapconfig"
	m, err := ReadFileToJson(path, file)
	if err != nil {
		t.Error(err)
	}
	if len(m) == 0 {
		t.Error(m)
	}
}

func TestMapToJson(t *testing.T) {
	var cc map[string]interface{}
	file := "test"
	path := "/mappings/"
	err := MapToJson(file, &cc, path)
	if err != nil {
		log.Error().Err(err).Msg("Mapping failed")
		t.Error(err)
	}
	err = os.Remove(UniversalConfigPath() + path + file + ".json")
	if err != nil {
		t.Error(err)
	}
}
