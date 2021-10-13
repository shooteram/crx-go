package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	crx "github.com/shooteram/crx-go/crx"
)

var (
	folder string
	key    string
)

func main() {
	flag.StringVar(&folder, "pack-extension", "", "")
	flag.StringVar(&key, "pack-extension-key", "", "")
	flag.Parse()

	if folder == "" {
		fmt.Println("the location of the extension is needed to proceed")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if key == "" {
		fmt.Println()
	}

	// todo: check if folder is a viable directory
	// todo: check if key is present. if its not, check a future env var

	zip, err := crx.Zip(folder)
	if err != nil {
		panic(err)
	}
	defer zip.Close()

	file, err := os.Create(path.Clean(fmt.Sprintf("%s/../extension.crx", folder)))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pem, err := crx.LoadKey(path.Clean(key))
	if err != nil {
		panic(err)
	}

	err = crx.Write(file, pem)
	if err != nil {
		panic(err)
	}
}
