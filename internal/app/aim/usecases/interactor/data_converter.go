package interactor

import (
	"encoding/json"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	bdr "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func convertApp(a app.Appliance) bdr.Appliance {
	var r bdr.Appliance
	var appID string = string(a.GetID())
	var devID string = string(a.GetDeviceID())
	var name string = string(a.GetName())

	switch a.GetType() {
	case app.TypeCustom:
		r = bdr.NewCustom(appID, name, devID)
		return r
	case app.TypeButton:
		r = bdr.NewButton(appID, name, devID)
		return r
	case app.TypeSwitch:
		r = bdr.NewSwitch(appID, name, devID)
	case app.TypeThermostat:
		var t = app.ThermostatOpt{}
		json.Unmarshal([]byte(a.GetOpt()), &t)
		r = bdr.NewThermostat(appID, name, devID,
			t.Scale,
			t.MinimumHeatingTemp, t.MaximumHeatingTemp,
			t.MinimumCoolingTemp, t.MaximumCoolingTemp)
		return r
	}
	return r
}
