MYSQL_USER=root
MYSQL_PASSWORD=123456
MYSQL_ADDRESS=host.docker.internal
MYSQL_ADDRESS_PORT=3306
MYSQL_DB=transferDB


##=======================Database================================
# Ｍysql 容器 
mysql:
	docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
# 建立資料庫
createdb:
	docker exec -it dev-host_mysql_1 mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" -e "create database "$(MYSQL_DB)""
# 刪除資料庫
dropdb:
	docker exec -it dev-host_mysql_1 mysql -u"$(MYSQL_USER)" -p"$(MYSQL_PASSWORD)" -e "drop database "$(MYSQL_DB)""
# 初始化migration使用文件
migrateinit:
	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate create -ext sql -dir migrations -seq init_schema
# 將migration文件進行處理（建立表格或是新增異動）
migrateup:
	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS):$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose up
# 將migration文件進行處理退回上一步處理（刪除表格或是刪除異動）
migratedown:
	docker run -v "$(shell pwd)"/db/migration:/migrations --rm migrate/migrate -path migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS):$(MYSQL_ADDRESS_PORT))/$(MYSQL_DB)" -verbose down 000001
##==================================================================


# 初始化sqlc文件 (sqlc.yaml)
# https://github.com/kyleconroy/sqlc
# https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html
sqlc-init:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc init

sqlc-generate:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

.PHONY: migrateinit migrateup migratedown dropdb sqlc-init sqlc-generate