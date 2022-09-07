package ldap

import (
	"strings"
	"testing"
)

func TestChangeRootDir(t *testing.T) {
	var path string = "/working"
	rootDir, err := ChangeRootDir(path)

	if err != nil {
		t.Error(err)
	}
	if strings.Contains(rootDir, path) {
		t.Error(rootDir)
	}
}
func TestParseConfig(t *testing.T) {
	path := ""
	fileName := "ldapconfig"

	_, err := ParseConfig(path, fileName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUniversalConfigPath(t *testing.T) {
	path := "working"
	configPath := UniversalConfigPath()

	if strings.Contains(configPath, path) {
		t.Error(configPath)
	}
}

func TestSetUpLogging(t *testing.T) {
	SetUpLogging()
}

/* func TestReadConData(t *testing.T) {
	fileName := "ldapconfig"
	path := ""
	gots, _ := ReadConData(path, fileName)
	want := connDataLdap{Bindusername: "",
		Bindpassword: "",
		Port:         22,
		Starttls:     true}
	for _, got := range *gots {
		for _, e := range got {

			v := reflect.ValueOf(*e)
			typeOfS := v.Type()
			vc := reflect.ValueOf(want)
			typeOfSc := vc.Type()
			//field check
			for i := 0; i < v.NumField(); i++ {
				if reflect.TypeOf(e.Port) != reflect.TypeOf(want.Port) {
					t.Errorf("got %s want %s", typeOfS.Field(i).Name, typeOfSc.Field(i).Name)
				}
			}
			//type check
			for i := 0; i < v.NumField(); i++ {
				if reflect.TypeOf(e.Port) != reflect.TypeOf(want.Port) {
					t.Errorf("got %s want %s", typeOfS.Field(i).Type, typeOfSc.Field(i).Type)
				}
			}

		}

	}
} */
