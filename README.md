go mod init
go mod tidy
go get github.com/gofiber/fiber/v2


gow run main.go


-----
for amd64
docker build . -f Dockerfile -t rakd/twinkle-img-server:tagname    
docker push rakd/twinkle-img-server:tagname
docker pull rakd/twinkle-img-server
docker run -p 80:3000 -d -it rakd/twinkle-img-server 



https://imgapi.twinkle.cat/img/cat/png/19/4/c/3/1
https://imgapi.twinkle.cat/img/cat/png/5/6/s/6/1
https://imgapi.twinkle.cat/img/cat/png/1/10/m/1/1
https://imgapi.twinkle.cat/img/cat/png/5/19/z/1/1

https://imgapi.twinkle.cat/img/cat/jpg/9/3/c/3/1
https://imgapi.twinkle.cat/img/cat/jpg/8/2/s/6/4
https://imgapi.twinkle.cat/img/cat/jpg/6/14/m/6/2
https://imgapi.twinkle.cat/img/cat/jpg/7/9/c/5/3