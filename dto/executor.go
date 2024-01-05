package dto

type ExecutorSettings struct {
	MaxGoroutines int

	InputFolder  string
	OutputFolder string
	Recursive    bool
}
