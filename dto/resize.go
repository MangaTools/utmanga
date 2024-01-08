package dto

const (
	ResizeModeBox             = "box"
	ResizeModeLanczos         = "lanczos"
	ResizeModeCubic           = "cubic"
	ResizeModeLinear          = "linear"
	ResizeModeNearestNeighbor = "nearest_neighbor"
)

var ResizeModeList = []string{ResizeModeLanczos, ResizeModeCubic, ResizeModeLinear, ResizeModeBox, ResizeModeNearestNeighbor}

type ResizeSettings struct {
	Mode string `json:"mode"`

	Coefficient  int `json:"coefficient"`
	TargetHeigth int `json:"target_heigth"`
	TargetWidth  int `json:"target_width"`

	Steps int `json:"steps"`
}
