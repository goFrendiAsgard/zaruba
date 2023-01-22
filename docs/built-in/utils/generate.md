<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Utils](README.md)
# generate
<!--endTocHeader-->

```
Generate a directory based on sourceTemplate and jsonMapReplacement

Usage:
  zaruba generate <sourceTemplatePath> <destinationPath> [{<jsonMapReplacement> | <key> <value>}] [flags]

Examples:

> mkdir -p template

# template/index.html
> echo '<h1>title</h1>' > template/index.html
> echo 'content' >> template/index.html
# template/start.sh
> echo 'python -m http.server port' > template/start.sh

> zaruba generate template web1 '{"port":3000, "title":"MyWeb", "content":"<p>Hello world!</p>"}'

> ls web1
index.html  start.sh
# web1/index.html
> cat web1/index.html
<h1>MyWeb</h1>
<p>Hello world!</p>
# web1/start.html
> cat web1/start.sh
python -m http.server 3000

> zaruba generate template web2 port 8000 title MySecondWeb content '<p>Hello world!</p>'

> ls web2
index.html  start.sh
# web2/index.html
> cat web2/index.html
<h1>MySecondWeb</h1>
<p>Hello world!</p>
# web2/start.html
> cat web2/start.sh
python -m http.server 8000


Flags:
  -h, --help   help for generate

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->