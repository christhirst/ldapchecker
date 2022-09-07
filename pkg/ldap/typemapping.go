package ldap

import (
	"github.com/go-ldap/ldap/v3"
)

func ExtractEntryAttribute(Attributes []*ldap.EntryAttribute) []ldap.Attribute {
	var aslice []ldap.Attribute
	var attr ldap.Attribute
	for _, e := range Attributes {
		attr.Type = e.Name
		attr.Vals = e.Values
		aslice = append(aslice, attr)
	}
	return aslice
}

func EntryAttrToMap(Attributes []*ldap.EntryAttribute) map[string][]string {
	m := make(map[string][]string)
	for _, e := range Attributes {
		m[e.Name] = e.Values
	}
	return m
}

func MapToEntryAttribute(FieldData map[string][]string) []ldap.Attribute {
	var attr ldap.Attribute
	var aslice []ldap.Attribute
	for key, entry := range FieldData {
		attr.Type = key
		attr.Vals = entry
		aslice = append(aslice, attr)
	}
	return aslice
}

func JSONToEntryAttribute(FieldData map[string][]string) []ldap.Attribute {
	var attr ldap.Attribute
	var aslice []ldap.Attribute
	for key, entry := range FieldData {
		attr.Type = key
		attr.Vals = entry
		aslice = append(aslice, attr)
	}
	return aslice
}

func EntryToDelRequestDN(e *ldap.Entry) *ldap.DelRequest {
	var dr = ldap.DelRequest{}
	if e == nil {
		return nil
	}
	dr.DN = e.DN
	return &dr
}
