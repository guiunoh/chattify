* run
```base
# run redis
docker compose up -d

# run api
go run ./cmd/thumbnailer-api 
```

* api test
```bash
http -f post http://localhost:8080/api/v1/thumbnails file@$PWD/testdata/original.jpg rate=RATE50 \
  | jq -r '.id' \
  | xargs -I {} sh -c 'http get http://localhost:8080/api/v1/thumbnails/{} -o thumbnail.png'
```
