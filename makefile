sdocker:
	sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0
startdocker: sdocker
	systemctl --user start docker-desktop
# staring from this till v4.17.0 it is docker command, after that it is to create migrate file u can see on their documentation
startmigration:
	docker run -it --rm --network host --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations -seq seen_timestamp
# this is to create mysql container if doen't exist else it will start the container
# Docker run command for local PostgreSQL container
rundocker:
	docker run --name my_postgres -p 5432:5432 -e POSTGRES_USER=mesay -e POSTGRES_PASSWORD=QmSnbi8qomVBhm09gdREAv4vTmbEXvRJ -e POSTGRES_DB=near_me -d postgres

# Create database (Optional, since it is created by default)
createdb:
	docker exec -it my_postgres psql -U mesay -d postgres -c "CREATE DATABASE near_me;"

# Migrate up with remote PostgreSQL
migrateup:
	docker run -it --rm --network host --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 -path /db/migrations -database "postgres://mesay:QmSnbi8qomVBhm09gdREAv4vTmbEXvRJ@dpg-cv02mrq3esus73e42rn0-a.oregon-postgres.render.com:5432/near_me?sslmode=require" up

# Migrate down with remote PostgreSQL
migratedown:
	docker run -it --rm --network host --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 -path /db/migrations -database "postgres://mesay:QmSnbi8qomVBhm09gdREAv4vTmbEXvRJ@dpg-cv02mrq3esus73e42rn0-a.oregon-postgres.render.com:5432/near_me?sslmode=require" down
