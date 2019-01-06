package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	log.Printf("My PID: %d\n", os.Getpid())

	cmd1 := exec.Command("powershell", "-File", "./child1.ps1")
	err := cmd1.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Child Process ID: %d", cmd1.Process.Pid)

        cmd2 := exec.Command("powershell", "-File", "./child2.ps1")
	err = cmd2.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Child Process ID: %d", cmd2.Process.Pid)

	log.Printf("Waiting for commands to finish...")
	err = cmd1.Wait()
	log.Printf("Command finished with error: %v", err)
	err = cmd2.Wait()
	log.Printf("Command finished with error: %v", err)



}
