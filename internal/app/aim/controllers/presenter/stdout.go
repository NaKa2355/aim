package presenter

import (
	"context"
	"fmt"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

type StdOut struct{}

func (o StdOut) AddCustom(ctx context.Context, out bdy.AddAppOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}

func (o StdOut) AddToggle(ctx context.Context, out bdy.AddAppOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}

func (o StdOut) AddButton(ctx context.Context, out bdy.AddAppOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}

func (o StdOut) AddThermostat(ctx context.Context, out bdy.AddAppOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}

func (o StdOut) AddCommand(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

// Read
func (o StdOut) GetCustom(ctx context.Context, out bdy.GetCustomOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}
func (o StdOut) GetToggle(ctx context.Context, out bdy.GetToggleOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}
func (o StdOut) GetButton(ctx context.Context, out bdy.GetButtonOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}
func (o StdOut) GetThermostat(ctx context.Context, out bdy.GetThermostatOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}
func (o StdOut) GetAppliances(ctx context.Context, out bdy.GetAppliancesOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}
func (o StdOut) GetCommand(ctx context.Context, out bdy.GetCommandOutput, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)
}

// Update
func (o StdOut) RenameAppliance(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
func (o StdOut) ChangeIRDevice(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
func (o StdOut) RenameCommand(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
func (o StdOut) SetIRData(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}

// Delete
func (o StdOut) DeleteAppliance(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
func (o StdOut) DeleteCommand(ctx context.Context, err error) {
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
