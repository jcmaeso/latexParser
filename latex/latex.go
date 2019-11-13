package latex

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)


var workingLatexDir string = ".\\latex\\latexResources\\"
var mainLatexFile string = ".\\latex\\latexResources\\main.tex"
var outputDirectory string = "\\latexOutput"

func CompilePdf(filename string)  {
	cmd :=  exec.Command("latexmk","-pdf","-cd",mainLatexFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Creating latex file failed with %s\n", err)
	}
	err = copyResFile(workingLatexDir+"main.pdf",fmt.Sprintf("%slatexOutput\\%s.pdf",workingLatexDir,filename))
	if err != nil{
		log.Fatalf("Copying latex file to outputdir filed with %s", err)
	}
	cleanCompilationFiles()
}

func copyResFile(inputFilename, outputFilename string) error {
	sourceFileStat, err := os.Stat(inputFilename)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", inputFilename)
	}

	source, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err := os.Stat(workingLatexDir+outputDirectory); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(workingLatexDir+outputDirectory,0755)
		}
	}

	destination, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer destination.Close()
	_ ,err = io.Copy(destination, source)
	return err
}

func cleanCompilationFiles(){
	files, err := ioutil.ReadDir(workingLatexDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && !strings.HasSuffix(f.Name(),".tex"){
			err = os.Remove(workingLatexDir+string(f.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}