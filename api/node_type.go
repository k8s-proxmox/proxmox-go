package api

type Node struct {
	Cpu            float32 `json:"cpu"`
	Disk           int     `json:"disk"`
	ID             string  `json:"id"`
	Level          string  `json:"level"`
	MaxCpu         int     `json:"maxcpu"`
	MaxDisk        int     `json:"maxdisk"`
	MaxMem         int     `json:"maxmem"`
	Mem            int     `json:"mem"`
	Node           string  `json:"node"`
	SSLFingerprint string  `json:"ssl_fingerprint"`
	Stauts         string  `json:"status"`
	Type           string  `json:"type"`
	UpTime         int     `json:"uptime"`
}
