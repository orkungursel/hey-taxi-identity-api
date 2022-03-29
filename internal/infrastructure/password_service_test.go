//go:generate mockgen -source password_service.go -destination mock/password_service_mock.go -package mock
package infrastructure

import (
	"context"
	"testing"

	. "github.com/orkungursel/hey-taxi-identity-api/pkg/logger/mock"
)

func TestPasswordService_HashAndCompare(t *testing.T) {
	ps := NewPasswordService(NewLoggerMock())

	type args struct {
		ctx      context.Context
		password string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should generate hashed password and compare it",
			args: args{
				ctx:      context.Background(),
				password: "password",
			},
			want:    "password",
			wantErr: false,
		},
		{
			name: "should generate hashed password and compare when password has spaces at the beginning and at the end",
			args: args{
				ctx:      context.Background(),
				password: "password",
			},
			want:    " password ",
			wantErr: false,
		},
		{
			name: "should error when password is empty",
			args: args{
				ctx:      context.Background(),
				password: "",
			},
			wantErr: true,
		},
		{
			name: "should error when password contains spaces only",
			args: args{
				ctx:      context.Background(),
				password: "  ",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ps.Hash(tt.args.ctx, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordService.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if ps.Compare(context.Background(), got, tt.args.password) != nil {
					t.Errorf("PasswordService.Hash() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
