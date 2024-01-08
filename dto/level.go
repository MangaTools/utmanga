package dto

type LevelSettings struct {
	LowInput   int     `json:"low_input"`
	HighInput  int     `json:"high_input"`
	Gamma      float64 `json:"gamma"`
	LowOutput  int     `json:"low_output"`
	HighOutput int     `json:"high_output"`

	DiapasonBlack int `json:"diapason_black"`
	DiapasonWhite int `json:"diapason_white"`
}
