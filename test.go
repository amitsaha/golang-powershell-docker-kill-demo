package main

import (
	"log"
	"os/exec"
	"syscall"
	"time"
)

func main() {

	cmd := exec.Command("powershell", "-File", "./test.ps1")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(30000 * time.Millisecond)

	log.Printf("Process ID: %d", cmd.Process.Pid)
	// This doesn't do anything
	log.Printf("Sending SIGTERM to PID: %d", cmd.Process.Pid)
	err = cmd.Process.Signal(syscall.SIGTERM)
	// This kills the process
	cmd.Process.Kill()

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()

	log.Printf("Command finished with error: %v", err)
}
