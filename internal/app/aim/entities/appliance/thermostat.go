package appliance

import (
	"fmt"
	"math"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

const HEATING_THRESHOLD_MIN = 0
const HEATING_THRESHOLD_MAX = 25

const COOLING_THRESHOLD_MIN = 10
const COOLING_THRESHOLD_MAX = 35

var _ Appliance = Thermostat{}

type Thermostat struct {
	*ApplianceData
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

func NewThermostat(name string, deviceID string,
	s float64, miht int, maht int, mict int, mact int) (t Thermostat, err error) {
	err = validate(s, miht, maht, mict, mact)
	if err != nil {
		return t, entities.NewError(entities.CodeInvaildInput, err)
	}

	a, err := NewApplianceData(name, deviceID,
		getCommands(s, miht, maht, mict, mact))

	if err != nil {
		return t, err
	}

	t = Thermostat{
		ApplianceData:      a,
		Scale:              s,
		MinimumHeatingTemp: miht,
		MaximumHeatingTemp: maht,
		MinimumCoolingTemp: mict,
		MaximumCoolingTemp: mact,
	}

	return t, err
}

func validate(s float64, miht int, maht int, mict int, mact int) error {
	//スケール(何度ずつ刻むかの値)のバリデーションチェック
	if !(0.1 <= s && s <= 1 && s-foor2ndDiminals(s) == 0) {
		return fmt.Errorf("scale need to be in between 0.1 and 5.0 available in 0.1")
	}

	if !(miht < maht) {
		return fmt.Errorf("maximum heating tempature need to be bigger then minimume one")
	}

	//暖房の最大温度のバリデーションチェック
	if !(HEATING_THRESHOLD_MIN <= maht && maht <= HEATING_THRESHOLD_MAX) {
		return fmt.Errorf("maximum heating tempature need to be in between 0 to 25")
	}

	//暖房の最小温度のバリデーションチェック
	if !(HEATING_THRESHOLD_MIN <= miht && miht <= HEATING_THRESHOLD_MAX) {
		return fmt.Errorf("minimum heating tempature need to be in between 0 to 25")
	}

	if !(mict < mact) {
		return fmt.Errorf("maximum cooling tempature need to be bigger then minimume one")
	}

	//冷房の最大温度のバリデーションチェック
	if !(COOLING_THRESHOLD_MIN <= mact && mact <= COOLING_THRESHOLD_MAX) {
		return fmt.Errorf("maximum cooling tempature need to be in between 10 to 35")
	}

	//暖房の最小温度のバリデーションチェック
	if !(COOLING_THRESHOLD_MIN <= mict && mict <= COOLING_THRESHOLD_MAX) {
		return fmt.Errorf("maximum cooling tempature need to be in between 10 to 35")
	}

	return nil
}

func getCommands(s float64, miht int, maht int, mict int, mact int) []command.Command {
	var temp float64
	var size int = int(float64(mact-mict)/s + 1 + float64(maht-miht)/s + 1 + 1)
	var commands = make([]command.Command, 0, size)

	temp = float64(miht)
	for temp <= float64(maht) {
		commands = append(commands, command.New(command.Name(fmt.Sprintf("h%.1f", temp)), nil))
		temp += s
		temp = round2ndDiminals(temp)
	}

	temp = float64(mict)
	for temp <= float64(mact) {
		commands = append(commands, command.New(command.Name(fmt.Sprintf("c%.1f", temp)), nil))
		temp += s
		temp = round2ndDiminals(temp)
	}

	commands = append(commands, command.New("off", nil))
	return commands
}

func foor2ndDiminals(f float64) float64 {
	r := math.Floor(f*10) / 10
	return r
}

func round2ndDiminals(f float64) float64 {
	r := math.Round(f*10) / 10
	return r
}

func (t Thermostat) ChangeCommandName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("thermostat appliance does not support changing command name"),
	)
}

func (t Thermostat) AddCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("thermostat appliance does not support adding command"),
	)
}

func (t Thermostat) RemoveCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("thermostat appliance does not support removing command"),
	)
}
