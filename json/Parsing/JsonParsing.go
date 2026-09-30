package parsing_data_from_json

import (
	"fmt"
	"strings"
)

type DataTypedMesh struct {
	GetData     func() string
	IsDataFinal bool
	Text        string
	Version     string
	Length      int
}

func (data DataTypedMesh) GetVersion() string {
	if data.Length > 0 {
		return fmt.Sprintf("%s", data.Version)
	}

	return fmt.Sprintf("%s", data.Version)
}

func (data DataTypedMesh) PrintElement() string {
	return fmt.Sprintf("%s", strings.ToUpper(data.Text))
}

func (typedData DataTypedMesh) GetMesh() string {

	typedData.GetData = func() string {
		return strings.ToUpper("Data Has Been DONE")
	}

	return typedData.GetData()
}
