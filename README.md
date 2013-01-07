## Sylvester

Sylvester is a Go library that imagines applications that handle streams of data
as a graph. Graphs contain nodes and edges: nodes are responsible for computation
and connection with outside data sources and sinks, and edges are responsible for
connections between nodes.

### Sylvester is:

* An attempt to sketch a Go framework for network applications.
* An abstraction for a certain class of applications.
* A chance to explore Go more deeply and work on a framework.
* Named after <a href="http://en.wikipedia.org/wiki/James_Joseph_Sylvester">James Joseph Sylvester</a>, who coined the term "graph," and was a general badass.

<a href="http://github.com/etsy/statsd">Statsd</a> is the canonical application that
Sylvester aims to prove a clean framework for.

### Statsd as a Graph:

```
[External TCP] -->
inputNode -->
aggregatorNode -->
outputNode -->
[External UDP]
```

### What's up now

Check out <a href="https://github.com/mrb/sylvester/blob/master/examples/syncasync/main.go">This example application</a> for the state of things now.
