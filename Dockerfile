# FROM phongthien/swagger as swagger
# ARG ACCESS_TOKEN_USER
# ARG ACCESS_TOKEN
# ARG APP
# ARG VERSION
# WORKDIR /src
# COPY . .
# RUN /root/swag init --parseDependency -d ./apps/$APP/src  -o apps/$APP/docs

# First stage: build the executable.
FROM golang:1.20-alpine AS builder
# It is important that these ARG's are defined after the FROM statement
WORKDIR /src

# Import the code from the context.
COPY . .
# COPY --from=swagger /src/apps/$APP/docs apps/$APP/docs
# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app ./cmd/cli/main.go



# Final stage: the running container.
FROM scratch AS final
WORKDIR /src
COPY --from=builder /app /app
CMD ["/app"]
