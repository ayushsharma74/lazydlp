package ytdlp

type Client struct {
	Binary string
}

func New(binary string) *Client {
	return &Client{Binary: binary}
}
