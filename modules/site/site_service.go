package site

import (
	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
)

type SiteService struct{}

var SiteIns = &SiteService{}

func (s *SiteService) Detail(id int) (model.SitesResp, error) {
	var site model.Sites
	if err := db.GetDb().First(&site, id).Error; err != nil {
		return model.SitesResp{}, err
	}
	return model.SitesResp{
		ID:     site.ID,
		City:   site.City,
		Status: site.Status,
	}, nil
}

func (s *SiteService) GetAll() ([]model.SitesResp, error) {
	var sites []model.Sites
	if err := db.GetDb().Find(&sites).Error; err != nil {
		return nil, err
	}

	var sitesResp []model.SitesResp
	for _, site := range sites {
		sitesResp = append(sitesResp, model.SitesResp{
			ID:     site.ID,
			City:   site.City,
			Status: site.Status,
		})
	}

	return sitesResp, nil
}

func (s *SiteService) Add(site model.SitesAddReq) error {
	newSite := model.Sites{
		City:   site.City,
		Status: site.Status,
	}
	return db.GetDb().Create(&newSite).Error
}

func (s *SiteService) Del(id int) error {
	return db.GetDb().Delete(&model.Sites{}, id).Error
}

func (s *SiteService) Edit(site model.SitesEditReq) error {
	return db.GetDb().Model(&model.Sites{}).Where("id = ?", site.ID).Updates(model.Sites{
		City:   site.City,
		Status: site.Status,
	}).Error
}
