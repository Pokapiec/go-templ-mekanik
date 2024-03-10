dev:
	@go run main.go

templ:
	@templ generate -watch

tailwind:
	@npx tailwindcss -i ./assets/styles/input.css -o ./assets/styles/output.css --watch

reload:
	@air -c .air.toml