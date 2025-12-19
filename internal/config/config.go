// Package config
// We load the initial configuration for the cli
// to run in the env in this package
package config

import "os"

type Config struct {
	YtDlpPath       string
	OutputDirectory string
}

func Load() Config {
	path := os.Getenv("YT_DLP_PATH")
	if path == "" {
		path = "yt-dlp"
	}

	out := os.Getenv("LAZYDLP_OUT")
	if out == "" {
		out = "./downloads"
	}

	return Config{
		YtDlpPath:       path,
		OutputDirectory: out,
	}
}

