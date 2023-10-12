package kpc

import "testing"

func TestInitConflict(t *testing.T) {
	conflict := InitConflict("math", "0.1.2")
	if *conflict.GetConflictName() != "math" && !conflict.ContainsVersion("0.1.2") {
		t.Errorf("name %s and version 0.1.2 not found", *conflict.GetConflictName())
	}
}

func TestAddVersion(t *testing.T) {
	conflict := InitConflict("math", "0.1.2")
	conflict.AddVersion("0.2")
	conflict.AddVersion("0.1.2")

	if !conflict.ContainsVersion("0.1.2") && !conflict.ContainsVersion("0.2") && conflict.NumberOfConflicts() == 2 {
		t.Error("cannot add versions to versions")
	}
}

func TestChangeVersion(t *testing.T) {
	conflict := InitConflict("math", "0.1.2")
	conflict.ChangeVersion("0.1.2", "0.2.1")
	if !conflict.ContainsVersion("0.2.1") && conflict.NumberOfConflicts() == 1 {
		t.Error("cannot change version")
	}
}

func TestGetConflictName(t *testing.T) {
	conflict := InitConflict("math", "0.1.2")
	conflict.SetName("test")
	if *conflict.GetConflictName() != "test" {
		t.Error("cannot change name")
	}
}
