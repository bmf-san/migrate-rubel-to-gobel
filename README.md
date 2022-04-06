# migrate-rubel-to-gobel
This is a data migration tool for rubel and gobel.

- [bmf-san/Rubel](https://github.com/bmf-san/Rubel)
- [bmf-san/gobel-api](https://github.com/bmf-san/gobel-api)

# Read Started
1. Copy an .env.example.json as .env.json, and edit it.
2. Run mysql containers.
`docker-compose up -d`
3. Restore data to rubel db.
4. Run migration
`go run main.go`
