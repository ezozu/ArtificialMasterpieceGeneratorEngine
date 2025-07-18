// cmd/artificialmasterpiecegeneratorengine/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorengine/internal/artificialmasterpiecegeneratorengine"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorengine.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
