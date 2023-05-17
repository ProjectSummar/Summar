# Summar

## Architecture

![https://github.com/ProjectSummar/Summar/assets/79556863/cd3e77da-0def-4bc6-92ab-a8c0c4d6b94d]

## Server

### Endpoints

#### Database Operations

- `/create_bookmark`
  - Params
    - url
- `/create_bookmarks`
  - Params
    - url[]
- `/get_bookmark`
  - Params
    - bookmark id
- `/get_bookmarks`
  - Params
    - bookmark id[]
- `/update_bookmark`
  - Params
    - bookmark id
- `/update_bookmarks`
  - Params
    - bookmark id[]
- `/delete_bookmark`
  - Params
    - bookmark id
- `/delete_bookmarks`
  - Params
    - bookmark id[]

#### Authentication

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

#### Summarising

- `summarise_bookmark`
  - Params
    - bookmark id
  - Description
    - Retrieves url from DB using id
    - Scrape the url for the body text
    - Send an API request to OpenAI API to summarise the body text
    - Format and store the result of the API request in the DB

## Database

![https://github.com/ProjectSummar/Summar/assets/79556863/72f274fe-0da0-4cee-a22d-6b0faad9c2d7]
