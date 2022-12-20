package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// /home/dinesh/asterix2/AddressBookSync
func main() {
	var appDir string
	var configFile string
	fmt.Println(strings.Repeat("=", 10), "start of", path.Base(os.Args[0]), strings.Repeat("=", 10))
	if l := len(os.Args); l != 3 {
		log.Fatalln("Please pass the configfile apprepo for ex: gradleconfig.txt /home/dinesh/asterix2/AddressBookSync")
	} else {
		configFile = os.Args[1]
		appDir = os.Args[2]
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Fatalln("Configfile doesnt exist -> " + configFile)
		}
		if _, err := os.Stat(appDir); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Fatalln("Directory doesnt exist -> " + appDir)
		}
	}
	// println(appDir)
	buildFile := appDir + string(os.PathSeparator) + "build.gradle"
	_ = buildFile
	// println(buildFile)

	b, err := os.ReadFile(configFile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	pluggableText := string(b) // convert content to a 'string'

	fmt.Println(pluggableText) // print the content as a 'string'

	fmt.Println(strings.Repeat("=", 10), "End of", path.Base(os.Args[0]), strings.Repeat("=", 10))
}
