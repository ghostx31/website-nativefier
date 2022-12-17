FROM golang:latest as builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o nativefier-downloader

FROM node:16 
WORKDIR /usr/src/app
COPY --from=builder /app/package*.json ./
RUN npm install 
COPY --from=builder /app ./
RUN ls -la node_modules/.bin/
EXPOSE 1323
ENTRYPOINT [ "./nativefier-downloader" ]
