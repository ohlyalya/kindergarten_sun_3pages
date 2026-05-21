package models

import "time"

type ContactInfo struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Hours   string `json:"hours"`
}

type Hero struct {
	Title       string      `json:"title"`
	Image       string      `json:"image"`
	ContactInfo ContactInfo `json:"contactInfo"`
}

type Advantage struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Room struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type Group struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Age         string `json:"age"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type Activity struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Teacher struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Image       string `json:"image"`
	Description string `json:"description"`
	IsMain      bool   `json:"isMain"`
}

type Review struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type News struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Application struct {
	ID        int       `json:"id"`
	Parent    string    `json:"parent"`
	Phone     string    `json:"phone"`
	ChildName string    `json:"childName"`
	ChildAge  int       `json:"childAge"`
	Comment   string    `json:"comment"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ApplicationRequest struct {
	Parent    string `json:"parent"`
	Phone     string `json:"phone"`
	ChildName string `json:"childName"`
	ChildAge  int    `json:"childAge"`
	Comment   string `json:"comment"`
}

type Question struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type QuestionRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Text  string `json:"text"`
}

type HomeResponse struct {
	Hero       Hero        `json:"hero"`
	Advantages []Advantage `json:"advantages"`
	Rooms      []Room      `json:"rooms"`
}

type ProgramsResponse struct {
	Groups     []Group    `json:"groups"`
	Activities []Activity `json:"activities"`
	Teachers   []Teacher  `json:"teachers"`
}

type ParentsResponse struct {
	Reviews []Review `json:"reviews"`
	Events  []Event  `json:"events"`
	News    []News   `json:"news"`
}

type SearchResult struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
