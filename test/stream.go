package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetEnabledStreams(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	streams, total, _, err := client.GetEnabledStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 1 {
		t.Fatalf("total == %d, wanted %d", total, 1)
	}

	stream := streams[0]

	stream.Disabled = true
	if _, err := server.UpdateStream(&stream); err != nil {
		t.Fatal(err)
	}
	streams, total, _, err = client.GetEnabledStreams()
	if err != nil {
		t.Fatal("Failed to GetStreams", err)
	}
	if total != 0 {
		t.Fatalf("total == %d, wanted %d", total, 0)
	}
}

func TestUpdateStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	stream.Description = "changed!"
	if _, err := client.UpdateStream(stream); err != nil {
		t.Fatal("Failed to UpdateStream", err)
	}
	stream.ID = ""
	if _, err := client.UpdateStream(stream); err == nil {
		t.Fatal("id is required")
	}
	stream.ID = "h"
	if _, err := client.UpdateStream(stream); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	if _, err := client.UpdateStream(nil); err == nil {
		t.Fatal("stream is nil")
	}
}

func TestPauseStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = client.PauseStream(stream.ID); err != nil {
		t.Fatal("Failed to PauseStream", err)
	}
	if _, err := client.PauseStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, err := client.PauseStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test pause
}

func TestResumeStream(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	_, stream, err := addDummyStream(server)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = client.ResumeStream(stream.ID); err != nil {
		t.Fatal("Failed to ResumeStream", err)
	}

	if _, err = client.ResumeStream(""); err == nil {
		t.Fatal("id is required")
	}

	if _, err = client.ResumeStream("h"); err == nil {
		t.Fatal(`no stream whose id is "h"`)
	}
	// TODO test resume
}
