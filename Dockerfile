FROM golang:1.19 AS downloader 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -la
RUN go build -o nativefier-downloader 


FROM node:16 AS nodeimage
WORKDIR /app
COPY --from=downloader /app/package*.json ./
RUN npm install
COPY --from=downloader /app .
RUN ls -la 

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=nodeimage /app .
EXPOSE 1323
ENTRYPOINT [ "/app/nativefier-downloader" ] 
