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
	Mode string

	Coefficient               int
	TargetHeigth, TargetWidth int

	Steps int
}
