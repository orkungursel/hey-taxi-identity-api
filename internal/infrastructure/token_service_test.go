package infrastructure

import (
	"context"
	"crypto/rsa"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	. "github.com/orkungursel/hey-taxi-identity-api/mock"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	issuer = "hey-taxi-identity-api-test"
)

func SetTokenServiceEnvForTesting(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	t.Setenv("AUTH_PRIVATE_KEY_FILE", filepath.Join(dir, "../../certs/private.pem"))
	t.Setenv("AUTH_PUBLIC_KEY_FILE", filepath.Join(dir, "../../certs/public.pem"))
	t.Setenv("AUTH_ISSUER", issuer)
}

func TestNewTokenService(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.NewConfig(), NewLoggerMock())

	if ts.accessTokenPrivateKey == nil {
		t.Errorf("privateKey is empty")
	}

	if ts.accessTokenPublicKey == nil {
		t.Errorf("publicKey is empty")
	}
}

func TestTokenService_GenerateAccessToken(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.NewConfig(), NewLoggerMock())

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should generate access token",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					UserID: primitive.NewObjectID(),
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when user id is empty",
			args: args{
				ctx:  context.Background(),
				user: &model.User{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.GenerateAccessToken(tt.args.ctx, tt.args.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.GenerateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got == "" {
				t.Errorf("TokenService.GenerateAccessToken() is empty")
			}
		})
	}
}

func TestTokenService_GenerateRefreshToken(t *testing.T) {
	type fields struct {
		TokenService          app.TokenService
		config                *config.Config
		logger                logger.ILogger
		accessTokenPrivateKey *rsa.PrivateKey
		accessTokenPublicKey  *rsa.PublicKey
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TokenService{
				TokenService:          tt.fields.TokenService,
				config:                tt.fields.config,
				logger:                tt.fields.logger,
				accessTokenPrivateKey: tt.fields.accessTokenPrivateKey,
				accessTokenPublicKey:  tt.fields.accessTokenPublicKey,
			}
			got, err := tr.GenerateRefreshToken(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.GenerateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TokenService.GenerateRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_ValidateAccessToken(t *testing.T) {
	type fields struct {
		TokenService          app.TokenService
		config                *config.Config
		logger                logger.ILogger
		accessTokenPrivateKey *rsa.PrivateKey
		accessTokenPublicKey  *rsa.PublicKey
	}
	type args struct {
		ctx   context.Context
		token string
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
			tr := &TokenService{
				TokenService:          tt.fields.TokenService,
				config:                tt.fields.config,
				logger:                tt.fields.logger,
				accessTokenPrivateKey: tt.fields.accessTokenPrivateKey,
				accessTokenPublicKey:  tt.fields.accessTokenPublicKey,
			}
			if err := tr.ValidateAccessToken(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("TokenService.ValidateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTokenService_ValidateRefreshToken(t *testing.T) {
	type fields struct {
		TokenService          app.TokenService
		config                *config.Config
		logger                logger.ILogger
		accessTokenPrivateKey *rsa.PrivateKey
		accessTokenPublicKey  *rsa.PublicKey
	}
	type args struct {
		ctx   context.Context
		token string
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
			tr := &TokenService{
				TokenService:          tt.fields.TokenService,
				config:                tt.fields.config,
				logger:                tt.fields.logger,
				accessTokenPrivateKey: tt.fields.accessTokenPrivateKey,
				accessTokenPublicKey:  tt.fields.accessTokenPublicKey,
			}
			if err := tr.ValidateRefreshToken(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("TokenService.ValidateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTokenService_ExtractFromRequest(t *testing.T) {
	type fields struct {
		TokenService          app.TokenService
		config                *config.Config
		logger                logger.ILogger
		accessTokenPrivateKey *rsa.PrivateKey
		accessTokenPublicKey  *rsa.PublicKey
	}
	type args struct {
		ctx context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TokenService{
				TokenService:          tt.fields.TokenService,
				config:                tt.fields.config,
				logger:                tt.fields.logger,
				accessTokenPrivateKey: tt.fields.accessTokenPrivateKey,
				accessTokenPublicKey:  tt.fields.accessTokenPublicKey,
			}
			got, err := tr.ExtractFromRequest(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.ExtractFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenService.ExtractFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_parseToken(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.NewConfig(), NewLoggerMock())

	u := &model.User{
		UserID: primitive.NewObjectID(),
	}
	validToken, err := ts.GenerateAccessToken(context.Background(), u)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to generate access token"))
	}

	invalidToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGkiLCJzdWIiOiI2MjM5MzgyOTUwYjMxODZkYjhkZDMxNDEifQ.D-9sK50mlj1-_PbHiN-3VsAnf2G-MF4w_JLRDTj8FbGoRDbM9UvvEuMuhldISYmm8m4YpPJ6j1U62cND0TNmVpE48q_d9bjcYVdKrWK_YHGvs4qu6-IdxoGNNzI03nY-A1M7J9yN9oxDjxq_CCm4Qm91I4clebEMaaD2Eozp5GutWNJdcaAdqkE7_g4yD7dUy5oOAbXLIMfgqXQCSyc1IISUuRf4W7_tfiXCoyQyDHwAOQ0EsVwUbim9C5HTX0oTd5q6N6yJlV0Tc_On-sigzrZp-EOWzvAN4kno4wGBeveeSju0TRQXPWirx07gHeq2Fiu_T8CNzmCTl2uhYg-77w"
	invalidSubjectToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdCIsInN1YiI6Im4vYSJ9.JbLkZI5o5P4BK90QVOFoZRd0hNwyoBV9H3ig_82SuEUlzrkmEtX7H8ewXrY89Dfpv1f7qoiaSVg7714r7ceevwWyaJPueqmgZwXlBj_XM_hl-gmGYbrp5gY9xIKaB4YRWZ_pYKBG-D6Znr8EXoTgPQBvD6CW-WfhGDYCkfP8Gdd876D9CDKatPaOWT7cnA_vkoY9yZKntxJwA2SBbmVuO8ctfBKPJBiCxdWse6DRZXoCzuhihXQQF9HRPZ2m3ZSDIxBEK_U6o3RYXnWkBabBt4QVKLtYj62X1HxBWx8yCa0FBcnlfzTJ_2Lt08dhKe7q2g0TptiFUb3aJwpdMGq1zQ"
	invalidIssuerToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdC0yIiwic3ViIjoiNjIzOTM4Mjk1MGIzMTg2ZGI4ZGQzMTQxIn0.d3KWJVJSUkogCWYyUOULCiGhXVCmWlIzm_OEL_Ubf4V1s8Hpi1x89eVyHpZgF-CgfRnNO8Nq7tV6xR0gCNe8Dti6NYO7-lMQoQBZqquurDR3B_Ynx7PPKbDS_BTHPXEkb0wJSserCY77TufT1PwZw7bAF6PnulLrtv7dRx-gh0lhjSpjsvDes3rwGPG5tR7mVx3_CHCOhnEQ7oXN4ioLQ5JwkT2BGnFZ00WWloBHHRnQ4RLQyNt22ptxiZNhbohTHLRgLEW0A1UWmgZOvUnDHkZXlq87HbL3_Y-sexbAxay3wEZHfTJk_87u97GFt0xbI1p1KrmpGj89QNDNPXe8rQ"
	expiredToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkxNzA5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdCIsInN1YiI6IjYyMzkzODI5NTBiMzE4NmRiOGRkMzE0MSJ9.WvTTWdrL_3DnF3BHtUfHpXzE1TdGOja5lAB99iYpYLADslse-epcZHk5VviooX5-yzMnNnxq_nmX3H3uswUgBCUSmhOmXsMBYIhQ5k7-U-A6ac2HExpG8gxMM7G-0zwNVx0eHhxMAewaLMABGHm6qnXJ9CNl5pLOipCIjZxT8-bu_ran7sBJaCTltGWic32lWd_k2pdra8Q3fGzPM6JR1EYy_DKlX6oC2uMqiyI5_AFNHqnra9bhrl5q-G2HBZHvNhJ3SpNmSJJDMN_QZxuvWCFCBsX8mOIFbWLravMUvdBErXHTpCNbePlYURccx9ZAYU7wpG80sewTG8XxzAdddg"

	type args struct {
		ctx   context.Context
		user  *model.User
		token string
	}

	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		{
			name: "should return claims with valid token",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: validToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.UserID.Hex(),
					Issuer:  issuer,
				},
			},
		},
		{
			name: "should fail because invalid token",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: invalidToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.UserID.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
		{
			name: "should fail because wrong subject",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: invalidSubjectToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.UserID.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
		{
			name: "should fail because wrong issuer",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: invalidIssuerToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.UserID.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
		{
			name: "should fail because token is expired",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: expiredToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.UserID.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.parseToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.parseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got == nil {
				t.Errorf("TokenService.parseToken() = %v, want %v", got, tt.want)
			} else {
				if got.Role != tt.want.Role {
					t.Errorf("TokenService.parseToken() = %v, want %v", got.Role, tt.want.Role)
				}
				if got.StandardClaims.Subject != tt.want.StandardClaims.Subject {
					t.Errorf("TokenService.parseToken() = %v, want %v", got.StandardClaims.Subject, tt.want.StandardClaims.Subject)
				}
				if got.StandardClaims.Issuer != tt.want.StandardClaims.Issuer {
					t.Errorf("TokenService.parseToken() = %v, want %v", got.StandardClaims.Issuer, tt.want.StandardClaims.Issuer)
				}
			}
		})
	}
}
