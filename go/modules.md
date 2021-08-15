# Go modules

## Contents
- [Create a new module](#create-a-new-module)
- [Module management basics](#module-management-basics)
- [Import a local package](#import-a-local-package)
- [Vendor dependencies](#vendor-dependencies)

### Create new module

Make the current directory the root of a module:
```
go mod init example.com/hello
```

List the current module and all its dependencies:
```
go list -m all
```

### Module management basics

Get a specific module version:
```
go get example.com/hello@v1.3.1
```

Get a branch:
```
go get example.com/hello@<branch-name>
```

Get a commit:
```
go get example.com/hello@<commit>
```

Add missing and remove unused modules:
```
go mod tidy
```

Explain why packages or modules are needed:
```
go mod why -m <module-path>
```

### Import a local package

#### Scenario
You have a local package `sample.com/math` that does not exist in a remote repo
that you would like to import into your `school` project. How can we accomplish
this?

First you must initialize the modules of each project:
```
# in the math dir
go mod init sample.com/math

# in the school dir
go mod init school
```

The projects should both contain `go.mod` files:
 ```
~/code/go/module-test  $ tree
.
├── math
│   ├── go.mod
│   └── math.go
└── school
    ├── go.mod
    └── main.go

2 directories, 4 files
 ```

Next, we need to manually update the `school` project's `go.mod` file using the
[replace directive](https://golang.org/ref/mod#go-mod-file-replace)
```
module school

go 1.16

replace sample.com/math => ../math
```

At this point, if you try to run the `main.go` file from the `school` project
you will likely see the following error:
```
~/code/go/module-test/school  $ go run main.go
main.go:6:2: module sample.com/math provides package sample.com/math and is replaced but not required; to add it:
	go get sample.com/math
```

The final step is to run [go mod tidy](https://golang.org/ref/mod#go-mod-tidy) which will find the local package:
 ```
~/code/go/module-test/school  v $ go mod tidy
go: found sample.com/math in sample.com/math v0.0.0-00010101000000-000000000000
 ```

Finally, we can successfully import the local package:
```
~/code/go/module-test/school  $ go run main.go
6
```

### Vendor dependencies

#### What is a vendor dir?
> Vendor directories serve two purposes. First, they specify by their contents
> the exact version of the dependencies to use during gobuild. Second, they
> ensure the availability of those dependencies, even if the original copies
> disappear. On the other hand, vendor directories are also difficult to manage
> and bloat the repositories in which they appear.

Ref: [Defining Go Modules (Go & Versioning, Part 6)](https://research.swtch.com/vgo-module)

#### Usage

Make a vendored copy of dependencies:
```
go mod vendor
```

Build dependencies from the vendor directory:
```
go build -mod vendor
```

By default `go build` will ignore the vendor directory.
