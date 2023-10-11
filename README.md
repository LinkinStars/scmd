# scmd
> Add a stop command to your application.

## Usage

```go
import _ "github.com/LinkinStars/scmd"
```

## Example

```bash
$ app stop

Find process [50406][app][app], stop it? yes/no
yes
The process has been stopped

```

- `app` is your application name.
- `50406` is pid.
- scmd find the process by name 'app', but you can specify the process name by `app stop app-server`.