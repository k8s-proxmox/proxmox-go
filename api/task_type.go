package api

type Task struct {
	Endtime int `json:"endtime"`
	// Id        string `json:"id,omitempty"`
	PID       int    `json:"pid"`
	PStart    int    `json:"pstart"`
	StartTime int    `json:"starttime"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UPID      string `json:"upid"`
	User      string `json:"user"`
}

type TaskStatus struct {
	Exitstatus string `json:"exitstatus"`
	Id         string `json:"id"`
}
