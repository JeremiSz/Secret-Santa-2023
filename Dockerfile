FROM golang
RUN ["touch","./saved.txt"]
COPY ./src/site ./site
COPY ./build/server ./build
EXPOSE 8080/tcp
EXPOSE 8080/udp
ENTRYPOINT [ "./build" ]