# Server

## Usage

1. Install ![docker](https://www.docker.com/)
2. Ensure current working directory is `./summar/server`
3. Build docker images and setup containers by running `docker compose up --build`
4. To shut down containers, run `docker compose down --remove-orphans --volumes`

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
