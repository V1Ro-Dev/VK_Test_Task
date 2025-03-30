# VK_Test_Task
Тестовое задание на стажировку VK

## Как запускать проект:

Для локального запуска необходимо поднять MatterMost в Docker- контейнере и при этом
запустить образ MatterMost'a в сети mattermost_network
Команда для создания сети и добавления образа приложения:
docker network create mattermost_network
docker network connect mattermost_network mattermost-app
Где mattermost-app - название контейнера
