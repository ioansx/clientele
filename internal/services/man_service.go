package services

import (
	"fmt"
	"os/exec"

	"github.com/ioansx/clientele/internal/models"
)

func GenerateManPage(arg string) (*models.ManGetOutdto, error) {
	cmd := exec.Command("man", "-P", "cat", arg)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Output for command '%s': %w", cmd.String(), err)
	}

	return &models.ManGetOutdto{Output: string(output)}, nil
}
