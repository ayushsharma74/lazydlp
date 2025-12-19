package domain

type Format struct {
	ID         string
	Ext        string
	Resolution string
	Size       int64
	FPS        int
	IsVideo    bool
	IsAudio    bool
}
