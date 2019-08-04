package auth

type AuthRepository interface {
	ValidateUser(email string, password string) (int, error)
}

type AuthService interface {
	AuthenticateUser (email string, password string) (string, error)
	ParseToken (tokenString string) (int, error)
	CreateToken (userID int) (string, error)
}

type authService struct {
	r AuthRepository
}

func NewAuthService(r AuthRepository) AuthService {
	return &authService{
		r: r,
	}
}

func (svc *authService) AuthenticateUser (email string, password string) (string, error) {
	userID, err := svc.r.ValidateUser(email, password)
	if err != nil {
		return "", err
	}

	tokenString, err := svc.CreateToken(userID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (svc *authService) ParseToken (tokenString string) (int, error) {
	return ParseToken(tokenString)
}

func (svc *authService) CreateToken (userID int) (string, error) {
	return CreateToken(userID)
}
