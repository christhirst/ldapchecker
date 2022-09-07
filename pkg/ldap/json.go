package ldap

import (
	"encoding/json"
	"os"

	_ "github.com/christhirst/ldapchecker/pkg/testing_init"
	"github.com/rs/zerolog/log"
)

func ReadFileToJson(path string, filename string) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	log.Info().Str(path, filename).Msg("path to file")
	byteValue, err := OpenFile(path, filename)
	if err != nil {
		log.Error().Err(err).Msg("Opening file failed")
	}
	err = json.Unmarshal(byteValue, &jsonMap)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal failed")
	}
	return jsonMap, err
}

func MapToJson(filename string, jsonMap *map[string]interface{}, pathMapping string) error {

	file, err := json.MarshalIndent(jsonMap, "", "   ")
	if err != nil {
		log.Error().Err(err).Msg("MarshalIndent failed")
	}

	err = os.WriteFile(UniversalConfigPath()+pathMapping+filename+".json", file, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Writing file failed")
	}
	return err
}

/*
func JsonParse(depth int) error {
	os.Chdir("")
	path, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("Path does not exist")
	}
	f, err := ioutil.ReadFile(path + "/ldapconfig")

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(f, &jsonMap)
	if err != nil {
		log.Error().Err(err).Str(vv.(map[string]interface{})["Hostname"].(string), "failed").Msg("Unmarshal failed")
	}
		for _, v := range jsonMap {
		for _, vv := range v.(map[string]interface{}) {


				for x, y := range jsonMap {
					var ss = connDataLdap{}
					for _, v := range y.(map[string]interface{}) {
						ss.Bindpassword = v.(map[string]interface{})["Bindpassword"].(string)
						ss.Bindusername = v.(map[string]interface{})["Bindusername"].(string)
						ss.Hostname = v.(map[string]interface{})["Hostname"].(string)
						ss.Port = int(v.(map[string]interface{})["Port"].(float64))
						ss.Starttls, _ = (v.(map[string]interface{})["Starttls"].(bool))
					}
					conf[x] = &ss
				}
				return &conf, err


	}
	return err
}
}*/
