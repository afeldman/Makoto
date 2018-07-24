# Makoto
This project utilizes a FANUC Karel package configuration structure and information.
The named after Makoto Mishimura the first japanese who developed a robot.


# structure
```YAML
name: "karel project name"
discription: "project description"
version: "Version number"
url: "url"
source: {url:"url to the source", type:"scm type (git for now)", version:"version id (master)"}
requirements: ["project name 1","project name 2"]
conflicts: "projekt name 1"
author: [{name:"mane of auther 1",email: "email address of author 1"}]
prefix: "installation prefix"
srcdir: "source directory"
typedir: ["type folder 1","type folder 2"]
constdir: ["const folder 2","const folder 2"]
includedir: ["header director 1"]
lib: "program file name"
types: "type file name"
consts: "constant file name"
includes: ["header name 1"]
```

# setting
The enviroment can set to

```bash
KPC_HOME=/usr/local/lib/karel
```
