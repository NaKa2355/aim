package custom

import (
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Custom struct {
	app.Appliance
}

func New(id app.ID, name app.Name, deviceID app.DeviceID) Custom {
	return Custom{
		app.NewAppliance(id, name, app.TypeCustom, deviceID, make([]command.Command, 0)),
	}
}
