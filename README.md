# Meditap Fullstack Engineer — Technical Test

## Contents
- `algorithms/TM_solution.cs`: three required C# algorithms in a single file.
- `aspnet-user-api/`: ASP.NET Core 8 minimal REST API + EF Core SQLite.
- `go-user-api/`: Go Gin REST API + GORM SQLite.

Both APIs implement `POST /api/users`, `GET /api/users`, `GET /api/users/{userId}`, `PUT /api/users/{userId}`, `DELETE /api/users/{userId}`.

## Run Algorithms
Prerequisite: .NET 8 SDK.
```bash
cd algorithms
dotnet run
```

## Run ASP.NET Core API
Prerequisite: .NET 8 SDK.
```bash
cd aspnet-user-api
dotnet restore
dotnet run --urls http://localhost:5000
```

## Run Golang API
Prerequisite: Go 1.23+, C compiler (required by go-sqlite3/CGO).
```bash
cd go-user-api
go mod tidy
go run .
```
Go server listens on `http://localhost:8081`.

## API example
```bash
curl -i -X POST http://localhost:5000/api/users \
 -H 'Content-Type: application/json' \
 -d '{"firstName":"Jane","lastName":"Doe","email":"jane@example.com"}'
curl http://localhost:5000/api/users
curl http://localhost:5000/api/users/1
curl -i -X PUT http://localhost:5000/api/users/1 \
 -H 'Content-Type: application/json' \
 -d '{"firstName":"Janet","lastName":"Doe","email":"janet@example.com"}'
curl -i -X DELETE http://localhost:5000/api/users/1
```
Replace port `5000` with `8081` to test the Go API.

## Behavior and design
- Request fields: `firstName`, `lastName`, `email` (all required). IDs are auto-generated.
- Valid email format and maximum lengths are checked; emails normalized to lowercase.
- Unique email index prevents duplicate users. Returns HTTP 409 on detected duplicates.
- HTTP 201 + Location for create, 200 for reads/updates, 204 for delete, 400 for invalid data, 404 for missing users.
- SQLite databases are independent: these two implementations do not share data.
- Algorithms: exact case-sensitive palindrome; string reversal uses UTF-16 code units (not Unicode grapheme clusters); sieve returns primes <= limit.

## Validation note
Source code is provided for review. The .NET SDK is unavailable in the authoring environment; ASP.NET build/runtime validation must be performed on a machine with .NET 8 installed. Go dependencies may need to be downloaded before running.
