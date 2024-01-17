# GoHFLabsParcer

- [GoHFLabsParcer](#gohflabsparcer)
  - [Описание](#описание)
  - [Инструкция по применению](#инструкция-по-применению)
  - [Сборка](#сборка)
  - [Структура проекта](#структура-проекта)
  - [Илюстрация работы](#илюстрация-работы)
  - [Разработчики](#разработчики)

## Описание
По техническому задания требовалось написать скрипт для парсинга таблицы с [сайта](https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999).

Я обнаружил API EndPoint позволяющий избежать прямого поиска нужных dom элементов в html коде.

[Ссылка на документ с результатом](https://docs.google.com/document/d/1ceHYcsZc3RGTz0X5zXY2vDdWeQXR3wiLWDhdPcD50XI/edit?usp=sharing)

## Инструкция по применению

- Для работы требуется иметь сервисный аккаунт в google cloud. 

*P.S: В конфиг файле, находится актуальные данные сервисного профиля, можете запустить и проверить. Не за что :)*

- Выставить переодичность выполнения можно в конфигурационной файле поле с именем "time_to_repeat", время указывается в минутах. (Стандартное время повторений 5 минут).

Более подробнее можно прочитать в файле **Info.md** лежащим в директории **/config/Info.md**.

## Сборка
**Сборка через docker:**
- docker-compose up --build
    
**Сборка без docker:**
- go build ./cmd/app/main.go (В этом случе потребуется переместить **config.json** в **/cmd/app**)

## Структура проекта
``` 
├── cmd
│   └── app
│       └── main.go
├── config
│   ├── config.go
│   ├── config.json
│   └── Info.md
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── run.go
│   ├── entity
│   │   └── hflabs.go
│   └── usecase
│       ├── gdoc.go
│       └── hflabs.go
└── Readme.md
```

## Илюстрация работы

[Watch the video](/gif/%D0%97%D0%B0%D0%BF%D0%B8%D1%81%D1%8C%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%20%D0%BE%D1%82%202023-02-14%2001-29-32.webm)
  
## Разработчики

- [xByte](https://github.com/xbytee)
