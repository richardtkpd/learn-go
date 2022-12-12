package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_HelloWorld(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "success",
			args: args{
				name: "Richard",
			},
			wantErr: false,
			want:    "Hello world, Richard",
		},
		{
			name: "fail",
			args: args{
				name: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HelloWorld(tt.args.name)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
