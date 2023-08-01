package interactor

import (
	"context"

	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) getRemotes(ctx context.Context) (out bdy.GetRemotesOutput, err error) {
	var remotes []*remote.Remote

	remotes, err = i.repo.ReadRemotes(ctx)
	if err != nil {
		return
	}

	out.Remotes = make([]bdy.Remote, len(remotes))
	for i, r := range remotes {
		out.Remotes[i] = bdy.Remote{
			ID:           string(r.ID),
			DeviceID:     string(r.DeviceID),
			Type:         convertType(r.Type),
			Name:         string(r.Name),
			CanAddButton: (r.AddButton() == nil),
		}
	}
	return
}

func (i *Interactor) GetRemotes(ctx context.Context) (bdy.GetRemotesOutput, error) {
	out, err := i.getRemotes(ctx)
	return out, wrapErr(err)
}
