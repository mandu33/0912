package mysql
import(
	"context"
	"log"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)
// var db = getDB() //连接池对象
// CREATE TABLE user (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL,
// 	profession VARCHAR(255),
// 	age INT
// 	following INT
// 	followed INT
// );
//用户数据结构
type User struct{
	gorm.Model 
	userName string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"name,omitempty"`
	passWord string `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	profession string `gorm:"type:varchar(256)" json:"profession,omitempty"`
	age uint `gorm:"default:0;not null" json:"age,omitempty"`
	introduction string `gorm:"type:varchar(256)" json:"introduction,omitempty"`
	following uint `gorm:"default:0" json:"following"`
	followed uint `gorm:"default:0" json:"followed"`
	favorite_count uint  `gorm:"default:0" json:"favorite_count"`
	total_favorited uint  `gorm:"default:0" json:"favorite_count"`
	//FavoriteVideos  []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos,omitempty"`
}
//新增用户
func insertUser(ctx context.Context, user []*User) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
//删除用户
func deleteUser(){
	log.Printf("to be made")
}
//更新用户
func updateUser(){
	log.Printf("to be made")
}
//查询用户
func queryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := db.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
//根据ID查询用户
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)
	if err := db.Clauses(dbresolver.Read).WithContext(ctx).First(&res, userID).Error; err == nil {
		return res, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}
//根据用户名查询用户
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
//根据用户名查询密码
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