package appliance

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewButton(t *testing.T) {
	type args struct {
		id       ID
		name     Name
		deviceID DeviceID
	}
	tests := []struct {
		name string
		args args
		want Button
	}{
		// TODO: Add test cases.
	}
	b := NewButton("", "hello", "test")
	b.SetID("ok")
	fmt.Println(b.ID)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewButton(tt.args.id, tt.args.name, tt.args.deviceID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewButton() = %v, want %v", got, tt.want)
			}
		})
	}
}
