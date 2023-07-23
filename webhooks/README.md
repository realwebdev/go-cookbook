# WEBHOOKS in GOLANG

In simple words a webhook is like a special bell attached to your magic mailbox, Whenever you friends send you a letter ot a surprise, the magic mailbox rings the bell to let you know that something new has arrived. You don't have to keep checking the mailbox every few minutes because the bell will tell you exactly when there's something new for you. So, webhooks are like a notification system that tells you when new things happen.

In this example , we'ww create a basic HTTP server that listens for incoming webhook requests and prints the received data.

# How it works

send the request from terminal
curl -X POST -d "hello world" localhost:8080/webhook

# Response from reciever

Webhook received successfully!
received the message: hello world%      