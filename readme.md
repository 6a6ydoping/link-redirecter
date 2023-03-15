<h3>Для выполнения задания я выбрал: Postgresql, Gorm</h3>
<h3>Soft: Goland, PgAdmin4, Postman</h3>

<h3>Выполненные обязательные требования - 100%</h3>
<h3>Выполненные опциональные требования:</h3>
    <li>При старте приложения кэш асинхронно "прогревается" </li>
    <li>Реализован механизм Time To Live для кэша </li>

<h3>Запросы к эндпоинтам:</h3>
<h4>GET/admin/redirects</h4>
Query: page, pageSize (default: 1, 100)
---
<h4>GET/admin/redirects/{id}</h4>

---
<h4>POST/admin/redirects</h4>
Request: JSON{active_link, history_link}
---
<h4>PATCH/admin/redirects/{id}</h4>
Request: JSON {active_link}
---
<h4>DELETE/admin/redirects/{id}</h4>