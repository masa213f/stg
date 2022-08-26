package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/masa213f/stg/pkg/manager"
)

var version string

var (
	debugOpt   bool
	versionOpt bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.BoolVar(&debugOpt, "debug", false, "show debug print")
	flag.BoolVar(&versionOpt, "version", false, "show version")
}

func main() {
	flag.Parse()
	if versionOpt {
		fmt.Println(version)
		os.Exit(0)
	}

	mgr := manager.New(debugOpt)
	if err := mgr.RunGame(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
}
