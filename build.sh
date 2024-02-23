IMG_NAME="dengrenjie31/shorturl"
rm -f shorturl.tar.xz
if [ -n "$(docker images | grep $IMG_NAME)" ];then
	docker rmi $IMG_NAME
fi
docker buildx build --platform=linux/amd64 -t $IMG_NAME .
docker save -o shorturl.tar $IMG_NAME
xz -9 -T 11 shorturl.tar