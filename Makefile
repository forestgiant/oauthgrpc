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
	docker run --name tick1 --network=tick_nw -e="NAME=ONE" -e="TPS=1" -d forestgiant/ticktick
	docker run --name tick2 --network=tick_nw -e="NAME=TWO" -e="TPS=2" -d forestgiant/ticktick
	docker run --name tick3 --network=tick_nw -e="NAME=THREE" -e="TPS=3" -d forestgiant/ticktick
	docker run --name tick4 --network=tick_nw -e="NAME=FOUR" -e="TPS=4" -d forestgiant/ticktick
	docker run --name tick5 --network=tick_nw -e="NAME=FIVE" -e="TPS=5" -d forestgiant/ticktick
	docker run --name tickweb --network=tick_nw -e="TICK1=tick1" -e="TICK2=tick2" -e="TICK3=tick3" -e="TICK4=tick4" -e="TICK5=tick5" -d forestgiant/tickweb

clean:
	-docker stop $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/ticktick" -q -a)
	-docker stop $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rm $$(docker ps -f="ancestor=forestgiant/tickweb" -q -a)
	-docker rmi forestgiant/ticktick
	-docker rmi forestgiant/tickweb
	-docker network rm tick_nw