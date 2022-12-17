# Durationparse

Parse simple duration strings such as `2d 3h`.

Code is implemented with regular expressions using the standard library.

## Usage

```go
import "github.com/adriansahlman/durationparse"

func main() {
    duration, err = durationparse.Parse("2 days, 3hr 5 min")
}
```
