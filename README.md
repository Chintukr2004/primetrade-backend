# PrimeTrade - Secure Backend API & Dashboard

A scalable RESTful API built with Go and PostgreSQL, featuring JWT-based authentication, role-based access control, and complete CRUD operations for task management. Includes a lightweight Vanilla JS frontend to interact with the API.

## Tech Stack
* **Backend:** Go (Golang 1.22+)
* **Database:** PostgreSQL
* **Authentication:** JWT (JSON Web Tokens) & Bcrypt Password Hashing
* **Frontend:** HTML/CSS & Vanilla JavaScript (Fetch API)

## Setup Instructions

### 1. Database Setup
Create a local PostgreSQL database named `intern_db` and execute the following SQL to create the tables:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
### 2. The Postman Collection (`postman_collection.json`)
Instead of manually clicking through Postman to export a file, I wrote the JSON format for you. Create a file named `Primetrade_API.postman_collection.json` in your project root and paste this in:

```json
{
	"info": {
		"name": "PrimeTrade API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Auth",
			"item": [
				{
					"name": "Register",
					"request": {
						"method": "POST",
						"header": [],
						"body": {"mode": "raw","raw": "{\n    \"email\": \"test@example.com\",\n    \"password\": \"password123\"\n}","options": {"raw": {"language": "json"}}},
						"url": {"raw": "http://localhost:8080/api/v1/register","protocol": "http","host": ["localhost"],"port": "8080","path": ["api","v1","register"]}
					}
				},
				{
					"name": "Login",
					"request": {
						"method": "POST",
						"header": [],
						"body": {"mode": "raw","raw": "{\n    \"email\": \"test@example.com\",\n    \"password\": \"password123\"\n}","options": {"raw": {"language": "json"}}},
						"url": {"raw": "http://localhost:8080/api/v1/login","protocol": "http","host": ["localhost"],"port": "8080","path": ["api","v1","login"]}
					}
				}
			]
		},
		{
			"name": "Tasks",
			"item": [
				{
					"name": "Create Task",
					"request": {
						"method": "POST",
						"header": [{"key": "Authorization","value": "Bearer YOUR_TOKEN_HERE"}],
						"body": {"mode": "raw","raw": "{\n    \"title\": \"Finish Assignment\"\n}","options": {"raw": {"language": "json"}}},
						"url": {"raw": "http://localhost:8080/api/v1/tasks","protocol": "http","host": ["localhost"],"port": "8080","path": ["api","v1","tasks"]}
					}
				},
				{
					"name": "Get Tasks",
					"request": {
						"method": "GET",
						"header": [{"key": "Authorization","value": "Bearer YOUR_TOKEN_HERE"}],
						"url": {"raw": "http://localhost:8080/api/v1/tasks","protocol": "http","host": ["localhost"],"port": "8080","path": ["api","v1","tasks"]}
					}
				}
			]
		}
	]
}

![Dashboard Preview](dashboard.png)