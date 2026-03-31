# Go Learning Resources

A curated guide to everything on [go.dev](https://go.dev), organized for learners and practitioners. Whether you're writing your first `Hello, World!` or digging into compiler internals, there's something here for you.

## Contents

- [Start Here](#start-here)
- [Tutorials](#tutorials)
- [Language Specification & Reference](#language-specification--reference)
- [Modules & Dependencies](#modules--dependencies)
- [Blog: Language Design & Features](#blog-language-design--features)
- [Blog: Generics](#blog-generics)
- [Blog: Concurrency](#blog-concurrency)
- [Blog: Performance & Internals](#blog-performance--internals)
- [Blog: Error Handling](#blog-error-handling)
- [Blog: Testing](#blog-testing)
- [Blog: Security](#blog-security)
- [Blog: Tooling](#blog-tooling)
- [Blog: Standard Library Deep Dives](#blog-standard-library-deep-dives)
- [Blog: WebAssembly](#blog-webassembly)
- [Blog: Modules & Versioning](#blog-modules--versioning)
- [Blog: Community & Ecosystem](#blog-community--ecosystem)
- [Blog: Release Announcements](#blog-release-announcements)
- [Blog: History & Milestones](#blog-history--milestones)
- [Case Studies](#case-studies)
- [Documentation](#documentation)
- [A Tour of Go](#a-tour-of-go)
- [Talks & Presentations](#talks--presentations)

## Start Here

If you're new to Go, this is the path:

- [Learn Go](https://go.dev/learn/) - Official learning hub with books, courses, and guided tutorials
- [A Tour of Go](https://go.dev/tour/) - Interactive introduction you can do in your browser
- [How to Write Go Code](https://go.dev/doc/code) - Setting up your workspace and writing your first program
- [Effective Go](https://go.dev/doc/effective_go) - The style and idioms that make Go code feel like Go
- [Frequently Asked Questions](https://go.dev/doc/faq) - Honest answers to "but why does Go...?"

## Tutorials

Hands-on, step-by-step guides. Best done in order if you're starting out.

- [Getting started](https://go.dev/doc/tutorial/getting-started) - Install Go and write your first program
- [Create a module](https://go.dev/doc/tutorial/create-module) - Build something reusable
- [Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces) - Work across multiple modules locally
- [Accessing a relational database](https://go.dev/doc/tutorial/database-access) - Connect to SQL, query rows, handle transactions
- [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin) - Build a web service
- [Getting started with generics](https://go.dev/doc/tutorial/generics) - Write type-safe generic functions
- [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz) - Find bugs automatically with fuzz testing
- [Find and fix vulnerable dependencies with govulncheck](https://go.dev/doc/tutorial/govulncheck) - Security scanning for your project
- [Vulnerability scanning in your IDE](https://go.dev/doc/tutorial/govulncheck-ide) - Real-time vulnerability detection

## Language Specification & Reference

The definitive source of truth for the language.

- [The Go Programming Language Specification](https://go.dev/ref/spec) - Every rule, every edge case, all of it
- [Go Modules Reference](https://go.dev/ref/mod) - Complete reference for the module system (`go.mod`, `go.sum`, commands)
- [Go Doc Comments](https://go.dev/doc/comment) - How to write documentation comments that tools understand
- [Go Memory Model](https://go.dev/ref/mem) - Guarantees about when goroutines can see each other's writes
- [Debugging with GDB](https://go.dev/doc/gdb) - Lower-level debugging when Delve isn't enough

## Modules & Dependencies

- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies) - Adding, upgrading, and removing dependencies
- [Developing and publishing modules](https://go.dev/doc/modules/developing) - Creating modules others can use
- [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow) - When to tag, how to version
- [Managing module source](https://go.dev/doc/modules/managing-source) - Repository layout and organization
- [Module version numbering](https://go.dev/doc/modules/version-numbers) - What v1.2.3 actually means
- [Major version updates](https://go.dev/doc/modules/major-version) - The v2+ story and import path changes
- [Publishing a module](https://go.dev/doc/modules/publishing) - Making your module available to the world
- [Module layout](https://go.dev/doc/modules/layout) - Recommended project structure
- [Module graph pruning](https://go.dev/doc/modules/pruning) - How Go keeps your dependency graph lean

## Blog: Language Design & Features

Posts that explain *why* Go works the way it does.

- **2010-07-07** [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax) - Why `x int` instead of `int x`
- **2014-08-25** [Constants](https://go.dev/blog/constants) - Untyped constants and why they matter more than you think
- **2013-10-23** [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) - The full story of text in Go
- **2013-09-26** [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices) - How slices really work under the hood
- **2013-02-06** [Go maps in action](https://go.dev/blog/maps) - Everything about maps
- **2011-01-05** [Go Slices: usage and internals](https://go.dev/blog/slices-intro) - A gentler intro to slices
- **2010-08-04** [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover) - Go's approach to cleanup and error recovery
- **2011-09-06** [The Laws of Reflection](https://go.dev/blog/laws-of-reflection) - Understanding `reflect` from first principles
- **2015-01-12** [Errors are values](https://go.dev/blog/errors-are-values) - Rob Pike on why `if err != nil` is a feature
- **2010-07-13** [Share Memory By Communicating](https://go.dev/blog/codelab-share) - The philosophy behind channels
- **2012-08-16** [Organizing Go code](https://go.dev/blog/organizing-go-code) - Project structure advice
- **2015-02-04** [Package names](https://go.dev/blog/package-names) - Naming conventions that make APIs pleasant
- **2015-05-07** [Testable Examples in Go](https://go.dev/blog/examples) - Examples that are also tests
- **2013-11-26** [Text normalization in Go](https://go.dev/blog/normalization) - Unicode normalization deep dive
- **2014-12-22** [Generating code](https://go.dev/blog/generate) - `go generate` and code generation patterns
- **2025-03-26** [Goodbye core types](https://go.dev/blog/coretypes) - Simplifying the type system in Go 1.25
- **2026-03-24** [Type Construction and Cycle Detection](https://go.dev/blog/type-construction-and-cycle-detection) - How the compiler handles recursive types
- **2024-02-22** [Robust generic functions on slices](https://go.dev/blog/generic-slice-functions) - Writing slice utilities that don't surprise you

## Blog: Generics

The multi-year journey from "maybe someday" to `func Map[T, U any]`.

- **2019-07-31** [Why Generics?](https://go.dev/blog/why-generics) - The problem statement, by Ian Lance Taylor
- **2020-06-16** [The Next Step for Generics](https://go.dev/blog/generics-next-step) - Updated design with type parameters
- **2021-01-12** [A Proposal for Adding Generics to Go](https://go.dev/blog/generics-proposal) - The accepted proposal
- **2021-12-14** [Go 1.18 Beta 1 is available, with generics](https://go.dev/blog/go1.18beta1) - It's happening
- **2022-03-22** [An Introduction To Generics](https://go.dev/blog/intro-generics) - Practical guide for the rest of us
- **2022-04-12** [When To Use Generics](https://go.dev/blog/when-generics) - And more importantly, when not to
- **2023-09-26** [Deconstructing Type Parameters](https://go.dev/blog/deconstructing-type-parameters) - Advanced patterns
- **2023-10-09** [Everything You Always Wanted to Know About Type Inference](https://go.dev/blog/type-inference) - How the compiler figures out your types
- **2023-02-17** [All your comparable types](https://go.dev/blog/comparable) - The nuances of `comparable`
- **2025-07-07** [Generic interfaces](https://go.dev/blog/generic-interfaces) - Combining generics with interfaces

## Blog: Concurrency

Go's bread and butter.

- **2010-09-23** [Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts) - Timeouts with channels
- **2014-03-13** [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines) - Building data pipelines
- **2014-07-29** [Go Concurrency Patterns: Context](https://go.dev/blog/context) - The `context` package explained
- **2013-05-23** [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency) - Talk by Sameer Ajmani
- **2021-02-24** [Contexts and structs](https://go.dev/blog/context-and-structs) - Where should context go?
- **2010-07-13** [Share Memory By Communicating](https://go.dev/blog/codelab-share) - Don't communicate by sharing memory
- **2013-06-26** [Introducing the Go Race Detector](https://go.dev/blog/race-detector) - Finding data races automatically
- **2022-09-26** [Go runtime: 4 years later](https://go.dev/blog/go119runtime) - How the scheduler evolved

## Blog: Performance & Internals

For when you need to understand what's happening underneath.

- **2011-06-24** [Profiling Go Programs](https://go.dev/blog/pprof) - CPU and memory profiling with pprof
- **2023-02-08** [Profile-guided optimization preview](https://go.dev/blog/pgo-preview) - Feeding runtime data back to the compiler
- **2023-09-05** [Profile-guided optimization in Go 1.21](https://go.dev/blog/pgo) - PGO goes GA
- **2025-10-29** [The Green Tea Garbage Collector](https://go.dev/blog/greenteagc) - Latest GC improvements
- **2015-08-31** [Go GC: Prioritizing low latency and simplicity](https://go.dev/blog/go15gc) - The Go 1.5 GC revolution
- **2018-07-12** [Getting to Go: The Journey of Go's GC](https://go.dev/blog/ismmkeynote) - GC history from Rick Hudson
- **2026-02-27** [Allocating on the Stack](https://go.dev/blog/allocation-optimizations) - Escape analysis and stack allocation
- **2025-02-26** [Faster Go maps with Swiss Tables](https://go.dev/blog/swisstable) - New map internals in Go 1.24
- **2016-08-18** [Smaller Go 1.7 binaries](https://go.dev/blog/go1.7-binary-size) - Binary size optimizations
- **2024-03-14** [More powerful Go execution traces](https://go.dev/blog/execution-traces-2024) - Runtime tracing improvements
- **2023-08-28** [Perfectly Reproducible, Verified Go Toolchains](https://go.dev/blog/rebuild) - Supply chain integrity
- **2025-08-20** [Container-aware GOMAXPROCS](https://go.dev/blog/container-aware-gomaxprocs) - Respecting container CPU limits
- **2025-09-26** [Flight Recorder in Go 1.25](https://go.dev/blog/flight-recorder) - Always-on diagnostics
- **2026-03-10** [//go:fix inline and the source-level inliner](https://go.dev/blog/inliner) - Source-level optimization

## Blog: Error Handling

- **2011-07-12** [Error handling and Go](https://go.dev/blog/error-handling-and-go) - The basics of Go error handling
- **2015-01-12** [Errors are values](https://go.dev/blog/errors-are-values) - A shift in perspective
- **2019-10-17** [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) - `errors.Is`, `errors.As`, and wrapping
- **2025-06-03** [On syntactic support for error handling](https://go.dev/blog/error-syntax) - The 2025 discussion on `?` and alternatives

## Blog: Testing

- **2013-12-02** [The cover story](https://go.dev/blog/cover) - Code coverage in Go
- **2016-10-03** [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests) - Table-driven tests done right
- **2023-03-08** [Code coverage for Go integration tests](https://go.dev/blog/integration-test-coverage) - Coverage beyond unit tests
- **2021-06-03** [Fuzzing is Beta Ready](https://go.dev/blog/fuzz-beta) - Introduction to Go's fuzzer
- **2025-04-02** [More predictable benchmarking with testing.B.Loop](https://go.dev/blog/testing-b-loop) - Better benchmark patterns
- **2025-02-19** [Testing concurrent code with testing/synctest](https://go.dev/blog/synctest) - Deterministic concurrency testing
- **2025-08-26** [Testing Time (and other asynchronicities)](https://go.dev/blog/testing-time) - Dealing with time in tests
- **2023-12-12** [Finding unreachable functions with deadcode](https://go.dev/blog/deadcode) - Static analysis for dead code

## Blog: Security

- **2025-05-19** [Go Cryptography Security Audit](https://go.dev/blog/tob-crypto-audit) - Third-party audit results
- **2025-07-15** [The FIPS 140-3 Go Cryptographic Module](https://go.dev/blog/fips140) - FIPS compliance in Go
- **2022-09-06** [Vulnerability Management for Go](https://go.dev/blog/vuln) - The vulnerability ecosystem
- **2023-07-13** [Govulncheck v1.0.0 is released!](https://go.dev/blog/govulncheck) - Dependency vulnerability scanning
- **2022-03-31** [How Go Mitigates Supply Chain Attacks](https://go.dev/blog/supply-chain) - Security design decisions
- **2021-01-19** [Command PATH security in Go](https://go.dev/blog/path-security) - Avoiding command injection
- **2021-09-15** [Automatic cipher suite ordering in crypto/tls](https://go.dev/blog/tls-cipher-suites) - TLS best practices
- **2025-03-12** [Traversal-resistant file APIs](https://go.dev/blog/osroot) - Safe file path handling
- **2024-05-02** [Secure Randomness in Go 1.22](https://go.dev/blog/chacha8rand) - ChaCha8 replaces the old PRNG
- **2025-03-06** [From unique to cleanups and weak](https://go.dev/blog/cleanups-and-weak) - New low-level efficiency tools

## Blog: Tooling

- **2026-02-17** [Using go fix to modernize Go code](https://go.dev/blog/gofix) - Automated code modernization
- **2019-03-21** [Debugging what you deploy in Go 1.12](https://go.dev/blog/debug-opt) - Better debug info in optimized builds
- **2013-12-12** [Inside the Go Playground](https://go.dev/blog/playground) - How play.golang.org works
- **2023-07-31** [Experimenting with project templates](https://go.dev/blog/gonew) - `gonew` for scaffolding
- **2021-02-01** [Gopls on by default in VS Code](https://go.dev/blog/gopls-vscode-go) - Language server improvements
- **2023-09-08** [Scaling gopls for the growing Go ecosystem](https://go.dev/blog/gopls-scalability) - Making the language server fast
- **2020-06-09** [The VS Code Go extension joins the Go project](https://go.dev/blog/vscode-go) - Official VS Code support
- **2023-08-22** [Structured Logging with slog](https://go.dev/blog/slog) - The new standard logging library
- **2024-09-03** [Go Telemetry](https://go.dev/blog/gotelemetry) - Opt-in usage analytics for Go tools
- **2022-01-14** [Two New Tutorials for 1.18](https://go.dev/blog/tutorials-go1.18) - Learning generics and fuzzing
- **2011-03-31** [Godoc: documenting Go code](https://go.dev/blog/godoc) - Documentation tooling

## Blog: Standard Library Deep Dives

- **2011-01-25** [JSON and Go](https://go.dev/blog/json) - Encoding and decoding JSON
- **2025-09-09** [A new experimental Go API for JSON](https://go.dev/blog/jsonv2-exp) - The v2 JSON API
- **2011-03-24** [Gobs of data](https://go.dev/blog/gob) - Go's binary serialization format
- **2020-03-02** [A new Go API for Protocol Buffers](https://go.dev/blog/protobuf-apiv2) - The protobuf v2 API
- **2024-12-16** [Go Protobuf: The new Opaque API](https://go.dev/blog/protobuf-opaque) - Next-gen protobuf performance
- **2011-09-21** [The Go image package](https://go.dev/blog/image) - Working with images
- **2011-09-29** [The Go image/draw package](https://go.dev/blog/image-draw) - Compositing images
- **2017-03-24** [HTTP/2 Server Push](https://go.dev/blog/h2push) - Server push in net/http
- **2024-02-13** [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements) - Pattern matching in the standard mux
- **2011-03-17** [C? Go? Cgo!](https://go.dev/blog/cgo) - Calling C code from Go
- **2024-08-27** [New unique package](https://go.dev/blog/unique) - Value interning for efficiency
- **2024-05-01** [Evolving the Standard Library with math/rand/v2](https://go.dev/blog/randv2) - Modernizing random number generation
- **2024-08-20** [Range Over Function Types](https://go.dev/blog/range-functions) - Iterator protocol in Go 1.23
- **2024-09-12** [Building LLM-powered applications in Go](https://go.dev/blog/llmpowered) - Go for AI applications

## Blog: WebAssembly

- **2023-09-13** [WASI support in Go](https://go.dev/blog/wasi) - Running Go in WebAssembly runtimes
- **2025-02-13** [Extensible Wasm Applications with Go](https://go.dev/blog/wasmexport) - Exporting Go functions to Wasm

## Blog: Modules & Versioning

The evolution of Go's dependency management.

- **2019-03-19** [Using Go Modules](https://go.dev/blog/using-go-modules) - Getting started with modules
- **2019-08-21** [Migrating to Go Modules](https://go.dev/blog/migrating-to-go-modules) - Moving from GOPATH
- **2019-09-26** [Publishing Go Modules](https://go.dev/blog/publishing-go-modules) - Making modules available
- **2019-11-07** [Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules) - Major version handling
- **2020-07-07** [Keeping Your Modules Compatible](https://go.dev/blog/module-compatibility) - API evolution guidelines
- **2018-12-19** [Go Modules in 2019](https://go.dev/blog/modules2019) - State of the module ecosystem
- **2019-08-29** [Module Mirror and Checksum Database Launched](https://go.dev/blog/module-mirror-launch) - proxy.golang.org and sum.golang.org
- **2021-02-18** [New module changes in Go 1.16](https://go.dev/blog/go116-module-changes) - Modules on by default
- **2019-08-01** [Experiment, Simplify, Ship](https://go.dev/blog/experiment) - The design process behind modules
- **2018-03-26** [A Proposal for Package Versioning in Go](https://go.dev/blog/versioning-proposal) - The original vgo proposal
- **2023-08-14** [Backward Compatibility, Go 1.21, and Go 2](https://go.dev/blog/compat) - The compatibility promise
- **2023-08-14** [Forward Compatibility and Toolchain Management](https://go.dev/blog/toolchain) - Automatic toolchain downloads

## Blog: Community & Ecosystem

- **2015-07-08** [Go, Open Source, Community](https://go.dev/blog/open-source) - Go's open-source philosophy
- **2019-11-13** [Go.dev: a new hub for Go developers](https://go.dev/blog/go.dev) - Launch of go.dev
- **2017-07-13** [Toward Go 2](https://go.dev/blog/toward-go2) - Russ Cox on the future of Go
- **2018-08-28** [Go 2 Draft Designs](https://go.dev/blog/go2draft) - Error handling, generics, error values drafts
- **2018-11-29** [Go 2, here we come!](https://go.dev/blog/go2-here-we-come) - Process for Go 2 proposals
- **2019-06-26** [Next steps toward Go 2](https://go.dev/blog/go2-next-steps) - Evaluating proposals
- **2020-03-25** [Go, the Go Community, and the Pandemic](https://go.dev/blog/pandemic) - 2020 and the community
- **2014-03-24** [The Go Gopher](https://go.dev/blog/gopher) - The story behind the mascot
- **2013-07-18** [The first Go program](https://go.dev/blog/first-go-program) - Where it all started
- **2018-01-22** [Hello, 中国!](https://go.dev/blog/hello-china) - Go documentation in Chinese
- **2020-12-17** [Go on ARM and Beyond](https://go.dev/blog/ports) - Platform support story
- **2014-09-26** [Deploying Go servers with Docker](https://go.dev/blog/docker) - Containerizing Go apps
- **2018-10-09** [Compile-time Dependency Injection With Go Cloud's Wire](https://go.dev/blog/wire) - DI without reflection
- **2018-07-24** [Portable Cloud Programming with Go Cloud](https://go.dev/blog/go-cloud) - Cloud abstraction libraries
- **2021-06-23** [The Go Collective on Stack Overflow](https://go.dev/blog/stackoverflow) - Community Q&A

## Blog: Release Announcements

Every major Go release, from the beginning.

| Version | Date | Post |
|---------|------|------|
| Go 1.26 | 2026-02 | [Go 1.26 is released](https://go.dev/blog/go1.26) |
| Go 1.25 | 2025-08 | [Go 1.25 is released](https://go.dev/blog/go1.25) |
| Go 1.24 | 2025-02 | [Go 1.24 is released!](https://go.dev/blog/go1.24) |
| Go 1.23 | 2024-08 | [Go 1.23 is released](https://go.dev/blog/go1.23) |
| Go 1.22 | 2024-02 | [Go 1.22 is released!](https://go.dev/blog/go1.22) |
| Go 1.21 | 2023-08 | [Go 1.21 is released!](https://go.dev/blog/go1.21) |
| Go 1.20 | 2023-02 | [Go 1.20 is released!](https://go.dev/blog/go1.20) |
| Go 1.19 | 2022-08 | [Go 1.19 is released!](https://go.dev/blog/go1.19) |
| Go 1.18 | 2022-03 | [Go 1.18 is released!](https://go.dev/blog/go1.18) |
| Go 1.17 | 2021-08 | [Go 1.17 is released](https://go.dev/blog/go1.17) |
| Go 1.16 | 2021-02 | [Go 1.16 is released](https://go.dev/blog/go1.16) |
| Go 1.15 | 2020-08 | [Go 1.15 is released](https://go.dev/blog/go1.15) |
| Go 1.14 | 2020-02 | [Go 1.14 is released](https://go.dev/blog/go1.14) |
| Go 1.13 | 2019-09 | [Go 1.13 is released](https://go.dev/blog/go1.13) |
| Go 1.12 | 2019-02 | [Go 1.12 is released](https://go.dev/blog/go1.12) |
| Go 1.11 | 2018-08 | [Go 1.11 is released](https://go.dev/blog/go1.11) |
| Go 1.10 | 2018-02 | [Go 1.10 is released](https://go.dev/blog/go1.10) |
| Go 1.9 | 2017-08 | [Go 1.9 is released](https://go.dev/blog/go1.9) |
| Go 1.8 | 2017-02 | [Go 1.8 is released](https://go.dev/blog/go1.8) |
| Go 1.7 | 2016-08 | [Go 1.7 is released](https://go.dev/blog/go1.7) |
| Go 1.6 | 2016-02 | [Go 1.6 is released](https://go.dev/blog/go1.6) |
| Go 1.5 | 2015-08 | [Go 1.5 is released](https://go.dev/blog/go1.5) |
| Go 1.4 | 2014-12 | [Go 1.4 is released](https://go.dev/blog/go1.4) |
| Go 1.3 | 2014-06 | [Go 1.3 is released](https://go.dev/blog/go1.3) |
| Go 1.2 | 2013-12 | [Go 1.2 is released](https://go.dev/blog/go1.2) |
| Go 1.1 | 2013-05 | [Go 1.1 is released](https://go.dev/blog/go1.1) |
| Go 1.0 | 2012-03 | [Go version 1 is released](https://go.dev/blog/go1) |

Detailed release notes for each version are at [go.dev/doc/devel/release](https://go.dev/doc/devel/release).

## Blog: History & Milestones

The anniversary posts are genuinely fun reads.

- **2025-11-14** [Go's Sweet 16](https://go.dev/blog/16years)
- **2024-11-11** [Go Turns 15](https://go.dev/blog/15years)
- **2023-11-10** [Fourteen Years of Go](https://go.dev/blog/14years)
- **2022-11-10** [Thirteen Years of Go](https://go.dev/blog/13years)
- **2021-11-10** [Twelve Years of Go](https://go.dev/blog/12years)
- **2020-11-10** [Eleven Years of Go](https://go.dev/blog/11years)
- **2019-11-08** [Go Turns 10](https://go.dev/blog/10years)
- **2018-11-10** [Nine years of Go](https://go.dev/blog/9years)
- **2017-11-10** [Eight years of Go](https://go.dev/blog/8years)
- **2016-11-10** [Seven years of Go](https://go.dev/blog/7years)
- **2015-11-10** [Six years of Go](https://go.dev/blog/6years)
- **2014-11-10** [Half a decade with Go](https://go.dev/blog/5years)
- **2013-11-10** [Four years of Go](https://go.dev/blog/4years)
- **2012-11-10** [Go turns three](https://go.dev/blog/3years)
- **2011-11-10** [Go turns two](https://go.dev/blog/2years)
- **2010-11-10** [Go: one year ago today](https://go.dev/blog/1year)

## Blog: Developer Surveys

What Go developers think, want, and struggle with.

- **2026-01-21** [Results from the 2025 Go Developer Survey](https://go.dev/blog/survey2025)
- **2024-12-20** [Go Developer Survey 2024 H2 Results](https://go.dev/blog/survey2024-h2-results)
- **2024-4-09** [Go Developer Survey 2024 H1 Results](https://go.dev/blog/survey2024-h1-results)
- **2023-12-05** [Go Developer Survey 2023 H2 Results](https://go.dev/blog/survey2023-h2-results)
- **2023-05-11** [Go Developer Survey 2023 Q1 Results](https://go.dev/blog/survey2023-q1-results)
- **2022-09-08** [Go Developer Survey 2022 Q2 Results](https://go.dev/blog/survey2022-q2-results)
- **2022-04-19** [Go Developer Survey 2021 Results](https://go.dev/blog/survey2021-results)
- **2021-03-09** [Go Developer Survey 2020 Results](https://go.dev/blog/survey2020-results)
- **2020-04-20** [Go Developer Survey 2019 Results](https://go.dev/blog/survey2019-results)
- **2019-03-28** [Go 2018 Survey Results](https://go.dev/blog/survey2018-results)
- **2018-02-26** [Go 2017 Survey Results](https://go.dev/blog/survey2017-results)
- **2017-03-06** [Go 2016 Survey Results](https://go.dev/blog/survey2016-results)

## Case Studies

Real companies, real production systems, real problems solved with Go.

### Cloud & Infrastructure
- [Google: Chrome](https://go.dev/solutions/google/chrome) - Go in Chrome infrastructure
- [Google: Core Data Solutions](https://go.dev/solutions/google/coredata) - Data processing at scale
- [Google: Firebase](https://go.dev/solutions/google/firebase) - Backend services for Firebase
- [Google: Site Reliability](https://go.dev/solutions/google/sitereliability) - SRE tooling in Go
- [Cloudflare](https://go.dev/solutions/cloudflare) - Edge computing and DNS services
- [CockroachDB](https://go.dev/solutions/cockroachlabs) - Distributed SQL database
- [Dropbox](https://go.dev/solutions/dropbox) - Open sourcing Go libraries

### Fintech & Payments
- [American Express](https://go.dev/solutions/americanexpress) - Payments and rewards
- [PayPal](https://go.dev/solutions/paypal) - Modernizing and scaling payment systems
- [Capital One](https://go.dev/solutions/capital-one) - Banking infrastructure
- [MercadoLibre](https://go.dev/solutions/mercadolibre) - Latin America's largest e-commerce
- [Curve](https://go.dev/solutions/curve) - Fintech card platform

### Consumer & Social
- [X (Twitter)](https://go.dev/solutions/x) - 5 billion sessions a day
- [Netflix](https://go.dev/solutions/netflix) - Streaming infrastructure
- [Twitch](https://go.dev/solutions/twitch) - Live streaming platform
- [Uber](https://go.dev/solutions/uber) - GPU-powered analytics engine
- [Bitly](https://go.dev/solutions/bitly) - URL shortening at scale
- [Monzo](https://go.dev/solutions/monzo) - Digital banking

### Enterprise & E-commerce
- [Microsoft](https://go.dev/solutions/microsoft) - Azure and internal services
- [Salesforce](https://go.dev/solutions/salesforce) - CRM platform services
- [Allegro](https://go.dev/solutions/allegro) - E-commerce marketplace
- [Trivago](https://go.dev/solutions/trivago) - Hotel search platform
- [SIXT](https://go.dev/solutions/sixt) - Car rental technology
- [Stream](https://go.dev/solutions/stream) - Activity feeds and chat API

### Gaming
- [Riot Games](https://go.dev/solutions/riotgames) - Game infrastructure
- [Wildlife Studios](https://go.dev/solutions/wildlifestudios) - Mobile game backends

### Other
- [ByteDance](https://go.dev/solutions/bytedance) - TikTok parent company
- [Facebook / Meta](https://go.dev/solutions/facebook) - Internal Go usage
- [Armut](https://go.dev/solutions/armut) - Local services marketplace

### Use Cases Overview
- [Why Go](https://go.dev/solutions/) - General overview
- [Go for Cloud & Network Services](https://go.dev/solutions/cloud) - Cloud-native development
- [Go for Web Development](https://go.dev/solutions/webdev) - Web services and APIs
- [Go for Command-line Interfaces](https://go.dev/solutions/clis) - CLI tooling
- [Go for DevOps & SRE](https://go.dev/solutions/devops) - Operations and reliability

## Documentation

### Database Access
- [Accessing relational databases](https://go.dev/doc/database/) - Overview and getting started
- [Opening a database handle](https://go.dev/doc/database/open-handle)
- [Executing SQL statements](https://go.dev/doc/database/change-data)
- [Querying for data](https://go.dev/doc/database/querying)
- [Using prepared statements](https://go.dev/doc/database/prepared-statements)
- [Executing transactions](https://go.dev/doc/database/execute-transactions)
- [Cancelling in-progress operations](https://go.dev/doc/database/cancel-operations)
- [Managing connections](https://go.dev/doc/database/manage-connections)
- [Avoiding SQL injection risk](https://go.dev/doc/database/sql-injection)

### Security
- [Security Best Practices](https://go.dev/doc/security/best-practices)
- [FIPS 140-3 Compliance](https://go.dev/doc/security/fips140)
- [Go Security Policy](https://go.dev/doc/security/policy)
- [Go Fuzzing](https://go.dev/doc/security/fuzz/)
- [Go Vulnerability Management](https://go.dev/doc/security/vuln/)
- [Go Vulnerability Database](https://go.dev/doc/security/vuln/database)
- [Go CNA Policy](https://go.dev/doc/security/vuln/cna)

### Build & Tools
- [Coverage profiling for integration tests](https://go.dev/doc/build-cover)
- [Profile-guided optimization](https://go.dev/doc/pgo)
- [Go Toolchains](https://go.dev/doc/toolchain)
- [Go Telemetry](https://go.dev/doc/telemetry)
- [Gopls settings](https://go.dev/gopls/doc/settings)

## A Tour of Go

The interactive tour at [go.dev/tour](https://go.dev/tour/). You can do it all from your browser, no install needed.

### 1. [Welcome](https://go.dev/tour/welcome/1)
- [Hello, World](https://go.dev/tour/welcome/1) - Your first Go program
- [Go local](https://go.dev/tour/welcome/3) - Installing Go on your machine

### 2. [Basics](https://go.dev/tour/basics/1) - Packages, variables, and functions
- [Packages](https://go.dev/tour/basics/1), [Imports](https://go.dev/tour/basics/2), [Exported names](https://go.dev/tour/basics/3)
- [Functions](https://go.dev/tour/basics/4), [Multiple results](https://go.dev/tour/basics/6), [Named return values](https://go.dev/tour/basics/7)
- [Variables](https://go.dev/tour/basics/8), [Short variable declarations](https://go.dev/tour/basics/10)
- [Basic types](https://go.dev/tour/basics/11), [Zero values](https://go.dev/tour/basics/12), [Type conversions](https://go.dev/tour/basics/13), [Type inference](https://go.dev/tour/basics/14)
- [Constants](https://go.dev/tour/basics/15), [Numeric constants](https://go.dev/tour/basics/16)

### 3. [Flow control](https://go.dev/tour/flowcontrol/1) - for, if, switch, defer
- [For](https://go.dev/tour/flowcontrol/1), [For is Go's "while"](https://go.dev/tour/flowcontrol/3), [Forever](https://go.dev/tour/flowcontrol/4)
- [If](https://go.dev/tour/flowcontrol/5), [If with short statement](https://go.dev/tour/flowcontrol/6), [If and else](https://go.dev/tour/flowcontrol/7)
- [Exercise: Loops and Functions](https://go.dev/tour/flowcontrol/8)
- [Switch](https://go.dev/tour/flowcontrol/9), [Switch with no condition](https://go.dev/tour/flowcontrol/11)
- [Defer](https://go.dev/tour/flowcontrol/12), [Stacking defers](https://go.dev/tour/flowcontrol/13)

### 4. [More types](https://go.dev/tour/moretypes/1) - Structs, slices, and maps
- [Pointers](https://go.dev/tour/moretypes/1), [Structs](https://go.dev/tour/moretypes/2), [Struct fields](https://go.dev/tour/moretypes/3), [Struct literals](https://go.dev/tour/moretypes/5)
- [Arrays](https://go.dev/tour/moretypes/6), [Slices](https://go.dev/tour/moretypes/7), [Slice literals](https://go.dev/tour/moretypes/9), [Slice length and capacity](https://go.dev/tour/moretypes/11), [Making slices](https://go.dev/tour/moretypes/13), [Appending to a slice](https://go.dev/tour/moretypes/15)
- [Exercise: Slices](https://go.dev/tour/moretypes/18)
- [Maps](https://go.dev/tour/moretypes/19), [Map literals](https://go.dev/tour/moretypes/20), [Mutating maps](https://go.dev/tour/moretypes/22)
- [Exercise: Maps](https://go.dev/tour/moretypes/23)
- [Function values](https://go.dev/tour/moretypes/24), [Function closures](https://go.dev/tour/moretypes/25)
- [Exercise: Fibonacci closure](https://go.dev/tour/moretypes/26)

### 5. [Methods and interfaces](https://go.dev/tour/methods/1)
- [Methods](https://go.dev/tour/methods/1), [Pointer receivers](https://go.dev/tour/methods/4), [Choosing a value or pointer receiver](https://go.dev/tour/methods/8)
- [Interfaces](https://go.dev/tour/methods/9), [Interfaces are implemented implicitly](https://go.dev/tour/methods/10), [Interface values](https://go.dev/tour/methods/11), [The empty interface](https://go.dev/tour/methods/14)
- [Type assertions](https://go.dev/tour/methods/15), [Type switches](https://go.dev/tour/methods/16)
- [Stringers](https://go.dev/tour/methods/17), [Exercise: Stringers](https://go.dev/tour/methods/18)
- [Errors](https://go.dev/tour/methods/19), [Exercise: Errors](https://go.dev/tour/methods/20)
- [Readers](https://go.dev/tour/methods/21), [Exercise: Readers](https://go.dev/tour/methods/22), [Exercise: rot13Reader](https://go.dev/tour/methods/23)
- [Images](https://go.dev/tour/methods/24), [Exercise: Images](https://go.dev/tour/methods/25)

### 6. [Generics](https://go.dev/tour/generics/1)
- [Type parameters](https://go.dev/tour/generics/1)
- [Generic types](https://go.dev/tour/generics/2)

### 7. [Concurrency](https://go.dev/tour/concurrency/1)
- [Goroutines](https://go.dev/tour/concurrency/1), [Channels](https://go.dev/tour/concurrency/2), [Buffered channels](https://go.dev/tour/concurrency/3)
- [Range and close](https://go.dev/tour/concurrency/4), [Select](https://go.dev/tour/concurrency/5), [Default selection](https://go.dev/tour/concurrency/6)
- [Exercise: Equivalent Binary Trees](https://go.dev/tour/concurrency/7)
- [sync.Mutex](https://go.dev/tour/concurrency/9)
- [Exercise: Web Crawler](https://go.dev/tour/concurrency/10)

## Talks & Presentations

Slide decks and presentations from Go team members at conferences and events. All available at [go.dev/talks](https://go.dev/talks/).

### 2010

- [Go](https://go.dev/talks/2010/go_talk-20100112.html) — January 12, 2010
- [Go, Networked](https://go.dev/talks/2010/go_talk-20100121.html) — January 21, 2010
- [Go Tech Talk](https://go.dev/talks/2010/go_talk-20100323.html) — March 23, 2010

### 2011

- [Lexical Scanning in Go](https://go.dev/talks/2011/lex.slide) — Rob Pike on writing a lexer by hand

### 2012

- [Concurrency is not Parallelism](https://go.dev/talks/2012/waza.slide) — Rob Pike's classic on the distinction
- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide) — Foundational patterns with goroutines and channels
- [Go at Google](https://go.dev/talks/2012/splash.slide) — Language design in the service of software engineering
- [The Path to Go 1](https://go.dev/talks/2012/go1.slide) — How Go reached its first stable release
- [Go: a simple programming environment](https://go.dev/talks/2012/simple.slide) — Simplicity as a feature
- [Go: code that grows with grace](https://go.dev/talks/2012/chat.slide) — Building a chat system live
- [Get started with Go](https://go.dev/talks/2012/tutorial.slide) — Tutorial walkthrough
- [10 things you (probably) don't know about Go](https://go.dev/talks/2012/10things.slide) — Fun surprises
- [Go for C programmers](https://go.dev/talks/2012/goforc.slide) — Bridging the gap from C
- [Go and the Zen of Python](https://go.dev/talks/2012/zen.slide) — Go through Python's lens
- [Go docs](https://go.dev/talks/2012/go-docs.slide) — Documentation tooling
- [Inside the "present" tool](https://go.dev/talks/2012/insidepresent.slide) — How these slides are built

### 2013

- [Advanced Go Concurrency Patterns](https://go.dev/talks/2013/advconc.slide) — Beyond the basics: pipelines, cancellation, and more
- [Twelve Go Best Practices](https://go.dev/talks/2013/bestpractices.slide) — Practical idioms from the Go team
- [Go, for Distributed Systems](https://go.dev/talks/2013/distsys.slide) — Building distributed services
- [Go for Pythonistas](https://go.dev/talks/2013/go4python.slide) — Go from a Python developer's perspective
- [Go Language for Ops and Site Reliability Engineering](https://go.dev/talks/2013/go-sreops.slide) — Go in production operations
- [What's new in Go 1.1](https://go.dev/talks/2013/go1.1.slide) — Release highlights
- [High Performance Apps with Go on App Engine](https://go.dev/talks/2013/highperf.slide) — Squeezing performance on GAE
- [dl.google.com: Powered by Go](https://go.dev/talks/2013/oscon-dl.slide) — Real-world Go at Google scale

### 2014

- [Go for gophers](https://go.dev/talks/2014/go4gophers.slide) — Deep dive for experienced Go users
- [Go for Javaneros](https://go.dev/talks/2014/go4java.slide) — Go from a Java perspective
- [Go, from C to Go](https://go.dev/talks/2014/c2go.slide) — Translating the Go compiler from C to Go
- [Go: Easy to Read, Hard to Compile](https://go.dev/talks/2014/compiling.slide) — Compiler internals
- [Go: 90% Perfect, 100% of the time](https://go.dev/talks/2014/gocon-tokyo.slide) — GoCon Tokyo keynote
- [Hello, Gophers!](https://go.dev/talks/2014/hellogophers.slide) — Welcome talk
- [What's in a name?](https://go.dev/talks/2014/names.slide) — Naming conventions in Go
- [Organizing Go code](https://go.dev/talks/2014/organizeio.slide) — Package layout and project structure
- [Inside the Go playground](https://go.dev/talks/2014/playground.slide) — How the playground works
- [When in Go, do as Gophers do](https://go.dev/talks/2014/readability.slide) — Idiomatic Go code
- [Cancellation, Context, and Plumbing](https://go.dev/talks/2014/gotham-context.slide) — Early context patterns
- [A Taste of Go](https://go.dev/talks/2014/taste.slide) — Quick introduction
- [Testing Techniques](https://go.dev/talks/2014/testing.slide) — How to test Go code effectively
- [Gophers With Hammers](https://go.dev/talks/2014/hammers.slide) — Code generation and tooling
- [Static analysis tools](https://go.dev/talks/2014/static-analysis.slide) — vet, lint, and friends
- [The Research Problems of Implementing Go](https://go.dev/talks/2014/research.slide) — Academic challenges
- [More Research Problems of Implementing Go](https://go.dev/talks/2014/research2.slide) — Part two
- [Toward Go 1.3](https://go.dev/talks/2014/go1.3.slide) — Upcoming release preview
- [The State of Go](https://go.dev/talks/2014/state-of-go.slide) — Where the project stands
- [The State of the Gopher](https://go.dev/talks/2014/state-of-the-gopher.slide) — Community update
- [Go on Android](https://go.dev/talks/2014/droidcon.slide) — Mobile development with Go
- [Camlistore: Android, ARM, App Engine, anywhere](https://go.dev/talks/2014/camlistore.slide) — Brad Fitzpatrick's personal storage system

### 2015

- [Simplicity is Complicated](https://go.dev/talks/2015/simplicity-is-complicated.slide) — Rob Pike on why simple is hard
- [The Evolution of Go](https://go.dev/talks/2015/gophercon-goevolution.slide) — How Go's design evolved
- [How Go was made](https://go.dev/talks/2015/how-go-was-made.slide) — Origin story
- [Go in Go](https://go.dev/talks/2015/gogo.slide) — The self-hosting compiler
- [Go for Java Programmers](https://go.dev/talks/2015/go-for-java-programmers.slide) — Migration guide
- [Go for C++ developers](https://go.dev/talks/2015/go4cpp.slide) — Comparing paradigms
- [Go on Mobile](https://go.dev/talks/2015/gophercon-go-on-mobile.slide) — GopherCon mobile talk
- [gRPC Go](https://go.dev/talks/2015/gotham-grpc.slide) — Introduction to gRPC in Go
- [JSON, interfaces, and go generate](https://go.dev/talks/2015/json.slide) — Working with JSON effectively
- [Go Dynamic Tools](https://go.dev/talks/2015/dynamic-tools.slide) — Race detector, profiler, and more
- [The Cultural Evolution of gofmt](https://go.dev/talks/2015/gofmt-en.slide) — How formatting shaped Go culture
- [Stupid Gopher Tricks](https://go.dev/talks/2015/tricks.slide) — Clever (ab)uses of Go
- [Keeping up with the Gophers](https://go.dev/talks/2015/keeping-up.slide) — Following Go's growth
- [The State of Go](https://go.dev/talks/2015/state-of-go.slide) — February update
- [The State of Go](https://go.dev/talks/2015/state-of-go-may.slide) — May update

### 2016

- [The Design of the Go Assembler](https://go.dev/talks/2016/asm.slide) — Rob Pike on Go's assembler
- [Program your next server in Go](https://go.dev/talks/2016/applicative.slide) — Server development
- [Stacks of Tokens](https://go.dev/talks/2016/token.slide) — A look at Go's tokenizer
- [The State of Go](https://go.dev/talks/2016/state-of-go.slide) — Update

### 2017

- [The State of Go](https://go.dev/talks/2017/state-of-go.slide) — February update
- [The State of Go](https://go.dev/talks/2017/state-of-go-may.slide) — May update

### 2019

- [Playground v3](https://go.dev/talks/2019/playground-v3/playground-v3.slide) — The redesigned Go Playground

