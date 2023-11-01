MYSQL_USER=root
MYSQL_PASSWORD=123456
MYSQL_ADDRESS=host.docker.internal
MYSQL_ADDRESS_PORT=3306
MYSQL_DB=transferDB


##=======================Database================================
# Ｍysql 容器 
mysql:
	docker-compose up -d mysql-test
# 建立資料庫
createdb:
	docker exec -it mysql-test mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" -e "create database "$(MYSQL_DB)""
# 刪除資料庫
dropdb:
	docker exec -it mysql-test mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" -e "drop database "$(MYSQL_DB)""
# 初始化migration使用文件
# 產生的up and down 文件將create 與 drop 語法準備好在文件內
migrateinit:
	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate create -ext sql -dir migrations -seq init_schema

# 將migration文件進行處理（建立表格或是新增異動）
#migrateup:
#	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql-test:$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose up
###
###
# [[ docker-compose 方式]]
migrateup:
	docker-compose run --rm migrateup

# 將migration文件進行處理退回上一步處理（刪除表格或是刪除異動）
#migratedown:
#	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS):$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose down 000001
###
###
# [[ docker-compose 方式]]
migratedown:
	docker-compose run --rm migratedown


# 建立下一步驟migration檔案
###
# [[ docker-compose 方式]]
migrateCreate:
	docker-compose run --rm migrateCreate


# 將migration文件進行處理（建立表格或是新增異動）
#migrateup:
#	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql-test:$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose up
###
###
# [[ docker-compose 方式]]
migrateup1:
	docker-compose run --rm migrateup1

# 將migration文件進行處理退回上一步處理（刪除表格或是刪除異動）
#migratedown:
#	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS):$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose down 000001
###
###
# [[ docker-compose 方式]]
migratedown1:
	docker-compose run --rm migratedown1




# 初始化sqlc文件 (sqlc.yaml)
# https://github.com/kyleconroy/sqlc
# https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html
sqlc-init:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc init

sqlc-generate:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate

##==================================================================

# 跑所有package的測試
test:
	go test -v -cover ./...

server:
	go run main.go

#mock tool
mockTool:
	go install go.uber.org/mock/mockgen@latest
mock:
	mockgen -package=mockdb -source=./db/sqlc/store.go -destination=./db/mock/store.go

.PHONY: migrateinit migrateup migratedown dropdb sqlc-init sqlc-generate test server mock mockTool migrateup1 migratedown1 migrateCreate