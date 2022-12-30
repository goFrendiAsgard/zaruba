<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Utils](README.md)
# generate
<!--endTocHeader-->

```
Generate a directory based on sourceTemplate and jsonMapReplacement

Usage:
  zaruba generate <sourceTemplatePath> <destinationPath> [jsonMapReplacement] [flags]

Examples:

> ls template
index.html  start.sh

> cat template/index.html
<h1>title</h1>
content

> cat template/start.sh
python -m http.server port

> zaruba generate template web '{"port":3000, "title":"MyWeb", "content":"<p>Hello world!</p>"}'

> ls web
index.html  start.sh

> cat web/index.html
<h1>MyWeb</h1>
<p>Hello world!</p>

> cat web/start.sh
python -m http.server 3000


Flags:
  -h, --help   help for generate

```

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->