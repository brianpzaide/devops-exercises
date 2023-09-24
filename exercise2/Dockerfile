FROM golang:1.19-alpine AS builder
RUN apk update && apk add make git
WORKDIR /home/exercise2
COPY . .
RUN make build

FROM golang:1.19-alpine AS deploy
RUN apk --no-cache add ca-certificates
WORKDIR /home/exercise2
COPY --from=builder /home/exercise2/bearsintheforest .
CMD ["./bearsintheforest"]