package ldap

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-ldap/ldap/v3"

	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/optional"
	"github.com/elimity-com/scim/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

func (c ConnsLdap) externalID(attributes scim.ResourceAttributes) optional.String {
	if eID, ok := attributes["externalId"]; ok {
		externalID, ok := eID.(string)
		if !ok {
			return optional.String{}
		}
		return optional.NewString(externalID)
	}
	return optional.String{}
}

func (c ConnLdap) Get(r *http.Request, id string) (scim.Resource, error) {
	baseDN := c.ConnData.Basedn
	DNPrefix := c.ConnData.Uid

	filterLDAP := "(&(objectclass=inetOrgPerson)(" + DNPrefix + "=" + id + "))"

	attributeList := []string{""}
	rs := scim.Resource{}

	result, err := c.SearchLDAP(baseDN, 1, filterLDAP, attributeList)

	if err != nil {
		log.Error().Err(err).Msg("Ldap Search failed")
	}
	entry := result.Entries

	entries := len(entry)
	if entries != 1 {
		err = errors.New("Id was not specific.")
	} else if entries == 1 {
		DN := entry[0].DN
		ab := entry[0].Attributes
		oo := scim.ResourceAttributes{}
		for _, v := range ab {
			oo[v.Name] = v.Values
		}
		rs := scim.Resource{ID: DN,
			Attributes: oo}
		return rs, err
	} else if entries == 0 {
		return rs, err
	}

	return rs, err
}

func (c ConnLdap) Delete(r *http.Request, id string) error {
	entry := ldap.Entry{DN: id}
	err := c.LdapDelete(&entry)
	return err
}

func (c ConnLdap) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	file := "SCIM"
	dnI := attributes["id"]
	dn := dnI.(string)

	ld, err := ScimToLdap(file, attributes)
	if err != nil {
		log.Error().Err(err).Msg("Converting Scim attributes to ldap attributes failed")
	}
	ld = append(ld, ldap.Attribute{Type: "objectClass",
		Vals: []string{"top", "inetOrgPerson"},
	})

	rs := ldap.AddRequest{
		DN:         dn,
		Attributes: ld,
	}

	err = c.LdapAdd(&rs)
	return scim.Resource{}, err
}

func (c ConnLdap) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	//from config
	ttt, _ := url.Parse("/Groups")
	ldapFilter := "(&(objectclass=inetOrgPerson))"
	if *r.URL == *ttt {
		fmt.Println("+++")
		fmt.Println(r.URL)
		ldapFilter = "(&(objectclass=groupOfNames))"

	}
	baseDN := c.ConnData.Basedn

	file := "SCIM"
	_, err := ScimToLdapFilter(params.Filter)
	if err != nil {
		log.Error().Err(err).Msg("Scim filter to Ldap filter failed")
	}
	attributeList := []string{"*"}
	result, err := c.SearchLDAP(baseDN, params.Count, ldapFilter, attributeList)
	var z []scim.Resource
	if err != nil {
		log.Error().Err(err).Msg("No Ldap result")
	}

	if *r.URL == *ttt {
		z, err = LdapGroupToScim(result.Entries, file)
	} else {
		z, err = LdapToScim(result.Entries, file)
	}

	if err != nil {
		log.Error().Err(err).Msg("Ldap-Con does not work")
	}
	fmt.Println(params.Count)
	if params.Count == 0 {
		return scim.Page{
			TotalResults: len(result.Entries),
		}, nil
	}

	sp := scim.Page{TotalResults: len(result.Entries), Resources: z[:1]}
	fmt.Println(sp)
	return sp, nil
}

func (c ConnLdap) Patch(r *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	//c.LdapModify
	// return stored resource
	return scim.Resource{}, nil
}
func (c ConnLdap) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	//delete attributes
	//c.LdapModify

	//add attributes
	//c.LdapModify

	// replace (all) attributes
	/* h.data[id] = testData{
		resourceAttributes: attributes,
	} */

	// return resource with replaced attributes
	return scim.Resource{}, nil
}

func ScimServed(c ConnLdap, scimPort int) {

	err := http.ListenAndServe(":"+strconv.Itoa(scimPort), Setup(c))
	if err != nil {
		log.Error().Err(err).Msg("Path does not exist")
	}
}

func Setup(c ConnLdap) *chi.Mux {

	userHandler := c
	config := scim.ServiceProviderConfig{
		DocumentationURI: optional.NewString("www.example.com/scim"),
	}

	schemas := schema.Schema{
		ID:          "urn:ietf:params:scim:schemas:core:2.0:User",
		Name:        optional.NewString("User"),
		Description: optional.NewString("User Account"),
		Attributes: []schema.CoreAttribute{
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name:       "userName",
				Required:   true,
				Uniqueness: schema.AttributeUniquenessServer(),
			})),
		},
	}

	groupSchemas := schema.Schema{
		ID:          "urn:ietf:params:scim:schemas:core:2.0:Group",
		Name:        optional.NewString("Group"),
		Description: optional.NewString("User Group"),
		Attributes: []schema.CoreAttribute{
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name:       "groupName",
				Required:   true,
				Uniqueness: schema.AttributeUniquenessServer(),
			})),
		},
	}

	extension := schema.Schema{
		ID:          "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name:        optional.NewString("EnterpriseUser"),
		Description: optional.NewString("Enterprise User"),
		Attributes: []schema.CoreAttribute{
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name: "employeeNumber",
			})),
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name: "organization",
			})),
		},
	}

	resourceTypes := []scim.ResourceType{
		{
			ID:          optional.NewString("User"),
			Name:        "User",
			Endpoint:    "/Users",
			Description: optional.NewString("User Account"),
			Schema:      schemas,
			SchemaExtensions: []scim.SchemaExtension{
				{Schema: extension},
			},
			Handler: userHandler,
		},
		{
			ID:          optional.NewString("Group"),
			Name:        "Group",
			Endpoint:    "/Groups",
			Description: optional.NewString("Groups of Users"),
			Schema:      groupSchemas,
			Handler:     userHandler,
		},
	}
	mux := chi.NewRouter()
	//auth.RegisterAPI(mux)
	mux.Use(cors.Handler(cors.Options{
		//AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type: application/json", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		//AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	var se = scim.Server{
		Config:        config,
		ResourceTypes: resourceTypes}

	mux.Get("/*", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		se.ServeHTTP(rw, r)

	})

	//port := strconv.Itoa(scimPort)

	//err := http.ListenAndServe(":"+port, mux)

	return mux

}
