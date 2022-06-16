package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) == 1 {
		JumpToTestDir()
	}
	cmd := os.Args[1]
	//fmt.Println("Cmd equals", cmd)
	k10tools := os.Getenv("K10TOOLS")
	k10cmdsubdir := os.Getenv("K10SUBDIR")
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
