package models

import (
	"encoding/base64"
	"path/filepath"
)

type Attachment struct {
	Mime    string
	Path    string
	Content []byte
}

func (a Attachment) Name() string {
	return filepath.Base(a.Path)
}

func (a Attachment) Base64Content() string {
	return base64.StdEncoding.EncodeToString(a.Content)
}
