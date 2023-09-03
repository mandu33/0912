package mysql
import(
	"context"
    "time"
	"gorm.io/gorm"
	// "gorm.io/plugin/dbresolver"
)
// var db = getDB() 

type Video struct {
	gorm.Model
	PublishTime     time.Time `gorm:"column:update_time;not null;index:idx_update" `
	Author        User      `gorm:"foreignkey:AuthorID"`
	UserID      int       `gorm:"index:idx_authorid;not null"`
	PlayUrl       string    `gorm:"type:varchar(255);not null"`
	CoverUrl      string    `gorm:"type:varchar(255)"`
	FavoriteCount int       `gorm:"default:0"`
	CommentCount  int       `gorm:"default:0"`
	Title         string    `gorm:"type:varchar(50);not null"`
}

func GetVideos(ctx context.Context, limit int, latestTime *int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == nil || *latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	conn := db.WithContext(ctx)

	if err := conn.Limit(limit).Order("update_time desc").Find(&videos, "update_time < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func CreateVideo(ctx context.Context, video *Video) error {
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func PublishList(ctx context.Context, userId int64) ([]*Video, error) {
	var pubList []*Video
	err := db.WithContext(ctx).Model(&Video{}).Where(&Video{UserID: int(userId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}

// 根据视频id得到视频
func GetVideoById(ctx context.Context, videoId int64) (Video, error) {
	var vid Video
	err := db.WithContext(ctx).Where("id=?", videoId).Find(&vid).Error
	if err != nil {
		return vid, err
	}
	return vid, nil
}