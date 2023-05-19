# Server

## Endpoints

### Database Operations

- `/create-bookmark`
  - Params
    - url
- `/create-bookmarks`
  - Params
    - url[]
- `/get-bookmark`
  - Params
    - bookmark id
- `/get-bookmarks`
  - Params
    - bookmark id[]
- `/update-bookmark`
  - Params
    - bookmark id
- `/update-bookmarks`
  - Params
    - bookmark id[]
- `/delete-bookmark`
  - Params
    - bookmark id
- `/delete-bookmarks`
  - Params
    - bookmark id[]

### Authentication

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
  - `summarise`
    - Params
      - bookmark id
    - Description
      - Retrieves url from DB using id
      - Scrape the url for the body text
      - Send an API request to OpenAI API to summarise the body text
      - Format and store the result of the API request in the DB
