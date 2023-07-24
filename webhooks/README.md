# WEBHOOKS in GOLANG

In simple words a webhook is like a special bell attached to your magic mailbox, Whenever you friends send you a letter or a surprise, the magic mailbox rings the bell to let you know that something new has arrived. You don't have to keep checking the mailbox every few minutes because the bell will tell you exactly when there's something new for you. So, webhooks are like a notification system that tells you when new things happen.
A webhook is a way for one application to tell another application somethings happened

In this example , we'll create a basic HTTP server that listens for incoming webhook requests and prints the received data.

# Question
Write a webhook by creating a function that listens for HTTP requests. When a request comes in, the function will parse the request body and do something with the information.

# How it works

send the request from terminal
curl -X POST -d "hello world" localhost:8080/webhook

# Response from reciever

Webhook received successfully!
received the message: hello world%      