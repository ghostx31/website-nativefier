FROM golang:1.19 as builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o nativefier-downloader

FROM node:latest AS nodeimage
WORKDIR /usr/src/app
COPY --from=builder /app/static/package*.json ./
RUN npm install
COPY --from=builder /app ./

FROM debian:latest
ENV DEBIAN_FRONTEND noninteractive
WORKDIR /app
COPY --from=nodeimage /usr/src/app ./
RUN apt update -y && \
  apt install -y wget ca-certificates software-properties-common gnupg2 nodejs npm tar
# Installing wine since we need it for packaging windows stuff
RUN echo "deb https://dl.winehq.org/wine-builds/debian/ bullseye main" >> /etc/apt/sources.list.d/wine.list && \
  wget -nc https://dl.winehq.org/wine-builds/winehq.key && \
  apt-key add winehq.key && \
  dpkg --add-architecture i386 && \
  apt update -y
RUN apt install -y --install-recommends winehq-staging
# Install wine-mono from their site since its not available in Debian repo.
RUN wget https://dl.winehq.org/wine/wine-mono/7.4.0/wine-mono-7.4.0-x86.tar.xz
RUN tar -xvf wine-mono-7.4.0-x86.tar.xz
RUN mv wine-mono-7.4.0 /usr/share/wine/
EXPOSE 1323
ENTRYPOINT [ "./nativefier-downloader" ]
