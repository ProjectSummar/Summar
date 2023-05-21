# Server

## Usage

Run postgres container

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

> - Use DB sessions
> - Use passwordless login? (magic email link)
> - Use OAuth? (google, github, ...)

- `/login`
  - Params
    - email
    - password?
- `/signup`
  - Params
    - email
    - password?
    - confirm password?
- `/me`
  - Cookie
    - session token
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
