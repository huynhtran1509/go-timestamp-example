### WillowTree is Hiring!

Want to write Go for mobile applications? Want to write anything else for mobile
applications? [Check out our openings!](http://willowtreeapps.com/careers/)

# Timestamp: An Example Go Mobile Project

This is an example project showing how business logic can be written in Go
and used on the iPhone.

It uses [rootx](http://github.com/willowtreeapps/rootx) to generate data access
methods against a Sqlite database.

## Installation

Make sure you're on at least Go version 1.5.1 and that you have properly set up
your `GOPATH`.

```
go get github.com/willowtreeapps/timestamp/...
```

### Usage

To run the Go example, use the included `gotimestamp-demo` command:

```
$ gotimestamp-demo
Found timestamps:
  Wed Dec 31 19:00:00 EST 1969
  ...
```

or run it manually from this project root:

```
$ go run ./cmd/gotimestamp-demo/main.go
Found timestamps:
  Wed Dec 31 19:00:00 EST 1969
  ...
```

To run the iOS example, make sure you have `gomobile` ready:

```
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init
```

Then create the framework and open the app for running.

```
gomobile bind -target=ios github.com/willowtreeapps/timestamp
open ios/Timestamps.xcodeproj/
```

## Output

![](https://github.com/willowtreeapps/go-timestamp-example/blob/master/sample-output.png)
