package boundary

type AddCustomInput struct {
	Name     string
	DeviceID string
}

type AddToggleInput struct {
	Name     string
	DeviceID string
}

type AddButtonInput struct {
	Name     string
	DeviceID string
}

type AddThermostatInput struct {
	Name               string
	DeviceID           string
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

type GetAppInput struct {
	AppID string
}

type RenameAppInput struct {
	AppID string
	Name  string
}

type ChangeIRDevInput struct {
	AppID    string
	DeviceID string
}

type DeleteAppInput struct {
	AppID string
}

type RenameCommandInput struct {
	AppID string
	ComID string
	Name  string
}

type AddCommandInput struct {
	AppID string
	Name  string
}

type DeleteCommandInput struct {
	AppID string
	ComID string
}

type GetCommandInput struct {
	AppID string
	ComID string
}

type SetIRDataInput struct {
	AppID string
	ComID string
	Data  IRData
}
