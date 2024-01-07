package dto

const (
	ResizeModeLanczos           = "lanczos"
	ResizeModeCatmullRom        = "catmullRom"
	ResizeModeMitchellNetravali = "mitchell_netravali"
	ResizeModeLinear            = "linear"
	ResizeModeBox               = "box"
	ResizeModeNearestNeighbor   = "nearest_neighbor"
)

var ResizeModeList = []string{ResizeModeLanczos, ResizeModeCatmullRom, ResizeModeMitchellNetravali, ResizeModeLinear, ResizeModeBox, ResizeModeNearestNeighbor}

type ResizeSettings struct {
	Mode string

	Coefficient               int
	TargetHeigth, TargetWidth int

	Steps int
}
