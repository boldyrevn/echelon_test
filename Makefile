server:
	docker compose up -d --build

client:
	go build ./cmd/client