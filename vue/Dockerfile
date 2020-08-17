FROM node:12.16.1
WORKDIR /usr/src/app
ADD . /usr/src/app

RUN  mkdir ~/.npm-global&&\
    npm config set prefix '~/.npm-global'

ENV PATH=~/.npm-global/bin:$PATH

RUN yarn global add @vue/cli&&\
    yarn install
CMD ["yarn","serve"]
