package api

type Version struct {
	Release string  `json:"release"`
	RepoID  string  `json:"repoid"`
	Version string  `json:"version"`
	Console Console `json:"console"`
}

type Console string
