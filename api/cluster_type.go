package api

type ClusterJoinConfig struct {
	ConfigDigest  string              `json:"config_digest"`
	NodeList      []ClusterNodeConfig `json:"nodelist"`
	PreferredNode string              `json:"preferred_node"`
	Totem         Totem               `json:"totem"`
}

type ClusterNodeConfig struct {
	Name        string `json:"name"`
	NodeID      string `json:"nodeid"`
	PVEAddr     string `json:"pve_addr"`
	PVEFP       string `json:"pve_fp"`
	QuorumVotes string `json:"quorum_votes"`
	Ring0Addr   string `json:"ring0_addr"`
}

type Totem struct {
	ClusterName   string      `json:"cluster_name"`
	ConfigVersion string      `json:"config_version"`
	Interface     interface{} `json:"interface"`
	IPVersion     string      `json:"ip_version"`
	LinkMode      string      `json:"link_mode"`
	SecAuth       string      `json:"secauth"`
	Version       string      `json:"version"`
}

type CorosyncLink struct {
	Link0 string `json:"link0"`
	Link1 string `json:"link1"`
	Link2 string `json:"link2"`
	Link3 string `json:"link3"`
	Link4 string `json:"link4"`
	Link5 string `json:"link5"`
	Link6 string `json:"link6"`
	Link7 string `json:"link7"`
	Link8 string `json:"link8"`
}
