# VK Test Task
Тестовое задание на стажировку VK


## Как запускать проект:

Перед запуском приложения необходимо поднять MatterMost в Docker-контейнере.  
При этом контейнер необходимо запустить в сети mattermost_network.  
Тут mattermost-app - название контейнера MatterMost'a
```sh
docker network create mattermost_network
docker network connect mattermost_network mattermost-app
```

Клонируем репозиторий:
```sh
git clone https://github.com/V1Ro-Dev/VK_Test_Task
```

Переходим в папку с конфигом MatterMost'a:
```sh
cd deploy/config/matter-most
```

Заполняем файл config.toml, в котором необходимо указать:
```sh
MM_TOKEN="your_bot_token"
MM_SERVER="mm_deploy_link" 
MM_USERNAME="your_bot_name"
```

Переходим в папку с конфигом Tarantool:
```sh
cd deploy
```

Заполняем .env файл, в котором необходимо указать:  
Где TARANTOOL_URL = имя контейнера с Tarantool + порт
```sh
TARANTOOL_URL=tarantool:3301
TARANTOOL_USER=admin
TARANTOOL_PASS=secret
```

Запускаем наше приложение + tarantool в Docker-контейнерах
```sh
docker-compose build
docker-compose up
```


