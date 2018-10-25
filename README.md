
# news_sample

Пример очень простого бекенда на Го для описания простого новостного сайта.

# установка

1. накатить схему БД и тестовые данные (PostgreSQL)
2. git clone https://github.com/stdpmk/news_sample.git в $GOPATH/src/github.com/stdpmk
3.  cd $GOPATH/src/github.com/stdpmk && dep ensure
4. cd cmd/news (оснновное приложение)
5. go build -v && news.exe 

Приложение запускается на 8081 порту

