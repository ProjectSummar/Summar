# Server

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
