package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: chronic COMMAND...")
		os.Exit(255)
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	b, err := cmd.CombinedOutput()
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", err, string(b))
	status := 1
	state := cmd.ProcessState
	if state != nil {
		if ws, ok := state.Sys().(syscall.WaitStatus); ok {
			if ws.Exited() {
				status = ws.ExitStatus()
			}
		}
	}
	os.Exit(status)
}
