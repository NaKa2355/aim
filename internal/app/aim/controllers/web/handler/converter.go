package handler

import (
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

func convertRemoteType(in bdy.RemoteType) (out aimv1.Remote_RemoteType) {
	switch in {
	case bdy.TypeCustom:
		return aimv1.Remote_CUSTOM
	case bdy.TypeButton:
		return aimv1.Remote_BUTTON
	case bdy.TypeToggle:
		return aimv1.Remote_TOGGLE
	case bdy.TypeThermostat:
		return aimv1.Remote_THERMOSTAT
	default:
		return aimv1.Remote_REMOTE_TYPE_UNKNOWN
	}
}

func convertThermostatScale(scale aimv1.AddThermostatRemoteRequest_Scale) bdy.ThermostatScale {
	switch scale {
	case aimv1.AddThermostatRemoteRequest_HALF:
		return bdy.Half
	case aimv1.AddThermostatRemoteRequest_ONE:
		return bdy.One
	default:
		return bdy.One
	}
}
