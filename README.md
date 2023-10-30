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
<code>
  POST 
</br>  
http://localhost:8000/api/auth/admin
</br>  
{
  "name": "Alex",
  "email": "aleksander@gmail.com",
  "password": "jfosdieipaoeik3558",
  "passwordConfirm": "jfosdieipaoeik3558",
  "role": "admin"
}
</code>

<h3>Login admin & users</h3>
<code>
POST
</br>  
http://localhost:8000/api/auth/login
</br>  
{
"email": "aleksander@gmail.com",
"password": "jfosdieipaoeik3558"
}
</br>  

User registration:
Only the administrator can register a student and a teacher
Authorization -> Bearer
Token: admin_token
</br>  
POST
</br>  
http://localhost:8000/api/auth/register
</br>  
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
</br>  


GET
</br>  
http://localhost:8000/api/users/profile
</br>  

GET
</br>  
http://localhost:8000/api/auth/logout

</code>