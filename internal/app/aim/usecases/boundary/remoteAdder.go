package boundary

import "context"

type AddRemoteOutput struct {
	Remote Remote
}

type AddRemoteInput struct {
	Name     string
	DeviceID string
	Tag      string
	Buttons  []struct {
		Name string
		Tag  string
	}
}

type RemoteAdder interface {
	AddRemote(ctx context.Context, i AddRemoteInput) (AddRemoteOutput, error)
}
