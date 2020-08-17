# HackU_Online_vol_1

## docker について

-   networks
    -   `vol1_network`
        -   external しているため，別ディレクトリの container からアクセスすることができる
-   containers
    -   `vol1_golang`
        -   go の API を動かすための container
        -   alpine を使うことで image の軽量化を図る
        -   ファイルのマウント処理を`docker-compose.yml`内でおこなうことにより，ファイルに変更を加える度に build する必要がなくなった．
    -   `vol1_mysql`
        -   mysql を動かすためのコンテナ
        -   `init.sql`に変更を加えた際はコンテナの削除が必要
-   commands
    -   `docker network create vol1_network`
        -   network を作成するコマンド
        -   初回起動時はこれを忘れると動かない
    -   `` docker container rm -f `docker container list -a -q` ``
        -   container を一括で削除するコマンド
        -   データベースのデータも消えるため注意が必要
    -   `` docker rmi -f `docker images | grep none | cut -b 40-54 | tr '\n' ' '` ``
        -   `<none>`な image を一括で削除できるコマンド
        -   build する度に`<none>`な image が生まれてしまうので定期的に実行することをオススメする
