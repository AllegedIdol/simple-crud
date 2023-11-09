DOCKER_IMAGE = mysql
DOCKER_CONTAINER = mysql-container

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p 3306:3306 -d --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)