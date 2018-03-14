package metrics

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// WithRootAt tweaks the root directory of the file system
func WithRootAt(p string) func(*Metrics) error {
	log.Debugf("Setting root directory to %s", p)
	return func(m *Metrics) error {
		m.root = p
		return nil
	}
}

// WithGPUInfoCommand tweaks the default gpu info command
func WithGPUInfoCommand(cmd *exec.Cmd) func(*Metrics) error {
	log.Debugf("Setting gpu info command to '%s'", cmd.Args)
	return func(m *Metrics) error {
		m.gpuInfoCmd = cmd
		return nil
	}
}

// WithScreenInfoCommand tweaks the default screen info command
func WithScreenInfoCommand(cmd *exec.Cmd) func(*Metrics) error {
	log.Debugf("Setting screen info command to '%s'", cmd.Args)
	return func(m *Metrics) error {
		m.screenInfoCmd = cmd
		return nil
	}
}

// WithSpaceInfoCommand tweaks the default space info command
func WithSpaceInfoCommand(cmd *exec.Cmd) func(*Metrics) error {
	log.Debugf("Setting space info command to '%s'", cmd.Args)
	return func(m *Metrics) error {
		m.spaceInfoCmd = cmd
		return nil
	}
}
