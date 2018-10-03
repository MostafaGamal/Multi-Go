package main

/*
Copyright 2018 TheRedSpy15

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/daviddengcn/go-colortext"
)

func main() {
	parser := argparse.NewParser("SecureMultiTool", "Runs multiple security orientated tasks")

	// TODO: use native go flags
	// Create flags
	t := parser.String("t", "Task", &argparse.Options{Required: true, Help: "Task to run"})
	r := parser.String("r", "Target", &argparse.Options{Required: false, Help: "Target to run task on"})

	// Error handling
	err := parser.Parse(os.Args)
	if err != nil {
		panic(parser.Usage(err))
	}

	// TODO: add 'bleach'
	// TODO: add 'compress'
	// TODO: add 'uncompress'
	// TODO: add 'pwnAccount'
	// TODO: add 'DOS' (IP based)
	// TODO: add 'Audit' (3rd party integrated system)
	// TODO: add 'toggleIncoming' (inbound connections)
	// TODO: add 'email'
	// Determine task
	switch *t {
	case "Hash":
		println("\nRunning task:", *t, "\nTarget:", *r)
		hashFile(*r)
	case "pwnAccount":
		println("\nRunning task:", *t, "\nTarget:", *r)
		pwnAccount(*r)
	case "encryptFile":
		println("\nRunning task:", *t, "\nTarget:", *r)
		encryptFileTask(*r)
	case "decryptFile":
		println("\nRunning task:", *t, "\nTarget:", *r)
		decryptFileTask(*r)
	case "Scrape":
		println("\nRunning task:", *t, "\nTarget:", *r)
		scrape(*r)
	case "Email":
		emailTask()
	case "About":
		about()
	case "List":
		listTasks()
	default:
		ct.Foreground(ct.Red, true)
		println("Invalid task - ", *t)
		ct.Foreground(ct.Yellow, false)
		println("Use '--help' or '-t List'")
	}
}
