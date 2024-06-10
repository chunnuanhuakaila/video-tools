goctl api go -api ./video.api -dir .

运行方式
go run video.go

请求方式
curl --location --request GET '127.0.0.1:8888/video/parse?url=https%3A%2F%2Fwww.xiaohongshu.com%2Fexplore%2F664410b1000000001e035f87%3Fapp_platform%3Dandroid&ignoreEngage=true&app_version=8.39.0&share_from_user_hidden=true&type=video&author_share=1&xhsshare=WeixinSession&shareRedId=N0xFNTQ3SEA2NzUyOTgwNjY0OTc7OEc6&apptime=1717900002&share_id=39fbb31463074546abadf74885fa59f1&exSource=' \
--header 'Content-Type: application/json'