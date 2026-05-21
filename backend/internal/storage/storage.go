package storage

import (
	"strings"
	"sync"
	"time"

	"kindergarten-sun-backend/internal/models"
)

type Store struct {
	mu             sync.RWMutex
	contact        models.ContactInfo
	hero           models.Hero
	advantages     []models.Advantage
	rooms          []models.Room
	groups         []models.Group
	activities     []models.Activity
	teachers       []models.Teacher
	reviews        []models.Review
	events         []models.Event
	news           []models.News
	applications   map[int]models.Application
	questions      map[int]models.Question
	nextAppID      int
	nextQuestionID int
}

func NewStore() *Store {
	contact := models.ContactInfo{
		Address: "г. Москва, бульвар Матроса Железняка, 26/11",
		Phone:   "8 (910) 899-71-65",
		Email:   "gbdou.15moskva@gov.ru",
		Hours:   "Пн-Пт: 8:00-20:00",
	}

	return &Store{
		contact: contact,
		hero: models.Hero{
			Title:       "Государственное бюджетное дошкольное образовательное учреждение города Москвы № 15 «Солнце»",
			Image:       "assets/1.png",
			ContactInfo: contact,
		},
		advantages: []models.Advantage{
			{ID: 1, Title: "Безопасность и комфорт", Description: "Мы создаём безопасную и комфортную среду для каждого ребёнка: постоянное наблюдение, чистота и забота о здоровье детей — наши приоритеты.", Image: "assets/2.png"},
			{ID: 2, Title: "Индивидуальный подход", Description: "Воспитатели учитывают индивидуальные особенности каждого ребёнка, помогая раскрыть их таланты и способности.", Image: "assets/3.png"},
			{ID: 3, Title: "Квалифицированные педагоги", Description: "Наши воспитатели — профессионалы с многолетним опытом работы, которые подходят к своей работе с любовью и заботой.", Image: "assets/4.png"},
			{ID: 4, Title: "Развитие и обучение через игру", Description: "Каждый день мы наполняем интересными и познавательными занятиями, способствующими развитию любознательности и творческих способностей.", Image: "assets/5.png"},
			{ID: 5, Title: "Здоровое питание", Description: "Мы обеспечиваем детей вкусным и сбалансированным питанием, которое отвечает всем необходимым стандартам.", Image: "assets/6.png"},
		},
		rooms: []models.Room{
			{ID: 1, Title: "Игровая комната", Image: "assets/7.png", Description: "Светлая игровая зона для занятий, свободной игры и развития социальных навыков."},
			{ID: 2, Title: "Столовая", Image: "assets/8.png", Description: "Уютное помещение для питания детей с соблюдением санитарных требований."},
		},
		groups: []models.Group{
			{ID: 1, Title: "Ясли", Age: "1.5-3 года", Icon: "pacifier", Description: "Группа мягкой адаптации для самых маленьких детей."},
			{ID: 2, Title: "Младшие группы", Age: "3-4 года", Icon: "baby", Description: "Игровое развитие, коммуникация и первые самостоятельные навыки."},
			{ID: 3, Title: "Средние группы", Age: "4-5 лет", Icon: "child", Description: "Развитие речи, внимания, моторики и творческого мышления."},
			{ID: 4, Title: "Старшие группы", Age: "5-7 лет", Icon: "kid", Description: "Подготовка к школе, занятия по интересам и проектная деятельность."},
			{ID: 5, Title: "Подготовительные", Age: "6-7 лет", Icon: "student", Description: "Закрепление базовых навыков перед поступлением в первый класс."},
		},
		activities: []models.Activity{
			{ID: 1, Title: "Изостудия", Description: "Занятия по рисованию и аппликации, развитие творческих навыков.", Image: "assets/9.png"},
			{ID: 2, Title: "Танцевальный кружок", Description: "Основы хореографии, развитие координации и чувства ритма.", Image: "assets/10.png"},
			{ID: 3, Title: "Кружок робототехники", Description: "Обучение основам программирования и работы с роботами.", Image: "assets/11.png"},
		},
		teachers: []models.Teacher{
			{ID: 1, Name: "Лезарова Ирина Александровна", Position: "Заведующий", Image: "assets/12.png", IsMain: true, Description: "Мы стремимся создать тёплую и безопасную атмосферу, где каждый ребёнок чувствует себя как дома, получает возможность развиваться, учиться новому и находить первых друзей."},
			{ID: 2, Name: "Иванова Ольга Сергеевна", Position: "Педагог по раннему развитию", Image: "assets/13.png", Description: "Работает с младшими группами и помогает детям пройти адаптацию."},
			{ID: 3, Name: "Петрова Мария Александровна", Position: "Педагог по музыкальному воспитанию", Image: "assets/14.png", Description: "Проводит музыкальные занятия, утренники и творческие выступления."},
			{ID: 4, Name: "Сидорова Елена Викторовна", Position: "Педагог по физической культуре", Image: "assets/15.png", Description: "Отвечает за физическое развитие и спортивные активности."},
		},
		reviews: []models.Review{
			{ID: 1, Name: "Анна К.", Text: "Нашему ребёнку очень нравится детский сад! Заботливые воспитатели и интересные занятия помогают ему развиваться и с удовольствием идти на каждый день. Спасибо вам за теплоту и профессионализм!", Avatar: "assets/16.png", CreatedAt: "2026-04-10"},
			{ID: 2, Name: "Максим С.", Text: "Очень довольны уровнем подготовки и подходом к детям. Особенно радует внимание к индивидуальным потребностям и развитие творческих способностей. Спасибо всему коллективу детского сада!", Avatar: "assets/17.png", CreatedAt: "2026-04-14"},
			{ID: 3, Name: "Екатерина Л.", Text: "Отличное место для детей! Видно, что коллектив действительно заботится о воспитании и образовании малышей. Мой сын всегда рад идти в садик и возвращается домой с новыми знаниями и впечатлениями.", Avatar: "assets/18.png", CreatedAt: "2026-04-21"},
		},
		events: []models.Event{
			{ID: 1, Title: "Праздник Осени", Date: "15 октября", Description: "Яркий праздник, посвящённый осеннему времени года. Дети примут участие в тематических играх, мастер-классах и конкурсах, а также представят свои творческие работы.", Image: "assets/19.png"},
			{ID: 2, Title: "Новогодний утренник", Date: "25 декабря", Description: "Праздничное мероприятие, на котором детей ждёт встреча с Дедом Морозом и Снегурочкой, весёлые конкурсы и танцы, а также подарки для каждого ребёнка.", Image: "assets/20.png"},
		},
		news: []models.News{
			{ID: 1, Title: "Мастер-класс по рисованию", Date: "25 апреля 2026 г.", Description: "Сегодня в нашем детском саду прошёл увлекательный мастер-класс по рисованию для детей старших групп. Ребята познакомились с различными художественными техниками, учились сочетать цвета и создавать собственные яркие работы.", Image: "assets/21.png"},
			{ID: 2, Title: "Спортивный праздник", Date: "18 апреля 2026 г.", Description: "В детском саду состоялся весенний спортивный праздник, посвящённый здоровому образу жизни. Дети участвовали в эстафетах, подвижных играх и командных соревнованиях.", Image: "assets/22.png"},
			{ID: 3, Title: "Праздник Светлой Пасхи", Date: "12 апреля 2026 г.", Description: "В преддверии праздника Пасхи для воспитанников были организованы тематические занятия и творческие мастерские. Дети украшали пасхальные яйца, знакомились с традициями праздника и создавали поделки своими руками.", Image: "assets/23.png"},
		},
		applications:   make(map[int]models.Application),
		questions:      make(map[int]models.Question),
		nextAppID:      1,
		nextQuestionID: 1,
	}
}

func (s *Store) Home() models.HomeResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return models.HomeResponse{Hero: s.hero, Advantages: append([]models.Advantage(nil), s.advantages...), Rooms: append([]models.Room(nil), s.rooms...)}
}

func (s *Store) Programs() models.ProgramsResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return models.ProgramsResponse{Groups: append([]models.Group(nil), s.groups...), Activities: append([]models.Activity(nil), s.activities...), Teachers: append([]models.Teacher(nil), s.teachers...)}
}

func (s *Store) Parents() models.ParentsResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return models.ParentsResponse{Reviews: append([]models.Review(nil), s.reviews...), Events: append([]models.Event(nil), s.events...), News: append([]models.News(nil), s.news...)}
}

func (s *Store) AddApplication(req models.ApplicationRequest) models.Application {
	s.mu.Lock()
	defer s.mu.Unlock()
	app := models.Application{ID: s.nextAppID, Parent: req.Parent, Phone: req.Phone, ChildName: req.ChildName, ChildAge: req.ChildAge, Comment: req.Comment, Status: "new", CreatedAt: time.Now()}
	s.applications[app.ID] = app
	s.nextAppID++
	return app
}

func (s *Store) Applications() []models.Application {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]models.Application, 0, len(s.applications))
	for _, item := range s.applications {
		items = append(items, item)
	}
	return items
}

func (s *Store) AddQuestion(req models.QuestionRequest) models.Question {
	s.mu.Lock()
	defer s.mu.Unlock()
	q := models.Question{ID: s.nextQuestionID, Name: req.Name, Phone: req.Phone, Text: req.Text, Status: "new", CreatedAt: time.Now()}
	s.questions[q.ID] = q
	s.nextQuestionID++
	return q
}

func (s *Store) Questions() []models.Question {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]models.Question, 0, len(s.questions))
	for _, item := range s.questions {
		items = append(items, item)
	}
	return items
}

func (s *Store) Search(query string) []models.SearchResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	q := strings.TrimSpace(strings.ToLower(query))
	if q == "" {
		return []models.SearchResult{}
	}

	results := make([]models.SearchResult, 0)
	add := func(itemType, title, description, url string) {
		text := strings.ToLower(title + " " + description)
		if strings.Contains(text, q) {
			results = append(results, models.SearchResult{Type: itemType, Title: title, Description: description, URL: url})
		}
	}

	add("page", s.hero.Title, s.contact.Address+" "+s.contact.Phone, "index.html")
	for _, item := range s.advantages {
		add("advantage", item.Title, item.Description, "index.html#advantages")
	}
	for _, item := range s.rooms {
		add("room", item.Title, item.Description, "index.html#rooms")
	}
	for _, item := range s.groups {
		add("group", item.Title, item.Age+" "+item.Description, "programs.html#groups")
	}
	for _, item := range s.activities {
		add("activity", item.Title, item.Description, "programs.html#activities")
	}
	for _, item := range s.teachers {
		add("teacher", item.Name, item.Position+" "+item.Description, "programs.html#teachers")
	}
	for _, item := range s.reviews {
		add("review", item.Name, item.Text, "parents.html#reviews")
	}
	for _, item := range s.events {
		add("event", item.Title, item.Date+" "+item.Description, "parents.html#events")
	}
	for _, item := range s.news {
		add("news", item.Title, item.Date+" "+item.Description, "parents.html#news")
	}
	return results
}
