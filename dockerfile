FROM Scratch
CMD ["touch","./saved.txt"]
COPY ./src/site ./site
COPY ./build/server ./build
RUN ./build