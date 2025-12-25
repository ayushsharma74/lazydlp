package ytdlp

import (
	"bufio"
	"io"
	"log"
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

func isNumericFormatID(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
func normalizeFormatID(formatID string) string {
	switch {
	case isNumericFormatID(formatID):
		// real yt-dlp format id like 137
		return formatID + "+bestaudio"

	case formatID == "mp4":
		return "bv*[ext=mp4]/bv*+ba[ext=m4a]/b"

	case formatID == "webm":
		return "bv*[ext=webm]/bv*+ba[ext=webm]/b"

	default:
		// unknown / empty / broken input
		return "best"
	}
}



func (c *Client) Download(
	url string,
	formatID string,
	outputDir string,
	progressCh chan<- domain.ProgressUpdate,
) error {

	format := normalizeFormatID(formatID)

	cmd := exec.Command(
		c.Binary,
		"-f", format,
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

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("LOG", line)

		if strings.Contains(line, "[download]") && strings.Contains(line, "%") {
			pct := extractPercent(line)
			speed := extractSpeed(line)

			if pct >= 0 {
				progressCh <- domain.ProgressUpdate{
					Percent: pct,
					Speed:   speed,
				}
			}
		}
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

