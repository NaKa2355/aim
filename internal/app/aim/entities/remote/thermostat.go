package remote

import (
	"errors"
	"fmt"
	"math"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

const HEATING_THRESHOLD_MIN = 0
const HEATING_THRESHOLD_MAX = 25

const COOLING_THRESHOLD_MIN = 10
const COOLING_THRESHOLD_MAX = 35

type thermostatController struct{}

func NewThermostat(name string, deviceID string,
	s float64, miht int, maht int, mict int, mact int) (t *Remote, err error) {

	ctr := thermostatController{}
	err = validate(s, miht, maht, mict, mact)
	if err != nil {
		return t, entities.NewError(entities.CodeInvaildInput, err)
	}

	coms := getButtons(s, miht, maht, mict, mact)

	return NewAppliance(name, deviceID, TypeThermostat, coms, ctr)
}

func LoadThermostat(id ID, name Name, deviceID DeviceID) *Remote {
	a := LoadAppliance(id, name, deviceID, TypeThermostat, thermostatController{})
	return a
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

func getButtons(s float64, miht int, maht int, mict int, mact int) []*button.Button {
	var temp float64
	var size int = int(float64(mact-mict)/s + 1 + float64(maht-miht)/s + 1 + 1)
	var commands = make([]*button.Button, 0, size)

	temp = float64(miht)
	for temp <= float64(maht) {
		commands = append(commands, button.New(button.Name(fmt.Sprintf("h%.1f", temp)), nil))
		temp += s
		temp = round2ndDiminals(temp)
	}

	temp = float64(mict)
	for temp <= float64(mact) {
		commands = append(commands, button.New(button.Name(fmt.Sprintf("c%.1f", temp)), nil))
		temp += s
		temp = round2ndDiminals(temp)
	}

	commands = append(commands, button.New("off", nil))
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

func (c thermostatController) ChangeButtonName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("thermotat appliance does not support changing the button name"),
	)
}

func (c thermostatController) AddButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("thermostat appliance does not support adding a button"),
	)
}

func (c thermostatController) RemoveButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("thermostat appliance does not support removing the button"),
	)
}
