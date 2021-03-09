package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/akijakya/gophercises/task/cmd"
	"github.com/akijakya/gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	// with this homedir package, we can give the db Init function the user's home path
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}