package dto

type LevelSettings struct {
	LowInput, HighInput   int
	Gamma                 float64
	LowOutput, HighOutput int

	DiapasonBlack, DiapasonWhite int
}
