# VK Test Task
Тестовое задание на стажировку VK

## Команды для бота:

```sh
/poll create        <question>   <option1, option2, ...>
/poll vote          <poll_id>    <option>
/poll check_results <poll_id>
/poll end           <poll_id>
/poll del           <poll_id>
```

## Пример использования:
```sh
/poll create test 1 2 3

Successfully created poll
PollID: po9dzjn3zjrab8rhe9yad5gwfw
Question: test
Answer Options:
    1
    2
    3
```

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
TARANTOOL_USER=your_user
TARANTOOL_PASS=your_pass
```

Запускаем наше приложение + tarantool в Docker-контейнерах
```sh
docker-compose build
docker-compose up
```


