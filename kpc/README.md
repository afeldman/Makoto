# KPC

Karel Package Configuration. this is a structure to setup youp project.

```toml
kpc_version = "0.2.0"
name = "project"
version = "0.0.1-alpha"
description = "test project"
url = "testproject.org"
source = "some git repo"
issue = "no issues known"
keywords = [ "demo", "project", "karel", "fanuc" ]
prefix = "installpath"
srcdir = "where is the search for the compiler?"
typedir = "where are your typefiles"
includedir = "where are your header files"
constdir = "where are your const definition files"
formdir = "where are your karel form files"
dictdir = "where are the dicionaries"
main = "main.kl"
dicts = [ "en.utx", "de.utx" ]
forms = [ "main.ftx" ]
types = [ "matrix.t.kl", "vector.t.kl" ]
includes = [ "matrix.h.kl", "vector.h.kl" ]
consts = [ "pi.c.kl" ]

[[requirements]]
name = "math"
version = "1.0"

[[conflicts]]
name = "math"
versions = [ "0.2.0", "0.2.1" ]

[[authors]]
name = "tester"
email = "tester@example.com"

[[authors]]
name = "test"
email = "test@example.com"
```