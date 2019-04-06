package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	Local = ".gitignore"
	Mode  = 0644
)

var Repo = "https://raw.githubusercontent.com/github/gitignore/master/"

// SaveTarget saves the file resource located at Repo + strings.Title(target) + ".gitignore"
// to local .gitignore file.
func SaveTarget(target string) {
	var (
		content  []byte
		response *http.Response
		err      error
		url      = fmt.Sprintf("%s%s.gitignore", Repo, strings.Title(target))
	)
	response, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Printf("Sorry no .gitignore found for '%s'. Maybe create one?\n", target)
		return
	}
	content, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(Local, content, Mode); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(fmt.Sprintf("Usage: %s target (ex: getignore python)", path.Base(os.Args[0])))
		return
	}
	SaveTarget(os.Args[1])
}
