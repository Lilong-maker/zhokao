package model

import "gorm.io/gorm"

type Show struct {
	gorm.Model
	Name  string  `grom:"type:varchar(30)"`
	Price float32 `grom:"type:decimal(10,2)"`
	Num   int     `grom:"type:varchar(30)"`
}

func (g *Show) FindGood(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}

func (g *Show) GoodsAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (g *Show) FIndGoods(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Find(&g).Error
}

func (g *Show) GoodsUpdate(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Update("goods", Show{}).Error
}

func (g *Show) DeleteGood(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Delete(&g).Error
}
