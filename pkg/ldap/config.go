package ldap

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ChangeRootDir(path string) (string, error) {
	err := os.Chdir(path)
	if err != nil {
		log.Error().Err(err).Str("Path:", path).Msg("Path does not exist")
	}

	path, err = os.Getwd()
	if err != nil {
		log.Error().Err(err).Str("Path:", path).Msg("Path does not exist")
	}
	return path, err
}

func ParseConfig(path string, fileName string) (map[string]map[string]ConnDataLdap, error) {
	SetUpLogging()

	file, err := OpenFile(path, fileName)
	fmt.Println(path, fileName)
	if err != nil {
		log.Error().Err(err).Msg("Path does not exist")
	}

	var jsonMap map[string]map[string]ConnDataLdap
	err = json.Unmarshal(file, &jsonMap)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Unmarshal file")
	}

	return jsonMap, err
}

func UniversalConfigPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("Universal Path does not exist")
	}
	concatPath := filepath.Join(path, "")
	return concatPath
}

func SetUpLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

}

/* func ReadConData(path string, fileName string) (*map[string]map[string]*connDataLdap, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	file, err := ReadFile(path, fileName)
	if err != nil {
		log.Error().Err(err).Str("File:", fileName).Msg("File does not exist")
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(file, &jsonMap)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Unmarshal file")
	}

	var conf = map[string]map[string]*connDataLdap{}
	for k, y := range jsonMap {
		var item = map[string]*connDataLdap{}
		for ki, v := range y.(map[string]interface{}) {
			item[ki] = &connDataLdap{
				Bindpassword: v.(map[string]interface{})["Bindpassword"].(string),
				Bindusername: v.(map[string]interface{})["Bindusername"].(string),
				Hostname:     v.(map[string]interface{})["Hostname"].(string),
				Port:         int(v.(map[string]interface{})["Port"].(float64)),
				Starttls:     v.(map[string]interface{})["Starttls"].(bool),
				Filter:       v.(map[string]interface{})["Filter"].(string),
				Basedn:       v.(map[string]interface{})["Basedn"].(string),
				Uid:          v.(map[string]interface{})["Uid"].(string),
				SyncMode:     v.(map[string]interface{})["SyncMode"].(string),
				Mapping:      v.(map[string]interface{})["Mapping"].(bool),
			}
		}

		conf[k] = item
	}
	return &conf, err
} */
