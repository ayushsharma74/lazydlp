package ytdlp

import (
	"bufio"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/ayushsharma74/lazydlp/internal/domain"
)

// Compile once globally for performance
var progressRegexp = regexp.MustCompile(`\s*(\d+(\.\d+)?)%`)
var speedRegexp = regexp.MustCompile(`at\s+([^\s]+(?:B|iB)/s)`)

func extractPercent(line string) float64 {
	matches := progressRegexp.FindStringSubmatch(line)
	if len(matches) >= 2 {
		// matches[1] is the first capture group: the digits and optional decimal
		p, err := strconv.ParseFloat(matches[1], 64)
		if err == nil {
			return p / 100.0
		}
	}
	return -1
}

func extractSpeed(line string) string {
	matches := speedRegexp.FindStringSubmatch(line)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}

func (c *Client) Download(url string, formatID string, outputDir string, progressCh chan<- domain.ProgressUpdate) error {
	cmd := exec.Command(
		c.Binary,
		"-f", formatID,
		"-P", outputDir,
		"--newline",
		"--progress",
		"--no-playlist",
		url,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "[download]") && strings.Contains(line, "%") {
			pct := extractPercent(line)
			speed := extractSpeed(line)

			if pct >= 0 {
				// Wrap the data in the struct before sending
				progressCh <- domain.ProgressUpdate{
					Percent: pct,
					Speed:   speed,
				}
			}
		}
	}

	return cmd.Wait()
}
