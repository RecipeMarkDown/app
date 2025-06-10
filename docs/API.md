# API Documentation

## Base URL http://localhost:8080/api

## UserAPI
It provides endpoints for managing the user accounts, all endpoints return in JSON responses

### Common Response Formats
``{
  "id": 1,
  "google_id": "google_123456789",
  "email": "brandon@example.com",
  "username": "brandon_shippy",
  "avatar": "https://example.com/avatar.jpg",
  "created_at": "2025-06-09T10:30:00Z",
  "updated_at": "2025-06-09T10:30:00Z"
}``
``{
  "error": "User not found"
}``

### User Endpoints
1. Create User
   Create a new user account.

  Method: `POST`
  Endpoint: `/users`
  Auth: None (will require OAuth soon)

  Request

  Headers:
    Content-Type: application/json

  Body Parameters:
  Field        Type    Required    Description
  email      string    Yes        Users email (unique)
  username   string    Yes        Users username (unique)
  google_id  string    Yes        Google OAuth identifier (unique)
  avataer    string    No         Profile Picture URL

  Example Request
  ``curl -X POST http://localhost:3001/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "brandon@example.com",
    "username": "brandon_shippy",
    "google_id": "google_123456789",
    "avatar": "https://lh3.googleusercontent.com/a/default-user"
  }'``

  Example Response
  Status: `201 Created`

  `{
    "id": 1,
    "google_id": "google_123456789",
    "email": "brandon@example.com",
    "username": "brandon_shippy",
    "avatar": "https://lh3.googleusercontent.com/a/default-user",
    "created_at": "2025-06-09T10:30:00Z",
    "updated_at": "2025-06-09T10:30:00Z"
  }`

  Error Responses
  Status                        Description                        Example Response
  400 Bad Request              Invalid JSON or missing fields      `{"error": "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`
  409 Conflict                 Email or GoogleID exists            `{"error": "User with this email already exists"}`
  500 Internal Server Error    Database error                      `{"error": "Failed to create user"}`

  
2. Get User by ID
   Retrieve a specific user by their unqiue ID

   Method: `GET`
   Endpoint: `/users/{id}`

   Path Parameters
   Parameter    Type     Description
   id           int      The users unique identifier

   Example Request
   Status: `200 OK`
   `{
    "id": 1,
    "google_id": "google_123456789",
    "email": "brandon@example.com",
    "username": "brandon_shippy",
    "avatar": "https://lh3.googleusercontent.com/a/default-user",
    "created_at": "2025-06-09T10:30:00Z",
    "updated_at": "2025-06-09T10:30:00Z"
  }`

  Error Responses
  Status              Description                    Example Response
  404 Not Found      User with ID doesnt exist        `{"error": "User not found"}`
  400 Bad Request    Invalid ID format                `{"error": "Invalid user ID"}`

3. Get Users (with filters)
   Query users by email or Google ID using query parameters

   Method: `GET`
   Endpoint: `/users`

   Query Parameters
   Parameter            Type      Required      Description
   email                string     No*           User's email address
   google_id            string     No*           User's Google OAuth ID
   *At least one parameter required

   Example Requests
   Get user by email:
   `curl "http://localhost:3001/api/users?email=brandon@example.com"`

   Get user by Google ID:
   `curl "http://localhost:3001/api/users?google_id=google_123456789"`

   URL encoding for special characters:
   `# If email contains special characters
    curl "http://localhost:3001/api/users?email=user%2Btest@example.com"`

   Example Response
   Status: `200 OK`
   `{
    "id": 1,
    "google_id": "google_123456789",
    "email": "brandon@example.com",
    "username": "brandon_shippy",
    "avatar": "https://lh3.googleusercontent.com/a/default-user",
    "created_at": "2025-06-09T10:30:00Z",
    "updated_at": "2025-06-09T10:30:00Z"
  }`

  Error Responses
  Status                Description                            Example Response
  400 Bad Request      No query parameters provided            `{"error": "Please specify email or google_id parameter"}`
  404 Not Found        User not found with given criteria      `{"error": "User not found"}`

4. Update User
   Update and existing user's information. Only provided fields will be updated.

   Method: `PUT`
   Endpoint: `/users/{id}`

   Path Parameters
   Parameter        Type      Description
   id               int        The user's unique identifier

   Request
   Headers:
     Content-Type: application/json'

   Body Parameters (all optional):
   Field        Type            Description
   email        string          User's email address (unique)
   username     string          User's username (unique)
   avatar       string          Profile picture URL

   Note: `google_id` cannot be updated for security reasons.

   Example Requests

   Update username only:
   `curl -X PUT http://localhost:3001/api/users/1 \
    -H "Content-Type: application/json" \
    -d '{
      "username": "new_username"
    }'`

   Update multiple fields:
   `curl -X PUT http://localhost:3001/api/users/1 \
    -H "Content-Type: application/json" \
    -d '{
      "username": "brandon_updated",
      "email": "brandon.updated@example.com",
      "avatar": "https://newavatar.com/photo.jpg"
    }'`

   Clear avatar (make empty/default):
   `curl -X PUT http://localhost:3001/api/users/1 \
    -H "Content-Type: application/json" \
    -d '{
      "avatar": ""
    }'`

   Example Response
   Status: `200 OK`

   `{
    "id": 1,
    "google_id": "google_123456789",
    "email": "brandon@example.com",
    "username": "new_username",
    "avatar": "https://lh3.googleusercontent.com/a/default-user",
    "created_at": "2025-06-09T10:30:00Z",
    "updated_at": "2025-06-09T15:45:00Z"
  }`

  Error Responses
  Status                        Description                            Example Response
  400 Bad Request               Invalid JSON or validation error       `{"error": "Invalid email format"}`
  404 Not Found                 User with ID doesnt exist              `{"error": "User not found"}`
  409 Conflict                  Email alr exists                       `{"error": "Email already taken"}`
  500 Internal Server Error     Database error                         `{"error": "Failed to update user"}`


  # Changelog
  - v1.0.0 - Initial User API implementation









   


  
