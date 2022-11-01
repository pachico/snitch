package domain

type File struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	IsDir      bool   `json:"is_dir"`
	Owner      string `json:"owner"`
	Permission string `json:"permission"`
}

type FSReport []File
