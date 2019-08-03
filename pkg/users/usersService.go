package users

type UsersService struct {
	r *UsersRepository
}

func NewUsersService(r *UsersRepository) *UsersService {
	return &UsersService{
		r: r,
	}
}

func (svc *UsersService) CreateUser(reqData PostUserRequestData) (*User, error) {
	if !emailRegExp.MatchString(reqData.Email) {
		return nil, ErrInvalidEmail
	}

	// validate password
	if err := validatePassword(reqData.Password); err != nil {
		return nil, err
	}

	user, err := svc.r.CreateUser(reqData)

	return user, err
}
