package toggle

import (
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Toggle struct {
	app.Appliance
}

func New(id app.ID, name app.Name, deviceID app.DeviceID) Toggle {
	return Toggle{
		app.NewAppliance(id, name, app.TypeToggle, deviceID, []command.Command{
			command.New("", "on", nil),
			command.New("", "off", nil),
		}),
	}
}
