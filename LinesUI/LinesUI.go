package linesui

type LinesUIState struct {
	LinesUI  string
	GetLines func() string
}

const (
	xPos = int(80/2) / 2
	yPos = int(10 / 5)
)

func (data LinesUIState) GetLinesLength() int {

	return int(xPos*yPos) / 2

}
