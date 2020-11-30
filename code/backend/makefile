build:
	docker build -t smartshopper .

pull:
	sudo docker pull {putaccounthere}/smartshopper

push:
	sudo docker push {put account here}/smartshopper

run:
	sudo docker run  --rm -d -p 8081:8081 -e PORT='8081' \
		--name smartshopper-backend smartshopper

kill:
	sudo docker kill wheypal
