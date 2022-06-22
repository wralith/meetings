package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wra/meetings/pkg/model"
	"wra/meetings/pkg/service"

	"github.com/go-chi/chi"
)

// TODO: DRY!!!!

type MeetingsAPI struct {
	MeetingsService service.MeetingsService
}

// New create new MeetingsApi instance
func New(m service.MeetingsService) MeetingsAPI {
	return MeetingsAPI{MeetingsService: m}
}

// FindAllMeetings implementation.
func (m *MeetingsAPI) FindAllMeetings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		meetings, err := m.MeetingsService.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(meetings)
	}
}

// FindByID implementation.
func (m *MeetingsAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getParamsID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		meeting, err := m.MeetingsService.FindByID(uint(*id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(meeting)
	}
}

// CreateMeeting implementation.
func (m *MeetingsAPI) CreateMeeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var meetingDTO model.MeetingDTO

		decodedInput, err := decodeBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		meetingDTO = *decodedInput

		newMeeting, err := m.MeetingsService.Save(model.ToMeeting(&meetingDTO))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(model.ToMeetingDTO(newMeeting))
	}
}

// UpdateMeeting implementation.
func (m *MeetingsAPI) UpdateMeeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := getParamsID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var meetingDTO model.MeetingDTO

		decodedInput, err := decodeBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		meetingDTO = *decodedInput

		meeting, err := m.MeetingsService.FindByID(uint(*id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		meeting.Title = meetingDTO.Title
		meeting.Body = meetingDTO.Body
		updatedMeeting, err := m.MeetingsService.Save(meeting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(model.ToMeetingDTO(updatedMeeting))

	}
}

// DeleteMeeting implementation.
func (m *MeetingsAPI) DeleteMeeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := getParamsID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		meeting, err := m.MeetingsService.FindByID(uint(*id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = m.MeetingsService.Delete(meeting.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Post Deleted Succsefully"})
	}
}

// Migrate implementation.
func (m *MeetingsAPI) Migrate() {
	err := m.MeetingsService.Migrate()
	if err != nil {
		log.Println(err)
	}
}

// getParamsID returns id in the URL as an integer type
func getParamsID(r *http.Request) (*int, error) {
	rawID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

// decodeBody decodes request body to meetingDTO object
func decodeBody(r *http.Request) (*model.MeetingDTO, error) {
	var meetingDTO model.MeetingDTO
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&meetingDTO); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	defer r.Body.Close()
	return &meetingDTO, nil
}
