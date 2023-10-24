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
(email, name, password, role)
VALUE ('alex@x.com', 'Alex', 'fjefiwjflskdjwi13549kde', 'admin');

User registration:
POST
http://localhost:8000/api/auth/register
{
  "email": "alex@x.com",
  "name": "Alex",
  "password": "hdjsewiuh",
  "passwordConfirm": "hdjsewiuh"
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
