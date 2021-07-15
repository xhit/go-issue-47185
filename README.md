Go issue 47185: https://github.com/golang/go/issues/47185

To get the build error, compile with go1.16.6 with command:

```bash
go build -tags "sqlite3driver"
```

This compile fine without the tag.

This compile fine with tag in from Go 1.13.5 to 1.16.5

The build error only happens in Windows 10 amd64. I can build successfully in Linux amd64, arm64 and macOS amd64.

I'm using gcc 9.2.0 from MSYS2.