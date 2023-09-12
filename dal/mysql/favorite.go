package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// 登录用户可以对视频点赞，在个人主页喜欢Tab下能够查看点赞视频列表

type Favorite struct {
	Video    Video     `gorm:"foreignkey:VideoID;" json:"video,omitempty"`
	VideoID  uint      `gorm:"index:idx_videoid;not null" json:"video_id"`
	User     User      `gorm:"foreignkey:UserID;" json:"user,omitempty"`
	UserID   uint      `gorm:"index:idx_userid;not null" json:"user_id"`
	LikeTime time.Time `gorm:"not null;index:idx_create" json:"created_at,omitempty"`
}

func (Favorite) TableName() string {
	return "favorite"
}

// 点赞
func CreateFavorite(ctx context.Context, userId int64, videoId int64) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 新增点赞
		err := tx.Create(&Favorite{UserID: uint(userId), VideoID: uint(videoId)}).Error
		if err != nil {
			return err
		}

		// video表中的点赞数量+1
		res := tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		// 改变点赞用户表中的点赞数量
		res = tx.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return err
		}
		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		//4.改变视频拥有者用户表中的获赞次数
		// res = tx.Model(&User{}).Where("id = ?", authorId).Update("total_favorited", gorm.Expr("total_favorited + ?", 1))
		// if res.Error != nil {
		// 	return err
		// }
		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		return nil
	})
	return err
}

// 取消点赞
func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		Favorite := new(Favorite)
		if err := tx.Where("user_id = ? and video_id = ?", userId, videoId).First(&Favorite).Error; err != nil {
			return err
		} else if err == gorm.ErrRecordNotFound {
			return nil
		}

		// 删除点赞表
		err := tx.Unscoped().Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favorite).Error
		if err != nil {
			return err
		}

		// 视频表中的点赞数量-1
		res := tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		// 改变用户的点赞次数
		res = tx.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return err
		}
		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		// 视频拥有者用户表中的获赞总数
		// res = tx.Model(&User{}).Where("id = ?", authorId).Update("total_favorited", gorm.Expr("total_favorited - ?", 1))
		// if res.Error != nil {
		// 	return err
		// }
		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		return nil
	})
	return err
}

// 根据id查看个人点赞视频列表
func GetFavoriteList(ctx context.Context, userId int64) ([]Video, error) {
	var FavoriteVideoList []Video
	if err := db.Clauses(dbresolver.Write).WithContext(ctx).Where("user_id = ?", userId).Find(&FavoriteVideoList).Error; err != nil {
		return nil, err
	}
	return FavoriteVideoList, nil
}
