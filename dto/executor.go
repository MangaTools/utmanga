package dto

type ExecutorSettings struct {
	MaxGoroutines int

	InputPath  string
	OutputPath string
	Recursive  bool
}
