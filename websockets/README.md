# Websockets

Websockers are a way for two applications to communicate with each other in real time. 

# Code

In this example we will write a websocket server by creating a function that listens for websocket conections. when a connection comes in, the function will create a new websocket object and start and start a goroutine to handle the connection.