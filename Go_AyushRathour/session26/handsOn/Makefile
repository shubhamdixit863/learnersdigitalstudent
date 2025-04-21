
migrate:
	 echo "Running migrations"
	 cd pkg && \
	 go run migration.go

start-mongo:
	go run cmd/main.go --dbtype=mongodb

start-mysql:
	go run cmd/main.go --dbtype=mysql

start-mysqlorm:
	go run cmd/main.go --dbtype=mysqlorm

start-inmemory:
	go run cmd/main.go --dbtype=inmemory



mongo:
	docker run -d \
		--name mongo-container \
		-p 27017:27017 \
		mongo