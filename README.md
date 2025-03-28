# bazel-query-guide-example

This repository contains a hypothetical Go microservices project structured to demonstrate Bazel's query capabilities. It's meant to serve as a companion to the Buildkite blog post [A Guide to Bazel Query](https://buildkite.com/blog/a-guide-to-bazel-query).

## Overview

The repository consists of:

- Three microservices (auth, user, payment)
- Shared libraries (logging, config, metrics)
- Database abstractions
- Environment-specific configuration

This structure allows you to explore how Bazel query can help you understand and navigate complex dependency relationships in a realistic project setup.

## Prerequisites

To use this repository, you'll need:

1. [Bazel](https://bazel.build/install) (we recommend installing with [Bazelisk](https://bazel.build/install/bazelisk)) 
1. [Go](https://golang.org/doc/install)
1. [GraphViz](https://graphviz.org/download/), if you want to generate graph visualizations

## Getting started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/bazel-query-guide-example.git
   cd bazel-query-guide-example
   ```

2. Build the project:
   ```bash
   bazel build //...
   ```

3. Run the tests:
   ```bash
   bazel test //...
   ```

## Using `bazel query`

The primary purpose of this repository is to demonstrate Bazel query commands. Here are some examples to try:

### Basic queries

List all targets in the workspace:
```bash
bazel query //...
```

Find all Go binaries:
```bash
bazel query "kind('go_binary', //...)"
```

### Dependency queries

Find what the payment service depends on:
```bash
bazel query "deps(//services/payment:payment)"
```

Find direct dependencies only (depth=1):
```bash
bazel query "deps(//services/payment:payment, 1)"
```

### Reverse dependency queries

Find what depends on the logging library:
```bash
bazel query "rdeps(//..., //lib/logging:logging)"
```

Find all binaries that depend on the logging library:
```bash
bazel query "kind('go_binary', rdeps(//..., //lib/logging:logging))"
```

### Path queries

Find a path between two targets:
```bash
bazel query "somepath(//services/payment:payment, @com_github_lib_pq//:go_default_library)"
```

### Combining queries

Find all tests that depend on the metrics library:
```bash
bazel query "kind('go_test', rdeps(//..., //lib/metrics:metrics))"
```

Find services that depend on both logging and config libraries:
```bash
bazel query "kind('go_binary', rdeps(//..., //lib/logging:logging)) intersect kind('go_binary', rdeps(//..., //lib/config:config))"
```

## Advanced usage with `aquery` and `cquery`

Try configuration-specific queries with cquery:
```bash
bazel cquery "deps(//services/payment:payment)" --define environment=production
```

Or explore the action graph with aquery:
```bash
bazel aquery //services/payment:payment
```

Generate a visual representation of dependencies:
```bash
bazel query 'deps(//services/payment/..., 2) ' --output graph --noimplicit_deps  | dot -Tpng -o graph.png
```

## Repository structure

```
.
├── lib
│   ├── config      # Configuration management
│   ├── logging     # Shared logging utilities
│   ├── metrics     # Metrics collection
│   └── tracing     # Environment-specific tracing
└── services
    ├── auth        # Authentication service
    ├── payment     # Payment processing service
    │   └── db      # Database interfaces
    └── user        # User management service
```

## Further reading

For more detailed explanations of Bazel query and how to leverage it in your own projects, refer to the accompanying blog post [A Guide to Bazel Query](https://buildkite.com/blog/a-guide-to-bazel-query) and the [official Bazel query docs](https://bazel.build/query/guide).

## License

This example is made available under the MIT License.