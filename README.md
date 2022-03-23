# About GO

- Go was designed by Robert Griesemer, Rob PIke and Ken Thompson.
####
- Go was designed by Google in 2007.But it was released publicly in 2012.
####
- Go is an open source programming language,
  - Which means it is not proprietary and publicly accessible for free that can be modified or built in such a manner that is open to the public.
####
- Go is a multipurpose programming language.
  - It can be used for web development,data science, cloud computing,artificial intelligence and microcontroller programming,devops,commanding tools ,robotics and more.
####
- Go is simple and easy to understand language
  - It is well known for its simplicity,readability and efficiency
####
- Go is an cross-platform programming language
  - which means a programming language that can be developed and executed in more than one platform. Platform includes both hardware(amd64,386) architecture and operating system(Linux,windows,Android and iOS )
  - [Note: Go achieves this by supporting cross compilation
  To do cross compilation we need to specify GOOS and GOARCH during compilation before build]
####
- Go is statically typed and compiled programming language
  - Source code is compiled to native language. Go compiler will not just work on compiling code successfully but also ensures on type conversions and compatibility it catches entire classes of bugs..
####
- Go provides unit testing feature by itself
  - A simple mechanism to write your unit testing parallel to your code which enhances  your code coverage by your own tests.
####
- Go has a rich and powerful standard Library.
  - The robust standard library is distributed as packages.
  - There are packages to support io support,cryptography,testing
####
- Go gives High Performance
  - It is designed for automation on a large scale which makes it easy to write high performance applications.
####
- Go uses garbage Collection.
  - Go uses garbage collection which takes care of memory management automatically.which in turn helps to write concurrent programs easily.
####

- Popular tools which use Go are Kubernetes,Docker,Prometheus,go-ethereum

## Go Paradigms
- Object-oriented programming is the programming paradigm which is based on the concept of objects, which contains data in the form of fields and code in the form of procedures.
- Objects in the OOP paradigm have attributes and methods.

### Paradigms of GO:
#### 1) Object-Oriented Language
- Go is not a pure object-oriented programming language.
- Go allows types and methods
- Like oo there is interfaces but no class support
- Methods are designed for all types including primitive data type
  - Object in go:
  Structs can store data and achieve methods using receivers.Struct is light weighted
  - Methods in Go:
  Methods are functions of a particular type. They have a receiver clause that mandates which type they are operated on.
  - Structs instead of classes:
  Go does not support class but struct is supported.Methods are allowed in struct.
  - Composition instead of Inheritance:
  Since go does not support classes there is no concept of inheritance.
  - Inheritance will takes place via struct embedding using composition
  - Anonymous field is composition where struct is embedded
  - In composition base struct can be embedded into child struct and methods on the base structs can be directly called on the child struct.
  - The new() function instead of constructors:
  - Polymorphism
  Polymorphism is achieved through interfaces.
  - Variable of type interface can hold any values which implements interface.


#### 2) Imperative
- It is an imperative language
- It has a loop,statements and selections.

#### 3) Concurrent Programming
- Go has native support for concurrent operations

#### 4) Encapsulation;
- In go encapsulation is achieved by capitalising fields,methods,struts and functions then it will become public. If it is public it is accessible at package level.

#### 5) Interfaces:
- Interfaces are types that have multiple methods
- Object that implements all the methods of the interface will implicitly implement the interface.

### Idiomatic feature:
#### 1) Orthogonality
- Orthogonality means pieces that are independent of each other.
- Changes to one part of types/programs/packages will have minimal effect to other parts

#### 2) Simplicity
- GO has no classes, methods can be added to any types
- It has no inheritance
- Interfaces are implicitly implemented
- GO does not have generics

#### 3) Readability

### Values:

##### 1) A good package starts with a good name
##### 2) Plan for failure, not success
##### 3) Return early rather than nesting deeply
##### 4) If you think it’s slow, prove it with a benchmark
##### 5) Before you launch a goroutine, know when it will stop
##### 6) Write tests to lock in the behaviour of your package’s API

