package userdata

import (
	timeofday "FileDirectoryCreation/TimeOfDay"
	"fmt"
)

type UserData struct {
	GetUsername func() string
	USERNAME    string
}

func (data UserData) Get() string {

	dayTime := timeofday.TimeDayMesh{}
	if len(data.USERNAME) > 0 {

		return fmt.Sprintf("%s %s", data.GetUsername(), dayTime.GetTimeofDay())
	}

	return fmt.Sprintf("%s %s", data.GetUsername(), dayTime.GetTimeofDay())
}
