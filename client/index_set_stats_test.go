package client_test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetIndexSetStats(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	iss, _, _, _, err := client.GetIndexSets(0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	if len(iss) == 0 {
		if _, err := client.CreateIndexSet(is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		// clean
		defer func(id string) {
			if _, err := client.DeleteIndexSet(id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	} else {
		is = &(iss[0])
	}

	if _, _, err := client.GetIndexSetStats(is.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.GetIndexSetStats(""); err == nil {
		t.Fatal("index set id is required")
	}
	// if _, _, err := client.GetIndexSetStats("h"); err == nil {
	// 	t.Fatal(`no index set whose id is "h"`)
	// }
}

func TestGetTotalIndexSetsStats(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is, f, err := testutil.GetIndexSet(client, server, u.String())
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(is.ID)
	}
	if _, _, err := client.GetTotalIndexSetsStats(); err != nil {
		t.Fatal(err)
	}
}
