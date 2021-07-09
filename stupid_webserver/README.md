# Stupid HTTP Server
This is an elementary HTTP server I made to practice basic Go socket
programming. It serves html files (if it finds one), or returns 404. Nothing
more.

To launch:
```
$ HTTP_LOCATION=$PWD/static go run cmd/main.go
```
