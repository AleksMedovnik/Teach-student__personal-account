docker-compose up -d
air 

POST
http://localhost:8000/api/auth/register
{
  "email": "admin@admin.com",
  "name": "Admin",
  "password": "password123",
  "passwordConfirm": "password123"
}

POST
http://localhost:8000/api/auth/login
{
  "email": "admin@admin.com",
  "password": "password123"
}

GET 
http://localhost:8000/api/users/me

GET 
http://localhost:8000/api/auth/logout