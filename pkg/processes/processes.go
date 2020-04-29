package processes

import (
	"os/exec"
	"syscall"
)

// Processes provides API to run processes under unix family.
type Processes interface {

	// Start the specified command but does not wait for it to complete.
	// It might retry running the process up to n-times, defined in a command input.
	Start(command CommandInput) (*exec.Cmd, *Output, error)

	// Read the process stdout and stderr
	Read(output *Output)

	// Stop identifies PGID of PID and stop all the processes which belong to it. 
	Stop(pid int, signum syscall.Signal) error
}

// CommandInput provides an interface to interact with parameters which are used
// to run unix processes.
type CommandInput interface {
	Name() string
	Arguments() []string
	SetArguments([]string)
	EnvironmentVariables() []string
	SetEnvironmentVariables([]string)
	Pgid() bool
	SetPgid(bool)
	NumberOfRetries() uint
	SetNumberOfRetries(uint)
}
