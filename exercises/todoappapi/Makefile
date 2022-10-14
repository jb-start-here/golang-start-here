up:
	docker run -d --rm -p 5432:5432 --name golang-start-here -e POSTGRES_PASSWORD=gopwd postgres:14.4-alpine > /dev/null
	@echo "launched a postgresql container..."
	@echo "port: 5432"
	@echo "user: postgres"
	@echo "password: gopwd"
	@echo "default db: postgres"
	@echo "container name: golang-start-here"

down:
	docker kill golang-start-here > /dev/null
	@echo "Postgres container 'golang-start-here' is killed and removed..."

migrate:
	docker cp ./initdb.sql golang-start-here:/etc > /dev/null
	docker exec golang-start-here bash -c "psql -U postgres -d postgres -f /etc/initdb.sql" > /dev/null
	@echo "todos table created in postgres db..."
	@echo "added two entries to todos table..."
