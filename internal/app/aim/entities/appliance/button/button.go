package button

import (
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Button struct {
	app.Appliance
}

func New(id app.ID, name app.Name, deviceID app.DeviceID) Button {
	return Button{
		app.NewAppliance(id, name, app.TypeButton, deviceID, []command.Command{command.New("", "push", nil)}),
	}
}
