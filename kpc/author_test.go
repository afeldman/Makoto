package kpc

import "testing"

func TestInitAuthor(t *testing.T) {
	author := InitAuthor("tester")
	if author.Email != "" || author.Name == "Tester" {
		t.Errorf("author incorrect name is %s", author.Name)
	}
}

func TestSetEmail(t *testing.T) {
	author := InitAuthor("tester")
	author.SetAuthorEmail("tester@example.com")
	if author.Email != "tester@example.com" {
		t.Errorf("email address %s is wrong", author.Email)
	}
}

func TestGetEmail(t *testing.T) {
	author := InitAuthor("tester")
	author.SetAuthorEmail("tester@example.com")
	if *author.GetAuthorEmail() != "tester@example.com" {
		t.Errorf("email address %s is wrong", author.Email)
	}
}

func TestSetName(t *testing.T) {
	author := InitAuthor("tester")
	author.SetAuthorName("not tester")
	if author.Name != "not tester" {
		t.Errorf("authors name %s did not change", author.Name)
	}
}

func TestGetName(t *testing.T) {
	author := InitAuthor("tester")
	if *author.GetAuthorName() != "tester" {
		t.Errorf("the name %s is wrong", author.Name)
	}
}
