Go is a new language whose working is diff from standard languages.

+ Here we're initially keeping the basics and file structure and naming conventions as they are traditionally used, we'll move forward to changing things only once we've better understanding of modules in Go.

+ get into a habit of reading the documentation as you learn about any new resource.

+ for Go documentation refer: go help <command for which documentation is required>

+ Walrus operator is allowed only inside the method, not outside. Walrus operator is used to declare a identifier without var/const and even without a datatype

+ You can define an identifier without mentioning a datatype since lexer helps to auto initialize a datatype based on the value being assigned.

+ Go does not put garbage values inside variables if they are not assigned any value at the time of initialization, the identifier is automatically assigned the zero value in this case.

+ For input and output we shall use bufio 

+ In golang, there is no try and catch, so if something goes wrong, there is comma ok syntax or err ok syntax

 + Bufio by default takes a string input

@Handling Time in GoLang:

time.Now() -> gives the current time
You can modify the format according to your need but you need to specify the exact mentioned syntax

Core Types in time Package
time.Time:
Represents an instant in time with nanosecond precision.
Used for most operations like getting the current time, formatting, and comparisons.
time.Duration:
Represents the elapsed time between two time.Time instances or a span of time (e.g., 5 hours).
It’s essentially an int64 representing nanoseconds.


Getting the Current Time
Use time.Now() to get the current time.

Formatting and Parsing Time
Go uses a unique reference time Mon Jan 2 15:04:05 MST 2006 for formatting and parsing. Replace parts of this reference time to create custom layouts

Manipulating Time
You can add or subtract durations to/from a time.Time

now := time.Now()
later := now.Add(2 * time.Hour)
earlier := now.Add(-30 * time.Minute)
start := time.Now()
    end := start.Add(10 * time.Second)
    duration := end.Sub(start)

    Time Zones
The time package supports time zones via the Location type.

Comparing Time
You can compare time.Time instances using methods like Before, After, and Equal.

Sleep
Pause execution for a specified duration using time.Sleep

time.Sleep(2 * time.Second)

All of the go related details are places inside GoLang's env file:

go env

under this, you'll find your current OS under GOOS: (not hostos)

you can change between the systems using 
GOOS="<linux/darwin/windows>" go build

Memory Allocation happens automatically in GoLang:
Methods used in MM:-> new() and make()


@STUDENTS API:

Start off by creating a new module at the beginning of a new project.

go mod init <module name/project name>

generic convention:

github.com/<username>/<project/repo name>

In go, the main entrypoint is the main.go file

database used in the project is sqllite, and it is a file base datastore, so we have to specify the path for the same.  
indentation is vimp in yaml files,




