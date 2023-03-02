# golang image
FROM golang:1.20-alpine

# set app folder as default working directory
WORKDIR /app

# copy to app folder
COPY . .

# create executable
RUN go build -o main .

# execute executable
CMD ["/app/main"]