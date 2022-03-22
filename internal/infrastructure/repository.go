package infrastructure

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	app.Repository
	config *config.Config
	logger logger.ILogger
	db     *mongo.Collection
}

func NewRepository(config *config.Config, logger logger.ILogger, db *mongo.Client) *Repository {
	return &Repository{
		config: config,
		logger: logger,
		db:     db.Database(config.Auth.DatabaseName, nil).Collection("users"),
	}
}

// GetUser returns a user by id
func (r *Repository) GetUser(ctx context.Context, id string) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return nil, errors.Wrap(err, "invalid id")
	}

	res := r.db.FindOne(ctx, bson.M{"_id": objectId}, nil)
	if err := res.Err(); err != nil {
		r.logger.Warnf("error while getting user: %s", err)
		return nil, errors.Wrap(err, "error while finding user")
	}

	u := &model.User{}
	if err := res.Decode(u); err != nil {
		r.logger.Warnf("error while decoding user: %s", err)
		return nil, err
	}

	return u, nil
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	if email == "" {
		r.logger.Warn("empty email")
		return nil, errors.New("empty email")
	}

	res := r.db.FindOne(ctx, bson.M{"email": email}, nil)
	if err := res.Err(); err != nil {
		r.logger.Warnf("error while getting user: %s", err)
		return nil, errors.Wrap(err, "error while finding user")
	}

	u := &model.User{}
	if err := res.Decode(u); err != nil {
		r.logger.Warnf("error while decoding user: %s", err)
		return nil, err
	}

	return u, nil
}

// CreateUser creates a new user
func (r *Repository) CreateUser(ctx context.Context, user *model.User) error {
	if _, err := r.db.InsertOne(ctx, user); err != nil {
		r.logger.Warnf("error while creating user: %s", err)
		return errors.Wrap(err, "error while creating user")
	}

	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, id string, user *model.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return errors.Wrap(err, "invalid id")
	}

	if _, err := r.db.UpdateOne(ctx, bson.M{"_id": objectId}, user); err != nil {
		r.logger.Warnf("error while updating user: %s", err)
		return errors.Wrap(err, "error while updating user")
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return errors.Wrap(err, "invalid id")
	}

	if _, err := r.db.DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
		r.logger.Warnf("error while deleting user: %s", err)
		return errors.Wrap(err, "error while deleting user")
	}

	return nil
}
