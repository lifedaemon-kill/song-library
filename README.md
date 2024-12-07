Онлайн библиотеки песен 🎶

Стэк технологий:

#### go, gin, sqlx, logrus, goose, postgres

Присутствует документация swagger

```http://localhost/swaqger/index.html```

Для подключения к базе данных и внешнего сервиса требуется .env файл в корне проекта

### .env

```
#Database parameters 
USER=           #postgres
DBNAME=         #song_lib_db
PASSWORD=       #not 123456
SSLMODE=        #enable / disable
HOST=           #127.0.0.1
PORT=           #5432

#For /info requests. used in clientService
EXTERNAL_API=   #host:port 
```
