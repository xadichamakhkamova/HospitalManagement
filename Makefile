# .env fayldan DB_URL o'qish
ifneq (,$(wildcard .env))
	include .env
	export
endif

migrate-up:
	migrate -path hospital/migrations -database "${DB_URL}" up

migrate-down:
	migrate -path hospital/migrations -database "${DB_URL}" down

migrate-force:
	migrate -path hospital/migrations -database "${DB_URL}" -verbose force 0
