<h3>Run database-server</h3>
<code>docker-compose up -d</code>

<h3>Stop postgres-server</h3>
<code>docker-compose down</code>

<h3>Run app</h3>
<code>air</code>

<h3>Viewing the database from the terminal</h3>
<code>docker exec -it {your postgres container name} psql -U {your db username} {your db name}</code>

<h3>Exit</h3>
<code>exit</code>

<h3>Admin registration:</h3>
<h4>POST</h4>
<code>http://localhost:8000/api/auth/admin</code>
<code>
{
  "name": "Alex",
  "email": "aleksander@gmail.com",
  "password": "jfosdieipaoeik3558",
  "passwordConfirm": "jfosdieipaoeik3558",
  "role": "admin"
}
</code>

<h3>Login admin & users</h3>
<h4>POST</h4>
<code>http://localhost:8000/api/auth/login</code>
<code> 
{
  "email": "aleksander@gmail.com",
  "password": "jfosdieipaoeik3558"
}
</code>

<h3>User registration:</h3>
Only the administrator can register a student and a teacher
<h4>Authorization -> Bearer</h4>
Token: admin_token

<h4>POST</h4>
<code>http://localhost:8000/api/auth/register</code>
<code>
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
</code>

<h4>GET</h4>
<code>http://localhost:8000/api/users/profile</code>

<h4>GET</h4>
<code>http://localhost:8000/api/auth/logout</code>