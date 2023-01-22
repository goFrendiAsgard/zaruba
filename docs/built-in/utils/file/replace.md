<!--startTocHeader-->
[🏠](../../../README.md) > [Built-in](../../README.md) > [Utils](../README.md) > [file](README.md)
# replace
<!--endTocHeader-->

```
Replace string by jsonMapReplacement

Usage:
  zaruba file replace <strFileName> [{<jsonMapReplacement> | <key> <value>}] [flags]

Examples:

> echo 'Capital of country is city' > example1.txt
> zaruba str replace example1.txt '{"country": "Indonesia", "city": "Jakarta"}'
> cat example1.txt
Capital of Indonesia is Jakarta

> echo 'Capital of country is city' > example2.txt
> zaruba str replace example2.txt country Japan city Tokyo
> cat example2.txt
Capital of Japan is Tokyo

> echo "def add(a):" > example.py
> echo "    pass" >> example.py
> echo "" >> example.py
> echo "def minus(" >> example.py
> echo "    a" >> example.py
> echo "):" >> example.py
> echo "    pass" >> example.py
> echo "" >> example.py
> echo "class Something():" >> example.py
> echo "    def __init__(a):" >> example.py
> echo "        pass" >> example.py
> zaruba file replace example.py '(?U)(?m)(?s)def (.*)\((.*)([\n\t ]*)\):' 'def $1($2, b$3):'
> cat example.py
def add(a, b):
    pass

def minus(
    a, b
):
    pass

class Something():
    def __init__(a, b):
        pass


Flags:
  -h, --help   help for replace

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->