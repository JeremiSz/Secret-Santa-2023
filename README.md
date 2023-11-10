# Secret Santa 2023
This server provides a simple method to deceminate secret santa targets and allow for sharing of wishlists without revealing shoppers.

# Building & Running

## Building
For all platforms, golang is required.
Use the offical guide to install Golang on your platform. https://go.dev/learn/

### Linux and MacOS
Run the build.sh to create an executable.

### Windows
Run the command contained in the build.sh file in your terminal. (Make sure it is run from the same folder as the script or adapt the paths accordingly)

` go build -o ./build/server ./src`

### Container 
**Linux only**

Run the build.sh then build the dockerfile.

`docker build -t \<container name\> .`

or

`podman build -t \<container name\> .`

## Running
The port number can be spesified by appending it to the executable.

`./server \<port\>`

### Linux and MacOS
Run start.sh to open the server on port 8080.

### Windows
Create a saved.txt file in the src folder. 

Run the build executable in the build folder from the src folder.

`../build/server.exe 8080`

### Container
Start the container.

`docker run \<container name\>`

or

`podman run \<container name\>`

# Customisation
To change the names and who gets who, edit **wishlists.go**. 
Change the mappings dictorary to change who gets gifts from who.

> 0:0 means index 0 (Admin) gets a gift from index 0.

Fill the names array with your members' names.

To change the amount of people, change the *PEOPLE_COUNT* variable in **handler.go**.



