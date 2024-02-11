# Leaderboard API

This is a simple leaderboard API written in Go. It uses the Gin web framework and an in-memory database.

### Motivations
- Code something simple but polished in Go.
- Play with the Gin framework (first time using it).
- Writing code for a leaderboard as a service seemed interesting.

### Use cases
- A web-hosted leaderboard e.g. a charity event tracking highest donations
- Not suitable for: an in-game leaderboard e.g. Valorant where the leaderboard needs to be immediately updated.

### Extensions
- Implement pagination when reading entries e.g. to fetch only the top 10 entries.
- Redis (or other form of persistence) implentation of the database.

## Running the Server Locally

To run the server locally, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the official website: https://golang.org/dl/

Once you have Go installed, follow these steps:

1. Clone the repository: `git clone <repository-url>`
2. Navigate to the project directory: `cd <project-directory>`
3. Run the server: `go run main.go`

The server will start on `localhost:8080`.

## API Endpoints

### `GET /ping`

A simple ping test. Returns `pong` if the server is running.

### `POST /leaderboards`

Creates a new leaderboard. The request body should be a JSON object with the following fields:

- `name`: The name of the leaderboard (string, required)
- `fieldNames`: An array of field names (array of strings, at least one required)
- `entries`: An array of entries - e.g. a player in an FPS game with their score would be a single entry (optional)

Each entry should be a JSON object with the following fields:

- `name`: The name of the entry (string, required)
- `fieldNamesToValues`: A map of field names to their values (map of string to int, required)

### `GET /leaderboards/:id`

Fetches a leaderboard by its ID. The ID should be a valid UUID and is passed as a path parameter.

### `POST /leaderboards/:id/entries`

Creates a new entry in a leaderboard. The request body should be a JSON object with the following fields:

- `name`: The name of the entry (string, required)
- `fieldNamesToValues`: A map of field names to their values (map of string to int, required)

The leaderboard ID should be a valid UUID and is passed as a path parameter.

### `PATCH /leaderboards/:lid/entries/:eid`

Updates an entry in a leaderboard. The request body should be a JSON object with the following fields:

- `fieldNamesToValues`: A map of field names to their new values (map of string to int, at least one required)

The leaderboard ID (`:lid`) and entry ID (`:eid`) should be valid UUIDs and are passed as path parameters.

### `GET /leaderboards/:id/entries`

Fetches all entries in a leaderboard. The leaderboard ID should be a valid UUID and is passed as a path parameter.

You can optionally pass the following query parameters:

- `order`: The order in which to sort the entries. Can be `asc` or `dsc`. Defaults to `dsc`.
- `sort_key`: The field name to sort the entries by. Must be a valid field name in the leaderboard. Defaults to the first field name specified in the leaderboard.
