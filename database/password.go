package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"github.com/jinzhu/gorm"
)

type Password interface {
	UpdatePassword(connection *gorm.DB, user *makeless_go_model.User, newPassword string) (*makeless_go_model.User, error)
}
