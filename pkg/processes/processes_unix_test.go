package processes

import (
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestProcesses_Start(t *testing.T) {
	commandInput := NewCommandInput("ls")

	processes := NewUnixProcesses()

	_, _, err := processes.Start(commandInput)
	if err != nil {
		t.Fatalf("execution failed [%v]", err)
	}
}

func TestProcesses_Arguments_Start(t *testing.T) {
	commandInput := NewCommandInput("ls")
	args := []string{"-altr"}
	commandInput.SetArguments(args)

	processes := NewUnixProcesses()
	_, _, err := processes.Start(commandInput)
	if err != nil {
		t.Fatalf("execution failed [%v]", err)
	}
}

func TestProcesses_EnvVariable_Start(t *testing.T) {
	commandInput := NewCommandInput("./test_scripts/error.sh")
	envVariable := []string{"IS_ERROR=false"}
	commandInput.SetEnvironmentVariables(envVariable)
	commandInput.SetNumberOfRetries(0)

	processes := NewUnixProcesses()

	cmd, _, _ := processes.Start(commandInput)
	err := cmd.Wait()
	if err != nil {
		t.Fatalf("execution should not have failed [%v]", err)
	}

	envVariable = []string{"IS_ERROR=true"}
	commandInput.SetEnvironmentVariables(envVariable)

	cmd, _, _ = processes.Start(commandInput)
	err = cmd.Wait()

	if err == nil {
		t.Fatalf("execution should have failed [%v]", err)
	}
}

func TestProcesses_Stop(t *testing.T) {
	var cmd *exec.Cmd
	go func() {
		cmd = exec.Command("./test_scripts/hello.sh")
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Run()
	}()

	time.Sleep(1 * time.Second)

	processes := NewUnixProcesses()
	processes.Stop(cmd.Process.Pid, syscall.SIGKILL)

	// If sig is 0, then no signal is sent, but error checking is still performed;
	// this can be used to check for the existence of a process ID or process
	// group ID.
	err := syscall.Kill(cmd.Process.Pid, syscall.Signal(0))
	if err != nil {
		t.Fatalf("process should have been killed [%v]", err)
	}
}
