build:
	docker build -t smartshopper-frontend .

pull:
	sudo docker pull eecs4312basedcode/smartshopper-frontend

push:
	sudo docker tag smartshopper-frontend eecs4312basedcode/smartshopper-frontend
	docker push eecs4312basedcode/smartshopper-frontend

run:
	sudo docker run  --rm -d -p 8080:8080 \
		--name smartshopper-frontend smartshopper-frontend

kill:
	sudo docker kill smartshopper-frontend
