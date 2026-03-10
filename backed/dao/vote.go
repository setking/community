package dao

import (
	"errors"
	"math"
	"myApp/global"
	"time"

	"github.com/go-redis/redis"
)

/*
	PostVote 为帖子投票

投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况

	1.之前没投过票，现在要投赞成票
	2.之前投过反对票，现在要改为赞成票

v=0时，有两种情况

	1.之前投过赞成票，现在要取消
	2.之前投过反对票，现在要取消

v=-1时，有两种情况

	1.之前没投过票，现在要投反对票
	2.之前投过赞成票，现在要改为反对票
*/
func PostVote(postId, userId string, v float64) (err error) {
	postTime := global.Redis.ZScore(KeyPostScoreZSet, postId).Val()
	if float64(time.Now().Unix())-postTime > WeekSeconds {
		return errors.New("已过投票时间")
	}
	key := KeyCommunityPostSetPrefix + postId
	ov := global.Redis.ZScore(key, userId).Val()
	diffAbs := math.Abs(ov - v)
	pipline := global.Redis.TxPipeline()
	pipline.ZAdd(key, redis.Z{Score: v, Member: userId})
	pipline.ZIncrBy(KeyPostScoreZSet, VoteScore*diffAbs*v, postId)
	switch math.Abs(ov) - math.Abs(v) {
	case 1:
		// 取消投票 ov=1/-1 v=0
		// 投票数-1
		pipline.HIncrBy(KeyPostInfoHashPrefix+postId, "votes", -1)
	case 0:
		// 反转投票 ov=-1/1 v=1/-1
		// 投票数不用更新
	case -1:
		// 新增投票 ov=0 v=1/-1
		// 投票数+1
		pipline.HIncrBy(KeyPostInfoHashPrefix+postId, "votes", 1)
	default:
		// 已经投过票了
		return errors.New("已经投过票了")

	}
	_, err = pipline.Exec()
	return
}
