---
layout: default
title: Home
---

[About](about)

**ScientificGo** is a collection of [Go](https://www.golang.org) packages for scientific computing.
Designed for maximal portability and high-performance, ScientificGo packages are written in _pure_ Go;
there are no hidden bindings to C or Fortran libraries and no platform-dependent assembly implementations
or build processes.

# Installation

You can use the `go get` command to install ScientificGo packages on any platform. It is *that* simple.

# Packages

{% for repository in site.github.public_repositories %}
  {% if repository.name != 'scientificgo.github.io' and repository.description %}
  * [{{ repository.name }}]({{ repository.name }}): {{ repository.description }}
  {% endif %}
{% endfor %}

<p align="center">
  <img src="gopher.png">
</p>
