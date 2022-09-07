package ldap

import (
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/rs/zerolog/log"
)

func SyncSingle(hostBundle string, c *ConnsLdap) (respCodes, error) {
	//var neu []ldap.Attribute
	var resCo respCodes
	//var Code68 int
	var err error
	dataIN := make(map[string][]*ldap.EntryAttribute)
	dataOUT := make(map[string][]*ldap.EntryAttribute)
	//data from conf
	baseDnOut := c.Conns["OUT"].ConnData.Basedn
	filterOut := c.Conns["OUT"].ConnData.Filter
	filterIN := c.Conns["IN"].ConnData.Filter
	baseDnIn := c.Conns["IN"].ConnData.Basedn

	inData, err := c.Conns["IN"].SearchLDAP(baseDnIn, 0, filterIN, []string{"*"})
	if err != nil {
		log.Error().Err(err).Msg("Getting data with ldapsearch - in - failed in sync")
	}
	//data to map
	for _, v := range inData.Entries {
		dataIN[dnPathClean(v.DN)] = v.Attributes
	}

	outData, err := c.Conns["OUT"].SearchLDAP(baseDnOut, 0, filterOut, []string{"*"})
	if err != nil {
		log.Error().Err(err).Msg("Getting data with ldapsearch - out - failed in sync")
	}
	for _, v := range outData.Entries {
		dataOUT[dnPathClean(v.DN)] = v.Attributes
	}

	for _, outE := range outData.Entries {
		dataInAttr := dataIN[outE.DN]
		for _, outAttr := range outE.Attributes {
			for _, inAttr := range dataInAttr {
				if outAttr.Name == inAttr.Name {
					for i, att := range inAttr.Values {
						if len(outAttr.Values) != len(inAttr.Values) || outAttr.Values[i] != att {
							ldapType := outAttr.Name
							newChange := []ldap.Change{{
								Operation: 2,
								Modification: ldap.PartialAttribute{
									Type: ldapType,
									Vals: outAttr.Values}}}

							testControl := []ldap.Control{}
							testChange := &ldap.ModifyRequest{DN: outE.DN,
								Changes: newChange, Controls: testControl}

							err := c.Conns["IN"].LdapModify(testChange)

							if err != nil && !strings.Contains(err.Error(), "Result Code 32") {
								log.Error().Err(err).Msg("Getting data with ldapsearch - in - failed in sync")
							}

						}
					}
				}

			}

		}
	}

	//s := strings.Split(outE.DN, ",")
	//dn := s[0] + "," + baseDnIn
	/* ess, err := AttributeMapping(outE.Attributes, c)
	if err != nil {
		log.Error().Err(err).Msg("Attribute mapping failed")
	} */
	//neu = ExtractEntryAttribute(ess)
	//deleteEntries
	addEntries, _, err := EntrysDiff(dataOUT, dataIN, baseDnOut, baseDnIn)
	for i, v := range addEntries {
		neu := ExtractEntryAttribute(v)
		dn := i + "," + baseDnIn
		err = c.Conns["IN"].LdapAdd(buildAddRequest(dn, neu, nil))
	}
	/* var inList []string
	for _, ee := range entryDiff {
		inList = append(inList, ee.DN)
	} */
	//todo: convert both to map[string-dn][[]*ldap.EntryAttribute]
	//e is outcoming target we need
	/* if slices.Contains(inList, outE.DN) {

					if err != nil {
			log.Error().Err(err).Msg("Diff entry failed")
		}

		err = c.Conns["IN"].LdapAdd(buildAddRequest(dn, neu, nil))

		if (err != nil) && (CountString(err.Error(), "Result Code 68", Code68) > 0) {
			Code68 = CountString(err.Error(), "Result Code 68", Code68)
		} else {
			log.Error().Err(err).Msg("Adding failed")
		}


	}
	resCo.Code68 = Code68

	if c.Conns["IN"].ConnData.SyncMode == "clone" {

	} */
	//}
	return resCo, err
}

func (c *AllLdap) Sync() (respCodes, error) {
	var resCo respCodes
	var err error

	for hostBundle, ConPointer := range c.AllConns {
		if hostBundle == "SCIM" {
			continue
		}
		savedConPointer := ConPointer
		c.BgTask(hostBundle, savedConPointer)

	}
	return resCo, err
}

func CountString(source string, search string, count int) int {
	if strings.Contains(source, search) {
		count += 1
	}
	return count
}

// deleteDiff

func dnPathClean(entry string) string {
	withDn := strings.Split(entry, ",")
	return withDn[0]
}

func EntrysDiff(disiredState map[string][]*ldap.EntryAttribute, checkedState map[string][]*ldap.EntryAttribute, baseDNout string, baseDNin string) (map[string][]*ldap.EntryAttribute, map[string][]*ldap.EntryAttribute, error) {
	copyState := disiredState
	for i := range checkedState {
		delete(disiredState, i)
	}

	for i := range copyState {
		delete(checkedState, i)
	}
	//add disired
	//Delete Checked

	return disiredState, checkedState, nil
}

func (c *AllLdap) SyncStop() error {
	for _, c := range c.AllConns {
		c.Control <- true
	}

	return nil
}
