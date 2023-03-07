package boundary

type AddApplianceReq struct {
	name string
}

type AddThermostatReq struct {
	scale              float64
	maximumHeatingTemp int
	minimumHeatingTemp int
	maximumCoolingTemp int
	minimumCoolingTemp int
}
