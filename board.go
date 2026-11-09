package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type DetectedBoard struct {
	Port string
	FQBN string
}

// ParseBoardList extracts ports and FQBN values from arduino-cli board list output.
func ParseBoardList(output string) []DetectedBoard {
	var boards []DetectedBoard

	for _, line := range strings.Split(output, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 || strings.EqualFold(fields[0], "Port") {
			continue
		}

		fqbn := ""
		for _, field := range fields {
			if strings.Count(field, ":") >= 2 {
				fqbn = field
				break
			}
		}
		if fqbn == "" {
			continue
		}

		boards = append(boards, DetectedBoard{
			Port: fields[0],
			FQBN: fqbn,
		})
	}

	return boards
}

// RunArduinoCLI executes an installed arduino-cli command and returns its output.
func RunArduinoCLI(args ...string) (string, error) {
	command := exec.Command("arduino-cli", args...)
	output, err := command.CombinedOutput()
	result := strings.TrimSpace(string(output))
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return result, fmt.Errorf("arduino-cli was not found; install it and make sure it is on PATH")
		}
		if result != "" {
			return result, fmt.Errorf("arduino-cli %s failed: %w", strings.Join(args, " "), err)
		}
		return result, fmt.Errorf("arduino-cli %s failed: %w", strings.Join(args, " "), err)
	}
	return result, nil
}
