

.PHONY: up
up:
	@docker-compose up -d

.PHONY: down
down:
	@docker-compose -f ./docker-compose.yml down

.PHONY: net
net:
	@docker inspect mysql | grep IPAddress

.PHONY: mysql
mysql:
	@docker exec -it mysql bash 
# mysql -uroot -p123456

.PHONY: redis
redis:
	@docker exec -it redis bash
# redis-cli

.PHONY: etcd
etcd:
	@docker exec -it etcd bash 
# etcdctl

# 替代
# docker-compose -f ./docker-compose.yml exec redis sh -c 'redis-cli'