FRONT_END_BINARY=frontApp
BROKER_BINARY=broker
ORDER_BINARY=order
LOGSERVICE_BINARY=logservice
PAYMENT_BINARY=payment
SHIPPING_BINARY=shipping
AUTH_BINARY=authApp
MAIL_BINARY=mailApp
LOGGER_BINARY=loggerApp
LISTENER_BINARY=listener
EVENTEMITTER_BINARY=eventEmitter
PROTO_VERSION ?= v1.0.17
## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build:  build_helper build_broker build_order build_payment build_shipping build_logservice build_listener build_eventemitter  ##build_broker build_auth build_logger build_mail build_listener
	@echo "Stopping docker images (if running...)"
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build  -d  
	
	@echo "Docker images built and started!"

build_all: build_helper build_broker build_order build_payment build_logservice build_shipping build_listener build_eventemitter  ##build_auth build_logger build_mail build_listener
	@echo "All binaries built!"

# git push
gitpush:
	@echo "Pushing to git..."
	bash ../push.sh
	@echo "Done!"


## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_proto:
	@echo "Building proto binary..."
	cd ../../microservices-proto &&  \
	sudo ./protoc.sh order $(PROTO_VERSION) nileshdt nileshdt@gmail.com && \
	sudo ./protoc.sh logservice $(PROTO_VERSION) nileshdt nileshdt@gmail.com && \
	sudo ./protoc.sh payment $(PROTO_VERSION) nileshdt nileshdt@gmail.com && \
	sudo ./protoc.sh shipping $(PROTO_VERSION) nileshdt nileshdt@gmail.com && \
	sudo ./protoc.sh eventEmitter $(PROTO_VERSION) nileshdt nileshdt@gmail.com  
	@echo "Done!"

build_protorefresh:
	@echo "Refreshing proto..."
	cd ../../microservices-helper && \
	go get github.com/ayuved/microservices-proto/golang/order@latest && \
	go get github.com/ayuved/microservices-proto/golang/logservice@latest && \
	go get github.com/ayuved/microservices-proto/golang/payment@latest && \
	go get github.com/ayuved/microservices-proto/golang/shipping@latest && \
	go mod tidy && \
	cd ../microservices/logservice && \
	go get github.com/ayuved/microservices-proto/golang/logservice@latest && go mod tidy && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	cd ../order && \
	go get github.com/ayuved/microservices-proto/golang/order@latest && go mod tidy && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	cd ../payment && \
	go get github.com/ayuved/microservices-proto/golang/payment@latest && go mod tidy && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	cd ../shipping && \
	go get github.com/ayuved/microservices-proto/golang/shipping@latest && go mod tidy && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	cd ../listener && \
	go mod tidy && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	go mod tidy  && \
	cd ../eventEmitter && \
	go mod tidy && \
	go get github.com/ayuved/microservices-proto/golang/eventEmitter@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest ;
	go mod tidy ;
	@echo "Done!"

build_helper:
	@echo "Building helper binary..."
	cd ../../microservices-helper && go mod tidy && \
	go get github.com/ayuved/microservices-proto/golang/order@latest && \
	go get github.com/ayuved/microservices-proto/golang/logservice@latest && \
	go get github.com/ayuved/microservices-proto/golang/payment@latest && \
	go get github.com/ayuved/microservices-proto/golang/shipping@latest && \
	go mod tidy && \
	env GOOS=linux CGO_ENABLED=0  && \
	bash ./push.sh	;
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker && go get github.com/ayuved/microservices-helper/domain && go get github.com/ayuved/microservices-helper/adapters && go mod tidy && env GOOS=linux CGO_ENABLED=0 go build  -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"


## build_order: builds the order binary as a linux executable
build_order:
	@echo "Building order binary..."
	cd ../order && go mod tidy && go get github.com/ayuved/microservices-helper/domain@latest && \
	go get github.com/ayuved/microservices-proto/golang/order@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	env GOOS=linux CGO_ENABLED=0  go build  -o ${ORDER_BINARY} ./cmd
	@echo "Done!"

## build_payment: builds the payment binary as a linux executable
build_payment:
	@echo "Building payment binary..."
	cd ../payment  && go mod tidy && go get github.com/ayuved/microservices-helper/domain@latest && \
	go get github.com/ayuved/microservices-proto/golang/payment@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && \
	go mod tidy && \
	env GOOS=linux CGO_ENABLED=0  go build  -o ${PAYMENT_BINARY} ./cmd
	@echo "Done!"

## build_shipping: builds the payment binary as a linux executable
build_shipping:
	@echo "Building shipping binary..."
	cd ../shipping && go mod tidy && \
	go get github.com/ayuved/microservices-proto/golang/shipping@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && go mod tidy && \
	env GOOS=linux CGO_ENABLED=0  go build  -o ${SHIPPING_BINARY} ./cmd
	@echo "Done!"

build_eventemitter:
	@echo "Building eventemitter binary..."
	cd ../eventEmitter && go mod tidy && \
	go get github.com/ayuved/microservices-proto/golang/eventEmitter@latest && \
	go get github.com/ayuved/microservices-helper/domain@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && go mod tidy && \
	env GOOS=linux CGO_ENABLED=0  go build  -o ${EVENTEMITTER_BINARY} ./cmd
	@echo "Done!"

## build_logservice gger: builds the logger binary as a linux executable
build_logservice:
	@echo "Building logservice binary..."
	cd ../logservice && go mod tidy && go get github.com/ayuved/microservices-helper/domain@latest && \
	go get github.com/ayuved/microservices-proto/golang/logservice@latest && \
	go get github.com/ayuved/microservices-helper/adapters@latest && go mod tidy && env GOOS=linux CGO_ENABLED=0 go build  -o ${LOGSERVICE_BINARY} ./cmd
	@echo "Done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener && env GOOS=linux CGO_ENABLED=0 go build  -o ${LISTENER_BINARY} .
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build  -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build  -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build  -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## build_front: builds the frone end binary
build_front_linux:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build  -o ${FRONT_BINARY} ./cmd/web
	@echo "Done!"
## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

## docpush: push to docker hub
docpush:
	@echo "Pushing to docker hub..."


	cd ../broker && docker build  -t 909007/broker:latest .
	# docker tag broker 909007/broker:latest
	docker push 909007/broker:latest

	cd ../order && docker build -t 909007/order:latest .
	# docker tag order 909007/order:latest
	docker push 909007/order:latest

	cd ../payment && docker build -t 909007/payment:latest .
	# docker tag payment 909007/payment:latest
	docker push 909007/payment:latest

	cd ../logservice && docker build -t 909007/logservice:latest .
	# docker tag logservice 909007/logservice:latest
	docker push 909007/logservice:latest

	cd ../shipping && docker build -t 909007/shipping:latest .
	# docker tag shipping 909007/shipping:latest
	docker push 909007/shipping:latest

	cd ../listener && docker build -t 909007/listener:latest .
	# docker tag listener 909007/listener:latest
	docker push 909007/listener:latest

	cd ../eventemitter && docker build -t 909007/eventemitter:latest .
	# docker tag eventemitter 909007/eventemitter:latest
	docker push 909007/eventemitter:latest
	@echo "Done!"

## kdelete: delete from kubernetes
kdelete:
	@echo "Deleting from kubernetes.."
	kubectl delete -f k8s
	@echo "Done!"

## kdeploy: push to docker and deploy to kubernetes
kdeploy:	  build_all docpush 
	@echo "Deploying to kubernetes..."
	kubectl apply -f k8s
	@echo "Done!"

## kredeploy: redeploy from kubernetes
kredeploy:  kdelete kdeploy
	@echo "redeploy to kubernetes..."
	@echo "Done!"

docker_db:
	docker-compose -f ./local/postgresLocal.yml up -d
	docker-compose -f ./local/mongoLocal.yml up -d
	docker-compose -f ./local/rabbitLocal.yml up -d