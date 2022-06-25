package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"os"
	"os/exec"
	"bufio"
	"bytes"
	// "encoding/json"
	"github.com/spf13/viper"
)

var Envars interface{}
var K10_config_file string = ".k10.env"


func main() {

	LoadSettings()

	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) == 1 {
		JumpToTestDir()
	}
	cmd := os.Args[1]
	//fmt.Println("Cmd equals", cmd)
	k10tools := viper.GetString("K10TOOLS")
	k10cmdsubdir := viper.GetString("K10SUBDIR")
	fmt.Println("k10 tools location =",k10tools+k10cmdsubdir)
	files, err := ioutil.ReadDir(k10tools+k10cmdsubdir+"/cmds")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		afile := fmt.Sprint(f.Name())
		//fmt.Println("afile=", afile)
		if afile[:4] == "k10." {
			if afile[4:] == cmd {
				fmt.Println("CMD FOUND:", cmd)
				logtxt := RunBash(k10tools+k10cmdsubdir+"/cmds/" + afile)
				fmt.Println("OUTPUT:", logtxt)
			}
		}
		//fmt.Println(f.Name())
	}
}

func LoadSettings() {
	viper.SetConfigName(K10_config_file) 	// name of config file (without extension)
	viper.SetConfigType("env")					// config file type
	viper.AddConfigPath("$HOME")   			// path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	// viper.AddConfigPath(".")               // optionally look for config in the working directory

	err := viper.ReadInConfig() 				// Find and read the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		fmt.Println("Config file " + os.Getenv("HOME") + K10_config_file + " successfully loaded.")
	}

// Set undefined variables
	// viper.SetDefault("database.dbname", "test_db")

	return
}

var settings struct {
    ServerMode bool `json:"serverMode"`
    SourceDir  string `json:"sourceDir"`
    TargetDir  string `json:"targetDir"`
}

func RunBash(scriptfn string) string {
	cmd, err := exec.Command("/bin/sh", scriptfn).Output()
	//    cmd, err := exec.Command( scriptfn, parms ).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	return output
}

func RunCmd(scriptfn string, parms string) string {
	cmd, err := exec.Command(scriptfn, parms).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	return output
}

func LoadEnviron() {
	homedir := os.Getenv("HOME")
	RunBash(homedir + "/.k10.env")
}

//
func JumpToTestDir() {
	k10tools := os.Getenv("K10TOOLS")
	if len(k10tools) == 0 {
			fmt.Println("k10-tools not installed.")
			fmt.Println("Run 'source ~.k10.env' at the command line to install.")
			os.Exit(1)
	}
	fmt.Println("k10-tools path is set to " + k10tools)
	testdir := os.Getenv("K10TESTDIR")
	if len(testdir) > 0 {
		result := RunBash("cd " + testdir)
		fmt.Println(result)
	}
}
