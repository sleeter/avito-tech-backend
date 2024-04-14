# Бэкенд для сервиса баннеров

## Описание
В этом репозитории представлена реализация бэкенда для сервиса баннеров. Взаимодействие с сервисом происходит следующим образом:
1. Администраторы загружают баннеры. Также можно изменять и удалять информацию о ранее загруженных баннерах.
2. Пользователи получают баннеры с необходимыми идентификаторами.

## Решение
Сервис написан на Golang с использованием gin-gonic, pgx, migrate, viper, squirrel, а также базовых библиотек. Для хранения данных используется PostgreSQL, в котором создана одна таблица для баннеров и индекс для эффективного поиска данных.

## Сборка и Деплой
Все необходимые операции осуществляются с помощью Makefile и Docker.
Чтобы запустить сервис, выполните:
```make up```

Чтобы остановить сервис, выполните:
```make down```

## Пример запроса
![Post_example](https://github.com/sleeter/avito-tech-backend/raw/master/post_example.png)
