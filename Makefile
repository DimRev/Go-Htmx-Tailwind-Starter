run: build
	@./bin/app

build: 
	@go build -o bin/app .

templ-watch:
	@templ generate --watch --proxy=http://localhost:3000

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch