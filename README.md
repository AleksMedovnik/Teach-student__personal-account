// run database-server
docker-compose up -d

// stop postgres-server
docker-compose down

// run app
air

// viewing the database from the terminal
docker exec -it {your postgres container name} psql -U {your db username} {your db name}
// exit
exit

Admin registration:
INSERT INTO users
(email, name, password, role, )
VALUE ('alex@x.com', 'Alex', 'fjefiwjflskdjwi13549kde', 'admin');

User registration:
POST
http://localhost:8000/api/auth/register
{
    "name": "Pasha",
    "email": "pasha@gmail.com",
    "password": "jfdskjfeiejijfo",
    "passwordConfirm": "jfdskjfeiejijfo",
    "role": "teacher",
    "surname": "Nikolaevna",
    "country": "Russia",
    "city": "Moscow",
    "contacts": "tel: +12345687952",
    "date_birth": "1983-01-01T00:00:00Z"
}

POST
http://localhost:8000/api/auth/login
{
  "email": "alex@x.com",
  "password": "hdjsewiuh"
}

GET
http://localhost:8000/api/users/profile

GET
http://localhost:8000/api/auth/logout
