package dao

import (
	"errors"
	"myApp/forms"
	"myApp/global"
	"myApp/models"
	"myApp/utils"
)

// 添加社区
func AddCommunity(community forms.CommunityForm) error {
	var communities models.Community
	// 检查用社区是否存在
	if global.DB.Where("community_name = ?", community.CommunityName).First(&communities).RowsAffected == 1 {
		return errors.New("此名称已存在")
	}
	var bizID int64
	global.DB.Exec("INSERT INTO seq_community VALUES (NULL)")
	global.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&bizID)
	rsp := &models.Community{
		CommunityId:   bizID,
		CommunityName: community.CommunityName,
		Introduction:  community.Introduction,
	}
	res := global.DB.Create(rsp)
	if res.Error != nil {
		return errors.New("内部错误")
	}
	return nil
}

// 删除社区
func DeleteCommunity(ID int64) error {
	var communities models.Community
	return global.DB.Delete(&communities, ID).Error
}

// 更新社区
func UpdateCommunity(community forms.CommunityForm, ID int64) error {
	var communities models.Community
	// 检查用社区是否存在
	if res := global.DB.First(&communities, ID); res.RowsAffected == 0 {
		return errors.New("没有这条数据")
	}
	communities.CommunityName = community.CommunityName
	communities.Introduction = community.Introduction
	res := global.DB.Save(&communities)
	if res.Error != nil {
		return errors.New("内部错误")
	}
	return nil
}

// 获取社区列表
func ListCommunity(page, pageSize int) ([]models.Community, int64, error) {
	var communities []models.Community
	result := global.DB.Find(&communities)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	Total := result.RowsAffected
	global.DB.Scopes(utils.Paginate(page, pageSize)).Find(&communities)
	return communities, Total, nil
}

// 通过id查询Community
func GetCommunityByID(id int64) (*models.Community, error) {
	var communities models.Community
	err := global.DB.Where("community_id =?", id).First(&communities).Error
	if err != nil {
		return nil, err
	}
	return &communities, nil
}
