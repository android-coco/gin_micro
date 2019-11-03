module gin_micro

go 1.13

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.1
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/parnurzeal/gorequest v0.2.16
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	gopkg.in/yaml.v2 v2.2.4
	moul.io/http2curl v1.0.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.1
	github.com/micro/go-plugins v1.4.0 => github.com/micro/go-plugins v1.4.0
)
