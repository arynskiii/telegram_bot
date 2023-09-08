Pocket bot
Pocket is a Telegram bot that allows you to save links in the Pocket app. You can say that it is a small client for this service.

To work with the Telegram API, I use - telegram-bot-api

To work with the Pocket API, I use SDK - go-pocket-sdk.

Bolt DB is used as the storage.

To implement user authorization, an HTTP server on port 80 is started along with the bot, to which a redirect from Pocket is made when the user is successfully authorized.

When the server receives the request, it generates an Access Token through the Pocket API for the user and stores it in the repository.
