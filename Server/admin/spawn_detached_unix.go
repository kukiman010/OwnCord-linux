//go:build !windows

package admin

import (
	"os"
	"os/exec"
	"syscall"
)

// spawnDetached starts a new process that is not attached to the current one.
func spawnDetached(exePath string, args []string) error {
	cmd := exec.Command(exePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Best-effort "detach" for Unix: run in a new process group so it won't
	// receive signals sent to the parent's group (e.g. Ctrl+C).
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	return cmd.Start()
}

