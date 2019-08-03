package users

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	AvatarURL string `json:"avatarURL"`
}
