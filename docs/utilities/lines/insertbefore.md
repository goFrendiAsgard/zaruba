<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🚈 Lines](README.md)
# InsertBefore
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba lines insertBefore --help
```
 
<details>
<summary>Output</summary>
 
```````
Insert new lines into a jsonStringList before a particular index.
The Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be 0.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to insert two 🍕 before 🍓.

--------------------------------------------------
Elements | Index  | Note
--------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert two🍕 before this
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍕","🍕","🍓","🍇"]

Usage:
  zaruba lines insertBefore <jsonStrList> <jsonStrListNewLines | strNewLine> [flags]

Examples:

> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '🍕'
["🍕","🍊","🍓","🍇"]

> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1
["🍊","🍕","🍕","🍓","🍇"]

> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕"]' \
  --index=-1
["🍊","🍓","🍕","🍇"]


Flags:
  -h, --help        help for insertBefore
  -i, --index int   desired pattern index
```````
</details>
<!--endCode-->

# Examples

<!--startTocSubTopic-->
<!--endTocSubTopic-->