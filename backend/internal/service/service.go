package service

import (
	"errors"
	"regexp"
	"strings"

	"kindergarten-sun-backend/internal/models"
	"kindergarten-sun-backend/internal/storage"
)

type Service struct {
	store *storage.Store
}

func New(store *storage.Store) *Service {
	return &Service{store: store}
}

func (s *Service) Home() models.HomeResponse                 { return s.store.Home() }
func (s *Service) Programs() models.ProgramsResponse         { return s.store.Programs() }
func (s *Service) Parents() models.ParentsResponse           { return s.store.Parents() }
func (s *Service) Search(query string) []models.SearchResult { return s.store.Search(query) }
func (s *Service) Applications() []models.Application        { return s.store.Applications() }
func (s *Service) Questions() []models.Question              { return s.store.Questions() }

func (s *Service) CreateApplication(req models.ApplicationRequest) (models.Application, error) {
	req.Parent = strings.TrimSpace(req.Parent)
	req.Phone = strings.TrimSpace(req.Phone)
	req.ChildName = strings.TrimSpace(req.ChildName)
	req.Comment = strings.TrimSpace(req.Comment)

	if req.Parent == "" || len([]rune(req.Parent)) < 2 {
		return models.Application{}, errors.New("укажите имя родителя")
	}
	if !validPhone(req.Phone) {
		return models.Application{}, errors.New("укажите корректный телефон")
	}
	if req.ChildName == "" || len([]rune(req.ChildName)) < 2 {
		return models.Application{}, errors.New("укажите имя ребёнка")
	}
	if req.ChildAge < 1 || req.ChildAge > 7 {
		return models.Application{}, errors.New("возраст ребёнка должен быть от 1 до 7 лет")
	}
	return s.store.AddApplication(req), nil
}

func (s *Service) CreateQuestion(req models.QuestionRequest) (models.Question, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Phone = strings.TrimSpace(req.Phone)
	req.Text = strings.TrimSpace(req.Text)
	if req.Name == "" || len([]rune(req.Name)) < 2 {
		return models.Question{}, errors.New("укажите имя")
	}
	if !validPhone(req.Phone) {
		return models.Question{}, errors.New("укажите корректный телефон")
	}
	if req.Text == "" || len([]rune(req.Text)) < 10 {
		return models.Question{}, errors.New("вопрос должен быть не короче 10 символов")
	}
	return s.store.AddQuestion(req), nil
}

func validPhone(phone string) bool {
	re := regexp.MustCompile(`^[+0-9()\-\s]{7,20}$`)
	return re.MatchString(phone)
}
