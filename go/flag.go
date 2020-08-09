package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var count int
var name string
var fork bool

// custom flag type
type timeInterval []time.Duration

var interval timeInterval

// String formats the interval's value.
// Required for type interval to implement the flag.Value interface.
func (i *timeInterval) String() string {
	return fmt.Sprint(*i)
}

// Set sets the interval's value.
// Set's argument is a string to be parsed to set the flag.
// Required for type interval to implement the flag.Value interface.
func (i *timeInterval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, val := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(val)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

func init() {
	flag.IntVar(&count, "count", 0, "the count of items")
	flag.IntVar(&count, "c", 0, "count (shorthand)")

	flag.StringVar(&name, "name", "", "the name")
	flag.StringVar(&name, "n", "", "name (shorthand)")

	flag.BoolVar(&fork, "fork", false, "fork as a daemon process")
	flag.BoolVar(&fork, "f", false, "fork (shorthand)")

	flag.Var(
		&interval,
		"interval",
		"comma-separated list of times\nvalid units are 'ns', 'us', 'ms', 's', 'm', 'h'",
	)
	flag.Var(&interval, "i", "interval (shorthand)")
}

func checkRequired() {
	if name == "" {
		fmt.Fprint(os.Stderr, "[error] name is required", "\nusage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	checkRequired()

	fmt.Println("parse CLI flags with the flag package")
	fmt.Printf("name: %s\tcount: %d\tfork: %v\tinterval: %v", name, count, fork, interval)

	args := flag.Args()
	if len(args) > 0 {
		fmt.Printf("\nargs not parsed as flags %v\n", args)
	}
}
