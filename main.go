package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/dineshr93/gradleosvconfig/model"
)

// /home/dinesh/asterix2/AddressBookSync
func main() {
	var appDir string
	var configFile string
	var sourcingscript string
	var lengthofargs int
	fmt.Println(strings.Repeat("=", 10), "start of", path.Base(os.Args[0]), strings.Repeat("=", 10))
	if lengthofargs = len(os.Args); !(lengthofargs >= 3 && lengthofargs <= 4) {
		fmt.Println("No of arguments given:", lengthofargs)
		fmt.Println(os.Args)
		log.Fatalln("Please pass the configfile apprepo for ex: gradleconfig.txt FULL_PATH_GRADLE_APP sourcing_fullpath_scriptsh")
	} else {
		fmt.Println("No of arguments given:", lengthofargs)
		configFile = os.Args[1]
		appDir = os.Args[2]

		if lengthofargs == 4 {
			sourcingscript = os.Args[3]
		}
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Fatalln("Configfile doesnt exist -> " + configFile)
		}
		if _, err := os.Stat(appDir); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Fatalln("Directory doesnt exist -> " + appDir)
		}

		if lengthofargs == 4 {
			if _, err := os.Stat(sourcingscript); os.IsNotExist(err) {
				// path/to/whatever does not exist
				log.Fatalln("Configfile doesnt exist -> " + configFile)
			}
		}
	}
	// println(appDir)
	buildFile := appDir + string(os.PathSeparator) + "build.gradle"

	// println(buildFile)

	c, err := os.ReadFile(configFile) // just pass the file name
	if err != nil {
		log.Fatalln(err)
	}

	pluggableText := string(c) // convert content to a 'string'

	fmt.Println(pluggableText) // print the content as a 'string'
	// read the whole file at once
	b, err := os.ReadFile(buildFile)
	if err != nil {
		panic(err)
	}
	s := string(b)
	// //check whether s contains substring text
	if !strings.Contains(s, pluggableText) {

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
	}

	var cmd *exec.Cmd
	if lengthofargs == 3 {
		cmd = exec.Command("bash", "-c", "cd "+appDir+"; gradle dependencies --configuration releaseRuntimeClasspath --write-locks ; osv --L gradle.lockfile --json > vuln2.json; jq \".results | length\" vuln2.json")
	} else {
		cmd = exec.Command("bash", "-c", "source "+sourcingscript+" ; cd "+appDir+"; gradle dependencies --configuration releaseRuntimeClasspath --write-locks ; osv --L gradle.lockfile --json > vuln2.json; jq \".results | length\" vuln2.json")
	}
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
		if lengthofargs == 3 {
			cmd = exec.Command("bash", "-c", "cd "+appDir+"; jq \".results[0].packages[]|{package,vulnerabilities}|del(.vulnerabilities[].affected[].versions)| del(.vulnerabilities[].database_specific)|del(.vulnerabilities[].affected[].database_specific)|del(.vulnerabilities[].references)\" vuln2.json")
		} else {
			cmd = exec.Command("bash", "-c", "source "+sourcingscript+" ; cd "+appDir+"; jq \".results[0].packages[]|{package,vulnerabilities}|del(.vulnerabilities[].affected[].versions)| del(.vulnerabilities[].database_specific)|del(.vulnerabilities[].affected[].database_specific)|del(.vulnerabilities[].references)\" vuln2.json")
		}

		bs, err = cmd.CombinedOutput()

		if err != nil {
			log.Fatalln(err)
		}
		cmdOutput = string(bs)

		s := &model.OSVData{}
		if err := s.Load(filepath.Join(appDir, "vuln2.json")); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		s.PrintVuls()
		e := os.Remove(filepath.Join(appDir, "gradle.lockfile"))
		if e != nil {
			log.Fatal(e)
		}
		e = os.Remove(filepath.Join(appDir, "settings-gradle.lockfile"))
		if e != nil {
			log.Fatal(e)
		}
		e = os.Remove(filepath.Join(appDir, "vuln2.json"))
		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(strings.Repeat("=", 10), "End of", path.Base(os.Args[0]), strings.Repeat("=", 10))
		// log.Fatalln(cmdOutput)
		os.Exit(1)

	} else {
		fmt.Println("NO Vulnerabilities found in OSV Database!")
		e := os.Remove(filepath.Join(appDir, "vuln2.json"))
		if e != nil {
			log.Fatal(e)
		}
	}
	// _ = exec.Command("bash", "-c", "cd "+appDir+"; git restore build.gradle")
	e := os.Remove(filepath.Join(appDir, "gradle.lockfile"))
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(filepath.Join(appDir, "settings-gradle.lockfile"))
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(strings.Repeat("=", 10), "End of", path.Base(os.Args[0]), strings.Repeat("=", 10))
}

// gradle dependencies --write-locks
// osv --L gradle.lockfile --json > vuln2.json
//  jq ".results | length" vuln2.json
// jq ".results[0].packages[]|{package,vulnerabilities}|del(.vulnerabilities[].affected[].versions)| del(.vulnerabilities[].database_specific)|del(.vulnerabilities[].affected[].database_specific)|del(.vulnerabilities[].references)" vuln2.json
