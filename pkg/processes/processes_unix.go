package processes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// NewUnixProcesses creates a new API to execute unix family processes.
func NewUnixProcesses() Processes {
	return &unixProcesses{}
}

type unixProcesses struct {}

// Output holds pipes for stdout and stderr.
type Output struct {
	Stdout io.ReadCloser
	Stderr io.ReadCloser
}

func (up *unixProcesses) Start(ci CommandInput) (*exec.Cmd, *Output, error) {
	cmd := exec.Command(ci.Name(), ci.Arguments()...)
	cmd.Env = append(os.Environ(), ci.EnvironmentVariables()...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: ci.Pgid()}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect stdout %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect stderr %v", err)
	}

	for i := 0; ; i++ {
		err = cmd.Start()
		if err != nil {
			if i == int(ci.NumberOfRetries()) {
				return nil, nil, fmt.Errorf("failed to run the process %v", err)
			}
		} else {
			output := &Output{
				Stdout: stdout,
				Stderr: stderr,
			}

			return cmd, output, nil
		}
	}
}

func (up *unixProcesses) Read(output *Output) {
	buf := new(bytes.Buffer)
	if output.Stdout != nil {
		buf.ReadFrom(output.Stdout)
		fmt.Printf("!%v\n", buf.String())
	} else {
		buf.ReadFrom(output.Stderr)
		fmt.Printf("process has failed with: [%v]\n", buf.String())
	}
}

func (up *unixProcesses) Stop(pid int, signum syscall.Signal) error {
	pgid, err := syscall.Getpgid(pid)

	if err != nil {
		return fmt.Errorf("failed to get a PGID process: [%v]", err)
	}
	syscall.Kill(-pgid, signum)

	return nil
}
