package repository

import (
	"wra/meetings/pkg/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// New creates new Repository instance
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// All returns all the meetings in the database
func (p *Repository) All() ([]model.Meeting, error) {
	meetings := []model.Meeting{}
	err := p.db.Find(&meetings).Error
	return meetings, err
}

// FindByID finds a single meeting by its id
func (p *Repository) FindByID(id uint) (*model.Meeting, error) {
	meeting := new(model.Meeting)
	err := p.db.Where("id = ?", id).First(&meeting).Error
	return meeting, err
}

// Save a meeting to database.
func (p *Repository) Save(m *model.Meeting) (*model.Meeting, error) {
	err := p.db.Save(&m).Error
	return m, err
}

// Delete deletes meeting by its id.
func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&model.Meeting{ID: id}).Error
	return err
}

// Migrate register go objects on database.
func (p *Repository) Migrate() error {
	err := p.db.AutoMigrate(&model.Meeting{})
	return err
}
