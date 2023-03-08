package appliance

import (
	"encoding/json"
	"fmt"
	"math"
	"unsafe"

	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

const HEATING_THRESHOLD_MIN = 0
const HEATING_THRESHOLD_MAX = 25

const COOLING_THRESHOLD_MIN = 10
const COOLING_THRESHOLD_MAX = 35

type ThermostatOpt struct {
	Scale              float64 `json:"scale"`
	MaximumHeatingTemp int     `json:"maximum_heating_temp"`
	MinimumHeatingTemp int     `json:"minimum_heating_temp"`
	MaximumCoolingTemp int     `json:"maximum_cooling_temp"`
	MinimumCoolingTemp int     `json:"minimum_cooling_temp"`
}

func NewThermostat(name Name, deviceID DeviceID,
	s float64, miht int, maht int, mict int, mact int) (Appliance, error) {

	var a *ApplianceData

	err := validateInput(s, miht, maht, mict, mact)
	if err != nil {
		return a, err
	}

	rawOpt, err := json.Marshal(
		ThermostatOpt{
			Scale:              s,
			MaximumHeatingTemp: maht,
			MinimumHeatingTemp: miht,
			MaximumCoolingTemp: mact,
			MinimumCoolingTemp: mict,
		},
	)
	if err != nil {
		return a, err
	}

	opt, err := NewOpt(*(*string)(unsafe.Pointer(&rawOpt)))
	if err != nil {
		return a, err
	}

	a, err = NewAppliance(name, AppTypeThermostat, deviceID, opt)
	if err != nil {
		return a, err
	}

	a.commands = getCommands(s, miht, maht, mict, mact)
	return a, nil
}

func validateInput(s float64, miht int, maht int, mict int, mact int) error {
	//スケール(何度ずつ刻むかの値)のバリデーションチェック
	if !(0.1 <= s && s <= 1 && s-foor2ndDiminals(s) == 0) {
		return fmt.Errorf("scale need to be in between 0.1 and 5.0 available in 0.1")
	}

	if !(miht < maht) {
		fmt.Printf("%d, %d \n", miht, maht)
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

	if !(miht < maht) {
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
	var name string
	var temp float64
	var size int = int(float64(mact-mict)/s + 1 + float64(maht-miht)/s + 1 + 1)
	var commands = make([]command.Command, 0, size)

	temp = float64(miht)
	for temp <= float64(maht) {
		name = fmt.Sprintf("h%.1f", temp)
		commands = append(commands, command.New(name))
		temp += s
		temp = round2ndDiminals(temp)
	}

	temp = float64(mict)
	for temp <= float64(mact) {
		name = fmt.Sprintf("c%.1f", temp)
		commands = append(commands, command.New(name))
		temp += s
		temp = round2ndDiminals(temp)
	}
	commands = append(commands, command.New("off"))
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
