# Build stage
FROM golang AS build-env

ADD ./app /src/order-service/app
ADD ./go.mod /src/order-service
ADD ./go.sum /src/order-service

ENV CGO_ENABLED=0

ARG API_PROTO_DEPENDENCIES_TOKEN

RUN git config --global url."https://oauth2:${API_PROTO_DEPENDENCIES_TOKEN}@gitlab.mapcard.pro/external-map-team/api-proto".insteadOf https://gitlab.mapcard.pro/external-map-team/api-proto
RUN go env -w GOPRIVATE=gitlab.mapcard.pro/external-map-team/*

RUN cd /src/order-service && go build -o /app ./app/cmd/main.go

# Production stage
FROM scratch

COPY --from=build-env /app /

ENTRYPOINT ["/app"]
