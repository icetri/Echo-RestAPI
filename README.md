# Приложение предоставляющее REST API на Echo Framework 
##### Список запросов:
* GET localhost:8080/users - Вывести список Users
* GET localhost:8080/users/id - Вывести User по ID
* POST localhost:8080/users и тело запроса в JSON {"name":"Name"} - Добавить нового User
* PUT localhost:8080/users/id и тело запроса в form-data (name:Name) - Редактировать User по ID
* DELETE localhost:8080/users/id - Удалить User по ID
---
По умолчанию имя JSON-файла с пользователями "json.json", если нет, тогда создает его.
---