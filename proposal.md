# Proposal: Modern Go language dialect for configuration templating and configuration as code

Status: Idea
Author: Bartek Płotka [@bwplotka](https://bwplotka.dev)
Date: 23.04.2021

### Problem Statement

Before we discuss any solution let's talk about, what problem engineers have in infrastructure, backend area. This problem space is targeted
to developers, sysadmins, devops, ops, SRE etc.

#### Why configuration language?

The main reason why these days we have many configuration languages (jsonnet, cue, starlark, pullumi) or templating engines (helm's go templating, kustomize, ...) is that the more configuration we maintain, the more we need abstractions that allows to have simpler view on things that matters as well as, simpler tooling and ability to change parts of configuration easier across many instances.

Some languages like cue also aims for certain amount of type safety and validation, allowing to test configuration before we deploy it. Unfortunately there are things that those languages are missing:

* Different projects have different formats to configure themselves. It can be certain structure of args, flags, YAML, JSON, INI, TOML, protobuf or its own DSL.
* Within each project the schema that is parsed in those languages can very between commits (software versions).

While we some languages allows us to validate some of it (although in practice not enough e.g cue simplification for creating structs from go code - it works only if basic structs tag based marshalling is used), validation is only part of story. The first problem that hits users of the software we configure... we cannot remember all the different names and meanings of fields. We require a quick navigation techniques to the documentation and ideally the code that is used from such configuration knob.

With certain amount of such configuration we need a single language that can rule them all, allowing consistent and rich configuration techniques for autocompletion, autodiscovery of fields and their documentation, type safety.

Existing configuration languages have serious problems then:

* Most of the languages are focused mostly on JSON or YAML.
* Not enough focus on tooling to link our configuration to software we configure.
* Most of the languages are not type safe at all and does not consider any validation techniques, not mentioning auto-completion
* Yet new languages with different approaches, syntax or even programming paradigms. the argument saying that "sys admins/ops" should not learn programming in order to be able to use programming language" is overused to create new config languages with more steep learning curve than existing languages.

#### Can we use Go?

At the startup I worked for (Improbable, UK), we had a vision to use Go for our infrastructure and configuration resources. It solves many problems that other languages don't. Primarily:

* Go is extremely easy to learn and read.
  It's typed, yet compilation is ultrafast (by design), so it feels like script. Furthermore script command is provided: `go run`, so you don't see build artifact.
* Tooling is mature and incredibly useful. Since most of the infrastructure projects are written in Go you can navigate to their structs that are used for parsing yaml, toml etc and import them directly. You can navigate to the code to understand each flag, you can read documentation and see the structs changing over time. Testing, benchmarking are stable and extremely useful.
* Mature dependency tooling.
* For other languages other tooling can be used like OpenAPI or protobuf to generate Go the code with documentation.
* Since it's a Turing complete language you can create arbitrary templates and templating logic. You can expose this as library too for others to use.
* With proper encoding code you can generate ANY DSL, not only JSON or YAML.

At the end it worked well for some engineers, yet portion of engineers from different than Go backgrounds was finding it hard (You can play with open source project which was ported from that code and is available here: https://github.com/bwplotka/mimic). Let's go through some issues that people reported about Go:

* Using Go "feels" like programming not configuring. There are voices that configuration is never static so should be easy to move and shape.
* There is certain amount of boilerplate e.g error handling or type creations (e.g using protobuf).
* The language patterns like channels, go routines, IO, syscalls etc are problematic if overused in templating language. It is harder to read and
  might obfuscate templating which has to be simple.
* While dependency (Go Modules) mechanism works well it downloads extreme amount of unnecesary dependencies. Since we mostly use structs and small Marshalling methods from other Go projects, it's just not needed to fight with dependency collisions (aka dependency hell), and compilation error in unrelated code to the actual structure.
* Composition is hard due to type system (e.g merging structs)

### Goals

* Type safe
* Not a new language, use good, existing programming language for a start allowing small learning curve.
* Rich auto-completion, navigation and code highlighting tooling from day 1.
* Lightweight, yet familiar dependency management.
* Provider easier composition techniques.

### Proposal

Create language which uses Go syntax, yet limits the language and simplify it in order to allow better configuration experience.
To leverage tooling we want tie in to Go as much as possible.

#### Ideas

* [ ] Leverage compatibility with Go tooling
* [ ] Extension to Go modules, allowing to extract types downloading the linked structs and marshalers only.
* [ ] Less boilerplate on type definitions
* [ ] Limit language features?
* [ ] Compiler that creates Go from the language or Go fork?

### Open Questions:

* Should this language be compatible with Go? Can we import Go code? Is Go code compilable?
* Should we use .go or some other extension?
* Two modes strict and non strict?