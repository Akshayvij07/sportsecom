# build stage
# FROM golang:1.18.1-alpine AS build-stage

# WORKDIR /app
# # download the dependancy
# COPY go.mod go.sum ./
# RUN  go mod download
# COPY . .
# # copy the source code and html files
# #COPY templates templates/

# # build the executable file
# RUN go build -v -o ./build/api ./cmd/api

# # final stage
# FROM apline:3. AS prod

# WORKDIR /app
# # copy the binay file and html files
# COPY --from=build-stage /app/build/api api

# CMD ["/app/api"]


# # build stage
# FROM golang:1.18.1-alpine AS build-stage

# WORKDIR /app
# # download the dependency
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# # copy the source code and HTML files
# # COPY templates templates/

# # build the executable file
# RUN go build -v -o ./build/api ./cmd/api

# # final stage
# FROM alpine:3.14 AS prod

# WORKDIR /app
# # Install any necessary runtime dependencies here

# RUN apk --no-cache add tzdata

# # copy the binary file and HTML files from the build stage
# COPY --from=build-stage /app/build/api api

# CMD ["/app/api"]



# build stage
FROM golang:1.18.1-alpine AS build-stage

WORKDIR /app
# download the dependancy
COPY go.mod go.sum ./
RUN  go mod download
# copy the source code and html files
#COPY . .
COPY cmd cmd/
COPY pkg pkg/
COPY templates templates/
# build the executable file
RUN go build -v -o ./build/api ./cmd/api

# final stage
FROM alpine:3.14 AS prod

WORKDIR /app
# copy the binay file and html files
COPY --from=build-stage /app/build/api api
COPY --from=build-stage /app/templates templates/
ENV test = 123

CMD ["/app/api"]
