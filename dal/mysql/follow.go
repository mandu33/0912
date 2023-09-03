package mysql
import(
	"context"
	"time"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)
// var db = getDB() //连接池对象
// CREATE TABLE follows (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	user_id VARCHAR(255) NOT NULL,
// 	dst_id VARCHAR(255) NOT NULL,
// 	follow_time DATETIME NOT NULL
// );
type Follow struct{
	gorm.Model
	userID uint `gorm:"index:idx_userid;not null" json:"user_id"`
	user User `gorm:"foreignkey:userID;" json:"user,omitempty"`
	dstID uint `gorm:"index:idx_userid;index:idx_userid_to;not null" json:"dst_id"`
	dstUser User `gorm:"foreignkey:dstID;" json:"dstUser,omitempty"`
	followTime time.Time `gorm:"not null;index:idx_create" json:"created_at,omitempty"`
}
func GetFollow(ctx context.Context, uid int64, tid int64) (*Follow, error) {
	follow := new(Follow)

	if err := db.Clauses(dbresolver.Write).WithContext(ctx).First(&follow, "user_id = ? and dst_id = ?", uid, tid).Error; err != nil {
		return nil, err
	}
	return follow, nil
}
//关注
func GetFollowing(ctx context.Context, uid int64, tid int64) error {
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增关注数据
		err := tx.Create(&Follow{userID: uint(uid), dstID: uint(tid)}).Error
		if err != nil {
			return err
		}

		// 2.改变 user 表中的 following
		res := tx.Model(new(User)).Where("id = ?", uid).Update("following", gorm.Expr("following + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		// 3.改变 user 表中的 followed
		res = tx.Model(new(User)).Where("id = ?", tid).Update("followered", gorm.Expr("followered + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		return nil
	})
	return err
}
//取关
func DeleteFollowing(ctx context.Context, uid int64, tid int64) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		follow := new(Follow)
		if err := tx.Where("user_id = ? AND dst_id=?", uid, tid).First(&follow).Error; err != nil {
			return err
		}

		// 1. 删除关注数据
		err := tx.Unscoped().Delete(&follow).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following
		res := tx.Model(new(User)).Where("id = ?", uid).Update("following", gorm.Expr("following - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		// 3.改变 user 表中的 follower
		res = tx.Model(new(User)).Where("id = ?", tid).Update("followered", gorm.Expr("followered - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	return errno.ErrDatabase
		// }

		return nil
	})
	return err
}
func GetFollowingList(ctx context.Context, uid int64) ([]*Follow, error) {
	var FollowList []*Follow
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Where("user_id = ?", uid).Find(&FollowList).Error
	if err != nil {
		return nil, err
	}
	return FollowList, nil
}

// FollowerList returns the Follower List.
func GetFollowerList(ctx context.Context, tid int64) ([]*Follow, error) {
	var FollowList []*Follow
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Where("dst_id = ?", tid).Find(&FollowList).Error
	if err != nil {
		return nil, err
	}
	return FollowList, nil
}