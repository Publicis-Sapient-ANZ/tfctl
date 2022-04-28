package exec

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

type CommandConfig struct {
	Command          string
	Paramaters       []string
	WorkingDirecotry string
}

type CommandResult struct {
	StdOut   string
	StdError string
}

type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w: w,
	}
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func ExecuteCommand(commandConfig CommandConfig) (commandResult CommandResult, commandError error) {

	// TODO perform some escaping on any string paramas that are supplied
	cmd := exec.Command(commandConfig.Command, commandConfig.Paramaters...)
	cmd.Dir = commandConfig.WorkingDirecotry
	cmd.Env = os.Environ()

	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		return commandResult, fmt.Errorf("command '%s' failed with %s\n", commandConfig.Command, err)
	}
	if errStdout != nil || errStderr != nil {
		return commandResult, fmt.Errorf("failed to capture stdout or stderr, the command may have run")
	}

	commandResult.StdOut = string(stdout.Bytes())
	commandResult.StdError = string(stderr.Bytes())

	return commandResult, nil

}
