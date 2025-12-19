package ytdlp

import (
	"encoding/json"
	"log"
	"os/exec"
	"github.com/ayushsharma74/lazydlp/internal/domain"
)

type ytDlpJSON struct {
	Formats []ytFormat `json:"formats"`
}

type ytFormat struct {
	ID             string  `json:"id"`
	Ext            string  `json:"ext"`
	Resolution     string  `json:"resolution"`
	Filesize       int64  `json:"filesize"`
	FilesizeApprox int64  `json:"filesize_approx"`
	VCodec         string  `json:"vcodec"`
	ACodec         string  `json:"acodec"`
	FPS            float64 `json:"fps"`
	FormatNote     string  `json:"format_note"`
}

func (c *Client) ListFormats(url string) ([]domain.Format, error) {
	cmd := exec.Command(
		c.Binary,
		"-J",
		"--no-warnings",
		"--no-progress",
		url,
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	
	var data ytDlpJSON

	if err := json.Unmarshal(out, &data); err != nil {
		return nil, err
	}

	for _, f := range data.Formats {
		log.Println(f)
	}

	var formats []domain.Format

	for _, f := range data.Formats {

		// skip storyboards & useless entries
		if f.VCodec == "none" && f.ACodec == "none" {
			continue
		}

		size := f.Filesize
		if size == 0 {
			size = f.FilesizeApprox
		}

		formats = append(formats, domain.Format{
			ID:         f.ID,
			Ext:        f.Ext,
			Resolution: f.Resolution,
			Size:       size,
			FPS:        int(f.FPS),
			IsVideo:    f.VCodec != "none",
			IsAudio:    f.ACodec != "none",
		})
	}

	return formats, nil

}
