package ldap

/* func TestBuildAddDelete(t *testing.T) {
	var dn string = "dc=example,dc=org"
	var attrName = "cn"
	var fakeData Fakedata
	err := gofakeit.Struct(&fakeData)
	if err != nil {
		t.Error(err)
	}
	addRequ := &ldap.AddRequest{
		DN: dn,
		Attributes: []ldap.Attribute{{
			Type: attrName,
			Vals: []string{fakeData.Name}}}}

	attributes := []ldap.Attribute{{
		Type: attrName,
		Vals: []string{fakeData.Name}}}

	vv := buildAddRequest(dn, attributes, nil)

	if vv.DN != addRequ.DN {
		t.Error()
	}

	if vv.Attributes[0].Type != addRequ.Attributes[0].Type {
		t.Error()
	}
	if vv.Attributes[0].Vals[0] != addRequ.Attributes[0].Vals[0] {
		t.Error()
	}
}
*/
