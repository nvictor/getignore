package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetIgnore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "Foo.gitignore") {
			fmt.Fprintln(w, "boo")
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()
	Repo = server.URL + "/"
	SaveTarget("foo")
	if _, err := os.Stat(Local); os.IsNotExist(err) {
		t.Errorf(fmt.Sprintf("expecting %s file to be created", Local))
	}
	os.Remove(Local)
	SaveTarget("bar")
	if _, err := os.Stat(Local); !os.IsNotExist(err) {
		t.Errorf(fmt.Sprintf("not expecting %s file to be created", Local))
	}
}
