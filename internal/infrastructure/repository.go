package infrastructure

import (
	"context"
	"time"

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
		db: db.Database(config.Auth.DatabaseName, nil).
			Collection(config.Auth.CollectionName),
	}
}

// GetUser returns a user by id
func (r *Repository) GetUser(ctx context.Context, id string) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return nil, app.ErrInvalidUserId
	}

	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	u := &model.User{}
	if err := r.db.FindOne(ctx, bson.M{"_id": objectId}).Decode(u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, app.ErrUserNotFound
		}

		r.logger.Warnf("error while finding user: %s", err)
		return nil, errors.New("error while finding user")
	}

	return u, nil
}

// GetUsersByIds returns users by ids
func (r *Repository) GetUsersByIds(ctx context.Context, ids []string) ([]*model.User, error) {
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			r.logger.Warnf("invalid id: %s", id)
			continue
		}
		objectIds = append(objectIds, oid)
	}

	if len(objectIds) == 0 {
		return nil, nil
	}

	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$in": objectIds}}
	res, err := r.db.Find(ctx, filter)
	if err != nil {
		r.logger.Warnf("error while getting users: %s", err)
		return nil, errors.New("error while finding users")
	}

	var users []*model.User
	if err := res.All(ctx, &users); err != nil {
		r.logger.Warnf("error while decoding user: %s", err)
		return nil, errors.New("error while decoding users")
	}

	return users, nil
}

// GetUserByEmail returns a user by email
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	if email == "" {
		r.logger.Warn("empty email")
		return nil, errors.New("empty email")
	}

	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	u := &model.User{}
	filter := bson.M{"email": email}

	if err := r.db.FindOne(ctx, filter).Decode(u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, app.ErrUserNotFound
		}

		r.logger.Warnf("error while finding user by email: %s", err)
		return nil, app.NewInternalServerError(err)
	}

	return u, nil
}

// CreateUser creates a new user
func (r *Repository) CreateUser(ctx context.Context, user *model.User) (string, error) {
	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	res, err := r.db.InsertOne(ctx, user)
	if err != nil {
		r.logger.Warnf("error while creating user: %s", err)
		return "", app.NewInternalServerError(err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *Repository) UpdateUser(ctx context.Context, id string, user *model.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return app.ErrInvalidUserId
	}

	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	result, err := r.db.UpdateOne(ctx, bson.M{"_id": objectId}, user)
	if err != nil {
		r.logger.Warnf("error while updating user: %s", err)
		return app.NewInternalServerError(errors.New("error while updating user"))
	}

	if result.MatchedCount == 0 {
		r.logger.Warnf("user not found to update: %s", id)
		return app.ErrUserNotFound
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Warnf("invalid id: %s", id)
		return app.ErrInvalidUserId
	}

	ctx, cancel := r.contextWithTimeout(ctx)
	defer cancel()

	result, err := r.db.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		r.logger.Warnf("error while deleting user: %s", err)
		return app.NewInternalServerError(errors.New("error while deleting user"))

	}

	if result.DeletedCount == 0 {
		r.logger.Warnf("user not found to delete: %s", id)
		return app.ErrUserNotFound
	}

	return nil
}

func (r *Repository) contextWithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(r.config.Mongo.SocketTimeout)*time.Second)
}
