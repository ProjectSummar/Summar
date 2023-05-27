# Server

## Usage

1. Install [docker](https://www.docker.com/)
2. Ensure current working directory is `./summar/server`
3. Build docker images and setup containers by running

   ```bash
   docker compose up --build
   ```

4. To shut down containers, run

   ```bash
   docker compose down --remove-orphans --volumes
   ```

## Endpoints

### Login

- `/login`
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

### Signup

- `/signup`
- POST request
- Uses the given credentials to create a user on the server

#### Body

```json
{
  "email": "...",
  "password": "..."
}
```

### Get current logged in user

- `/me`
- GET request
- Validate session using cookie and responds with associated user credentials

### Create bookmark

- `/bookmark`
- POST request
- Uses the given details to create a bookmark for the current logged in user

#### Body

```json
{
  "url": "..."
}
```

### Get bookmark

- `/bookmark/{id}`
- GET request
- Validate that current logged in user is authorised to view this bookmark
- Return the bookmark with the given id

### Updated bookmark

- `/bookmark/{id}`
- PATCH request
- Validate that current logged in user is authorised to modify this bookmark
- Update the bookmark with the given details

#### Body

```json
{
  "url": "... (optional)",
  "summary": "... (optional)"
}
```

### Delete bookmark

- `/bookmark/{id}`
- DELETE request
- Validate that current logged in user is authorised to modify this bookmark
- Delete the bookmark with the given id

### Todo

- `/bookmark`
  - `/summarise`
    - Params
      - bookmark id
    - Description
      - Retrieves url from DB using id
      - Scrape the url for the body text
      - Send an API request to OpenAI API to summarise the body text
      - Format and store the result of the API request in the DB
