FROM    golang

EXPOSE  8080

ADD     app.go app.go

ENTRYPOINT  ["go", "run", "app.go"]     
