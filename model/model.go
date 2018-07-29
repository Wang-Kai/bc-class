package model

// deployment config
type DeploymentConf struct {
	Name     string            `json:"name"`
	Labels   map[string]string `json:"labels"`
	Strategy string            `json:"strategy"`
	Pod      *PodConf          `json:"pod"`
}

// pod config
type PodConf struct {
	Labels     map[string]string `json:"labels"`
	Containers []*ContainerConf  `json:"containers"`
}

// container config
type ContainerConf struct {
	Name           string               `json:"name"`
	Image          string               `json:"image"`
	Command        []string             `json:"command,omitempty"`
	Args           []string             `json:"args,omitempty"`
	ContainerPorts []*ContainerPortConf `json:"containerPorts"`
}

// container port config
type ContainerPortConf struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int32  `json:"container_port"`
}

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
