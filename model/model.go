package model

// ListDeployment
type Deployment struct {
	Name     string `json:"name"`
	Avalable int32  `json:"available"`
}

type ListDeploymentResp struct {
	Code    int32         `json:"code"`
	Message string        `json:"message"`
	Data    []*Deployment `json:"data"`
}

// HandleAccess
type HandleAccessResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    *Pod   `json:"data"`
}

type Pod struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}
