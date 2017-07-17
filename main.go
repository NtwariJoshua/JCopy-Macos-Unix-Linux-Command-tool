package main

import (
	"runtime"
	_"os"
	"log"
	"path/filepath"
	"os/exec"
	"io/ioutil"
	"os"
)

func main(){
	arch := runtime.GOOS
	if(len(os.Args) < 1){
		log.Fatal("Supply the required arguments")
	}

	args := os.Args[1]
	fp , err := filepath.Abs(args)
	if(err != nil){
		log.Fatal(err)
	}
	//Open and read the file
	file,err := os.Open(fp)
	if(err != nil){
		log.Fatal(err)
	}
	defer file.Close()
	fileContents,err := ioutil.ReadAll(file)
	if(err != nil){
		log.Fatal(err)
	}
	toClipboard(fileContents,arch)
}

func toClipboard(fileContents []byte, arch string){

	var copyCmd *exec.Cmd

	//Mac Os
	if(arch == "darwin"){
		copyCmd = exec.Command("pbcopy")
	}
	//Linux Os
	if(arch == "linux"){
		copyCmd = exec.Command("xclip", "-selection", "c")
	}

	in,err := copyCmd.StdinPipe()
	if(err != nil){
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil{
		log.Fatal(err)
	}

	if _,err := in.Write([]byte(fileContents)); err != nil {
		log.Fatal(err)
	}
	if err := in.Close(); err != nil{
		log.Fatal(err)
	}
	copyCmd.Wait()

}