package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/koopa0/blog-service/internal/model"
	"github.com/koopa0/blog-service/pkg/app"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}

	return tag.Delete(d.engine)
}
