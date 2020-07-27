package main

import (
	"os"
	"flag"
	"syscall"
	"os/exec"

	"github.com/qydysky/part"
)

var (
	path = flag.String("path", "", "exe path.")
	file *os.File
)

func main(){

	flag.Parse()
	
	if *path == "" {
		file, _ = os.OpenFile("README.Use.txt",os.O_CREATE|os.O_WRONLY, 0666)
		file.WriteAt([]byte("-path {your *.exe path}"), 0)
		file.Close()
		return
	}

    if part.Checkfile().IsOpen(".lock.loop") {
		file, _ = os.OpenFile("README.Use.txt",os.O_CREATE|os.O_WRONLY, 0666)
		file.WriteAt([]byte("u must shutdown the exe and then remove .lock.loop file"), 0)
		file.Close()
		return
	}

	file, _ = os.Create(".lock.loop")

	cmd := exec.Command(*path)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	cmd.Start()
	cmd.Wait()
	
	file.Close()
	os.Remove(".lock.loop")

	return
}