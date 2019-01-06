package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var (
	libkernel32                  = syscall.MustLoadDLL("kernel32")
	procGenerateConsoleCtrlEvent = libkernel32.MustFindProc("GenerateConsoleCtrlEvent")
)

const (
	createNewProcessGroupFlag = 0x00000200
)

func main() {

	log.Printf("My PID: %d\n", os.Getpid())

	cmd1 := exec.Command("powershell", "-File", "./spawnDockerContainer.ps1")
	cmd1.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_UNICODE_ENVIRONMENT | createNewProcessGroupFlag,
	}
	
	err := cmd1.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Root Process ID: %d. Waiting for docker container to spawn\n", cmd1.Process.Pid)
	time.Sleep(30 * time.Second)
	log.Printf("Sending CTRL_BREAK_EVENT to root process\n")

	r1, _, err := procGenerateConsoleCtrlEvent.Call(syscall.CTRL_BREAK_EVENT, uintptr(cmd1.Process.Pid))
	if r1 == 0 {
		log.Fatal(err)
	}

    log.Printf("Waiting for commands to finish...")
	err = cmd1.Wait()
	log.Printf("Command finished with error: %v", err)
}
