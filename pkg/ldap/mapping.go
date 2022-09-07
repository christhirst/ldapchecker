package ldap

import (
	"fmt"
	"regexp"

	"github.com/elimity-com/scim"
	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
	"github.com/scim2/filter-parser/v2"
)

func AttributeMapping(ea []*ldap.EntryAttribute, c *ConnsLdap) ([]*ldap.EntryAttribute, error) {
	b := []*ldap.EntryAttribute{}
	var err error
	for i, v := range c.Mapping {
		var outstring []string
		for _, l := range v.([]interface{}) {
			outstring = append(outstring, l.(string))
		}
		for _, vv := range ea {
			re := regexp.MustCompile(`{{` + vv.Name + `}}`)
			for _, vvv := range outstring {
				if len(re.FindAllString(vvv, 3)) > 0 {
					outstring = nil
					for _, vvvv := range vv.Values {
						bb := re.ReplaceAllString(vvv, vvvv)
						outstring = append(outstring, bb)
					}
					/*
						outstring[i] = bb */

				}
			}
		}
		b = append(b, &ldap.EntryAttribute{Name: i, Values: outstring})
	}
	return b, err
}

func AttrStatValAdded(str string) {
	re := regexp.MustCompile(`r.t`)
	fmt.Println(re.ReplaceAllString(str, "ram"))
	// prints "ram cat ram dog"

}

func AttrReplaceVariable() {

}

func AttributeMappingMem(filename string) (map[string]interface{}, error) {
	path := "mappings/"

	fileMI, err := ReadFileToJson(path, filename)
	if err != nil {
		log.Error().Err(err).Msg("ReadFileToJson failed")
	}
	return fileMI, err
}

func ScimToLdap(filename string, scimMap scim.ResourceAttributes) ([]ldap.Attribute, error) {
	path := "mappings/"
	fileMI, err := ReadFileToJson(path, filename)
	if err != nil {
		log.Error().Err(err).Msg("ReadFileToJson failed")
	}

	basenameOpts := []ldap.Attribute{}
	for k, v := range scimMap {
		switch x := v.(type) {
		case string:
			basenameOpts = append(basenameOpts, ldap.Attribute{Type: fileMI[k].(string),
				Vals: []string{v.(string)},
			})
		case map[string]interface{}:
			for kk, vv := range v.(map[string]interface{}) {
				basenameOpts = append(basenameOpts, ldap.Attribute{Type: fileMI[k].(map[string]interface{})[kk].(string),
					Vals: []string{vv.(string)}},
				)
			}
		default:
			log.Info().Err(err).Interface("v", x).Msg("Interface")
		}
	}
	return basenameOpts, nil
}

func MapToMap(filename string) (map[string]interface{}, error) {
	path := "mappings/"
	fileMI, err := ReadFileToJson(path, filename)
	if err != nil {
		log.Error().Err(err).Msg("ReadFileToJson failed")
	}
	switchedMap := make(map[string]interface{})
	switchedMaps, _ := Mapswitch(switchedMap, fileMI)
	return switchedMaps, err

}
func Mapswitch(m map[string]interface{}, in map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	for k, value := range in {
		switch v := value.(type) {
		case string:
			(m)[value.(string)] = k
			delete(in, k)
			if len(in) < 1 {
				return m, in
			}
			return Mapswitch(m, in)
		case map[string]interface{}:
			for kk, vv := range value.(map[string]interface{}) {
				(m)[vv.(string)] = map[string]string{k: kk}
				delete(in[k].(map[string]interface{}), kk)
				return Mapswitch(m, in)
			}
		default:
			log.Info().Str("", "").Msg(v.(string))
		}
	}
	fmt.Println(m)
	return m, nil
}

func ScimToLdapFilter(scimfilter filter.Expression) (string, error) {

	return "basenameOpts", nil
}

func LdapToScim(ldapEntries []*ldap.Entry, file string) ([]scim.Resource, error) {
	var scimEntries []scim.Resource

	mapping, err := MapToMap(file)
	if err != nil {
		log.Error().Err(err).Msg("ReadFileToJson failed")
	}
	//.Entries[0].GetAttributeValues("member")
	for _, v := range ldapEntries {
		var scimEntry = scim.Resource{ID: v.DN}
		foods := map[string]interface{}{}
		nestedmap := make(map[string]string)
		for _, vv := range v.Attributes {
			switch vvv := mapping[vv.Name].(type) {
			case string:
				foods[mapping[vv.Name].(string)] = vv.Values[0]
			case map[string]string:
				var h string
				for s, g := range vvv {
					nestedmap[g] = vv.Values[0]
					h = s
				}
				foods[h] = nestedmap

			case nil:
			default:
				log.Info().Err(err).Interface("vvv", vvv).Msg("Interface")
			}
		}
		scimEntry.Attributes = foods
		scimEntries = append(scimEntries, scimEntry)
	}
	return scimEntries, nil
}

func LdapGroupToScim(ldapEntries []*ldap.Entry, file string) ([]scim.Resource, error) {
	var scimEntries []scim.Resource
	var scimEntry = scim.Resource{}

	for _, v := range ldapEntries {
		members := map[string]interface{}{}
		var gg []interface{}
		for _, i := range v.GetAttributeValues("member") {
			gg = append(gg, map[string]string{"value": i, "display": "eee"})
		}
		members["members"] = gg
		members["displayName"] = "displayName_test"
		members["type"] = "Profile"
		scimEntry.ID = v.DN
		scimEntry.Attributes = members
		scimEntries = append(scimEntries, scimEntry)
	}
	return scimEntries, nil
}

/* func ScimMapping(filename string, lookup string) (string, map[string]interface{}, error) {
	path := "mappings"
	fileMI, err := ReadFileToJson(path, filename)
	fmt.Println(fileMI)
	if err != nil {
		log.Error().Err(err).Msg("ReadFileToJson failed")
	}

	value := fileMI[lookup]

	switch v := value.(type) {
	case string:
		return value.(string), nil, err
	case int:
		log.Info().Err(err).Int("v", v).Msg("Int")
	case float32:
		log.Info().Err(err).Float32("v", v).Msg("Float32")
	case float64:
		log.Info().Err(err).Float64("v", v).Msg("Float64")
	case map[string]interface{}:
		return "nil", value.(map[string]interface{}), err
	default:
		log.Info().Err(err).Str("", "").Msg("")
	}

	return "nil", nil, err
} */
