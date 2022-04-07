package infrastructure

import (
	"context"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	. "github.com/orkungursel/hey-taxi-identity-api/pkg/logger/mock"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	issuer                    = "hey-taxi-identity-api-test"
	invalidToken              = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGkiLCJzdWIiOiI2MjM5MzgyOTUwYjMxODZkYjhkZDMxNDEifQ.D-9sK50mlj1-_PbHiN-3VsAnf2G-MF4w_JLRDTj8FbGoRDbM9UvvEuMuhldISYmm8m4YpPJ6j1U62cND0TNmVpE48q_d9bjcYVdKrWK_YHGvs4qu6-IdxoGNNzI03nY-A1M7J9yN9oxDjxq_CCm4Qm91I4clebEMaaD2Eozp5GutWNJdcaAdqkE7_g4yD7dUy5oOAbXLIMfgqXQCSyc1IISUuRf4W7_tfiXCoyQyDHwAOQ0EsVwUbim9C5HTX0oTd5q6N6yJlV0Tc_On-sigzrZp-EOWzvAN4kno4wGBeveeSju0TRQXPWirx07gHeq2Fiu_T8CNzmCTl2uhYg-77w"
	invalidSAlgorithmToken    = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MzY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdC0yIiwic3ViIjoiNjIzOTM4Mjk1MGIzMTg2ZGI4ZGQzMTQxIn0.Nt4mV8u2ByiqSXrQbLaFwoejyDfmq9smUVPlFymrixIEmZclVx9oPMepUkuI7YaJXRpMpQEJm5CK6LPxryrRysaKl4RElBht5LRA-3NNzIkIN05TJOGFbBQdeVYAHEBg8E2SmCh_TUwf7bJGN3R-L3u-hPYeQkZZee4lyt2JvwVKeH4FzStxRY4UUV6cSOX71tgUGig05rzyxYeFXwFOc_qrZF55l_v0pkTGEx848-cp_yg00Ihlw15Z1SMC_FhYMEDPErKvBIvsVG4nLf2lLon-MA_on6JyCxbtsadn2kyWgLsZzAbi7rN_bYNikENy3xCuoOqPV5gLCIKGjHc6GQ"
	invalidSubjectToken       = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MzY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdCIsInN1YiI6IjYyMzkzODI5NTBiMzE4NmRiOGRkMzE0MiJ9.jtrJWDKrYhqX1SiT2vIJmOGC_ucGzhEaYSusIW3uo1ThazGCmOuJdqyLG6bG4N18H4TLrd40507FF25LFsJC62KnNCngkZ6deMKPUI_ksSlO9RSP1b57sb86u4l5W-TRX9DjvAV-OtRBWQqEfLyfuMqGzllSir7Kj0g05MLuLnyn5nJKU1mJ6jtpBl2Qw2irvkGA124rg3RFMBcW-SGksx-3FKRXu9v2gREgQ1pqF6eu1mQbAEiYaW5yjLwTGwvkxk2wHkpwEDVIcI4OoAXQy_dCtuKo3jCNTFwOVcYxf7IWrP3tEpvSrJgxEh4u9paGsftSDnJFIx2DMrPqoH_i6Q"
	invalidIssuerToken        = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MzY0NzkyMDY5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdC0yIiwic3ViIjoiNjIzOTM4Mjk1MGIzMTg2ZGI4ZGQzMTQxIn0.EUE1dus4Aph6se9lmXd0MSX4DJEHF17-T6WCzssbyDPwUpqjZAGuEZEnGbalv7x2nBqxuGAL8xRGReJs0nvii9F1lU9ArNXNTrkJXRgLd6nOllopHxTXSJK5WiaJt7GboHqw87EfzhCQAvapap47WskbEwMhTkLO76Q0AxnyZjJuFTKJdms6sqNHe6SI3oa0daW73dmLWzM4-zzz_hsPb5f2NBbb-jIF2YVFSubtyy05fd0rfArymPtS3brB5NxKm8TT5zLlbwvUkry2dABUofs_pN5R961E09MW4MSx1U5SNwxsvPAyNmsETwUX4hzAl6YciUh3jvfN5U74tfDSzg"
	expiredToken              = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImV4cCI6MTY0NzkxNzA5NywiaWF0IjoxNjQ3OTE3MDk3LCJpc3MiOiJoZXktdGF4aS1pZGVudGl0eS1hcGktdGVzdCIsInN1YiI6IjYyMzkzODI5NTBiMzE4NmRiOGRkMzE0MSJ9.WvTTWdrL_3DnF3BHtUfHpXzE1TdGOja5lAB99iYpYLADslse-epcZHk5VviooX5-yzMnNnxq_nmX3H3uswUgBCUSmhOmXsMBYIhQ5k7-U-A6ac2HExpG8gxMM7G-0zwNVx0eHhxMAewaLMABGHm6qnXJ9CNl5pLOipCIjZxT8-bu_ran7sBJaCTltGWic32lWd_k2pdra8Q3fGzPM6JR1EYy_DKlX6oC2uMqiyI5_AFNHqnra9bhrl5q-G2HBZHvNhJ3SpNmSJJDMN_QZxuvWCFCBsX8mOIFbWLravMUvdBErXHTpCNbePlYURccx9ZAYU7wpG80sewTG8XxzAdddg"
	validRefreshToken         = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI2NDg1MTk2ODMsImp0aSI6ImN1c3RvbS1pZCIsImlhdCI6MTY0ODQzMzI4MywiaXNzIjoiaGV5LXRheGktaWRlbnRpdHktYXBpLXRlc3QiLCJzdWIiOiI2MjNiM2QyYzJkYTllNTBhNzQ4NmRlN2YifQ.sJgEbJryAY4pKGFTDKtCN09Q3eH6vgalqHVcw13jcOKTrBAwVQ3_-50PNNedafdIkbFOWHWCdjGd_afGPcc5nasZw1cax3l3Fmev522NQ-OubfR5s2y_MVQjwBdAY9LhIqWLpb9_BgnFXcRJFkF0JAuU0aq17ssMzwUy7mePD68I21VvdrEqfG4hbmw15T-ZPsLwWXVCzkki6FtZp3nCP8h7QSMC4Sk0UqizDC3esfWDCH6o6VdHjMTESGbxMA8UVCZL9lljlQp5DG0g_Eol4ZS5NPdt55DUtE5XH8ZimCUnKOoCptUdqRsNNu5YTd0vyuzAWmRC6iqoPKuqtGDAeg"
	expiredRefreshToken       = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDg0MzMyODMsImp0aSI6ImN1c3RvbS1pZCIsImlhdCI6MTY0ODQzMzI4MywiaXNzIjoiaGV5LXRheGktaWRlbnRpdHktYXBpLXRlc3QiLCJzdWIiOiI2MjNiM2QyYzJkYTllNTBhNzQ4NmRlN2YifQ.hsLVMLAV5aB7hrQxBSsodsgPxTcyhevIgXou5dGDXRsWcbgJx5ck32RZpGR5ceeK0ekewKF9PZ5APwNtKN1xfn2ED7noG0OBRtcKiUz4XPzo4Ynyun0EXbYedhewXTL6teug0cxHFNlY5_-samoz4ypomvKRzdxdwwO4nyduCOcEtDMxGD7HV-7Jw7H_XspvvsoiIaB4yZBnYB43NcTgYmiKDvAKy6xhgi7jyom5Fyp7BPtXx6FDm3PpYW--JIVz0JL-HuuDjpOc0dlSweMjxt6-9zis6o0m-5NaHdebSXEQwNtjygE7_6uSS7nWvJr-Mwo3JlXjKmvnvCb-p235xA"
	invalidIssuerRefreshToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI2NDg1MTk2ODMsImp0aSI6ImN1c3RvbS1pZCIsImlhdCI6MTY0ODQzMzI4MywiaXNzIjoiaGV5LXRheGktaWRlbnRpdHktYXBpIiwic3ViIjoiNjIzYjNkMmMyZGE5ZTUwYTc0ODZkZTdmIn0.DJBYpmrpRjzQc2uhYYYGU1e7RW8wBsAnUvKxiv52fOT-jg0L7TaQWBiLgnlFTznIqqK4EUi6OnyWb0iMRuNFIYEeOw-u-BWCgLw2-yprQ4ZUUGw5RRT7p9gWpqnPbWhrdKc-lSKkg5MRyo520e7ljPWaiArLE1JtdNnXbLCXKgsiWFjRmLnWJiYnZwUbP0tDwnJPUa2-BQy50M4oWILioHwgDN5F7jN2meioqeRhV1NH6QezkU2U-mNOFVaY8SN5_3mm6Om872AWFmEmJnqw5zAKR3KdzrlQaxls8ClIRnmC1FzhpHWekGdg_N-oDC0cD_I6Z15vGvotnH3pcTdqyQ"
)

func SetTokenServiceEnvForTesting(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	t.Setenv("JWT_ACCESS_TOKEN_PRIVATE_KEY_FILE", filepath.Join(dir, "../../certs/private.pem"))
	t.Setenv("JWT_ACCESS_TOKEN_PUBLIC_KEY_FILE", filepath.Join(dir, "../../certs/public.pem"))
	t.Setenv("JWT_REFRESH_TOKEN_PRIVATE_KEY_FILE", filepath.Join(dir, "../../certs/private.pem"))
	t.Setenv("JWT_REFRESH_TOKEN_PUBLIC_KEY_FILE", filepath.Join(dir, "../../certs/public.pem"))
	t.Setenv("JWT_ISSUER", issuer)
}

func TestNewTokenService(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())

	if ts.accessTokenPrivateKey == nil {
		t.Errorf("access token privateKey is empty")
	}

	if ts.accessTokenPublicKey == nil {
		t.Errorf("access token publicKey is empty")
	}

	if ts.refreshTokenPrivateKey == nil {
		t.Errorf("refresh token privateKey is empty")
	}

	if ts.refreshTokenPublicKey == nil {
		t.Errorf("refresh token publicKey is empty")
	}
}

func TestTokenService_GenerateAccessToken(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())

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
					Id: primitive.NewObjectID(),
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
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.GenerateRefreshToken(tt.args.ctx, tt.args.user)
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

func TestTokenService_ValidateAccessTokenFromRequest(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())

	type args struct {
		ctx context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ValidateAccessTokenFromRequest(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.ValidateAccessTokenFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenService.ValidateAccessTokenFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_parseToken(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())

	u := &model.User{
		Id: primitive.NewObjectID(),
	}
	validToken, err := ts.GenerateAccessToken(context.Background(), u)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to generate access token"))
	}

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
					Subject: u.Id.Hex(),
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
					Subject: u.Id.Hex(),
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
					Subject: u.Id.Hex(),
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
					Subject: u.Id.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
		{
			name: "should fail because wrong algorithm",
			args: args{
				ctx:   context.Background(),
				user:  u,
				token: invalidSAlgorithmToken,
			},
			want: &Claims{
				Role: "user",
				StandardClaims: jwt.StandardClaims{
					Subject: u.Id.Hex(),
					Issuer:  issuer,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ParseToken(tt.args.ctx, tt.args.token)

			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.parseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("TokenService.parseToken() = %v, want %v", got, tt.want)
				}
			} else {
				if (err != nil) != tt.wantErr && got.GetRole() != tt.want.Role {
					t.Errorf("TokenService.parseToken() = %v, want %v", got.GetRole(), tt.want.Role)
				}
				if (err != nil) != tt.wantErr && got.GetIssuer() != tt.want.StandardClaims.Issuer {
					t.Errorf("TokenService.parseToken() = %v, want %v", got.GetIssuer(), tt.want.StandardClaims.Issuer)
				}
			}
		})
	}
}

func TestTokenService_ValidateRefreshToken(t *testing.T) {
	SetTokenServiceEnvForTesting(t)

	ts := NewTokenService(config.New(), NewLoggerMock())
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return user id",
			args: args{
				ctx:   context.Background(),
				token: validRefreshToken,
			},
			want: "623b3d2c2da9e50a7486de7f",
		},
		{
			name: "should fail because expired token",
			args: args{
				ctx:   context.Background(),
				token: expiredRefreshToken,
			},
			wantErr: true,
		},
		{
			name: "should fail because invalid issuer",
			args: args{
				ctx:   context.Background(),
				token: invalidIssuerRefreshToken,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ValidateRefreshToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenService.ValidateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TokenService.ValidateRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
