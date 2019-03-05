package services

import (
	"cms/backend/db"
	"cms/backend/models"

	"github.com/go-xorm/xorm"
)

type IVideoServerice interface {
	GetById(id string) (models.Vod, bool)
}

type videoServerice struct {
	dbengine *xorm.Engine
}

func NewVideoServerice() IVideoServerice {
	return &videoServerice{
		dbengine: db.XORM,
	}
}

func (this *videoServerice) GetById(id string) (models.Vod, bool) {
	video := models.Vod{
		Id: id,
	}

	if ok, _ := this.dbengine.Get(&video); ok {
		return video, true
	}
	return models.Vod{}, false
}
