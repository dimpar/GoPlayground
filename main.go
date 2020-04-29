package main

import (
	"log"

	"github.com/dimpar/go-processes/pkg/processes"
)

func main() {
	process := processes.NewUnixProcesses()
	// input := processes.NewCommandInput("sleep")
	// args := []string{"5"}
	input := processes.NewCommandInput("ls")
	args := []string{"-altr"}
	input.SetArguments(args)
	envVariable := []string{"HELLO=Percona"}
	input.SetEnvironmentVariables(envVariable)

	cmd, output, err := process.Start(input)
	if err != nil {
		log.Fatalf("error occured whil running the process %v", err)
		return
	}

	process.Read(output)
	
	cmd.Wait()
	
	// process.Stop(cmd.Process.Pid, syscall.SIGKILL)
}
