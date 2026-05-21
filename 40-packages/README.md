# How to run 

```bash
cd myapp
go run ./cmd/api
```

# Packages in Go

This is one of the MOST important Go topics because Go’s philosophy around packages is VERY different from:

* Java
* C++
* Python

Go strongly emphasizes:

# simplicity + modularity + clear ownership

And a LOT of Go’s elegance comes from its package system.

---

# Core Philosophy

In Go:

# package = unit of organization + encapsulation

NOT:

* class
* namespace only
* inheritance boundary

---

# Every `.go` File Belongs To A Package

Example:

```go id="packageexample"
package main
```

OR:

```go id="anotherpackage"
package auth
```

---

# IMPORTANT RULE

All `.go` files in SAME directory:

# MUST belong to SAME package

Example:

```text id="validpackagefolder"
auth/
 ├── login.go     -> package auth
 ├── jwt.go       -> package auth
 └── user.go      -> package auth
```

VALID.

---

# INVALID

```text id="invalidpackagefolder"
auth/
 ├── login.go     -> package auth
 └── helper.go    -> package utils
```

NOT allowed.

Huge beginner gotcha.

---

# Package Name Convention

Usually:

# short

# lowercase

# singular

# meaningful

Examples:

```text id="goodpackagenames"
auth
cache
db
config
http
user
```

---

# BAD Package Names

```text id="badpackagenames"
AuthUtils
Helpers
CommonStuff
MyPackage
```

Go prefers:

# simple names

---

# Package Name Should Match Folder Name

Convention:

```text id="folderpackage"
auth/
  └── package auth
```

Very important idiomatic convention.

---

# `package main`

Special package.

Means:

# executable application

Must contain:

```go id="mainfunc"
func main()
```

---

# Non-main Packages

Used as:

* libraries
* reusable modules

---

# Export Rules (VERY IMPORTANT)

Go visibility is:

# capitalization-based

---

# Uppercase = Exported/Public

```go id="publicfield"
type User struct {
	Name string
}
```

Accessible from other packages.

---

# Lowercase = Package Private

```go id="privatefield"
type User struct {
	name string
}
```

ONLY accessible within same package.

---

# HUGE Difference From Java

Go has NO:

* public/private keywords
* protected
* friend classes

Visibility is intentionally:

# extremely simple

---

# Importing Packages

Example:

```go id="importexample"
import "fmt"
```

---

# Multiple Imports

```go id="multipleimports"
import (
	"fmt"
	"time"
)
```

---

# Using Imported Package

```go id="fmtprintln"
fmt.Println("hello")
```

---

# Aliased Imports

```go id="aliasimport"
import f "fmt"
```

Use:

```go id="aliasusage"
f.Println()
```

Rarely needed.

---

# Blank Imports

```go id="blankimport"
import _ "some/package"
```

Import for:

# side effects only

Usually:

* DB drivers
* init registration

Advanced topic.

---

# Dot Imports (Avoid)

```go id="dotimport"
import . "fmt"
```

Now:

```go id="dotusage"
Println()
```

Very discouraged.

Pollutes namespace.

---

# Package Initialization

Go supports:

```go id="initfunc"
func init()
```

Runs automatically before main.

---

# Example

```go id="initexample"
func init() {
	fmt.Println("Initializing...")
}
```

---

# IMPORTANT

Multiple init functions allowed.

But:

# overusing init is discouraged

Can create:

* hidden behavior
* startup confusion

---

# Go Modules (`go.mod`)

Modern Go dependency system.

---

# Create Module

```bash id="gomodinit"
go mod init github.com/tirthraj/project
```

Creates:

```text id="gomodfile"
go.mod
```

---

# Example `go.mod`

```go id="gomodexample"
module github.com/tirthraj/project

go 1.25
```

---

# VERY IMPORTANT

Module path becomes:

# import root

---

# Example Structure

```text id="modulelayout"
project/
 ├── go.mod
 ├── main.go
 └── auth/
      └── auth.go
```

Import:

```go id="localimport"
import "github.com/tirthraj/project/auth"
```

---

# Folder Structure Conventions

Now the important part.

---

# Small Projects

Simple structure preferred.

```text id="smallproject"
project/
 ├── go.mod
 ├── main.go
 ├── auth/
 ├── db/
 └── utils/
```

Perfectly fine.

---

# Medium/Large Projects

Common idiomatic structure:

```text id="largeproject"
project/
 ├── cmd/
 ├── internal/
 ├── pkg/
 ├── api/
 ├── configs/
 ├── scripts/
 ├── test/
 ├── go.mod
 └── README.md
```

---

# VERY IMPORTANT

Go culture:

# prefers flatter structures

NOT:

* deeply nested Java-style package trees

---

# `cmd/`

Entrypoints.

---

# Example

```text id="cmdexample"
cmd/
 ├── api-server/
 │    └── main.go
 └── worker/
      └── main.go
```

Each subfolder:

* separate executable

Very common.

---

# `internal/`

VERY important Go feature.

Packages inside:

# only importable within same module

---

# Example

```text id="internalexample"
internal/auth
```

Cannot be imported externally.

Excellent encapsulation mechanism.

---

# `pkg/`

Historically:

# reusable public libraries

But:

* somewhat controversial now

Modern Go often avoids unnecessary `pkg/`.

---

# Many Modern Projects Prefer

```text id="modernlayout"
internal/
```

over:

```text id="oldlayout"
pkg/
```

unless truly reusable externally.

---

# `api/`

Usually:

* protobufs
* OpenAPI specs
* API definitions

---

# `configs/`

Configuration files.

---

# `scripts/`

Automation scripts.

---

# `test/`

Integration tests sometimes.

Though Go often colocates tests.

---

# TEST FILES

VERY important Go convention.

---

# Test File Naming

```text id="testnaming"
user_test.go
```

---

# Test Functions

```go id="testfunc"
func TestUser(t *testing.T)
```

---

# Tests Usually Live Beside Code

Example:

```text id="colocatedtests"
auth/
 ├── auth.go
 └── auth_test.go
```

VERY idiomatic Go.

---

# Circular Imports Forbidden

Huge rule.

---

# BAD

```text id="circularimports"
auth imports db
db imports auth
```

Compiler error.

Go forces:

# acyclic dependency graph

VERY important architectural constraint.

---

# Why This Is Good

Prevents:

* tangled architectures
* hidden coupling
* giant dependency messes

This shapes Go project design heavily.

---

# Interface Placement Convention

VERY important Go idiom.

---

# BAD Java-Style

```text id="badinterfaces"
interfaces/
```

Go avoids giant interface folders.

---

# GOOD Go Style

Interfaces usually live:

# near consumer

NOT implementation.

---

# Example

Suppose:

* service uses storage

Then:

```go id="consumerinterface"
type UserStore interface {
	GetUser(id int)
}
```

belongs in:

* service package

NOT:

* db package

Huge Go design philosophy.

---

# Why?

Because:

# consumer defines required behavior

NOT producer.

Very elegant architecture principle.

---

# File Naming Conventions

Usually:

# lowercase

# snake_case optional

# descriptive

Examples:

```text id="goodfilenames"
auth.go
jwt.go
user_service.go
cache.go
```

---

# Avoid

```text id="badfilenames2"
AuthHelpers.go
Utils.go
Common.go
```

---

# “utils” Package Anti-Pattern

VERY important.

Go discourages:

# generic helper dumping grounds

Example:

```text id="utilsantipattern"
utils/
helpers/
common/
misc/
```

Usually signals:

* poor organization
* unclear ownership

---

# Prefer Domain-Oriented Packages

```text id="domainpackages"
auth/
billing/
cache/
storage/
```

Much cleaner.

---

# Package APIs Matter A LOT

Go packages should expose:

# small clean APIs

Go values:

* discoverability
* readability
* explicitness

---

# Idiomatic Go Packages Often Read Nicely

Example:

```go id="cleanapi"
cache.Set(...)
auth.Login(...)
db.Connect(...)
```

Short package names help readability.

---

# Package Cycles Force Better Architecture

One of Go’s underrated strengths.

You naturally end up designing:

* cleaner layers
* explicit dependencies
* better ownership

because compiler forces it.

---

Go packages are intentionally:

# simple but opinionated

The ecosystem heavily values:

* clear ownership
* explicit dependencies
* shallow hierarchies
* modular composition

The package system shapes Go architecture MUCH more than inheritance/class hierarchies do in OOP-heavy languages.

---
# Go generally does NOT encourage heavy: layer-first/package-by-technical-role

structures like:

```text id="layeredjava"
controller/
service/
repository/
dto/
model/
```

That style is VERY common in:

* Java Spring
* enterprise C#
* classic layered architecture

But idiomatic Go tends to prefer:

# package-by-domain/feature

This is a HUGE philosophical difference.

---

# Why Go Dislikes The Java-Style Structure

Because it often creates:

* artificial separation
* giant packages
* weak cohesion
* circular dependency pressure
* anemic domain models
* “DTO hell”

And Go strongly values:

# locality + cohesion + simplicity

---

# Example Problem

Suppose you add:

# Auth feature

Now files spread across:

```text id="spreadproblem"
controller/auth_controller.go
service/auth_service.go
repository/auth_repository.go
dto/auth_request.go
model/user.go
```

To understand ONE feature:

* must jump everywhere

Go generally dislikes this fragmentation.

---

# Idiomatic Go Often Prefers

```text id="featurepackages"
auth/
 ├── handler.go
 ├── service.go
 ├── repository.go
 ├── model.go
 └── dto.go
```

OR even flatter:

```text id="smallfeaturepackage"
auth/
 ├── auth.go
 ├── jwt.go
 └── store.go
```

Everything related to:

# auth

lives together.

Much better cohesion.

---

# VERY IMPORTANT Go Principle

Packages should represent:

# ownership boundaries

NOT:

# technical categories

Huge distinction.

---

# Java-Style Layered Architecture Optimizes For

```text id="javaoptimization"
separation by technical responsibility
```

Go tends to optimize for:

```text id="gooptimization"
feature cohesion + simplicity
```

---

# Go Encourages

```text id="gopackagephilosophy"
what belongs together?
```

NOT:

```text id="technicalsplitting"
what abstract layer is this?
```

---

# Example — Idiomatic Go Backend

Instead of:

```text id="springstyle"
controller/
service/
repository/
```

you’ll often see:

```text id="gostylebackend"
internal/
 ├── auth/
 ├── billing/
 ├── orders/
 └── users/
```

Each feature contains:

* handlers
* logic
* storage
* models

related to that domain.

---

# Why This Works Better In Go

Because Go packages are:

# lightweight

Unlike Java:

* no class explosion
* no inheritance ceremony
* no annotations everywhere

So colocating things becomes much cleaner.

---

# Another HUGE Difference

Go interfaces are usually:

# consumer-defined

So giant:

```text id="repositoryinterfaces"
repository/
```

layers often become awkward.

---

# Example Java Pattern

```java id="javarepository"
UserService
    ↓
UserRepository interface
    ↓
UserRepositoryImpl
```

Go often simplifies dramatically.

---

# Idiomatic Go Might Just Do

```go id="gosimplification"
type Store struct {
	db *sql.DB
}
```

Sometimes:

# no interface at all

if only one implementation exists.

Very important Go philosophy.

---

# Go Avoids Abstraction-First Design

Go tends to prefer:

# concrete-first

Only introduce interfaces when:

* multiple implementations
* testing needs
* abstraction genuinely useful

Huge philosophical difference from Java.

---

# DTOs In Go

Go usually avoids:

# DTO explosion

Unless:

* external APIs differ significantly
* transport/domain separation needed

Often simple structs sufficient.

---

# Example

Instead of:

```text id="dtoexplosion"
UserDTO
UserEntity
UserRequest
UserResponse
UserView
```

Go may simply use:

```go id="simpleuserstruct"
type User struct {
	...
}
```

across multiple layers.

Less ceremony.

---

# “Model” Package Is Often Discouraged

Because:

```text id="modelpackageproblem"
model/
```

becomes:

# giant dumping ground

containing unrelated business entities.

Go prefers:

# domain ownership

---

# Example Better Structure

```text id="domainownership"
auth/user.go
billing/invoice.go
orders/order.go
```

instead of:

```text id="centralmodel"
model/user.go
model/invoice.go
model/order.go
```

---

# But IMPORTANT

Go is NOT dogmatic.

Your proposed structure is NOT “wrong”.

Especially if:

* team already uses it
* project large enterprise system
* architecture standardized

Many Go companies STILL use layered structures.

---

# Realistic Hybrid Structure

Many mature Go codebases use something like:

```text id="hybridstructure"
internal/
 ├── auth/
 │    ├── handler.go
 │    ├── service.go
 │    ├── store.go
 │    └── types.go
 │
 ├── billing/
 │    ├── handler.go
 │    ├── service.go
 │    └── store.go
```

This gives:

* feature cohesion
* still some layer separation internally

VERY common sweet spot.

---

# Another Important Go Philosophy

Go prioritizes:

# simplicity over purity

So many “enterprise architecture” patterns become:

* intentionally simplified
* flattened
* pragmatic

---

# Example

Java might have:

```text id="javaexplosion"
Controller
Service
Manager
Facade
Mapper
Repository
Factory
Builder
DTO
Entity
```

Go engineers often look at this and ask:

```text id="gophilosophyquestion"
"Do we actually need all these abstractions?"
```

Very different engineering culture.

---

# VERY IMPORTANT

Go codebases usually optimize for:

# readability by average engineer

not:

# maximal abstraction flexibility

Huge cultural distinction.

---

# Insights

The biggest architectural shift moving from Java-style systems to Go is:

# from abstraction-oriented design

to

# cohesion-oriented design

Go tends to favor:

* fewer layers
* fewer abstractions
* stronger feature locality
* simpler dependency graphs
* concrete implementations first

This often results in systems that are:

* easier to navigate
* easier to reason about
* easier to refactor
* operationally simpler

especially at scale.

