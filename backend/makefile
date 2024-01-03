db:
	make db

migrate_up:
	migrate -path database/migration/ -database "postgresql://todoapp:password@localhost:5432/todoapp?sslmode=disable" -verbose up

migrate_down:
	migrate -path database/migration/ -database "postgresql://todoapp:password@localhost:5432/todoapp?sslmode=disable" -verbose down

migrate_force:
	migrate -path database/migration/ -database "postgresql://todoapp:password@localhost:5432/todoapp?sslmode=disable" -verbose force VERSION