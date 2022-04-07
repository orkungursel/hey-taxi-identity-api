package infrastructure

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app/mock"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	. "github.com/orkungursel/hey-taxi-identity-api/pkg/logger/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	dummyUser = &model.User{
		Id:       primitive.NewObjectID(),
		Email:    "foo@bar.com",
		Password: "password",
	}
	dummyUser2 = &model.User{
		Id:       primitive.NewObjectID(),
		Email:    "bar@baz.com",
		Password: "123456",
	}
)

func TestNewUserService(t *testing.T) {
	type args struct {
		config *config.Config
		logger logger.ILogger
		repo   app.Repository
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.config, tt.args.logger, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := config.New()
	logger := NewLoggerMock()

	repo := mock.NewMockRepository(ctrl)
	service := NewUserService(config, logger, repo)

	ctx := context.Background()
	repo.EXPECT().GetUsersByIds(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, ids []string) ([]*model.User, error) {
			if len(ids) == 0 {
				return nil, nil
			}

			var users []*model.User
			var addedIds = make(map[string]bool)

			for _, id := range ids {
				if _, ok := addedIds[id]; ok {
					continue
				}

				if id == dummyUser.GetIdString() {
					addedIds[id] = true
					users = append(users, dummyUser)
				} else if id == dummyUser2.GetIdString() {
					addedIds[id] = true
					users = append(users, dummyUser2)
				}
			}

			return users, nil
		}).AnyTimes().MinTimes(1)

	type args struct {
		ctx context.Context
		ids []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*app.UserResponse
		wantErr bool
	}{
		{
			name: "should return users",
			args: args{
				ctx: ctx,
				ids: []string{dummyUser.GetIdString(), dummyUser2.GetIdString()},
			},
			want: []*app.UserResponse{
				app.UserResponseFromUser(dummyUser),
				app.UserResponseFromUser(dummyUser2),
			},
		},
		{
			name: "should return users when duplicate ids",
			args: args{
				ctx: ctx,
				ids: []string{dummyUser.GetIdString(), dummyUser.GetIdString(), dummyUser2.GetIdString()},
			},
			want: []*app.UserResponse{
				app.UserResponseFromUser(dummyUser),
				app.UserResponseFromUser(dummyUser2),
			},
		},
		{
			name: "should return empty when no users found",
			args: args{
				ctx: ctx,
				ids: []string{"not-found-user-id"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.UsersByIds(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UsersByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UsersByIds() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
