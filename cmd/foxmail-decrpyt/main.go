package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chuchiy/foxmail-decrypt/pkg/foxmail"
)

func main() {

	var bVersion6 bool

	flag.BoolVar(&bVersion6, "v6", false, "password from foxmail 6")
	hexpwd := flag.String("p", "", "hex encrypt password")
	flag.Parse()
	if *hexpwd == "" {
		flag.Usage()
		os.Exit(1)
	}

	pwd, err := foxmail.DecryptPassword(*hexpwd, bVersion6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hex pwd %s decrpyt error %s", *hexpwd, err)
		os.Exit(1)
	}

	fmt.Println(pwd)
}
