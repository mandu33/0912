package mysql

import (
	"context"
	"log"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// 用户数据结构
type User struct {
	gorm.Model
	UserName       string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"name,omitempty"`
	PassWord       string `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	Profession     string `gorm:"type:varchar(256)" json:"profession,omitempty"`
	Age            uint   `gorm:"default:0;not null" json:"age,omitempty"`
	Introduction   string `gorm:"type:varchar(256)" json:"introduction,omitempty"`
	Following      uint   `gorm:"default:0" json:"following"`
	Followed       uint   `gorm:"default:0" json:"followed"`
	FavoriteCount  uint   `gorm:"default:0" json:"favorite_count"`
	TotalFavorited uint   `gorm:"default:0" json:"total_favorited"`
	//FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos,omitempty"`
}

func (User) TableName() string {
	return "users"
}

// 新增用户
// func CreateUser(ctx context.Context, users User) error {
// 	return db.Create(&users).Error
// }

func InsertUser(ctx context.Context, user User) error {
	return db.WithContext(ctx).Create(&user).Error
}

// 删除用户
func DefaulteleteUser() {
	log.Printf("to be made")
}

// 更新用户
func UpdateUser() {
	log.Printf("to be made")
}

// 查询用户
func QueryUser(ctx context.Context, userName string) (*User, error) {
	res := new(User)
	if err := db.WithContext(ctx).Where("user_name = ?", userName).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByID(ctx context.Context, userID int) (*User, error) {
	res := new(User)
	if err := db.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 根据ID查询用户
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)
	if err := db.WithContext(ctx).First(&res, userID).Error; err == nil {
		return res, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}

// 根据用户名查询用户
func GetUserByName(ctx context.Context, userName int64) (*User, error) {
	res := new(User)
	if err := db.Clauses(dbresolver.Read).WithContext(ctx).Select("id, user_name, password").Where("user_name = ?", userName).First(&res).Error; err == nil {
		return res, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}

// 根据用户名查询密码
func GetPasswordByUsername(ctx context.Context, userName string) (*User, error) {
	user := new(User)
	if err := db.Clauses(dbresolver.Read).WithContext(ctx).
		Select("password").Where("user_name = ?", userName).
		First(&user).Error; err == nil {
		return user, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
