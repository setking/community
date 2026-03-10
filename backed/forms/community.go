package forms

// 社区
type CommunityForm struct {
	CommunityName string `form:"community_name" json:"community_name" binding:"required"`
	Introduction  string `form:"introduction" json:"introduction" binding:"required"`
}
