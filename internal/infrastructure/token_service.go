package infrastructure

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/pkg/errors"
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

type TokenService struct {
	app.TokenService
	config                *config.Config
	logger                logger.ILogger
	accessTokenPrivateKey *rsa.PrivateKey
	accessTokenPublicKey  *rsa.PublicKey
}

func NewTokenService(config *config.Config, logger logger.ILogger) (s *TokenService) {
	s = &TokenService{
		config: config,
		logger: logger,
	}

	s.init()

	return
}

func (s *TokenService) init() *TokenService {
	// get private key from file
	b, err := ioutil.ReadFile(s.config.Auth.PrivateKeyFile)
	if err != nil {
		panic(err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	if privKey == nil {
		panic("private key is nil")
	}

	s.accessTokenPrivateKey = privKey

	// get public key from file
	b, err = ioutil.ReadFile(s.config.Auth.PublicKeyFile)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	if publicKey == nil {
		panic("private key is nil")
	}

	s.accessTokenPublicKey = publicKey

	return s
}

// GenerateAccessToken generates a new access token
func (t *TokenService) GenerateAccessToken(ctx context.Context, user *model.User) (string, error) {
	sub := user.GetUserIDString()

	if sub == "" {
		return "", errors.New("user id is empty")
	}

	now := time.Now().UTC()

	claims := Claims{
		Role: user.GetRole(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.config.Auth.Issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Duration(t.config.Auth.AccessTokenExp) * time.Second).Unix(),
			Subject:   sub,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(t.accessTokenPrivateKey)
	return token, err
}

// GenerateRefreshToken generates a new refresh token
func (t *TokenService) GenerateRefreshToken(ctx context.Context, user *model.User) (string, error) {
	sub := user.GetUserIDString()

	if sub == "" {
		return "", errors.New("user id is empty")
	}

	now := time.Now().UTC()

	claims := Claims{
		Role: user.GetRole(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.config.Auth.Issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Duration(t.config.Auth.RefreshTokenExp) * time.Second).Unix(),
			Subject:   sub,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(t.accessTokenPrivateKey)
	return token, err
}

// ValidateAccessToken validates an access token
func (t *TokenService) ValidateAccessToken(ctx context.Context, token string) error {
	claims, err := t.parseToken(ctx, token)
	if err != nil {
		return err
	}

	if !claims.VerifyIssuer(t.config.Auth.Issuer, true) {
		return errors.New("invalid issuer")
	}

	if !claims.VerifyExpiresAt(time.Now().UTC().Unix(), true) {
		return errors.New("token expired")
	}

	return nil
}

// ValidateRefreshToken validates a refresh token
func (t *TokenService) ValidateRefreshToken(ctx context.Context, token string) error {
	panic("not implemented") // TODO: Implement
}

// ExtractToken extracts token from http request
func (t *TokenService) ExtractFromRequest(ctx context.Context, r *http.Request) (map[string]interface{}, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("token is empty")
	}

	// remove bearer prefix
	token = token[7:]

	claims, err := t.parseToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id": claims.Subject,
		"role":    claims.Role,
	}, nil
}

// parseToken parses a token
func (t *TokenService) parseToken(ctx context.Context, token string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, t.provideAccessTokenPublicKey)

	if err != nil {
		return nil, err
	}

	if tkn.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return nil, errors.New("invalid algorithm")
	}

	if !tkn.Valid {
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}

// provideAccessTokenPublicKey provides access token public key to veriy token
func (t *TokenService) provideAccessTokenPublicKey(token *jwt.Token) (interface{}, error) {
	return t.accessTokenPublicKey, nil
}
