# Personal Blogging Platform API

This is a robust, RESTful API built in Go (Golang) with PostgreSQL persistence to support a personal blogging platform. It follows a clean, layered architecture for scalability and maintainability, implementing full CRUD (Create, Read, Update, Delete) functionality for blog posts, including flexible search capabilities. 

## Features

- RESTful CRUD: Complete endpoints for creating, reading (single and list), updating, and deleting blog posts.

- Database Persistence: Uses PostgreSQL for reliable data storage.

- Structured Architecture: Separates concerns into handler, store, and model packages.

- Search/Filtering: Allows filtering of posts based on a search term across title, content, and category fields.

- Input Validation: Ensures required fields are present before processing requests.

## Project Structure

The application follows a standard, scalable Go project layout:

```
blog-api/
├── cmd/
│   └── main.go           # Application entry point, server setup, and routing.
└── internal/
    ├── handler/
    │   ├── post.go       # HTTP endpoint handlers (API layer).
    │   └── utils.go      # JSON and error response helpers.
    ├── model/
    │   └── post.go       # Data structures for requests, responses, and database.
    └── store/
        └── postgres.go   # Data access layer (PostgreSQL connection and CRUD operations).
```

## Setup and Running

### Prerequisites

- Go: Go 1.20 or newer.

- PostgreSQL: A running PostgreSQL instance (Local or remote).

## Steps

### Initialize Go Module

```
go mod init blog-api
go mod tidy
```

### Configure Database
You must set the DATABASE_URL environment variable before running the application. This URL tells the application how to connect to your PostgreSQL server.

```
# Example for Linux/macOS
export DATABASE_URL="postgres://user:password@localhost:5432/blog_db?sslmode=disable"
```

Run the Server
The application will automatically connect to the database and create the posts table if it doesn't already exist.

```
go run cmd/main.go
```

The API will start on http://localhost:8080.

## API Endpoints

All endpoints use the base path /posts.

### 1. Create Blog Post (POST)

Method          POST

Endpoint        /posts

Description     Creates a new blog post.

Request Body Example:

```
{
  "title": "A Look at Go's Concurrency",
  "content": "Exploring goroutines and channels in Go.",
  "category": "Technology",
  "tags": ["Go", "Concurrency", "Programming"]
}
```

Responses:

201 Created: Returns the newly created post, including id, createdAt, and updatedAt timestamps.

400 Bad Request: If required fields (title, content) are missing.

### 2. Get All Blog Posts (GET)

Method          GET

Endpoint        /posts

Description     Retrieves all blog posts.

```
GET /posts?term={search_term}
```

Filters posts by a search term in title, content, or category.

Request Example (Search):

```
GET /posts?term=Concurrency
```

Response:

200 OK: Returns an array of blog post objects.

### 3. Get Single Blog Post (GET)

Method          GET

Endpoint        /posts/{id}

Description     Retrieves a single post by its ID.


Response:

200 OK: Returns the blog post object.

404 Not Found: If the post ID does not exist.

### 4. Update Blog Post (PUT)

Method          PUT

Endpoint        /posts/{id}

Description     Updates all fields of an existing post.


Request Body Example (Same structure as POST):

```
{
  "title": "The Updated Go Concurrency Guide",
  "content": "Updated content focusing on Go's scheduler.",
  "category": "Technology",
  "tags": ["Go", "Concurrency", "Advanced"]
}
```

Responses:

200 OK: Returns the updated post object with a new updatedAt timestamp.

400 Bad Request: If required fields are missing.

404 Not Found: If the post ID does not exist.

### 5. Delete Blog Post (DELETE)

Method          DELETE

Endpoint        /posts/{id}

Description     Deletes a post by its ID.

Response:

204 No Content: If the post was successfully deleted.

404 Not Found: If the post ID does not exist.

## Acknowledgements
https://roadmap.sh/projects/url-shortening-service