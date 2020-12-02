build:
	DOCKER_BUILDKIT=1 docker build -t smartshopper-backend .

pull:
	sudo docker pull eecs4312basedcode/smartshopper-backend

push:
	sudo docker tag smartshopper-backend eecs4312basedcode/smartshopper-backend
	sudo docker push eecs4312basedcode/smartshopper-backend

run:
	sudo docker run  --rm -d -p 8081:8081 -e PORT='8081' \
		--name smartshopper-backend smartshopper-backend

kill:
	sudo docker kill smartshopper-backend
