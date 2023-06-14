# App

## Usage

1. Ensure current working directory is `./summar/app` and the
   [server is running](/server/README.md)

2. Install dependencies by running

   ```bash
   npm install
   ```

3. Configure the base URL for the server in `app.json` under `extra.baseUrl`

4. Start the development build of the app by running

   ```bash
   npx expo start
   ```

   To start the production build, run

   ```bash
   npx expo start --no-dev --minify
   ```

## Sitemap

### Authentication pages

#### Login page

- `/auth/login`

#### Signup page

- `/auth/signup`

### Main pages

#### Bookmarks page

- `/main/bookmark`
- Displays all current user bookmarks
- Search bookmarks by title
- Manage bookmarks through context menus

##### Create bookmark modal

- `/main/bookmark/create`
- Modal in the bookmarks page to create a bookmark

##### Update bookmark modal

- `/main/bookmark/update?id=...`
- Modal in the bookmarks page to update a bookmark's title

#### Individual bookmark page

- `/main/bookmark/{id}`
- Context menu options to summarise bookmark and toggle webview and summary view

#### Settings page

- `/main/settings`
