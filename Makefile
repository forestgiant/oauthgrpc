build:
	-docker network create --driver bridge tick_nw
	-docker stop $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker stop $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rmi forestgiant/ticktick
	-docker rmi forestgiant/tickweb
	docker build -t forestgiant/ticktick ticktick/server
	docker build -t forestgiant/tickweb webserver
	docker run --name tick1 --network=tick_nw -d forestgiant/ticktick
	docker run --name tick2 --network=tick_nw -d forestgiant/ticktick
	docker run --name tick3 --network=tick_nw -d forestgiant/ticktick
	docker run --name tick4 --network=tick_nw -d forestgiant/ticktick
	docker run --name tick5 --network=tick_nw -d forestgiant/ticktick
	docker run --name tickweb --network=tick_nw -d forestgiant/tickweb

clean:
	-docker stop $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker stop $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rmi forestgiant/ticktick
	-docker rmi forestgiant/tickweb
	-docker network rm tick_nw