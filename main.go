package main
 
import (
    "fmt"
    "os"
    "strings"
    "flag"

    "github.com/flosch/pongo2"
)
 
func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(2)
}



func parseline(line string) (key string, value string) {
    splitString := strings.SplitN(line, "=", 2)
    key = splitString[0]
    value = splitString[1]

    return
}

func getenv() (envMap pongo2.Context, err error) {
    var env []string
    env = os.Environ()
    envMap = make(pongo2.Context)

    for _, line := range env {
        key, value := parseline(line)
        if err == nil {
            envMap[key] = value
        }
    }
    return
}

func main() {
    if len(os.Args) < 2 {
        usage()
    }

    var tpl = pongo2.Must(pongo2.FromFile(os.Args[1]))

    envMap, err := getenv()

    renderedTemplate, err := tpl.Execute(envMap)
    if err != nil {
        panic(err)
    }
    fmt.Println(renderedTemplate)
}

