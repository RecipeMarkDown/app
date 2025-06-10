# Database Documentation

## Overview

The application uses PostgreSQL with GORM as the ORM for database operations.

## Connection

Database connection is configured in backend/internal/config/database.go using environment variables

## Models

### User Model

Located at: backend/internal/models/user.go

### Recipe Model

## GORM Operations

### Create

`db.Create(&user)`

#### Generated SQL

INSERT INTO users (google_id, first_name, last_name, email, created_at, updated_at)
VALUES ('google_123', 'Brandon', 'Shippy', 'brandon@example.com', NOW(), NOW());

### Read

#### Find by ID

`db.First(&user, 1)`

#### Find by email

`db.Where("email = ?", "brandon@example.com").First(&user)`

#### Find by Google ID

`db.Where("google_id = ?", "google_123").First(&user)`

##### Generated SQL

-- By ID
SELECT \* FROM users WHERE id = 1 LIMIT 1;

-- By email
SELECT \* FROM users WHERE email = 'brandon@example.com' LIMIT 1;

-- By Google ID
SELECT \* FROM users WHERE google_id = 'google_123' LIMIT 1;

### Update

#### Update specific fields

db.Model(&user).Updates(User{FirstName: "New Name", Email: "new@example.com"})

#### Update all fields

user.FirstName = "Updated"
db.Save(&user)

##### Generated SQL

-- Partial update
UPDATE users SET first_name = 'New Name', email = 'new@example.com', updated_at = NOW()
WHERE id = 1;

-- Full save
UPDATE users SET google_id = '...', first_name = 'Updated', ... , updated_at = NOW()
WHERE id = 1;

### Delete

#### Soft delete

db.Delete(&user, 1)

#### Hard delete

db.Unscoped().Delete(&user, 1)

## Database Setup

1. Install PostgreSQL
2. Create database and user:
   CREATE USER recipeuser WITH PASSWORD 'recipepass';
   CREATE DATABASE recipeapp OWNER recipeuser;
   GRANT ALL PRIVILEGES ON DATABASE recipeapp TO recipeuser;
3. Set environment variables in config/.env (you might need to make this)

## Auto-Migration

GORM automatically creates/updates tables when the application starts:
db.AutoMigrate(&models.User{})

This will:

- Create users table if it doesn't exist
- Add missing columns if model is updated
- Never removes columns or data
