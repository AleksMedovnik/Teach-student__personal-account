<h3>Run database-server</h3>
<code>docker-compose up -d</code>

<h3>Stop postgres-server</h3>
<code>docker-compose down</code>

<h3>Run app</h3>
<code>air</code>

<h3>Viewing the database from the terminal</h3>
<code>docker exec -it {your postgres container name} psql -U {your db username} {your db name}</code>

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
Only the administrator can register and get a student and a teacher
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
<code>http://localhost:8000/api/users/</code>

<h3>Get own profile:</h3>
<h4>GET</h4>
<code>http://localhost:8000/api/users/profile</code>

<h4>GET</h4>
<code>http://localhost:8000/api/auth/logout</code>

<h3>Create, get, update, delete groups:</h3>
Only the administrator can create, get, update and delete student groups.
<h4>Authorization -> Bearer</h4>
Token: admin_token
<h4>POST</h4>
<code>http://localhost:8000/api/groups/</code>
<code>
{
  "number": 130,
  "users": [
      {
          "id": "a3af738f-378a-44f2-9c41-c9ec1a9059d4",
          "name": "Alex",
          "email": "aleksander@gmail.com",
          "role": "admin",
          "photo": "default.png",
          "date_birth": "0001-01-01T00:00:00Z",
          "created_at": "2023-11-09T19:34:32.775224+03:00",
          "updated_at": "2023-11-09T19:34:32.775224+03:00"
      },
      {
          "id": "dbee3c3f-cad4-4680-9ec5-561f94022656",
          "name": "Pasha",
          "email": "pasha@gmail.com",
          "role": "teacher",
          "date_birth": "1983-01-01T00:00:00Z",
          "surname": "Nikolaevich",
          "country": "Russia",
          "city": "Moscow",
          "contacts": "tel: +12345687952",
          "created_at": "2023-11-09T19:35:17.505857+03:00",
          "updated_at": "2023-11-09T19:35:17.505857+03:00"
      },
      {
          "id": "a563435b-6a1d-458d-8edf-e7b9b242554a",
          "name": "Masha",
          "email": "masha@gmail.com",
          "role": "teacher",
          "date_birth": "1983-01-01T00:00:00Z",
          "surname": "Nikolaevna",
          "country": "Russia",
          "city": "Moscow",
          "contacts": "tel: +12345687952",
          "created_at": "2023-11-09T19:35:17.505857+03:00",
          "updated_at": "2023-11-09T19:35:17.505857+03:00"
      }
  ]
}
</code>

<h4>GET</h4>
<code>http://localhost:8000/api/groups/</code>

<h4>GET</h4>
<code>http://localhost:8000/api/groups/1/</code>

<h4>PUT</h4>
<code>http://localhost:8000/api/groups/1/</code>

<h4>DELETE</h4>
<code>http://localhost:8000/api/groups/1/</code>

<h4>GET</h4>
<code>http://localhost:8000/api/groups/1/users</code>