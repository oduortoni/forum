# forum

## description

The objective of this project is to create a **web forum** that allows users to register, login, post content, comment on posts, like/dislike posts, and filter posts based on categories, created date, and liked status.

## usage

The server itself can run on different ports dependong on the PORT envirenment variable. If it is not defined, the default port will be used. If you need to run the server on a specific port, then define the environment variable PORT.

```bash
export PORT=9000
```

The server will now listen on port 9000. This is inline with shared hosting providers who provide the port through an enveironment variable.

**run the program**

```bash
go run .
```

You will sees a message informing you on which port the server is listening for incoming calls

