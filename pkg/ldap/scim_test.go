package ldap

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/elimity-com/scim"
	"github.com/rs/zerolog/log"
)

func TestSetup(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}

	Setup(*ls)
	//ScimServed(ls, 8080)
}

func TestGet(t *testing.T) {
	result, err := LdapAddHelper()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", err)
	}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v %+v", ls, err)
	}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/User", nil)
		uid := result[0].GetAttributeValue("cn")
		scimResult, err := ls.Get(request, uid)

		if err != nil {
			t.Errorf("Get request to Scim: %+v %+s", ls, err)
		}
		got := scimResult.ID
		want := result[0].DN

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func TestDelete(t *testing.T) {

	result, err := LdapAddHelper()
	if err != nil {
		log.Error().Err(err).Msg("Can't read json body")
	}
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v %+v", ls, err)
	}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, "/Users", nil)

		err := ls.Delete(request, result[0].DN)

		if err != nil {
			t.Errorf("Read connection data from file: %+v", ls)
		}
		got := err

		if got != nil {
			t.Errorf("got %q, not nil", got)
		}
	})

}

func TestCreate(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}

	js := `{"name":{"familyName":"Schrute","givenName":"Dwight"}}`
	var result map[string]interface{}
	err = json.Unmarshal([]byte(js), &result)
	if err != nil {
		t.Errorf("%+v", err)
	}

	DN := "uid=Mekhiz,dc=example,dc=org"

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/Create", nil)
		genScimAttr := make(scim.ResourceAttributes)

		genScimAttr["id"] = DN
		//genScimAttr["externalId"] = "CCMekhi"
		genScimAttr["userName"] = "Mekhi"
		genScimAttr["name"] = result["name"]

		scimRes, err := ls.Create(request, genScimAttr)

		if err != nil && !strings.Contains(err.Error(), "Result Code 68") {
			t.Errorf("Read connection data from file: %+v", err)
		}
		got := scimRes
		want := genScimAttr

		gotAttr := got.Attributes

		for k, v := range gotAttr {
			w := want[k].(string)
			g := v.(string)
			if g != w {
				t.Errorf("%+v    %+v", g, w)
			}
		}

	})
}
func TestGetAll(t *testing.T) {
	ls, err := HardCon()
	if err != nil {
		t.Errorf("Read connection data from file: %+v", ls)
	}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/Groups", nil)
		params := scim.ListRequestParams{Count: 5000, StartIndex: 2}
		scimPage, err := ls.GetAll(request, params)

		for _, v := range scimPage.Resources {
			if v.ID == "" {
				t.Errorf("Read connection data from file: %+v", ls)
			}
		}
		if err != nil {
			t.Errorf("Read connection data from file: %+v", ls)
		}
		/* 	got := scimResult.ID
		want := scim.Resource{ID: DN}.ID */

		//if got != want {
		//	t.Errorf("got %q, want %q", got, want)
		//}
	})
}
