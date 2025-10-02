
# Movie Reservation System

A simple backend service in Go for managing movies, showtimes and seat reservations.

## Features

- User signup, login and role-based access via JWT
- CRUD operations for movies (admin only)
- CRUD operations for showtimes (admin only)
- List movies and their showtimes
- Check available seats and make/cancel reservations

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ticketmonster
   ```
2. Copy `.env.example` to `.env` and fill in suas configurações.
- Ajuste `DB_HOST=db` para uso com Docker Compose.
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Start the application locally:
   ```bash
   go run cmd/main.go
   ```

The server listens by default on `:8080`.

## Docker

To run both the API and a PostgreSQL database with Docker:

```bash
docker-compose up -d
```

This will build the Go binary, start the service on port `8080`,
and spin up Postgres on port `5432` (database `movie_reservation` is created automatically).

To connect to the DB container or recreate the database manually:

```bash
docker-compose exec db createdb --username=postgres movie_reservation
```

Stop and remove containers:

```bash
docker-compose down
```

## API Endpoints

> **Public**
>
>| Method | Endpoint     | Description        |
>| ------ | ------------ | ------------------ |
>| POST   | `/api/signup`| Register a new user|
>| POST   | `/api/login` | Authenticate user  |
>| GET    | `/api/movies`| List all movies    |
>
> **User (authenticated)**
>
>| Method | Endpoint                                     | Description                     |
>| ------ | -------------------------------------------- | ------------------------------- |
>| GET    | `/api/user/reservations`                     | List current user's reservations|
>| POST   | `/api/reservations`                          | Create a new reservation        |
>| DELETE | `/api/reservations/:reservationId`           | Cancel a reservation            |
>| GET    | `/api/showtimes/:showtimeId/seats`           | Get available seats             |
>| GET    | `/api/movies/:movieID/showtimes`             | List showtimes for a movie      |
>
> **Admin**
>
>| Method | Endpoint                                     | Description                     |
>| ------ | -------------------------------------------- | ------------------------------- |
>| POST   | `/api/admin/movies`                          | Create a movie                  |
>| PUT    | `/api/admin/movies/:movieId`                 | Update a movie                  |
>| DELETE | `/api/admin/movies/:movieId`                 | Delete a movie                  |
>| GET    | `/api/admin/reservations`                    | List all reservations           |
>| POST   | `/api/admin/users/:userId/promote`           | Promote a user to admin         |
>| POST   | `/api/admin/showtimes`                       | Create a showtime               |
>| PUT    | `/api/admin/showtimes/:showtimeId`           | Update a showtime               |
>| DELETE | `/api/admin/showtimes/:showtimeId`           | Delete a showtime               |
