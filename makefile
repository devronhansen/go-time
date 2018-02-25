build:
	cd consignment-cli && make build
	cd consignment-service && make build
	cd vessel-service && make build
	cd user-service && make build
	cd user-cli && make build

	docker-compose build
# build:
# 	DIRS = dir1 dir2 dir3 dir4 ...
# 	OBJS = $(wildcard $(DIRS:=/*.o))

# 	GENERATE_OBJECT_FILES:
# 		@echo $(OBJS);