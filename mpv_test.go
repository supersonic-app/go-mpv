package mpv

import (
	"testing"
)

func TestCreate(t *testing.T) {
	mpv := Create()
	if mpv == nil {
		t.Fatal("Create() returned nil")
	}
}

func TestClientName(t *testing.T) {
	mpv := Create()
	if err := mpv.Initialize(); err != nil {
		t.Fatalf("Initialize() error: %v", err)
	}
	name := mpv.ClientName()
	if name == "" {
		t.Fatal("ClientName() returned an empty string")
	}
}
