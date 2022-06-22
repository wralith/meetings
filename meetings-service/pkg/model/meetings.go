package model

import (
	"time"
)

type Meeting struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
}

type MeetingDTO struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// ToMeeting converts MeetingDTO to Meeting
func ToMeeting(m *MeetingDTO) *Meeting {
	return &Meeting{
		Title: m.Title,
		Body:  m.Body,
	}
}

// ToMeetingDTO converts Meeting to MeetingDTO
func ToMeetingDTO(m *Meeting) *MeetingDTO {
	return &MeetingDTO{
		ID:    m.ID,
		Title: m.Title,
		Body:  m.Body,
	}
}
