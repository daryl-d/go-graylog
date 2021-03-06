package client_test

import (
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func TestGetLDAPSetting(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			statusCode int
			resp       string
			setting    *graylog.LDAPSetting
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			statusCode: 200,
			resp: `{
  "enabled": true,
  "system_username": "CN=admin",
  "system_password": "***",
  "ldap_uri": "ldap://ldap.example.com:389/",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "OU=user,OU=foo,DC=example,DC=com",
  "search_pattern": "(cn={0})",
  "display_name_attribute": "displayname",
  "default_group": "Reader",
  "group_mapping": {
    "foo": "Reader"
  },
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": ""
}`,
			setting: &graylog.LDAPSetting{
				Enabled:                 true,
				SystemUsername:          "admin",
				SystemPassword:          "***",
				LDAPURI:                 "ldap://ldap.example.com:389/",
				SearchBase:              "OU=user,OU=foo,DC=example,DC=com",
				SearchPattern:           "(cn={0})",
				DisplayNameAttribute:    "displayname",
				DefaultGroup:            "Reader",
				GroupMapping:            map[string]string{"foo": "Reader"},
				AdditionalDefaultGroups: set.StrSet{},
			},
			checkErr: assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Get("/api/system/ldap/settings").
				MatchType("json").Reply(d.statusCode).BodyString(d.resp)
			m, _, err := client.GetLDAPSetting()
			if err != nil {
				assert.Equal(t, d.setting, m)
			}
			d.checkErr(t, err)
		}
	}
}

func TestUpdateLDAPSetting(t *testing.T) {
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")

	if endpoint == "" {
		defer gock.Off()
		endpoint = "http://example.com/api"
		client, err := client.NewClient(endpoint, authName, authPass)
		assert.Nil(t, err)
		data := []struct {
			statusCode int
			body       string
			setting    *graylog.LDAPSetting
			checkErr   func(assert.TestingT, interface{}, ...interface{}) bool
		}{{
			statusCode: 204,
			body: `{
  "enabled": true,
  "system_username": "admin",
  "system_password": "***",
  "ldap_uri": "ldap://ldap.example.com:389/",
  "search_base": "OU=user,OU=foo,DC=example,DC=com",
  "search_pattern": "(cn={0})",
  "display_name_attribute": "displayname",
  "default_group": "Reader",
  "group_mapping": {
    "foo": "Reader"
  }
}`,
			setting: &graylog.LDAPSetting{
				Enabled:              true,
				SystemUsername:       "admin",
				SystemPassword:       "***",
				LDAPURI:              "ldap://ldap.example.com:389/",
				SearchBase:           "OU=user,OU=foo,DC=example,DC=com",
				SearchPattern:        "(cn={0})",
				DisplayNameAttribute: "displayname",
				DefaultGroup:         "Reader",
				GroupMapping:         map[string]string{"foo": "Reader"},
			},
			checkErr: assert.Nil,
		}}
		for _, d := range data {
			gock.New("http://example.com").
				Put("/api/system/ldap/settings").
				MatchType("json").BodyString(d.body).Reply(d.statusCode)
			_, err := client.UpdateLDAPSetting(d.setting)
			d.checkErr(t, err)
		}
	}
}
