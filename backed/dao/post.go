package dao

import (
	"errors"
	"fmt"
	"myApp/forms"
	"myApp/global"
	"myApp/initialize"
	"myApp/models"
	"myApp/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	KeyPostInfoHashPrefix     = "sky:post:"
	KeyPostTimeZSet           = "sky:post:time"
	KeyPostScoreZSet          = "sky:post:score"
	KeyPostVotedZSetPrefix    = "sky:post:voted:"
	KeyCommunityPostSetPrefix = "sky:community:"
)
const (
	WeekSeconds         = 7 * 24 * 3600
	VoteScore   float64 = 432
	PostPerAge          = 20
)

// 添加社区
func AddPost(post forms.PostForm) (err error) {
	var posts models.Post
	// 检查帖子标题是否重复
	if global.DB.Where("title = ?", post.Title).First(&posts).RowsAffected == 1 {
		return errors.New("帖子标题已存在")
	}
	CommunityID, err := strconv.ParseInt(post.CommunityID, 10, 64)
	if err != nil {
		return err
	}
	rsp := &models.Post{
		Title:       post.Title,
		Content:     post.Content,
		CommunityID: CommunityID,
		AuthorID:    post.AuthorID,
		Status:      post.Status,
	}
	rsp.PostID, err = initialize.GetID()
	if err != nil {
		return err
	}
	res := global.DB.Create(rsp)
	if res.Error != nil {
		return res.Error
	}
	Community, err := GetCommunityByID(CommunityID)
	if err != nil {
		return err
	}
	if err = CreateRedisPost(
		fmt.Sprint(rsp.PostID),
		fmt.Sprint(rsp.AuthorID),
		post.Title,
		utils.TruncateByWords(post.Content, 120),
		Community.CommunityName,
	); err != nil {
		zap.L().Error("Create Redis Post error: ", zap.Error(err))
		return err
	}
	return
}

// 同步到redis
func CreateRedisPost(postID, userID, title, summary, communityName string) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPrefix + postID
	communityKey := KeyCommunityPostSetPrefix + communityName
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}
	pipeline := global.Redis.Pipeline()
	pipeline.ZAdd(votedKey, redis.Z{
		Score:  now,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*WeekSeconds)
	pipeline.HMSet(KeyPostInfoHashPrefix+postID, postInfo)
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})
	pipeline.SAdd(communityKey, postID)
	_, err = pipeline.Exec()
	return
}

// 删除帖子
func DeletePost(ID int64) error {
	var posts models.Post
	return global.DB.Where("post_id =?", ID).Delete(&posts).Error
}

// 更新社区
func UpdatePost(post forms.PostForm, id int64) error {
	var posts models.Post
	// 检查用帖子是否存在
	if res := global.DB.First(&posts, id); res.RowsAffected == 0 {
		return errors.New("没有这条数据")
	}
	posts.Title = post.Title
	posts.Content = post.Content
	res := global.DB.Save(&posts)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 获取帖子列表
func ListPost(page, pageSize int) ([]models.Post, int64, error) {
	var posts []models.Post
	result := global.DB.Find(&posts)
	if result.RowsAffected == 0 {
		return nil, 0, errors.New("没有数据")
	}
	if result.Error != nil {
		return nil, 0, result.Error
	}
	Total := result.RowsAffected
	global.DB.Scopes(utils.Paginate(page, pageSize)).Find(&posts)
	return posts, Total, nil
}
func ListPostByRedis(order string, page int64) ([]map[string]string, int64, error) {
	key := KeyPostTimeZSet
	if order == "time" {
		key = KeyPostTimeZSet
	}
	start := (page - 1) * PostPerAge
	end := start + PostPerAge - 1
	ids := global.Redis.ZRange(key, start, end).Val()
	postlist := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		postData := global.Redis.HGetAll(KeyPostInfoHashPrefix + id).Val()
		postData["id"] = id
		postlist = append(postlist, postData)
	}
	return postlist, int64(len(postlist)), nil

}
func GetPostDetail(id int64) (*models.Post, error) {
	var post models.Post
	result := global.DB.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
