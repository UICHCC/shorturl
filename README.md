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
The default configuration file is `etc/config.json`. If your configuration is exactly the same as it, you can run the application as below:
```shell
docker run --name shorturl --restart always -p 3000:3000 -d dengrenjie31/shorturl
```

If you want to modify the configuration file, edit the file, put it in a certain path and run the application as below:
```shell
docker run --name shorturl --restart always -p 3000:3000 -v /path/to/config.json:/app/config.json -d dengrenjie31/shorturl
```
