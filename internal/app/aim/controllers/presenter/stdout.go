package presenter

import (
	"context"
	"errors"
	"fmt"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type StdOut struct{}

func (o StdOut) AddCustom(ctx context.Context, out bdy.AddAppOutput, err error) {
	fmt.Println(out, " ", err)
}

func (o StdOut) AddToggle(ctx context.Context, out bdy.AddAppOutput, err error) {
	fmt.Println(out, " ", err)
}

func (o StdOut) AddButton(ctx context.Context, out bdy.AddAppOutput, err error) {
	fmt.Println(errors.Is(err, repository.ErrInvaildArgs))
	fmt.Println(out, " ", err)
}

func (o StdOut) AddThermostat(ctx context.Context, out bdy.AddAppOutput, err error) {
	fmt.Println(out, " ", err)
}

func (o StdOut) AddCommand(ctx context.Context, err error) {
	fmt.Println(err)
}

// Read
func (o StdOut) GetCustom(ctx context.Context, out bdy.GetCustomOutput, err error) {
	fmt.Println(out, " ", err)
}
func (o StdOut) GetToggle(ctx context.Context, out bdy.GetToggleOutput, err error) {
	fmt.Println(out, " ", err)
}
func (o StdOut) GetButton(ctx context.Context, out bdy.GetButtonOutput, err error) {
	fmt.Println(out, "", err)
}
func (o StdOut) GetThermostat(ctx context.Context, out bdy.GetThermostatOutput, err error) {
	fmt.Println(out, " ", err)
}
func (o StdOut) GetAppliances(ctx context.Context, out bdy.GetAppliancesOutput, err error) {
	fmt.Println(out, " ", err)
}
func (o StdOut) GetCommand(ctx context.Context, out bdy.GetCommandOutput, err error) {
	fmt.Println(out, " ", err)
}

// Update
func (o StdOut) RenameAppliance(ctx context.Context, err error) {
	fmt.Println(err)
}
func (o StdOut) ChangeIRDevice(ctx context.Context, err error) {
	fmt.Println(err)
}
func (o StdOut) RenameCommand(ctx context.Context, err error) {
	fmt.Println(err)
}
func (o StdOut) SetRawIRData(ctx context.Context, err error) {
	fmt.Println(err)
}

// Delete
func (o StdOut) DeleteAppliance(ctx context.Context, err error) {
	fmt.Println(err)
}
func (o StdOut) DeleteCommand(ctx context.Context, err error) {
	fmt.Println(err)
}
