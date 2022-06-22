package service

import (
	"wra/meetings/pkg/model"
	"wra/meetings/pkg/repository"
)

type MeetingsService struct {
	MeetingsRepository *repository.Repository
}

// New creates new MeetingService instance
func New(r *repository.Repository) *MeetingsService {
	return &MeetingsService{MeetingsRepository: r}
}

// All ..
func (m *MeetingsService) All() ([]model.Meeting, error) {
	return m.MeetingsRepository.All()
}

// FindByID implementation.
func (m *MeetingsService) FindByID(id uint) (*model.Meeting, error) {
	return m.MeetingsRepository.FindByID(id)
}

// Save implementation.
func (m *MeetingsService) Save(meeting *model.Meeting) (*model.Meeting, error) {
	return m.MeetingsRepository.Save(meeting)
}

// Delete implementation.
func (m *MeetingsService) Delete(id uint) error {
	return m.MeetingsRepository.Delete(id)
}

// Migrate implementation.
func (m *MeetingsService) Migrate() error {
	return m.MeetingsRepository.Migrate()
}
