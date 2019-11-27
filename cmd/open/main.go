package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// cribbed from https://github.com/tj/go-editor/blob/master/editor.go

func Open(editor string, path string) error {

	cmd := exec.Command("sh", "-c", editor+" "+path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {

	editor := flag.String("editor", "", "The editor to open a WOF ID with. If empty the value of the 'EDITOR' environment variable with be used.")
	root := flag.String("root", "", "The path where the path for each WOF ID lives. If empty the current directory is used.")
	is_repo := flag.Bool("repo", false, "Indicates that -root is a whosonfirst style repo and appends a 'data' folder to its path.")

	flag.Parse()

	if *editor == "" {
		*editor = os.Getenv("EDITOR")
	}
	
	if *editor == "" {
		log.Fatal("Missing EDITOR environment variable")
	}

	if *root == "" {

		cwd, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		*root = cwd
	}

	if *is_repo {
		*root = filepath.Join(*root, "data")
	}

	for _, str_id := range flag.Args() {

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		rel_path, err := uri.Id2RelPath(id)

		if err != nil {
			log.Fatal(err)
		}

		abs_path := filepath.Join(*root, rel_path)

		err = Open(*editor, abs_path)

		if err != nil {
			log.Printf("Failed to open '%s', %v\n", abs_path, err)
		}
	}
}
