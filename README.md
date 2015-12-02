# goj2
CLI tool for Jinja2 template rendering based on [pongo2](https://github.com/flosch/pongo2).
Loosely inspired by [j2](https://bitbucket.org/cavanaug/j2).

DISCLAIMER: first app with Golang, so please excuse the copy-pasta or big mistakes.
I wanted a lightweight j2 tool to render my environment variables on my Ansible (jinja2) templates.

## Usage
```
usage: goj2 [inputfile]
```
It will output a rendered Jinja2-compatible template using the environment variables (*I'm looking at you, docker*).

e.g.
```
$ cat > /tmp/mytemplate.txt.j2 << EOF
Hello {{ WORLD }}
EOF
$ ./goj2 /tmp/mytemplate.txt.j2
Hello

$ WORLD='World!' ./goj2 /tmp/mytemplate.txt.j2
Hello World!

```
