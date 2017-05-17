# le-z
Rule based torrent service, server, and client. Making it easy to be lazy as fuck and download great torrents.

### Installing

```
go get github.com/kris-nova/le-z
```

### Running the server

The server is designed to act as a functioning torrent seedbox.
Based on arbitrary rules the server will look for *magnet* URLs in various places.

```
le-z serve <flags>
```

### Running the client

The client is designed to run as a thin CLI client that communicates with the server.
All communication should be encrypted.
The client will be able to manage the server.
The client will be able to retrieve files from the server.

```
le-z <flags>
```