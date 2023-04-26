
* message publisher api
```bash
http GET http://localhost:9002/api/v1/messages \
  lastID==01GY1AZN62M9PDQDWGYG3VW6DJ topic==01GXZE2X4BMVE0G2H50KZ8REQD limit==10 forward==true
  
http GET http://localhost:9002/api/v1/messages/01GY06XKXY3QAFCCAWKB5X9G06 \
  topic==01GXZE2X4BMVE0G2H50KZ8REQD

http POST http://localhost:9002/api/v1/publish \
  topic=01GXZE2X4BMVE0G2H50KZ8REQD content="Hello Houston." sender="alice"

```

* subscribe api
```bash
http POST http://localhost:9002/api/v1/subscribes \
  topic=01GXZE2X4BMVE0G2H50KZ8REQD connectionID=01GYZ8RAHSE45A2JCVNVA6BTJ2
  
http DELETE http://localhost:9002/api/v1/subscribes \
  topic=01GXZE2X4BMVE0G2H50KZ8REQD connectionID=01GY1AZN62M9PDQDWGYG3VW6DJ

http DELETE http://localhost:9002/api/v1/subscribes \
  connectionID=01GY1AZN62M9PDQDWGYG3VW6DJ  
```

* connection notifier api
```bash
wscat -c 127.0.0.1:9001/api/v1/ws --slash

http POST http://localhost:9090/api/v1/notification connectionID=01GYPX6Q25TE5P0ZXZVBN4ZYTB payload=test
```

* consumer
```bash
docker run --rm redis \
  redis-cli -h host.docker.internal -p 6379 publish forwarder-channel "{\"connectionID\" : \"01GYSXHBJC5VQV6RENWY0JCD21\", \"payload\": \"Hello Houston\"}"
```




### message api
* 메세지 생성
* 생성 이벤트 publish

### notify ws
* client websocket 연결관리
* websocket 에 연결 되어 있는 client 에 메세지 전송 api (connectionID, payload)  

### management api
* websocket 연결된 서버를 찾아서 메세지 전송 api 호출

### event consumer
* 메세지 생성 이벤트를 처리
* 이벤트를 management api 에 전달

### message broker(redis)

### database(redis)



```bash
http GET http://localhost:9001/api/v1/messages limit==10 forward==true
http GET http://localhost:9001/api/v1/topics/01GXZE2X4BMVE0G2H50KZ8REQD/messages/01GY06XKXY3QAFCCAWKB5X9G06
http POST http://localhost:9001/api/v1/topics/01GXZE2X4BMVE0G2H50KZ8REQD/messages content="Hello Houston." author="alice"


http GET http://localhost:8080/api/v1/topics/01GXZE2X4BMVE0G2H50KZ8REQD
http POST http://localhost:8080/api/v1/topics creator="alice" name="virtual topic."

http POST http://localhost:8080/api/v1/connection
http DELETE http://localhost:8080/api/v1/connection/01GY7G2K4H8Y3CW0FD95E6WBFC
http GET http://localhost:8080/api/v1/connection/01GY7G2K4H8Y3CW0FD95E6WBFC

http PUT http://localhost:8080/api/v1/connection/01GY7G2K4H8Y3CW0FD95E6WBFC/ttl ttl:=3600



# query
http POST http://localhost:8080/api/v1/connection \
  | jq -r '.id' \
  | xargs -I {} http GET http://localhost:8080/api/v1/connection/{}

# delete  
http POST http://localhost:8080/api/v1/connection \
  | jq -r '.id' \
  | xargs -I {} http DELETE http://localhost:8080/api/v1/connection/{}
 
# extend ttl
http POST http://localhost:8080/api/v1/connection \
  | jq -r '.id' \
  | xargs -I {} echo "http PUT http://localhost:8080/api/v1/connection/{}/ttl ttl:=3600"
```

```bash
wscat -c 127.0.0.1:8080/api/v1/ws --slash
{"type" : "subscribe", "topic": "topic-10001"}


http POST http://localhost:9090/api/v1/notification connectionID=01GYPX6Q25TE5P0ZXZVBN4ZYTB payload=test
```