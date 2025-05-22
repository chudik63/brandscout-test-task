# Brandscout Test Task
RESTful API for storing quotes with the following options:
1. Добавление новой цитаты (POST /quotes)​
2. Получение всех цитат (GET /quotes) ​
3. Получение случайной цитаты (GET /quotes/random) ​
4. Фильтрация по автору (GET /quotes?author=Confucius) ​
5. Удаление цитаты по ID (DELETE /quotes/{id}) ​

## Installation
```
git clone https://github.com/chudik63/brandscout-test-task.git
cd brandscout-test-task
```

## Building
Using Docker:
```
docker build -t brandscout_test_task .
docker run -p 8080:8080 brandscout_test_task
```

## Documentation

### Add a quote
Request:
```bash
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{
    "author":"Confucius", 
    "quote":"Life is simple, but we insist on making it complicated."
}'
```
Responses:
- 201 Created
- 400 Invalid input

### Get all quotes
Request:
```bash
curl http://localhost:8080/quotes
```

Responses:
- 200 OK
```json
{
    {
        "id": 1,
        "author": "Author1",
        "quote": "Quote1",
        "created_at": "2023-01-01"
    },
}
```
- 404 Not found

### Get a random quote
Request:
```bash
curl http://localhost:8080/quotes/random ​
```

Responses:
- 200 OK
```json
{
    "id": 1,
    "author": "Author1",
    "quote": "Quote1",
    "created_at": "2023-01-01"
},
```
- 404 Not found

### Get a quote by Author
Request:
```bash
curl http://localhost:8080/quotes?author=Author2 ​
```

Responses:
- 200 OK
```json
{
    "id": 1,
    "author": "Author2",
    "quote": "Quote2",
    "created_at": "2023-01-02"
},
```
- 404 Not found

### Delete a quote
Request:
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

Responses:
- 200 OK
- 404 Not found
- 404 Invalid ID