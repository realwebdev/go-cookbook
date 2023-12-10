I am still learning Go whenever I have time so take this with some doubts:

learn concurrency patterns

learn about the sync package

build some service that uses something else than HTTP (like websockets, gRPC, raw TCP sockets)

build some command line tools as it is widely used for this

build some service that use message brokers to exchange data (MQTT, NATs, RabbitMQ)

Add cache to your HTTP servers both in memory and maybe Redis

Basically extend your knowledge beyond web development to lower level things.


# Go interview questions or candidates background

When I worked in hiring, I largely glanced at projects on resumes and mostly just cared about prior experience and their ability to answer technical questions.

So I'd ask stuff like:

how does Go handle concurrency? How is that different from C threads? How is it different from JavaScript? Explain some pros and cons of each.

what constructs does Go provide to help with concurrency? When might you use one over another?

coding challenge: demonstrate the following: single producer, multiple consumer; multiple producer, single consumer, multiple producer, multiple consumer; explain some gotchas with each and explain how you'd architect around them

explain Go's error handling; compare and contrast panics and exceptions in our languages; describe some pros and cons of Go's error handling vs other languages

describe the select control structure and explain how you'd use it in a project; how do these behave when you close a channel?

Some of these are senior level questions, though I'd expect a mid level to be able to get at least a partial answer to all of them. If you're not confident in fielding questions like these, you should apply at a junior level and learn more about concurrency in the meantime.

Most people pick up on the other concepts like interfaces, loops, etc, I'm more interested in what makes Go especially unique, and that's concurrency. I'd start with these more basic questions, but the above is what separates a junior from a mid from a senior in my mind.

# learning

https://docs.google.com/document/u/0/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/mobilebasic

# open source contribution

https://goodfirstissue.dev/language/go/

