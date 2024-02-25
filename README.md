# ShortURL

<a href="https://uich.cc"><img alt="Uptime Robot ratio (7 days)" src="https://img.shields.io/uptimerobot/ratio/7/m787678797-d17e32f3520e4c4b32dc820a"></a>
<a href="https://uich.cc"><img alt="Uptime Robot status" src="https://img.shields.io/uptimerobot/status/m787678797-d17e32f3520e4c4b32dc820a"></a>

A simple Short URL service.


## Run application
### Docker
#### Get image from Docker Hub
```shell
docker pull dengrenjie31/shorturl
```


#### Deployment
```shell
docker run --name shorturl --restart always -p 3000:3000 -v /path/to/config.json:/app/config.json -d dengrenjie31/shorturl
```
