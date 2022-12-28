run:
	DATABASE_URL=postgres://eeeudyfz:SPDShKf1CshaNKgs_yM4XNiVlELJkabr@tiny.db.elephantsql.com/eeeudyfz PORT=:2565 go run server.go

test-integration:
	go test ./expense -tags=integration