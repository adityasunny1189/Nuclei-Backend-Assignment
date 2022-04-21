package user

type DBMethods interface {
	SaveUserDetails([]User)
	ShowUserDetails([]User)
}
