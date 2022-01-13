// vim: tabstop=4 shiftwidth=4 expandtab
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/pantoniou/go-fyaml"
)

var copyp = flag.Bool("copy", false, "Copy file to memory")
var debugp = flag.Bool("debug", false, "debugging")
var verbosep = flag.Bool("verbose", false, "verbose")
var resolvep = flag.Bool("resolve", false, "resolve")

func main() {

    flag.Parse()

    if len(flag.Args()) < 1 {
        log.Fatal("Missing file argument")
    }
    file := flag.Args()[0]

    debug := ""
    if *debugp {
        debug = "debug"
    }
    verbose := ""
    if *verbosep {
        verbose = "verbose"
    }
    resolve := ""
    if *resolvep {
        resolve = "resolve"
    } else {
        resolve = "noresolve"
    }

    if *debugp {
        fmt.Printf("test yaml file %s\n", file)
        fmt.Printf("libfyaml library version: %v\n", fyaml.LibraryVersion())
        fmt.Printf("Supported versions: %v\n", fyaml.VersionsSupported())
        fmt.Printf("Supported schemas: %v\n", fyaml.ListSchema())
    }

    var v interface{}
    var err error
    var data []byte

    json := ""
    if strings.HasSuffix(file, ".json") {
        json = "json=force"
    }

    if *debugp {
        fmt.Printf("Unmarshaling via Unmarshal()\n")
    }
    if data, err = os.ReadFile(file); err != nil {
        log.Fatal(fmt.Sprintf("os.ReadFile failed: %s", err.Error()))
    }

    if err = fyaml.Unmarshal(data, &v, debug, verbose, json, resolve); err != nil {
        log.Fatal(fmt.Sprintf("Unmarshal failed: %s", err.Error()))
    }

    // and print out
    if *debugp {
        fmt.Printf("\nunmarshal-output:\n")
        fmt.Printf("%+v\n", v)

        // and print out type
        fmt.Printf("\nunmarshal-output-type:\n")
        fmt.Printf("%T\n", v)
    }

    // output, in pretty mode always
    data, err = fyaml.Marshal(&v, debug, verbose, json, resolve, "output-mode=pretty")
    if err != nil {
        log.Fatal(fmt.Sprintf("Marshal failed: %s", err.Error()))
    }

    if *debugp {
        fmt.Printf("\nunmarshal-output:\n")
    }
    fmt.Printf("%s", string(data))
}
