package model

// ListDeployment
type Deployment struct {
	Name     string `json:"name"`
	Avalable int32  `json:"available"`
}

type ListDeploymentResp struct {
	Code    int32         `json:"code"`
	Message string        `json:"message,omitempty"`
	Data    []*Deployment `json:"data"`
}

// HandleAccess
type HandleAccessResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message,omitempty"`
	Data    *Pod   `json:"data"`
}

type Pod struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

// List
type ListPodResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message,omitempty"`
	Data    []*Pod `json:"data"`
}

type CommonResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message,omitempty"`
}
