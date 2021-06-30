package Controller

import (
	"CeylonPaperOnline-BackEnd/Middleware"
	"CeylonPaperOnline-BackEnd/Module"
	"CeylonPaperOnline-BackEnd/View"
)

type AuthGateway struct {
	qc Middleware.QueryConsole
}

func (a* AuthGateway) Init(console Middleware.QueryConsole)  {
	a.qc = console
}

func (a AuthGateway) GetUserById(userid string) View.GetUserViewModule {
	var origin Module.User
	stmt, err := a.qc.Db.Prepare("SELECT user_id, user_group, user_major, user_status FROM users WHERE user_id=?")
	defer stmt.Close()

	if err != nil {
		Middleware.Console.Log(err, "Prepare SQL Script")
		return View.GetUserViewModule{}
	}

	queryResult, err := stmt.Query(userid)
	if err != nil {
		Middleware.Console.Log(err, "Query Database")
		return View.GetUserViewModule{}
	}

	for queryResult.Next() {
		err := queryResult.Scan(&origin.UserId, &origin.UserGroup, &origin.UserMajor, &origin.UserStatus)
		if err != nil {
			Middleware.Console.Log(err, "Generate User ViewModule")
			return View.GetUserViewModule{}
		} else {
			break
		}
	}

	return View.GetUserViewModule{
		UserId: origin.UserId,
		UserGroup: origin.UserGroup,
		UserMajor: origin.UserMajor,
		UserStatus: origin.UserStatus,
	}
}

func (a AuthGateway) GetUserWithAuth(userid, userpass string) View.UserViewModule {
	var origin Module.User
	stmt, err := a.qc.Db.Prepare("SELECT user_id, user_name, user_major, user_grade, user_class, user_status, user_group, user_sex, user_pass, user_email, user_phone FROM users WHERE user_id=?")
	defer stmt.Close()

	if err != nil {
		Middleware.Console.Log(err, "Prepare SQL Script")
		return View.UserViewModule{}
	}

	queryResult, err := stmt.Query(userid)
	if err != nil {
		Middleware.Console.Log(err, "Query Database")
		return View.UserViewModule{}
	}

	for queryResult.Next() {
		err := queryResult.Scan(&origin.UserId, &origin.UserName, &origin.UserMajor, &origin.UserGrade, &origin.UserClass,
			&origin.UserStatus, &origin.UserGroup, &origin.UserSex, &origin.UserPass, &origin.UserEmail, &origin.UserPhone)
		if err != nil {
			Middleware.Console.Log(err, "Generate User ViewModule")
			return View.UserViewModule{}
		} else if origin.UserPass != Module.EncodingPassword(userpass, origin.UserName) {
			return View.UserViewModule{}
		} else {
			break
		}
	}

	return View.UserViewModule{
		UserId: origin.UserId,
		UserName: origin.UserName,
		UserStatus: origin.UserStatus,
		UserMajor: origin.UserMajor,
		UserGroup: origin.UserGroup,
		UserClass: origin.UserClass,
		UserGrade: origin.UserGrade,
		UserSex: origin.UserSex,
		UserEmail: origin.UserEmail,
		UserPhone: origin.UserPhone,
	}
}