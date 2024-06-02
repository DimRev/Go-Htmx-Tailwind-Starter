run: build
	@./bin/app

build: 
	@go build -o bin/app .

templ:
	@templ generate

css:
	tailwindcss -i app/css/app.css -o public/styles.css --watch