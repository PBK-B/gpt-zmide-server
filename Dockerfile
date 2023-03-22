FROM archlinux:base-devel-20220626.0.64095 AS builder
# RUN echo 'Server = http://mirrors.tuna.tsinghua.edu.cn/archlinux/$repo/os/$arch' > /etc/pacman.d/mirrorlist

LABEL stage=gobuilder

ENV CGO_ENABLED 0

WORKDIR /build

RUN pacman-db-upgrade
RUN pacman -Sy --noconfirm go nodejs wget npm curl openssl icu

# npm install
COPY package.json package-lock.json ./
RUN npm install

# go mod download
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# # build react and golang
COPY . .
RUN npm run build
RUN go build -v -o /build/gptzmideserver .

# CMD [ "/build/gptzmideserver" ]


FROM alpine

EXPOSE 8091

COPY --from=builder /build/gptzmideserver /bin/gptzmideserver

WORKDIR /bin

CMD [ "/bin/gptzmideserver" ]
