package infrastructure

import (
	"context"
	"reflect"
	"testing"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		config *config.Config
		logger logger.ILogger
		db     *mongo.Client
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.config, tt.args.logger, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetUser(t *testing.T) {
	type fields struct {
		Repository app.Repository
		config     *config.Config
		logger     logger.ILogger
		db         *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Repository: tt.fields.Repository,
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				db:         tt.fields.db,
			}
			got, err := r.GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetUserByEmail(t *testing.T) {
	type fields struct {
		Repository app.Repository
		config     *config.Config
		logger     logger.ILogger
		db         *mongo.Collection
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Repository: tt.fields.Repository,
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				db:         tt.fields.db,
			}
			got, err := r.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreateUser(t *testing.T) {
	type fields struct {
		Repository app.Repository
		config     *config.Config
		logger     logger.ILogger
		db         *mongo.Collection
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Repository: tt.fields.Repository,
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				db:         tt.fields.db,
			}
			if err := r.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateUser(t *testing.T) {
	type fields struct {
		Repository app.Repository
		config     *config.Config
		logger     logger.ILogger
		db         *mongo.Collection
	}
	type args struct {
		ctx  context.Context
		id   string
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Repository: tt.fields.Repository,
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				db:         tt.fields.db,
			}
			if err := r.UpdateUser(tt.args.ctx, tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_DeleteUser(t *testing.T) {
	type fields struct {
		Repository app.Repository
		config     *config.Config
		logger     logger.ILogger
		db         *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Repository: tt.fields.Repository,
				config:     tt.fields.config,
				logger:     tt.fields.logger,
				db:         tt.fields.db,
			}
			if err := r.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repository.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
