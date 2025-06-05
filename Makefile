#!makefile

APP_NAME="iris-swift"
CMD_PATH="./cmd/${APP_NAME}"

tidy:
	@go mod tidy
	@go fmt ./...
	@echo "Done!"

generate:
	@echo "Generateing Templ Code"
	templ generate ./views/
	@echo "Generating SQLC code"
	sqlc generate -f ./sqlc.yml
	@echo "Generating Tailwind CSS code"
	tailwindcss -i ./views/static/custom.css -o ./views/static/styles.css --minify


run: generate
	go run ./cmd/api/main.go


watch:
	air
