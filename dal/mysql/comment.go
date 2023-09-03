package mysql
import (
	"context"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	
)
// 评论属于视频和用户
type Comment struct {
	gorm.Model   
	Video      Video          `gorm:"foreignkey:VideoID" json:"video,omitempty"`
	VideoID    uint           `gorm:"index:idx_videoid;not null" json:"video_id"`
	User       User           `gorm:"foreignkey:UserID" json:"user,omitempty"`
	UserID     uint           `gorm:"index:idx_userid;not null" json:"user_id"`
	Content    string         `gorm:"type:varchar(255);not null" json:"content"`

}

func (Comment) TableName() string {
	return "comment"
}
//增加评论
func AddComment(ctx context.Context, comment *Comment) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增评论数据
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		// 2.对 Video 表中的评论数+1
		res := tx.Model(&Video{}).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		// if res.RowsAffected != 1 {
		// 	// 影响的数据条数不是1
		// 	return res.ErrDatabase
		// }

		return nil
	})
	return err
}
// 删除评论
func DeleteComment(ctx context.Context, commentId int64, videoId int64) error {
	err := db.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.First(&comment, commentId).Error; err != nil {
			return err
		} else if err == gorm.ErrRecordNotFound {
			return nil
		}
        // 删除评论表
		err := tx.Where("id = ?", commentId).Delete(&Comment{}).Error
		if err != nil {
			return err
		}
		// 将视频对于评论数-1
		res := tx.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
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

// 获得当前视频的所有评论
func GetVideoCommentList(ctx context.Context, videoID int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Clauses(dbresolver.Read).WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoID: uint(videoID)}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
