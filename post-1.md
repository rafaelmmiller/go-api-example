# Go journey

## first mistake:

PS: Take care of the field names. I forgot to use `id` as the field name for the primary key and instead used `ID` (or something like this) in the db which caused some panic errors that I thought that were related to zog schema validation. I was lost in the error handling and, used to do JavaScript trycatch logic, but I knew that this isn't how go works. Then I found this quote:

> A WORD OF CAUTION. ZOG & PANICS Zog will never panic due to invalid input but will always panic if invalid destination is passed to the Parse function (i.e if the destination does not match the schema).

Which made me realize that the panic was due to the wrong field name in the db.

---

## folder structure:

following the https://go-blueprint.dev/ folder structure

```
/my_project
  /cmd
    /api
      main.go
  /internal
    /database
      database.go
    /server
      routes.go
      server.go
  /tests
    handler_test.go
  go.mod
  go.sum
  .env
  air.toml
  .gitignore
  Makefile
  README.md
```

in our case we are not adding all server logic inside the `server.go` file. We're going to use DDD so server will only have only server related logic, in this example routes and controllers for rest requests.
- Controllers -> responsible for validating requests (with zod/zog package) and returning responses.

outside `server` we'll have `db`, `models`, `repositories` and `usecases` folders.

- `db` -> database related logic.
- `models` -> data types.
- `repositories` -> database repositories.
- `usecases` -> business logic.

and finally we have tests/e2e folder for end to end tests.

---

## Air:

Air is "like" nodemon but for go.

- auto restart server on file changes.
- auto run tests on file changes.
- builds the project in the `tmp` folder.

---

Next steps:

- AWS lambda function.
