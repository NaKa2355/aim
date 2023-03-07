package appliance

import (
	"fmt"
	"math"

	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

const HEATING_THRESHOLD_MIN = 0
const HEATING_THRESHOLD_MAX = 25

const COOLING_THRESHOLD_MIN = 10
const COOLING_THRESHOLD_MAX = 35

type Thermostat struct {
	*ApplianceData

	scale              float64
	maximumHeatingTemp int
	minimumHeatingTemp int
	maximumCoolingTemp int
	minimumCoolingTemp int
}

func NewThermostat(name string, deviceID string,
	s float64, miht int, maht int, mict int, mact int) (*Thermostat, error) {

	var t *Thermostat

	err := validateInput(s, miht, maht, mict, mact)
	if err != nil {
		return t, err
	}

	a, err := NewAppliance(name, AppTypeThermostat, deviceID)
	if err != nil {
		return t, err
	}

	t = &Thermostat{
		ApplianceData:      a,
		scale:              s,
		maximumHeatingTemp: maht,
		minimumHeatingTemp: miht,
		maximumCoolingTemp: mact,
		minimumCoolingTemp: mict,
	}

	t.addCommands()
	return t, nil
}

func (t *Thermostat) GetScale() float64 {
	return t.scale
}

func (t *Thermostat) GetMaximumHeatingTemp() int {
	return t.maximumHeatingTemp
}

func (t *Thermostat) GetMinimumHeatingTemp() int {
	return t.minimumHeatingTemp
}

func (t *Thermostat) GetMaximumCoolingTemp() int {
	return t.maximumCoolingTemp
}

func (t *Thermostat) GetMinimumCoolingTemp() int {
	return t.minimumCoolingTemp
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

func (t *Thermostat) addCommands() {
	var name string
	var temp float64

	temp = float64(t.minimumHeatingTemp)
	for temp <= float64(t.maximumHeatingTemp) {
		name = fmt.Sprintf("h%.1f", temp)
		t.commands = append(t.commands, command.New(name))
		temp += t.scale
		temp = round2ndDiminals(temp)
	}

	temp = float64(t.minimumCoolingTemp)
	for temp <= float64(t.maximumCoolingTemp) {
		name = fmt.Sprintf("c%.1f", temp)
		t.commands = append(t.commands, command.New(name))
		temp += t.scale
		temp = round2ndDiminals(temp)
	}

	t.commands = append(t.commands, command.New("off"))
}

func foor2ndDiminals(f float64) float64 {
	r := math.Floor(f*10) / 10
	return r
}

func round2ndDiminals(f float64) float64 {
	r := math.Round(f*10) / 10
	return r
}
