package repo

import "github.com/jinzhu/gorm"

type Repo struct {
	DB *gorm.DB
}
