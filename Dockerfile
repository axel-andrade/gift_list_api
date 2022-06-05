FROM golang
WORKDIR /app/src/go_gift_list_api
ENV GOPATH=/app
COPY . /app/src/go_gift_list_api
RUN go get -u github.com/gin-gonic/gin 
RUN go get -u github.com/joho/godotenv 
RUN go get -u github.com/satori/go.uuid 
RUN go get -u gorm.io/driver/mysql
RUN go get -u gorm.io/gorm
RUN go build -o main .
CMD [ "./main" ]