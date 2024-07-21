# Q:= What are design Patterns?
Ans:= Design patterns are reusable solutioins to common problems in software. the term reusable solution is important because the it's the best practise of writing a piece of code where it can be reused in multiple similar scenarios.

# Q:= What are the types of design patterns?
Ans:= There are three types of design patterns:
1. Creational Desing Patterns
2. Structural Design Patterns
3. Behavioral Design Patterns

# Q:= What is Creational Design Patterns?
Ans:= Creational design patterns are design patterns that deal with object creation mechanisms, trying to create objects in a manner suitable to the situation. The basic form of object creation could result in design problems or added complexity to the design. Creational design patterns solve this problem by somehow controlling this object creation process or simply managing it in a better way.

Singletons, Factories, Builders, Prototypes, and Abstract Factories are some of the design patterns that fall under this category.

# Q:= What is a Singleton Design Pattern?
Ans:= Ensures a class has only one instance and provides a global point of access to it. It is a creational pattern as it deals with object creation mechanisms.

# Q:= What is a Factory Design Pattern?
Ans:= Factory pattern is one of the most used design patterns in Java. This type of design pattern comes under creational pattern as this pattern provides one of the best ways to create an object. It defines an interface for creating an object, but let subclasses decide which class to instantiate. The Factory method lets a class defer instantiation to subclasses.

# Q:= What is a Builder Design Pattern?
Ans:= The Builder pattern is a design pattern designed to provide a flexible solution to various object creation problems in object-oriented programming. The intent of the Builder design pattern is to separate the construction of a complex object from its representation. It is one of the Gang of Four design patterns. 

# Q:= What is a Prototype Design Pattern?
Ans:= The prototype pattern is a creational design pattern in software development. It is used when the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects. This pattern is used to: avoid subclasses of an object creator in the client application, like the abstract factory pattern does.

# Q:= What is an Abstract Factory Design Pattern?
Ans:= The abstract factory pattern provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes. In normal usage, the client software creates a concrete implementation of the abstract factory and then uses the generic interface of the factory to create the concrete objects that are part of the theme.

# Q:= What is Structural Design Patterns?
Ans:= Structural design patterns are design patterns that ease the design by identifying a simple way to realize relationships between entities. These patterns increase the flexibility of the structure in which they are used. They are different from creational design patterns as they deal with how objects are composed to form larger structures. 

Adapter, Bridge, Composite, Decorator, Facade, Flyweight, and Proxy are some of the design patterns that fall under this category.

# Q:= What is a Adapter Design Pattern?
Ans:= The adapter pattern is a design pattern that allows the interface of an existing class to be used as another interface. It is often used to make existing classes work with others without modifying their source code. The adapter pattern is also referred to as the wrapper pattern.

# Q:= What kind of problems does the Adapter Design Pattern solve?
Ans:= The adapter pattern solves the problem of integrating different interfaces that are not compatible with each other. It allows objects with incompatible interfaces to collaborate. 

# Q:= What is a interface in Adapter Design Pattern?
Ans:= An interface is a contract that defines the signature of a functionality. An interface is a reference type in Java, similar to a class, that can contain only constants, method signatures, default methods, static methods, and nested types. Method bodies exist only for default methods and static methods. Interfaces cannot be instantiated—they can only be implemented by classes or extended by other interfaces.

# Q:= What is a Bridge Design Pattern?
Ans:= The bridge pattern is a design pattern used in software engineering that is meant to "decouple an abstraction from its implementation so that the two can vary independently", introduced by the Gang of Four. The bridge uses encapsulation, aggregation, and can use inheritance to separate responsibilities into different classes.

# Q:= What is a Composite Design Pattern?
Ans:= The composite pattern is a partitioning design pattern. The composite pattern describes that a group of objects is to be treated in the same way as a single instance of an object. The intent of a composite is to "compose" objects into tree structures to represent part-whole hierarchies. Implementing the composite pattern lets clients treat individual objects and compositions uniformly.

# Q:= What is a Decorator Design Pattern?
Ans:= The decorator pattern is a design pattern that allows behavior to be added to individual objects, dynamically, without affecting the behavior of other objects from the same class. The decorator pattern is often useful for adhering to the Single Responsibility Principle, as it allows functionality to be divided between classes with unique areas of concern.

# Q:= What is a Facade Design Pattern?
Ans:= A facade is an object that provides a simplified interface to a larger body of code, such as a class library. A facade can: make a software library easier to use, understand, and test, since the facade has convenient methods for common tasks; make the library more readable, for the same reason.

# Q:= What is a Flyweight Design Pattern?
Ans:= The flyweight pattern is a software design pattern. A flyweight is an object that minimizes memory usage by sharing as much as possible with related objects; it is a way to use objects in large numbers when a simple repeated representation would use an unacceptable amount of memory.

# Q:= What is a Proxy Design Pattern?
Ans:= A proxy, in its most general form, is a class functioning as an interface to something else. The proxy could interface to anything: a network connection, a large object in memory, a file, or some other resource that is expensive or impossible to duplicate. In short, a proxy is a wrapper or agent object that is being called by the client to access the real serving object behind the scenes.

# Q:= What is Behavioral Design Patterns?
Ans:= Behavioral design patterns are design patterns that identify common communication patterns between objects and realize these patterns. By doing so, these patterns increase flexibility in carrying out this communication.

Chain of Responsibility, Command, Interpreter, Iterator, Mediator, Memento, Observer, State, Strategy, Template Method, Visitor are some of the design patterns that fall under this category.

# Q:= What is a observer Design Pattern?
Ans:= The observer pattern is a software design pattern in which an object, named the subject, maintains a list of its dependents, called observers, and notifies them of any state changes, usually by calling one of their methods. It is mainly used to implement distributed event handling systems.