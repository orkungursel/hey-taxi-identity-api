package infrastructure

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app/mock"
	. "github.com/orkungursel/hey-taxi-identity-api/internal/app/mock"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	. "github.com/orkungursel/hey-taxi-identity-api/pkg/logger/mock"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	dummyAuthUser = &model.User{
		Id:       primitive.NewObjectID(),
		Email:    "foo@bar.com",
		Password: "password",
	}
	dummyAuthUser2 = &model.User{
		Id:       primitive.NewObjectID(),
		Email:    "foo2@bar.com",
		Password: "123456",
	}
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := NewMockTokenService(ctrl)
	pws := NewMockPasswordService(ctrl)

	type args struct {
		config *config.Config
		logger logger.ILogger
		repo   app.Repository
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.config, tt.args.logger, tt.args.repo, ts, pws); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	config := config.New()
	logger := NewLoggerMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepository(ctrl)
	ts := NewMockTokenService(ctrl)
	pws := NewMockPasswordService(ctrl)

	service := NewAuthService(config, logger, repo, ts, pws)
	ctx := context.Background()

	repo.EXPECT().GetUserByEmail(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, email string) (*model.User, error) {
			if email == dummyAuthUser.Email {
				return dummyAuthUser, nil
			}
			return nil, errors.New("not found")
		}).AnyTimes()

	ts.EXPECT().GenerateAccessToken(ctx, dummyAuthUser).
		Return("access_token", nil).AnyTimes().MinTimes(1)

	ts.EXPECT().GenerateRefreshToken(ctx, dummyAuthUser).
		Return("refresh_token", nil).AnyTimes().MinTimes(1)

	pws.EXPECT().Compare(ctx, dummyAuthUser.Password, gomock.Any()).
		DoAndReturn(func(_ context.Context, hashedPassword string, password string) error {
			if password == hashedPassword {
				return nil
			}
			return errors.New("not match")
		}).AnyTimes()

	type args struct {
		ctx context.Context
		req *app.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		svc     app.AuthService
		want    *app.SuccessAuthResponse
		wantErr bool
	}{
		{
			name: "should return token",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: dummyAuthUser.Email, Password: dummyAuthUser.Password},
			},
			want: &app.SuccessAuthResponse{
				UserDto:               *app.UserResponseFromUser(dummyAuthUser),
				AccessToken:           "access_token",
				RefreshToken:          "refresh_token",
				AccessTokenExpiresIn:  config.Jwt.AccessTokenExp,
				RefreshTokenExpiresIn: config.Jwt.RefreshTokenExp,
			},
		},
		{
			name: "should return error when password is wrong",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.LoginRequest{Email: dummyAuthUser.Email, Password: "123456"},
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
				req: &app.LoginRequest{Email: "foo3@bar.com", Password: "123456"},
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

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := config.New()
	logger := NewLoggerMock()

	repo := mock.NewMockRepository(ctrl)
	ts := NewMockTokenService(ctrl)
	pws := NewMockPasswordService(ctrl)

	service := NewAuthService(config, logger, repo, ts, pws)

	ctx := context.Background()
	repo.EXPECT().CreateUser(ctx, gomock.AssignableToTypeOf(&model.User{})).
		Return(dummyAuthUser2.Id.Hex(), nil).Times(1)

	repo.EXPECT().GetUserByEmail(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, email string) (*model.User, error) {
			if email == dummyAuthUser.Email {
				return dummyAuthUser, nil
			}
			return nil, errors.New("not found")
		}).AnyTimes()

	ts.EXPECT().GenerateAccessToken(ctx, gomock.Any()).
		Return("access_token", nil).AnyTimes().MinTimes(1)

	ts.EXPECT().GenerateRefreshToken(ctx, gomock.Any()).
		Return("refresh_token", nil).AnyTimes().MinTimes(1)

	pws.EXPECT().Hash(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, password string) (string, error) {
			return password, nil
		}).AnyTimes().MinTimes(1)

	type args struct {
		ctx context.Context
		req *app.RegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		svc     app.AuthService
		want    *app.SuccessAuthResponse
		wantErr bool
	}{
		{
			name: "should return token",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: dummyAuthUser2.Email, Password: dummyAuthUser2.Password},
			},
			want: &app.SuccessAuthResponse{
				UserDto:               *app.UserResponseFromUser(dummyAuthUser2),
				AccessToken:           "access_token",
				RefreshToken:          "refresh_token",
				AccessTokenExpiresIn:  config.Jwt.AccessTokenExp,
				RefreshTokenExpiresIn: config.Jwt.RefreshTokenExp,
			},
		},
		{
			name: "should error when email is already registered",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: dummyAuthUser.Email, Password: "123456"},
			},
			wantErr: true,
		},
		{
			name: "should error when email is empty",
			svc:  service,
			args: args{
				ctx: ctx,
				req: &app.RegisterRequest{Email: "foo@bar", Password: "123456"},
			},
			wantErr: true,
		},
		{
			name: "should error when password is empty",
			svc:  service,
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

func TestAuthService_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := config.New()
	logger := NewLoggerMock()

	repo := mock.NewMockRepository(ctrl)
	ts := NewMockTokenService(ctrl)
	pws := NewMockPasswordService(ctrl)

	service := NewAuthService(config, logger, repo, ts, pws)

	ctx := context.Background()
	repo.EXPECT().GetUser(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, uid string) (*model.User, error) {
			if uid == dummyAuthUser.GetIdString() {
				return dummyAuthUser, nil
			}
			return nil, errors.New("not found")
		}).AnyTimes().MinTimes(1)

	type args struct {
		ctx context.Context
		req string
	}
	tests := []struct {
		name    string
		args    args
		svc     app.AuthService
		want    *app.UserResponse
		wantErr bool
	}{
		{
			name: "should return user",
			svc:  service,
			args: args{
				ctx: ctx,
				req: dummyAuthUser.GetIdString(),
			},
			want: app.UserResponseFromUser(dummyAuthUser),
		},
		{
			name: "should error when user not found",
			svc:  service,
			args: args{
				ctx: ctx,
				req: dummyAuthUser2.GetIdString(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.Me(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Me() = %v, want %v", got, tt.want)
			}
		})
	}
}
