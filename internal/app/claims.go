package app

type Claims interface {
	GetSubject() string
	GetRole() string
	GetIssuer() string
}
