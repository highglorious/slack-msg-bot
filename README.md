# Message Sender to Slack

Программа принимает на входе файл example.json и отправляет текстовые сообщения в каналы из этого json-файла.
- Зарегистрироваться в Slack
- Создать workspace
- Создать три открытых канала “test1”, ”test2”, “test3”
- Создать приложение (бота) для отправки сообщений из json-файла в каналы (добавить bot token scopes: chat:write, chat:write.public)
- Добавить Bot User OAuth Token от своего workspace в файл example.json (заменить "your_bot_token").
- Запустить $ go run main.go
