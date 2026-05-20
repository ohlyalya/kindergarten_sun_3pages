const API_BASE = localStorage.getItem('SUN_API_BASE') || 'http://localhost:8080/api';
const app = document.getElementById('app');
const currentPage = app?.dataset.page || 'home';

const icons = {
  search: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><circle cx="11" cy="11" r="7"/><path d="m16.5 16.5 4 4"/></svg>`,
  eye: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><path d="M3 12s3-6 9-6 9 6 9 6-3 6-9 6-9-6-9-6Z"/><circle cx="12" cy="12" r="2.8"/></svg>`,
  pin: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><path d="M12 21s7-6.2 7-12a7 7 0 0 0-14 0c0 5.8 7 12 7 12Z"/><circle cx="12" cy="9" r="2"/></svg>`,
  clock: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 2"/></svg>`,
  phone: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><path d="M22 16.9v3a2 2 0 0 1-2.2 2 19.8 19.8 0 0 1-8.6-3.1 19.5 19.5 0 0 1-6-6A19.8 19.8 0 0 1 2.1 4.2 2 2 0 0 1 4.1 2h3a2 2 0 0 1 2 1.7c.1 1 .4 1.9.7 2.8a2 2 0 0 1-.4 2.1L8.1 9.9a16 16 0 0 0 6 6l1.3-1.3a2 2 0 0 1 2.1-.4c.9.3 1.8.6 2.8.7a2 2 0 0 1 1.7 2Z"/></svg>`,
  mail: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><rect x="3" y="5" width="18" height="14" rx="2"/><path d="m3 7 9 6 9-6"/></svg>`,
  chevronRight: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><path d="m9 18 6-6-6-6" stroke-width="2"/></svg>`,
  chevronLeft: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><path d="m15 18-6-6 6-6" stroke-width="2"/></svg>`
};

const groupIconFiles = {
  pacifier: 'yasli.svg',
  baby: 'mladshie.svg',
  child: 'srednie.svg',
  kid: 'starshie.svg',
  student: 'podgotovitilnie.svg',

  yasli: 'yasli.svg',
  mladshie: 'mladshie.svg',
  srednie: 'srednie.svg',
  starshie: 'starshie.svg',
  podgotovitilnie: 'podgotovitilnie.svg',
};

function renderGroupIcon(iconKey, title) {
  const fileName = groupIconFiles[iconKey] || 'srednie.svg';

  return `
    <img src="assets/${fileName}" alt="${esc(title)}">
  `;
}

function esc(value) {
  return String(value ?? '').replace(/[&<>'"]/g, char => ({'&':'&amp;','<':'&lt;','>':'&gt;',"'":'&#39;','"':'&quot;'}[char]));
}

function api(path, options = {}) {
  return fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...(options.headers || {}) },
    ...options,
  }).then(async response => {
    const data = await response.json().catch(() => ({}));
    if (!response.ok) throw new Error(data.error || 'Ошибка запроса');
    return data;
  });
}

function renderHeader() {
  const nav = [
    ['index.html', 'Главная', 'home'],
    ['programs.html', 'Образовательные программы', 'programs'],
    ['parents.html', 'Для родителей', 'parents'],
    ['#contacts', 'Контакты', 'contacts']
  ];

  document.getElementById('siteHeader').innerHTML = `
    <header class="header">
      <div class="header__inner">
        <div class="header__top">
          <a href="index.html" class="brand" aria-label="На главную">
            <img class="brand__logo" src="assets/logo.png" alt="Логотип ГБДОУ Солнце">
            <div>
              <h1 class="brand__name">ГБДОУ «Солнце»</h1>
              <div class="brand__place">г. Москва, р-н Коптево</div>
            </div>
          </a>
          <form class="search" id="searchForm">
            ${icons.search}
            <input id="searchInput" type="search" placeholder="Поиск" aria-label="Поиск по сайту">
          </form>
          <a class="accessibility" href="#" aria-label="Версия для слабовидящих">${icons.eye}<span>Версия для слабовидящих</span></a>
        </div>
        <div class="header__bottom">
          <nav class="nav">
            ${nav.map(([href, label, key]) => `<a class="${key === currentPage ? 'active' : ''}" href="${href}">${label}</a>`).join('')}
          </nav>
          <button class="btn" data-open-application>Записаться</button>
        </div>
      </div>
    </header>`;
}

function renderFooter() {
  document.getElementById('siteFooter').innerHTML = `
    <footer class="footer" id="contacts">
      <div class="footer__inner">
        <div class="footer__col">
        <a href="index.html">Главная</a>
        <a href="programs.html">Образовательные программы</a>
        <a href="#" class="accessibility-link">
         <img src="assets/eyes.svg" alt="" class="accessibility-icon">
         <span>Версия для слабовидящих</span>
         </a>
         </div>
        <div class="footer__col">
          <a href="parents.html">Для родителей</a>
          <a href="#">Документы</a>
        </div>
        <div class="footer__col">
          <a href="#contacts">Контакты</a>
          <a href="#">Политика конфиденциальности</a>
          <a href="#">Пользовательское соглашение</a>
        </div>
        <div class="footer__col footer__social">
          <a href="#" class="social-vk">
          <img src="assets/vk.svg" alt="VK">
          </a>
          <span class="counter">5 3 1<br>1 9 8</span>
        </div>
      </div>
    </footer>`;
}

function showLoader() {
  app.innerHTML = `<div class="main-space"><div class="container"><div class="loader">Загрузка данных...</div></div></div>`;
}

function showError(error) {
  app.innerHTML = `<div class="main-space"><div class="container"><div class="error-block">Не удалось загрузить данные с backend.<br>${esc(error.message)}<br><br>Проверьте, что Go-сервер запущен на http://localhost:8080</div></div></div>`;
}

async function renderHome() {
  showLoader();
  const data = await api('/home');
  const hero = data.hero;
  const contact = hero.contactInfo;
  const advantages = data.advantages;
  const rooms = data.rooms;

  app.innerHTML = `
    <div class="main-space">
      <div class="container">
        <section class="card hero-card">
          <div class="hero-card__content">
            <h2 class="hero-card__title">${esc(hero.title)}</h2>
            <div class="info-grid">
              <div class="info-item">${icons.pin}<div><div class="info-label">Адрес</div><div class="info-text">${esc(contact.address)}</div></div></div>
              <div class="info-item">${icons.clock}<div><div class="info-label">Режим работы</div><div class="info-text">${esc(contact.hours)}</div></div></div>
              <div class="info-item">${icons.phone}<div><div class="info-label">Телефон</div><div class="info-text">${esc(contact.phone)}</div></div></div>
              <div class="info-item">${icons.mail}<div><div class="info-label">Почта</div><div class="info-text">${esc(contact.email)}</div></div></div>
            </div>
            <button class="btn btn--wide" data-open-application>Записаться в детский сад</button>
          </div>
          <img class="hero-card__image" src="${esc(hero.image)}" alt="Дети в детском саду">
        </section>

        <section class="section card advantages-card" id="advantages">
          <div class="advantages-grid">
            ${advantages.map((item, index) => `
              <article class="advantage ${index < 2 ? 'advantage--big' : 'advantage--small'}">
                <img src="${esc(item.image)}" alt="${esc(item.title)}">
                <h3>${esc(item.title)}</h3>
                <p>${esc(item.description)}</p>
              </article>`).join('')}
          </div>
        </section>

        <section class="section" id="rooms">
          <h2 class="section__title">Наши помещения</h2>
          <div class="carousel-row">
            <button class="arrow" aria-label="Назад">${icons.chevronLeft}</button>
            <div class="room-grid">
              ${rooms.map(room => `
                <article class="room-card">
                  <img src="${esc(room.image)}" alt="${esc(room.title)}">
                  <h3>${esc(room.title)}</h3>
                </article>`).join('')}
            </div>
            <button class="arrow" aria-label="Вперёд">${icons.chevronRight}</button>
          </div>
        </section>
      </div>
    </div>`;
}

async function renderPrograms() {
  showLoader();
  const data = await api('/programs');
  const director = data.teachers.find(t => t.isMain);
  const teachers = data.teachers.filter(t => !t.isMain);

  app.innerHTML = `
    <div class="main-space">
      <div class="container">
        <section id="groups" class="section--small-top">
          <h2 class="section__title">Группы детского сада</h2>
          <div class="group-grid">
            ${data.groups.map(group => `
              <article class="group-card" title="${esc(group.description)}">
                <div class="group-card__icon">
                ${renderGroupIcon(group.icon, group.title)}
                </div>
                <div><h3>${esc(group.title)}</h3><p>${esc(group.age)}</p></div>
                <div class="group-card__arrow">${icons.chevronRight}</div>
              </article>`).join('')}
          </div>
        </section>

        <section id="activities" class="section">
          <h2 class="section__title">Дополнительные занятия</h2>
          <div class="activities-grid">
            ${data.activities.map(item => `
              <article class="activity-card">
                <img src="${esc(item.image)}" alt="${esc(item.title)}">
                <div class="activity-card__body"><h3>${esc(item.title)}</h3><p>${esc(item.description)}</p></div>
              </article>`).join('')}
          </div>
          <div class="center-actions"><button class="btn btn--wide">Посмотреть все занятия</button></div>
        </section>

        <section id="teachers" class="section">
          <h2 class="section__title">Наш коллектив</h2>
          ${director ? `
            <article class="card director-card">
              <img src="${esc(director.image)}" alt="${esc(director.name)}">
              <div>
                <h3 class="director-name">${director.name.split(' ').map(esc).join('<br>')}</h3>
                <div class="director-position">${esc(director.position)}</div>
              </div>
              <div class="director-message">
                <p>Уважаемые родители и гости!</p>
                <p>Рада приветствовать вас в нашем детском саду. Мы стремимся создать тёплую и безопасную атмосферу, где каждый ребёнок чувствует себя как дома, а также получает возможность развиваться, учиться новому и находить первых друзей.</p>
                <p>${esc(director.description)}</p>
                <button class="btn" data-open-question>Задать вопрос</button>
              </div>
            </article>` : ''}
          <div class="teacher-grid">
            ${teachers.map(teacher => `
              <article class="teacher-card">
                <div class="teacher-card__top">
                  <img src="${esc(teacher.image)}" alt="${esc(teacher.name)}">
                  <h3>${teacher.name.split(' ').map(esc).join('<br>')}</h3>
                </div>
                <p>${esc(teacher.position)}</p>
                <button class="btn btn--outline">Читать полностью</button>
              </article>`).join('')}
          </div>
          <div class="center-actions"><button class="btn btn--wide">Показать всех</button></div>
        </section>
      </div>
    </div>`;
}

async function renderParents() {
  showLoader();
  const data = await api('/parents');

  app.innerHTML = `
    <div class="main-space">
      <div class="container">
        <section id="reviews" class="section--small-top">
          <h2 class="section__title">Отзывы</h2>
          <div class="review-grid">
            ${data.reviews.map(review => `
              <article class="review-card">
                <div class="review-card__head"><img src="${esc(review.avatar)}" alt="${esc(review.name)}"><h3>${esc(review.name)}</h3></div>
                <p>${esc(review.text)}</p>
              </article>`).join('')}
          </div>
          <div class="center-actions"><button class="btn btn--wide">Читать все отзывы</button></div>
        </section>

        <section id="events" class="section">
          <h2 class="section__title">Мероприятия</h2>
          <div class="event-carousel">
            <button class="arrow arrow--blue" aria-label="Назад">${icons.chevronLeft}</button>
            <div class="event-grid">
              ${data.events.map(event => `
                <article class="event-card">
                  <img src="${esc(event.image)}" alt="${esc(event.title)}">
                  <h3>${esc(event.title)}</h3>
                  <div class="date">${esc(event.date)}</div>
                  <p>${esc(event.description)}</p>
                </article>`).join('')}
            </div>
            <button class="arrow arrow--blue" aria-label="Вперёд">${icons.chevronRight}</button>
          </div>
        </section>

        <section id="news" class="section">
          <h2 class="section__title">Новости и события</h2>
          <div class="news-grid">
            ${data.news.map(item => `
              <article class="news-card">
                <img src="${esc(item.image)}" alt="${esc(item.title)}">
                <div class="news-card__body">
                  <h3>${esc(item.title)}</h3>
                  <div class="date">${esc(item.date)}</div>
                  <p>${esc(item.description)}</p>
                  <button class="btn btn--outline">Читать далее</button>
                </div>
              </article>`).join('')}
          </div>
          <div class="center-actions"><button class="btn btn--wide">Посмотреть все новости</button></div>
        </section>
      </div>
    </div>`;
}

function openApplicationModal() {
  openModal(`
    <div class="modal__head"><h2>Запись в детский сад</h2><button class="modal__close" data-close-modal>×</button></div>
    <form class="form-grid" id="applicationForm">
      <div class="field"><label>Имя родителя</label><input name="parent" required placeholder="Анна Иванова"></div>
      <div class="field"><label>Телефон</label><input name="phone" required placeholder="+7 900 123-45-67"></div>
      <div class="field"><label>Имя ребёнка</label><input name="childName" required placeholder="София"></div>
      <div class="field"><label>Возраст ребёнка</label><input name="childAge" required type="number" min="1" max="7" placeholder="4"></div>
      <div class="field"><label>Комментарий</label><textarea name="comment" placeholder="Например: хотим попасть в среднюю группу"></textarea></div>
      <button class="btn" type="submit">Отправить заявку</button>
      <div id="modalAlert" class="alert"></div>
    </form>`);

  document.getElementById('applicationForm').addEventListener('submit', async event => {
    event.preventDefault();
    const form = new FormData(event.target);
    const payload = Object.fromEntries(form.entries());
    payload.childAge = Number(payload.childAge);
    await submitModal('/applications', payload, 'Заявка сохранена. Администратор свяжется с вами.');
  });
}

function openQuestionModal() {
  openModal(`
    <div class="modal__head"><h2>Задать вопрос</h2><button class="modal__close" data-close-modal>×</button></div>
    <form class="form-grid" id="questionForm">
      <div class="field"><label>Имя</label><input name="name" required placeholder="Мария"></div>
      <div class="field"><label>Телефон</label><input name="phone" required placeholder="+7 900 987-65-43"></div>
      <div class="field"><label>Вопрос</label><textarea name="text" required placeholder="Напишите вопрос администрации"></textarea></div>
      <button class="btn" type="submit">Отправить вопрос</button>
      <div id="modalAlert" class="alert"></div>
    </form>`);

  document.getElementById('questionForm').addEventListener('submit', async event => {
    event.preventDefault();
    const payload = Object.fromEntries(new FormData(event.target).entries());
    await submitModal('/questions', payload, 'Вопрос сохранён. Мы свяжемся с вами.');
  });
}

async function submitModal(path, payload, successText) {
  const alert = document.getElementById('modalAlert');
  alert.className = 'alert';
  alert.textContent = 'Отправка...';
  try {
    await api(path, { method: 'POST', body: JSON.stringify(payload) });
    alert.className = 'alert alert--ok';
    alert.textContent = successText;
  } catch (error) {
    alert.className = 'alert alert--error';
    alert.textContent = error.message;
  }
}

function openModal(content) {
  const root = document.getElementById('modalRoot');
  root.innerHTML = `<div class="modal-backdrop"><div class="modal">${content}</div></div>`;
  root.querySelector('[data-close-modal]')?.addEventListener('click', closeModal);
  root.querySelector('.modal-backdrop')?.addEventListener('click', event => {
    if (event.target.classList.contains('modal-backdrop')) closeModal();
  });
}

function closeModal() {
  document.getElementById('modalRoot').innerHTML = '';
}

function bindGlobalEvents() {
  document.addEventListener('click', event => {
    if (event.target.closest('[data-open-application]')) openApplicationModal();
    if (event.target.closest('[data-open-question]')) openQuestionModal();
  });

  document.getElementById('searchForm')?.addEventListener('submit', async event => {
    event.preventDefault();
    const query = document.getElementById('searchInput').value.trim();
    if (!query) return;
    try {
      const results = await api(`/search?q=${encodeURIComponent(query)}`);
      openModal(`
        <div class="modal__head"><h2>Результаты поиска</h2><button class="modal__close" data-close-modal>×</button></div>
        ${results.length ? results.map(item => `<a class="search-result" href="${esc(item.url)}"><b>${esc(item.title)}</b><span>${esc(item.description)}</span></a>`).join('') : '<p>Ничего не найдено.</p>'}`);
    } catch (error) {
      openModal(`<div class="modal__head"><h2>Поиск</h2><button class="modal__close" data-close-modal>×</button></div><p>${esc(error.message)}</p>`);
    }
  });
}

async function init() {
  renderHeader();
  renderFooter();
  bindGlobalEvents();

  try {
    if (currentPage === 'home') await renderHome();
    if (currentPage === 'programs') await renderPrograms();
    if (currentPage === 'parents') await renderParents();
  } catch (error) {
    showError(error);
  }
}

init();
