# Ordering

A small web app for ordering drinks at a home bar. Guests pick a name, browse the menu and send orders from their phone; the bartender sees incoming orders and marks them done. There's also a leaderboard, because why not.

Built with Go, [templ](https://templ.guide), HTMX and PostgreSQL.

## Running it

You need Go 1.24+ and a PostgreSQL database.

1. Create a database and run the migrations in `sql/schema` (they're written for [goose](https://github.com/pressly/goose)):

   ```sh
   goose -dir sql/schema postgres "$DB_URL" up
   ```

2. Put the connection string in a `.env` file:

   ```
   DB_URL=postgres://user:password@localhost:5432/ordering?sslmode=disable
   ```

3. Start the server:

   ```sh
   go run .
   ```

The app runs on `http://localhost:8080`. Guests use `/`, the bartender view is at `/bar`.

## Development

If you change a `.templ` file, regenerate with `go tool templ generate`. If you change a query in `sql/queries`, regenerate with `sqlc generate`. For live reload there's an [air](https://github.com/air-verse/air) config, so just run `air`.

## Credits

Illustrations by [Irasutoya](https://www.irasutoya.com/), used under their terms of use.

## License

The code is MIT licensed, see [LICENSE](LICENSE). The license doesn't cover the illustrations in `assets/`.
