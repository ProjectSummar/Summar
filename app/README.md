# App

## Usage

1. Ensure current working directory is `./summar/app`

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

- `/auth/login`
  - Login page
- `/auth/signup`
  - Signup page
- `/main/bookmark`
  - Bookmarks page
  - Displays all current user bookmarks
- `/main/bookmark/create`
  - Modal in the bookmarks page to create a bookmark
- `/main/bookmark/update?id=...`
  - Modal in the bookmarks page to update a bookmark's title
- `/main/bookmark/{id}`
  - Individual bookmark page
  - Contains webview and summary of the bookmark
- `/main/settings`
  - Settings page
