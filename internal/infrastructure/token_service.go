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

type TokenService struct {
	config                 *config.Config
	logger                 logger.ILogger
	accessTokenPrivateKey  *rsa.PrivateKey
	accessTokenPublicKey   *rsa.PublicKey
	refreshTokenPrivateKey *rsa.PrivateKey
	refreshTokenPublicKey  *rsa.PublicKey
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
	b, err := ioutil.ReadFile(s.config.Jwt.AccessTokenPrivateKeyFile)
	if err != nil {
		panic(err)
	}

	atprk, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.accessTokenPrivateKey = atprk

	// get public key from file
	b, err = ioutil.ReadFile(s.config.Jwt.AccessTokenPublicKeyFile)
	if err != nil {
		panic(err)
	}

	atpuk, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.accessTokenPublicKey = atpuk

	// get private key from file
	b, err = ioutil.ReadFile(s.config.Jwt.RefreshTokenPrivateKeyFile)
	if err != nil {
		panic(err)
	}

	rtprk, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.refreshTokenPrivateKey = rtprk

	// get public key from file
	b, err = ioutil.ReadFile(s.config.Jwt.RefreshTokenPublicKeyFile)
	if err != nil {
		panic(err)
	}

	rtpuk, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.refreshTokenPublicKey = rtpuk

	return s
}

// GenerateAccessToken generates a new access token
func (t *TokenService) GenerateAccessToken(ctx context.Context, user *model.User) (string, error) {
	sub := user.GetIdString()

	if sub == "" {
		return "", errors.New("user id is empty")
	}

	now := time.Now().UTC()

	claims := Claims{
		Role: user.GetRole(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.config.Jwt.Issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Duration(t.config.Jwt.AccessTokenExp) * time.Second).Unix(),
			Subject:   sub,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(t.accessTokenPrivateKey)
}

// ValidateAccessTokenFromRequest validates access token from request
func (t *TokenService) ValidateAccessTokenFromRequest(ctx context.Context, r *http.Request) (app.Claims, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("token is empty")
	}

	// remove bearer prefix
	token = token[7:]

	claims, err := t.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// GenerateRefreshToken generates a new refresh token
func (t *TokenService) GenerateRefreshToken(ctx context.Context, user *model.User) (string, error) {
	sub := user.GetIdString()

	if sub == "" {
		return "", errors.New("user id is empty")
	}

	now := time.Now().UTC()

	claims := jwt.StandardClaims{
		Issuer:    t.config.Jwt.Issuer,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Duration(t.config.Jwt.RefreshTokenExp) * time.Second).Unix(),
		Subject:   sub,
		Id:        "custom-id",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(t.refreshTokenPrivateKey)
	if err != nil {
		return "", err
	}

	//TODO add to redis

	return token, err
}

func (t *TokenService) ValidateRefreshToken(ctx context.Context, token string) (string, error) {
	c, err := jwt.Parse(token, t.provideRefreshTokenPublicKey)
	if err != nil {
		return "", err
	}

	claims, ok := c.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("claims is not jwt.MapClaims")
	}

	if !c.Valid {
		return "", errors.New("token is not valid")
	}

	if !claims.VerifyIssuer(t.config.Jwt.Issuer, true) {
		return "", errors.New("token issuer is not valid")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return "", errors.New("token expired")
	}

	sub := claims["sub"].(string)
	if sub == "" {
		return "", errors.New("subject is empty")
	}

	jti := claims["jti"].(string)
	if jti == "" {
		return "", errors.New("jti is empty")
	}

	//TODO check in redis

	return sub, nil
}

// ParseToken parses a token
func (t *TokenService) ParseToken(ctx context.Context, token string) (app.Claims, error) {
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

	if !claims.VerifyIssuer(t.config.Jwt.Issuer, true) {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}

// provideAccessTokenPublicKey provides access token public key to veriy token
func (t *TokenService) provideAccessTokenPublicKey(_ *jwt.Token) (interface{}, error) {
	return t.accessTokenPublicKey, nil
}

// provideRefreshTokenPublicKey provides access token public key to veriy token
func (t *TokenService) provideRefreshTokenPublicKey(_ *jwt.Token) (interface{}, error) {
	return t.refreshTokenPublicKey, nil
}
