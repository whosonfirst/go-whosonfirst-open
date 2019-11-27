# go-whosonfirst-open

## Tools

### open

Open one or more Who's On First (WOF) documents specified only by their IDs.

```
> go run -mod vendor cmd/open/main.go -h
  -editor string
    	The editor to open a WOF ID with. If empty the value of the 'EDITOR' environment variable with be used.
  -repo
    	Indicates that -root is a whosonfirst style repo and appends a 'data' folder to its path.
  -root string
    	The path where the path for each WOF ID lives. If empty the current directory is used.
```

For example:

```
go run -mod vendor cmd/open/main.go -root /usr/local/data/sfomuseum-data-exhibition -repo 1159159407
```

Will cause `/usr/local/data/sfomuseum-data-exhibition/data/115/915/940/7/1159159407.geojson` to be opened by the application defined by the `-editor` flag.

If you pass multiple Who's On First IDs each ID will be opened (by the application defined by	the `-editor` flag) in sucession.

## To do

* Integrate with the [go-reader packages](https://github.com/whosonfirst?utf8=%E2%9C%93&q=go-reader&type=&language=) so that it is possible to open remote WOF documents.