package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	var sources flags.MultiString
	flag.Var(&sources, "source", "One or more filesystem based sources to use to read WOF ID data, which may or may not be part of the sources to graph. This is work in progress.")

	flag.Parse()

	editor, ok := os.LookupEnv("EDITOR")

	if !ok {
		log.Fatal("Unknown editor")
	}
	
	r, err := reader.NewMultiReaderFromStrings(sources...)

	if err != nil {
		log.Fatal(err)
	}

	paths := make([]string, 0)

	for _, str_id := range flag.Args() {

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		rel_path, err := uri.Id2RelPath(id)

		if err != nil {
			log.Fatal(err)
		}

		abs_path := r.URI(rel_path)

		_, err = os.Stat(abs_path)

		if err != nil {
			log.Fatal(err)
		}

		paths = append(paths, abs_path)
	}

	cmd := exec.Command(editor, paths...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
