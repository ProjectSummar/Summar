# Server

## Usage

Run Postgres container

```bash
sudo docker run --name summar-db -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres
```

Interact with Postgres container

```bash
sudo docker exec -it summar-db psql -U postgres
```

Build and run the server

```bash
make run
```

## Endpoints

### `/login`

- POST request
- Validates user credentials and creates a session on the server
- Sets the session token as a cookie on the client

#### Body

```json
{
  "email": "...",
  "password": "..."
}
```

### `/signup`

- POST request
- Uses the given credentials to create a user on the server

#### Body

```json
{
  "email": "...",
  "password": "..."
}
```

### `/me`

- GET request
- Validate session using cookie and responds with associated user credentials

### Todo

- `/bookmark`
  - `/create`
    - Params
      - url
    - Cookie
      - session token
  - `/get`
    - Params
      - bookmark id
    - Cookie
      - session token
  - `/update`
    - Params
      - bookmark id
      - Fields to update
    - Cookie
      - session token
  - `/delete`
    - Params
      - bookmark id
    - Cookie
      - session token
  - `/summarise`
    - Params
      - bookmark id
    - Description
      - Retrieves url from DB using id
      - Scrape the url for the body text
      - Send an API request to OpenAI API to summarise the body text
      - Format and store the result of the API request in the DB
