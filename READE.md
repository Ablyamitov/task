# Форма регистрации

## Запуск приложения

1. **Настройка БД:**
    - По умолчанию используется БД `task`.
    - URL к БД можно задать в `config.yaml`:
      ```yaml
      db:
        url: "postgresql://<username>:<password>@<host>:<port>/<dbname>"
      ```
    - Необходимо запустить `init.sql` для создания таблицы и вставки админа 

2. **Запуск приложения через Makefile:**
    - Сборка:
      ```bash
      make build
      ```
    - Запуск:
      ```bash
      make run
      ```

---

## Основные API-эндпоинты

### 1. Регистрация пользователя

**POST** `/register`

**Пример запроса:**
```json
{
  "last_name": "Doe",
  "first_name": "John",
  "gender": "Male",
  "birth_date": "1990-01-01",
  "phone": "+1234567890"
}
```
**Пример ответа:**
```json
{
    "data": {
        "Status": true
    },
    "status": true,
    "errors": null
}
```

### 2. Логин (без пароля)

**POST** `/login`

**Пример запроса:**
```json
{
   "phone"  : "+79785859202"
}
```
**Пример ответа:**
```json
{
   "data": {
      "Status": true
   },
   "status": true,
   "errors": null
}
```

> **Примечание:** В случае успешного логина в *хедере* ответа приходит jwt токен, который необходим для проверки роли

### 3. Получение всех участников (только для администратора)

**GET** `/users`

**Пример запроса:**

Тип авторизации: Bearer

В хедер запроса должен быть включен jwt токен в виде: 
```
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI4MTcyNDcsImlkIjozLCJyb2xlIjoiUm9sZV9Vc2VyIn0.E9nzLIsrcIlGWVJ4WSTdl7nxKzKTVDOXT2rVCWSrs3k
```

**Пример ответа:**
```json
{
   "data": {
      "Users": [
         {
            "id": 2,
            "last_name": "Ablyamitov",
            "first_name": "Enver",
            "gender": "Male",
            "birth_date": "14-08-2003",
            "phone": "+79785859101",
            "role": "Role_User"
         },
         {
            "id": 3,
            "last_name": "Ivanod",
            "first_name": "Alexey",
            "gender": "Male",
            "birth_date": "14-08-2010",
            "phone": "+79785859202",
            "role": "Role_User"
         }
      ]
   },
   "status": true,
   "errors": null
}
```

> **Примечание:** телефон админа: +79787678178

