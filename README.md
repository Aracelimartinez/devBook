# DevBook Application Documentation

## Introduction
This document provides an overview of how to run the DevBook application and the PostgreSQL database service. DevBook is a social network where users can share thoughts, interact with posts, comments, and manage their profiles.

### How to Use
To run the DevBook application using Docker and Docker Compose, follow the steps below:

  1. Clone the DevBook application repository:
      ```bash
      git clone https://github.com/Aracelimartinez/devBook.git
      ```
  2. Navigate to the cloned directory
      ```bash
      cd devBook
      ```
  3. Configure the env variables for the database container creating a db.env file in the ./sql directory with the following variables.
      ```bash
      POSTGRES_DB = "devbook"
      POSTGRES_USER = "your_db_username"
      POSTGRES_PASSWORD = "your_db_password"
      ```
  4. Configure the env variables for the API container creating a .env file in the ./api directory with the following variables.
      ```bash
      # Database environment variables
      DB_NAME = "devbook" # Name of the database
      ROL_NAME = "your_db_username" # Name of the database user
      ROL_PASSWORD = "your_db_password" # Password to access the database
      DB_HOST = "db" # Name of the service in the docker-compose file
      DB_PORT = "5432" # Port of your database service

      #  Environment variables of the API
      API_PORT= "8000" #Port of the API
      SECRET_KEY = "your_secret_key" # Secret key used to generate the access token for users
      ```

  5. Execute Docker Compose in the devBook directory:
      ```bash
      docker-compose up
      ```
  3. The application and database will be initialized. The API will be accessible at http://localhost:8000.

## Docker Images
You can find The DevBook application images in the links bellow:

1. **DevBook API Application:** Image containing the environment, dependencies and code needed to run the DevBook application API.
**Docker Hub Link**: [https://hub.docker.com/repository/docker/aracelimartinez/devbook-api/general]
2. **PostgreSQL:** Official PostgreSQL image configured for use with the DevBook application.
**Docker Hub Link**: [https://hub.docker.com/repository/docker/aracelimartinez/devbook-db/general]

There you can find instructions of how to use each image.

## DevBook API Endpoints

The DevBook API facilitates interactions with the DevBook social network platform. Below is a list of the core API endpoints, their purposes, and how to use them.

### Authentication

- **POST `/login`**
  - Description: Authenticate a user and return an authentication token.
  - Payload: `{ "email": "user@example.com", "password": "password" }`
  - Endpoint with NO authentication required

### Users

- **POST `/users`**
  - Description: Register a new user with their information.
  - Payload: `{ "name": "Some name", "nick": "example94", "email": "example@gmail.com", "password": "123456"}`
  - Endpoint with NO authentication required

- **GET `/users`**
  - Description: Retrieve all users profiles.
  - Authorization: Required (Bearer Token)

- **GET `/users/{userId}`**
  - Description: Retrieve a user's profile information by their ID.
  - Authorization: Required (Bearer Token)

- **PUT `/users/{userId}`**
  - Description: Update a user's profile information.
  - Payload: `{ "name": "New name", "nick": "newnick", "email": "example2@gmail.com" }`
  - Authorization: Required (Bearer Token)

- **DELETE `/users/{userId}`**
  - Description: Delete a user account.
  - Authorization: Required (Bearer Token)

- **POST `/users/{userId}/follow`**
  - Description: Allow a user to follow another user.
  - Authorization: Required (Bearer Token)

- **DELETE `/users/{userId}/unfollow`**
  - Description: Allow a user to unfollow another user.
  - Authorization: Required (Bearer Token)

- **GET `/users/{userId}/followers`**
  - Description: Retrieve the followers of a user by their ID.
  - Authorization: Required (Bearer Token)

- **GET `/users/{userId}/following`**
  - Description: Retrieve the users that a user follow.
  - Authorization: Required (Bearer Token)

- **POST `/users/{userId}/update-password`**
  - Description: Update a user's password.
  - Authorization: Required (Bearer Token)

### Posts

- **GET `/posts`**
  - Description: Retrieve a list of recent posts from all users.
  - Authorization: Required (Bearer Token)

- **POST `/posts`**
  - Description: Create a new post.
  - Payload: `{ "title": "This is a title", "content": "This is a new post." }`
  - Authorization: Required (Bearer Token)

- **GET `/posts/{postId}`**
  - Description: Retrieve a single post by its ID.
  - Authorization: Required (Bearer Token)

- **PUT `/posts/{postId}`**
  - Description: Update a post by its ID.
  - Authorization: Required (Bearer Token, must be the post's author)

- **DELETE `/posts/{postId}`**
  - Description: Delete a specific post by its ID.
  - Authorization: Required (Bearer Token, must be the post's author)

- **GET `/users/{userId}/posts`**
  - Description: Retrieve all posts created by a specific user.
  - Authorization: Required (Bearer Token)

- **POST `/posts/{postId}/like`**
  - Description: Allow a user to like a post.
  - Authorization: Required (Bearer Token)

- **POST `/posts/{postId}/like`**
  - Description: Allow a user to like a post.
  - Authorization: Required (Bearer Token)

- **POST `/posts/{postId}/unlike`**
  - Description: Allow a user to unlike a post.
  - Authorization: Required (Bearer Token)

### Error Handling

All endpoints should properly handle errors and return appropriate status codes and messages, e.g., `404 Not Found` for non-existent resources, `401 Unauthorized` for access violations, etc.

Please ensure that all requests to endpoints requiring authorization include a valid authentication token, passed as an `Authorization` header with the `Bearer token` schema.
