#build stage
FROM golang:alpine AS builder

#Configure
ENV REPO_URL=github.com/dung997bn/bookstore_items_api
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORK_PATH=${APP_PATH}/src

#Copy the entire source code from the current directory (src) to ${WORK_PATH}
COPY src ${WORK_PATH}
WORKDIR ${WORK_PATH}

# RUN go env
RUN go build -o items-api.exe .

#expose port 8082
EXPOSE 8082

CMD [ "./items-api.exe" ]

# WORKDIR /go/src/app
# COPY . .
# RUN apk add --no-cache git
# RUN go get -d -v ./...
# RUN go install -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT ./app
# LABEL Name=bookstoreitemsapi Version=0.0.1
# EXPOSE 8082
