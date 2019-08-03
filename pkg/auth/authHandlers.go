package auth

type AuthHandlers struct {
	svc *AuthService
}

func NewAuthHandlers(svc *AuthService) *AuthHandlers {
	return &AuthHandlers{
		svc: svc,
	}
}

