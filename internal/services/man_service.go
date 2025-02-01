package services

import (
	"log/slog"
	"os/exec"
)

func GenerateManPage(arg string) ([]byte, error) {
	cmd := exec.Command("man", "-P", "cat", arg)

	slog.Info("Generate man page", "command", cmd.String())
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
