package main

import (
	"embed"
	"fmt"
	"os"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/sys/windows"
)

//go:embed ressources
var res embed.FS

//go:embed config.yml
var config []byte

func main() {
	if !amAdmin() {
		runMeElevated("")
	}

	fmt.Println("Welcome to the Flex Installer.")

	cfg := ParseConfig(config)

	color.Cyan("Now running actions..")
	CreateTemp()
	for _, a := range cfg.Actions {
		a.Run(&res)
	}
	color.Cyan("Done!")
}

func runMeElevated(args string) {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
