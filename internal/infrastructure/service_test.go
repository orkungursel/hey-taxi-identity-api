package infrastructure

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app/mock"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	. "github.com/orkungursel/hey-taxi-identity-api/mock"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewService(t *testing.T) {
	type args struct {
		config *config.Config
		logger logger.ILogger
		repo   app.Repository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.config, tt.args.logger, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	t.Parallel()

	config := config.NewConfig()
	logger := NewLoggerMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepository(ctrl)
	service := NewService(config, logger, repo)

	u := &model.User{
		UserID:   primitive.NewObjectID(),
		Email:    "foo@bar.com",
		Password: "password",
	}
	ctx := context.Background()
	repo.EXPECT().GetUserByEmail(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, email string) (*model.User, error) {
			if email == u.Email {
				return u, nil
			}
			return nil, errors.New("not found")
		}).AnyTimes()

	type args struct {
		ctx context.Context
		req *app.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		svc     app.Service
		want    *app.LoginResponse
		wantErr bool
	}{
		{
			name: "should return token",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: u.Email, Password: u.Password},
			},
			want: &app.LoginResponse{
				Token:        "token",
				RefreshToken: "refresh_token",
				ExpiresIn:    config.Auth.AccessTokenExp,
			},
		},
		{
			name: "should return error when password is wrong",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: u.Email, Password: "123456"},
			},
			wantErr: true,
		},
		{
			name: "should return error when email is wrong",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: "foo@bar", Password: ""},
			},
			wantErr: true,
		},
		{
			name: "should return error when user not found",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: "foo2@bar.com", Password: "123456"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := config.NewConfig()
	logger := NewLoggerMock()

	repo := mock.NewMockRepository(ctrl)
	service := NewService(config, logger, repo)

	ctx := context.Background()
	repo.EXPECT().CreateUser(ctx, gomock.AssignableToTypeOf(&model.User{})).Return(nil).Times(1)

	type args struct {
		ctx context.Context
		req *app.RegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		svc     app.Service
		want    *app.LoginResponse
		wantErr bool
	}{
		{
			name: "register",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: "foo@bar.com", Password: "123456"},
			},
			want: &app.LoginResponse{
				Token:        "token",
				RefreshToken: "refresh_token",
				ExpiresIn:    config.Auth.AccessTokenExp,
			},
		},
		{
			name: "register",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: "foo@bar", Password: "123456"},
			},
			wantErr: true,
		},
		{
			svc:  service,
			name: "register",
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: "foo@bar.com"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
