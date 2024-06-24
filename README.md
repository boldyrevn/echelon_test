# Тестовое задание в компанию "Эшелон Технологии"

## Описание
В данном репозитории хранится решение тестового задания на позицию 
Junior Golang разработчика в компанию "Эшелон Технологии". По итогам работы
удалось реализовать gRPC прокси сервер для загрузки превью для видео с платформы
YouTube и консольный клиент. 

## Структура проекта
Структура проекта выстроена в соответствии с принятыми при разработке проектов на Golang
практиками. В корне есть 3 директории:
- `/cmd` - исходники для сборки проекта (`client` и `server` для клиента и сервера соответственно)
- `/internal` - основной код программы
- `/pkg` - здесь помещены proto спецификации, т.к. в последствии они могут быть использованы
другим сервисом  

Подробнее про содержимое `/internal`:
- `/adapters/ytbclient` - здесь хранится реализация сервиса для загрузки превью
с YouTube. Имеется простой HTTP клиент и кэширующий слой. Сам кеш хранится в Redis.
- `/model` - директория с доменными моделями сервиса
- `/ports/grpc` - реализация gRPC сервера и клиента

## Запуск приложения
Для запуска клиента и сервера можно использовать Makefile, либо запустить указанные
в нем команды вручную
### Сервер
Поднимается вместе с Redis через Docker Compose. Перед запуском нужно указать необходимые
для работы приложения переменные окружения в `.env` файле в директории с `compose.yaml` файлом:
- `REDIS_PORT` - порт, который будет пробрасываться из докера для Redis
- `SERVICE_PORT` - порт, на который будет пробрасываться порт сервера. Его нужно будет
указывать при запуске CLI приложения  

После указания всех необходимых переменных нужно ввести команду
> docker compose up -d --build
### Клиент
Клиент нужно скомпилировать из директории `/cmd/client` командой
> go build ./cmd/client  

Запуск клиента для скачивания файлов осуществляется следующей командой (для Windows
к названию бинарника нужно добавить `.exe`):  
>./client --address server_address [--async] [--quality thumbnail_quality] video_url [...]  

Все возможные опции можно посмотреть введя
>./client -h

Для скачивания нескольких превью нужно ввести ссылки на видео через пробел. Сами превью
сохраняются в директории, из которой запускается клиент

## Тестирование
Для приложения написаны юнит-тесты и интеграционный тест для Redis. Сами тесты
прогоняются в Github Actions, для чего написан Workflow файл в директории
`/.github/workflows`. Для тестов применены стандартная библиотека `testing`, а также
библиотека `testify` и генератор для моков `gomock`.