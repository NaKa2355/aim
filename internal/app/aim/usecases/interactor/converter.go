package interactor

import (
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func convertType(in remote.RemoteType) (out bdy.RemoteType) {
	switch in {
	case remote.TypeCustom:
		return bdy.TypeCustom
	case remote.TypeButton:
		return bdy.TypeButton
	case remote.TypeToggle:
		return bdy.TypeToggle
	case remote.TypeThermostat:
		return bdy.TypeThermostat
	default:
		return bdy.RemoteType(in)
	}
}
