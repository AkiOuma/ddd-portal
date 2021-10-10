package valobj

type User struct {
	userid   string
	username string
	avatar   string
	email    string
	mobile   string
}

func NewUser(
	userid string,
	username string,
	avatar string,
	email string,
	mobile string,
) *User {
	return &User{
		userid:   userid,
		username: username,
		avatar:   avatar,
		email:    email,
		mobile:   mobile,
	}
}

func (vo *User) UserID() string {
	return vo.userid
}

func (vo *User) Username() string {
	return vo.username
}

func (vo *User) Avatar() string {
	return vo.avatar
}

func (vo *User) Email() string {
	return vo.email
}

func (vo *User) Mobile() string {
	return vo.mobile
}
