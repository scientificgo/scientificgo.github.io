---
layout: default
---

**ScientificGo** is a collection of [Go](https://www.golang.org) packages for scientific computing.
Designed for maximal portability and high-performance, ScientificGo packages are written in _pure_ Go;
there are no hidden bindings to C or Fortran libraries and no architecture-dependent assembly implementations
or complex build processes.

# Installation

You can use `go get` to download and install ScientificGo packages on any platform. It couldn't be simpler. To install the package [`fft`]({{ site.url }}/fft), for example:

```
go get github.com/scientificgo.org/fft
```

# Packages

{% for repository in site.github.public_repositories %}
  {% if repository.name != 'scientificgo.github.io' %}
  * [{{ repository.name }}]({{ site.url }}/{{ repository.name }}): {{ repository.description }}
  {% endif %}
{% endfor %}

<p align="center">
<img src="gopher.png">
</p>

