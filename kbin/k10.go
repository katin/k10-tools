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
	fmt.Println("Cmd equals", cmd)
	files, err := ioutil.ReadDir("./cmds")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		afile := fmt.Sprint(f.Name())
		fmt.Println("afile=", afile)
		if afile[:4] == "k10." {
			if afile[4:] == cmd {
				fmt.Println("CMD FOUND:", cmd)
			}
		}
		fmt.Println(f.Name())
		logtxt := RunBash("./cmds/" + afile)
		fmt.Println("OUTPUT:", logtxt)
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

//
func JumpToTestDir() {
	k10tools := os.Getenv("K10TOOLS")
	if len(k10tools) == 0 {
		homedir := os.Getenv("HOME")
		result := RunBash(homedir + "/.k10.env")
		fmt.Println("source cmd result:", result)
		k10tools := os.Getenv("K10TOOLS")
		if len(k10tools) == 0 {
			fmt.Println("k10-tools not installed.")
		}
	}
	fmt.Println("k10-tools path is set to " + k10tools)
	testdir := os.Getenv("K10TESTDIR")
	if len(testdir) > 0 {
		result := RunBash("cd " + testdir)
		fmt.Println(result)
	}
	os.Exit(0)
}
