package site

import (
	"github.com/changwei4869/wedding/model"
	"gorm.io/gorm"
)

type SiteService struct {
	db *gorm.DB
}

func NewSiteService(db *gorm.DB) *SiteService {
	return &SiteService{db: db}
}

func (s *SiteService) Detail(id int) (model.SitesResp, error) {
	var site model.Sites
	if err := s.db.First(&site, id).Error; err != nil {
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
	if err := s.db.Find(&sites).Error; err != nil {
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
	return s.db.Create(&newSite).Error
}

func (s *SiteService) Del(id int) error {
	return s.db.Delete(&model.Sites{}, id).Error
}

func (s *SiteService) Edit(site model.SitesEditReq) error {
	return s.db.Model(&model.Sites{}).Where("id = ?", site.ID).Updates(model.Sites{
		City:   site.City,
		Status: site.Status,
	}).Error
}
