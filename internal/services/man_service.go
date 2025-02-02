package services

import (
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/ioansx/clientele/internal/models"
)

func GenerateManPage(arg string) (*models.ManGetOutdto, error) {
	cmd := exec.Command("man", "-P", "cat", arg)

	slog.Debug("Generate man page", "command", cmd.String())
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Output: %w", err)
	}

	return &models.ManGetOutdto{Output: string(output)}, nil
}
