package View

type GetUserViewModule struct {
	UserId string `json:"user_id"`
	UserGroup string `json:"user_group"`
	UserMajor string `json:"user_major"`
	UserStatus int `json:"user_status"`
}

type UserViewModule struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserMajor string `json:"user_major"`
	UserGrade int `json:"user_grade"`
	UserClass int `json:"user_class"`
	UserStatus int `json:"user_status"`
	UserGroup string `json:"user_type"`
	UserSex int `json:"user_sex"`
	UserEmail string `json:"user_email"`
	UserPhone string `json:"user_phone"`
}