package main

import (
	"flag"

	"github.com/a-skua/example-wasi/read-web"
	"github.com/a-skua/example-wasi/read-web/internal/gen/wasi/cli/run"
)

var (
	url = flag.String("url", "", "Read URL")
)

func init() {
	flag.Parse()
	read.URL = *url
	run.Exports.Run = read.Run
}

func main() {}
