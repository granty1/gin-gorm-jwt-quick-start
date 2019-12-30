package migrate

import (
	"github.com/gin-cli/model"
	"github.com/gin-cli/tools"
)

//RegisterTables migrate tables that necessary
func RegisterTables() {
	db := model.GetMysqlConn()
	db.AutoMigrate(
		&model.BaseRole{},
		&model.BaseUser{},
	)
	initRoles := []model.BaseRole{
		{
			RoleId:   "1",
			RoleName: "Administrator",
		},
		{
			RoleId:   "2",
			RoleName: "NormalUser",
		},
	}
	for _, role := range initRoles {
		db.Where(role).FirstOrCreate(&role)
	}

	initUsers := []model.BaseUser{
		{
			Username: "grant",
			Password: tools.Md5("grant"),
			Email:    "granty1@163.com",
			NickName: "Granty1",
			Phone:    "15316368801",
			RoleID:   "1",
		},
		{
			Username: "karl",
			Password: tools.Md5("karl"),
			Email:    "karl@163.com",
			NickName: "KralW",
			Phone:    "15316368801",
			RoleID:   "2",
		},
	}
	for _, user := range initUsers {
		db.Where(user).FirstOrCreate(&user)
	}
}
