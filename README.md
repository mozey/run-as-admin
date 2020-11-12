# run-as-admin
Golang example of using a manifest file to prompt "Run as administrator"

## Build app with manifest

Install [rsrc](https://github.com/akavel/rsrc)
tool for embedding binary resources in Go programs

    go get github.com/akavel/rsrc
    
Generates a .syso file with embedded resources

    rsrc -manifest app.exe.manifest -o app.syso

Build executable

    GOOS="windows" GOARCH="386" go build -o app.exe


## Reference

[how-to-ask-for-administer-privileges-on-windows-with-go](http://stackoverflow.com/questions/31558066/how-to-ask-for-administer-privileges-on-windows-with-go)

[check-if-application-is-running-as-administrator-in-golang](http://stackoverflow.com/questions/27366298/check-if-application-is-running-as-administrator-in-golang)


# Without a manifest

See [this technique](https://stackoverflow.com/a/59147866/639133), "...run as a standard user in most cases, and only elevate when needed. I use this in command line tools where most functions don't need admin rights"

