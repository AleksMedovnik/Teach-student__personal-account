<h3>Run database-server</h3>
docker-compose up -d

<h3>Stop postgres-server</h3>
docker-compose down

<h3>Run app</h3>
air

<h3>Viewing the database from the terminal</h3>
docker exec -it {your postgres container name} psql -U {your db username} {your db name}

<h3>Exit</h3>
exit

<h3>Admin registration:</h3>
POST
http://localhost:8000/api/auth/admin
<code>
{
  "name": "Alex",
  "email": "aleksander@gmail.com",
  "password": "jfosdieipaoeik3558",
  "passwordConfirm": "jfosdieipaoeik3558",
  "role": "admin"
}
</code>

Login admin & users
POST
http://localhost:8000/api/auth/login
{
  "email": "aleksander@gmail.com",
  "password": "jfosdieipaoeik3558"
}

User registration:
Only the administrator can register a student and a teacher
Authorization -> Bearer
Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg2MTI2OTYsImlhdCI6MTY5ODYwOTA5NiwibmJmIjoxNjk4NjA5MDk2LCJzdWIiOiIzMjkzYWE1OC0wODJjLTQ1ZjItYTkxNC05NWFjZjNkYWJjMjcifQ.-jrv8EMA0J6luTW-48aEMzCoccaZlqgxkqm5ir_9kds

POST
http://localhost:8000/api/auth/register
{
  "name": "Masha",
  "email": "masha@gmail.com",
  "password": "jfosdieipaoeik3558",
  "passwordConfirm": "jfosdieipaoeik3558",
  "role": "teacher",
  "surname": "Nikolaevna",
  "country": "Russia",
  "city": "Moscow",
  "contacts": "tel: +12345687952",
  "date_birth": "1983-01-01T00:00:00Z"
}

GET
http://localhost:8000/api/users/profile

GET
http://localhost:8000/api/auth/logout
