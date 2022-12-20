package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
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

	// println(buildFile)

	b, err := os.ReadFile(configFile) // just pass the file name
	if err != nil {
		log.Fatalln(err)
	}

	pluggableText := string(b) // convert content to a 'string'

	fmt.Println(pluggableText) // print the content as a 'string'

	f, err := os.OpenFile(buildFile, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln(err)
	}

	// write in next line else it throws error
	n, err := f.WriteString("\n" + pluggableText)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	_ = n

	cmd := exec.Command("bash", "-c", "source ~/devenvs/asterix2-int-devenv-1.7.0/main/env_setup.sh ; cd "+appDir+"; gradle dependencies --configuration releaseRuntimeClasspath --write-locks ; osv --L gradle.lockfile --json > vuln2.json; jq \".results | length\" vuln2.json")
	bs, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err)
	}
	cmdOutput := string(bs)
	fmt.Println(cmdOutput)
	lenth := len(cmdOutput)
	time.Sleep(2 * time.Second)

	// print(string(cmdOutput[lenth-2]))
	if !(string(cmdOutput[lenth-2]) == "0") {
		fmt.Println("Vulnerabilities found in OSV Database!")
		cmd = exec.Command("bash", "-c", "source ~/devenvs/asterix2-int-devenv-1.7.0/main/env_setup.sh ; cd "+appDir+"; jq \".results[0].packages[]|{package,vulnerabilities}|del(.vulnerabilities[].affected[].versions)| del(.vulnerabilities[].database_specific)|del(.vulnerabilities[].affected[].database_specific)|del(.vulnerabilities[].references)\" vuln2.json")
		bs, err = cmd.CombinedOutput()

		if err != nil {
			log.Fatalln(err)
		}
		cmdOutput = string(bs)
		fmt.Println(cmdOutput)

	} else {
		fmt.Println("NO Vulnerabilities found in OSV Database!")
	}

	fmt.Println(strings.Repeat("=", 10), "End of", path.Base(os.Args[0]), strings.Repeat("=", 10))
}

// gradle dependencies --write-locks
// osv --L gradle.lockfile --json > vuln2.json
//  jq ".results | length" vuln2.json
// jq ".results[0].packages[]|{package,vulnerabilities}|del(.vulnerabilities[].affected[].versions)| del(.vulnerabilities[].database_specific)|del(.vulnerabilities[].affected[].database_specific)|del(.vulnerabilities[].references)" vuln2.json
