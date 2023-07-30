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

type TermProxy struct {
	Port   string `json:"port"`
	Ticket string `json:"ticket"`
	UPID   string `json:"upid"`
	User   string `json:"user"`
}

type TermProxyOption struct {
	CMD     string `json:"cmd,omitempty"`
	CMDOpts string `json:"cmd-opts,omitempty"`
}

type VNCShellOption struct {
	TermProxyOption
	Height    int  `json:"height,omitempty"`
	Websocket bool `json:"websocket,omitempty"`
	Width     int  `json:"width,omitempty"`
}

type VNCWebSocket struct {
	Port      string `json:"port,omitempty"`
	VNCTicket string `json:"vncticket,omitempty"`
}
