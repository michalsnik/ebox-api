package users

type UsersService interface {
	CreateUser(reqData PostUserRequestData) (*User, error)
	GetUserById(userID int) (*User, error)
}

type usersService struct {
	r UsersRepository
}

func NewUsersService(r UsersRepository) UsersService {
	return &usersService{
		r: r,
	}
}

func (svc *usersService) CreateUser(reqData PostUserRequestData) (*User, error) {
	if !emailRegExp.MatchString(reqData.Email) {
		return nil, ErrInvalidEmail
	}

	if err := validatePassword(reqData.Password); err != nil {
		return nil, err
	}

	user, err := svc.r.CreateUser(reqData)

	return user, err
}

func (svc *usersService) GetUserById(userID int) (*User, error) {
	user, err := svc.r.GetUserById(userID)
	return user, err
}
