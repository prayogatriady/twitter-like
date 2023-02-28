# golang image
FROM golang:latest

# create app folder
RUN mkdir /app

# copy to app folder
ADD . /app

# set app folder as default working directory
WORKDIR /app

# create executable
RUN go build -o main .

# execute executable
CMD ["/app/main"]