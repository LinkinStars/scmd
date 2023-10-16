# scmd
> Add a stop command for your Golang application.

## Usage

```go
import _ "github.com/LinkinStars/scmd"
```

## Example

```bash
$ ./app stop
Try to find process by name [app]...
Find process [50406][app][./app run], stop it? yes(y)/no(n)
yes
The process has been stopped
```

- `app` is your application name.
- `50406` is pid.
- scmd will find the process via filename by default, but you can specify the process name by `./app stop app-server`.