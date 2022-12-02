
# Mobile Validator Register Service

GoLang REST API. Part of transaction gateway. Responsibility for register mobile validator devices.
This is using clean architecture inspired by [https://github.com/evrone/go-clean-tempLast](https://github.com/evrone/go-clean-tempLast)
## Development
For Development at Local Machine
#### Prerequisite
* Docker must be installed

#### Set-up Redis

1. Pull redis image
`docker pull redis`
2. Start Redis
`docker run --name some-redis -d -p 6379:6379 redis redis-server --save 60 1 --loglevel warning`
Redis stores your persisted data in the `VOLUME` /data location. 

#### Set-up Mongo
1. Pull mongo image
`docker pull mongo`
2. Start Mongo
`docker run --name validator-mongo -d -p 27017:27017 mongo`

#### Preparing Local Dev. Environment
1. Copy file .env.sample as .env
`cp .env.sample .env`

2. Edit file .env as following:
    ```
    APP_NAME: "mobile-validator-register-service"
    APP_VERSION: "1.0.0"
    HTTP_PORT: "8001"
    LOG_LEVEL: "debug"
    MONGO_URI: "mongodb://localhost:27017/?authSource=evice-register-services"
    MONGO_DB: "device-register-services"
    MONGO_DEVICE_COLLECTION_NAME: "devices"
    REDIS_ADDR: "127.0.0.1:6379"
    REDIS_PASSWORD: ""
    REDIS_DEVICE_DB: "1"
    PUSHER_BEAM_INSTANCE_ID: "eb3d4c21-3ff6-4378-8117-f0acf095d6b3"
    PUSHER_BEAM_SECRET_KEY: "5B7C86F49099A5132BCA42A60BD0B61A9144A835C81CCD483290C5E21BB9B4E8"
    ```

## Run

```bash
go mod tidy

go run cmd/app/main.go
```
**Output**
```
Starting App...
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (4 handlers)
[GIN-debug] GET    /healthz                  --> nextclan/validator-register/mobile-validator-register-service/internal/controller/http/v1.NewRouter.func1 (4 handlers)
[GIN-debug] GET    /metrics                  --> github.com/gin-gonic/gin.WrapH.func1 (4 handlers)
[GIN-debug] POST   /v1/devices               --> nextclan/validator-register/mobile-validator-register-service/internal/controller/http/v1.(*deviceRoutes).doRegisterDevice-fm (4 handlers)
[GIN-debug] POST   /v1/devices/:deviceId/status --> nextclan/validator-register/mobile-validator-register-service/internal/controller/http/v1.(*deviceRoutes).doUpdateDeviceStatus-fm (4 handlers)
[GIN-debug] GET    /v1/devices/auth          --> nextclan/validator-register/mobile-validator-register-service/internal/controller/http/v1.(*deviceRoutes).doAuth-fm (4 handlers)

```

## Run with Docker Compose
### Prerequisite
- Redis already run on another container
- Mongo already run on another container
- Must have basic knownledge about Docker Compose

### Step to run docker compose
1. Copy file .env.sample as .env
    `cp .env.sample .env`

2. Edit file .env as following:
    ```
    APP_NAME: "mobile-validator-register-service"
    APP_VERSION: "1.0.0"
    HTTP_PORT: "8001"
    LOG_LEVEL: "debug"
    MONGO_URI: "mongodb://host.docker.internal:27017/?authSource=evice-register-services"
    MONGO_DB: "device-register-services"
    MONGO_DEVICE_COLLECTION_NAME: "devices"
    REDIS_ADDR: "host.docker.internal:6379"
    REDIS_PASSWORD: ""
    REDIS_DEVICE_DB: "1"
    PUSHER_BEAM_INSTANCE_ID: "eb3d4c21-3ff6-4378-8117-f0acf095d6b3"
    PUSHER_BEAM_SECRET_KEY: "5B7C86F49099A5132BCA42A60BD0B61A9144A835C81CCD483290C5E21BB9B4E8"
    ```
3. Run Docker Compose
   `docker compose up`

   Output:
   ```
  
   ```

## Build

```bash
go build cmd/app/main.go
```

## Test
```bash
go test -v ./...
```

## Deploy
### With Helm 
```bash
helm upgrade --set deployment.image="{image:tag}" -f helm/values.yaml {release name} ./helm -n {namespace} --install
```
### Example
```bash
helm upgrade --set deployment.image="registry.gitlab.com/nextdb-project/digital-reality-foundation/mobile-validator/validator-register/mobile-validator-register-service:develop" -f helm/values.yaml mobile-validator-register-service ./helm -n transaction-gateway-be --install
```


## Environment
```
APP_NAME: "mobile-validator-register-service"
APP_VERSION: "1.0.0"
HTTP_PORT: "8001"
LOG_LEVEL: "debug"
MONGO_URI: "mongodb://{username}:{password}@{host}:{port}/?authSource=device-register-services"
MONGO_DB: "device-register-services"
MONGO_DEVICE_COLLECTION_NAME: "devices"
REDIS_ADDR: ""
REDIS_PASSWORD: ""
REDIS_DEVICE_DB: "1"
PUSHER_BEAM_INSTANCE_ID: ""
PUSHER_BEAM_SECRET_KEY: ""
```
