package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

type Program struct {
	Name       string   `yaml:"Name"`
	Kind       string   `yaml:"Kind"` //exe, powershell, bat, msi
	Args       []string `yaml:"Args"`
	Ressources []string `yaml:"Ressources"`
	Path       string   `yaml:"Path"`
}

func (p Program) extract(fs *embed.FS) {
	fmt.Printf("Extracting %s, please wait..\n", p.Name)
	//extract main
	mainRessource, err := fs.ReadFile("ressources/" + p.Path)
	if err != nil {
		panic(err)
	}
	createSubDir(p.Name)
	CreateFile(p.Name+"\\"+p.Path, mainRessource)

	//Now for ressources

	if p.Ressources != nil {
		for _, v := range p.Ressources {
			res, err := fs.ReadFile("ressources/" + v)
			if err != nil {
				panic(err)
			}
			CreateFile(p.Name+"\\"+v, res)

		}
	}
	fmt.Printf("Successfully extracted %s.\n", p.Name)
}

func (p Program) Run(fs *embed.FS) {
	p.extract(fs)

	fmt.Printf("Now running %s, this may take a while..\n", p.Name)

	switch p.Kind {
	case "exe":
		p.runExe()
	case "msi":
		p.runMsi()
	case "pwsh":
		p.runPwsh()
	}

	fmt.Printf("Finished running %s. \n", p.Name)

}

func (p Program) runPwsh() {
	cmd := &exec.Cmd{
		Path: "powershell",
		Args: append([]string{
			"powershell",
			BASEPATH + p.Name + "\\" + p.Path,
		}, p.Args...),
		Stdout: os.Stdout,
	}

	if runErr := cmd.Run(); runErr != nil {
		fmt.Printf("Couldn't run %s: %s", p.Name, runErr)
	}

}

func (p Program) runExe() {
	cmd := &exec.Cmd{
		Path:   BASEPATH + p.Name + "\\" + p.Path,
		Args:   append([]string{"p.Path"}, p.Args...),
		Stdout: os.Stdout,
	}

	if runErr := cmd.Run(); runErr != nil {
		fmt.Printf("Couldn't run %s: %s", p.Name, runErr)
	}
}

func (p Program) runMsi() {
	cmd := &exec.Cmd{
		Path: "msiexec",
		//QN for NO UI /QB For basic UI
		Args:   append([]string{"msiexec", "/qn", "/i", BASEPATH + p.Name + "\\" + p.Path}, p.Args...),
		Stdout: os.Stdout,
	}

	if runErr := cmd.Run(); runErr != nil {
		fmt.Printf("Couldn't run msiexec: %s", runErr)
	}
}
