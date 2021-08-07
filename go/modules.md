# Go modules

### Create new module

Make the current directory the root of a module:
```
go mod init example.com/hello
```

List the current module and all its dependencies:
```
go list -m all
```

### Module management

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
