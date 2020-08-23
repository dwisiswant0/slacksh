package slacksh

import "os/exec"

// Run will execute shell command with bash
func Run(text string) (string, error) {
	out, err := exec.Command("/bin/bash", "-c", text).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
