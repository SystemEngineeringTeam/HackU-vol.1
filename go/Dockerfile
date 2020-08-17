FROM golang:alpine

WORKDIR /app
EXPOSE 80
EXPOSE 443

ADD ./wait /wait 

RUN chmod +x /wait&&\
    ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

CMD ["go","run","main.go"]