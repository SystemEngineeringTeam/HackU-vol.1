FROM golang:alpine

WORKDIR /app
EXPOSE 80
EXPOSE 443

ADD ./wait /wait 

RUN chmod +x /wait&&\
    apk --update add tzdata&&\
    cp  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime&&\
    apk del tzdata&&\
    rm -rf /var/cache/apk/*

CMD ["go","run","main.go"]