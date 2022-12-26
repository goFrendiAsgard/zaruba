<!--startTocHeader-->
[🏠](../README.md)
# 🔧 Utilities
<!--endTocHeader-->

Zaruba provide several utilities to deal with environment, files, tasks, and JSON/YAML.

<!--startCode-->
```bash
zaruba
```
 
<details>
<summary>Output</summary>
 
```````
MMM"""AMV                               *MM              db      
M'   AMV                                 MM             ;MM:     
'   AMV    ,6"Yb.  '7Mb,od8 '7MM  '7MM   MM,dMMb.      ,V^MM.    
   AMV    8)   MM    MM' "'   MM    MM   MM    'Mb    ,M  'MM    
  AMV   ,  ,pm9MM    MM       MM    MM   MM     M8    AbmmmqMA   
 AMV   ,M 8M   MM    MM       MM    MM   MM.   ,M9   A'     VML  
AMVmmmmMM 'Moo9^Yo..JMML.     'Mbod"YML. P^YbmdP'  .AMA.   .AMMA.
--.. .- .-. ..- -... .-    .--. .-.. . .- ... .    ... - .- .-. - 
                                      Task runner and CLI utility
v0.9.0-alpha-2-00cddc3987ef0dec6d5e534993a25901bcd9e3f0

Usage:
  zaruba [command]

Available Commands:
  advertisement Advertisement related utilities
  base64        Base64 manipulation utilities
  completion    Generate the autocompletion script for the specified shell
  env           Environment manipulation utilities
  file          File manipulation utilities
  generate      Generate a directory based on sourceTemplate and jsonMapReplacement
  help          Help about any command
  install       Install third party tools
  json          Json manipulation utilities
  lines         JsonStrList manipulation utilities
  list          JsonList manipulation utilities
  map           JsonMap manipulation utilities
  num           Number manipulation utilities
  path          Path manipulation utilities
  please        Run Task(s)
  project       Project manipulation utilities
  serve         Serve static website
  str           String manipulation utilities
  task          Task manipulation utilities
  version       Show current version
  yaml          Yaml manipulation utilities

Flags:
  -h, --help   help for zaruba

Use "zaruba [command] --help" for more information about a command.
```````
</details>
<!--endCode-->


# JSON Data Structures

Several complex data structures are not supported by `bash` out of the box.

Zaruba can help you to deal with `map`, `list`, and `lines`.

## Map

A `map` is a valid JSON object.

The following is a valid `map`:

```json
{  
  "name": "Stark Industry",
  "public": false,
  "valuation": 10000000,
  "tags": [
    "tech",
    "military"
  ],
  "employees": [
    {  
      "name": "Jon",   
      "salary": 56000,   
      "married": false  
    },
    {  
      "name": "Arya",   
      "salary": 56000,   
      "married": false  
    }
  ]
}  
```

To see how `map` utitilites works, please visit [map utilities documentation](map/README.md). 

## List

A `list` is a valid JSON array.

The following is a valid `list`:

```json
[
  "Stark Industry",
  false,
  10000000,
  [
    "tech",
    "military"
  ],
  {
    "employees":  [
      {  
        "name": "Jon",   
        "salary": 56000,   
        "married": false  
      },
      {  
        "name": "Arya",   
        "salary": 56000,   
        "married": false  
      }
    ]
  }
]
```

To see how `list` utitilites works, please visit [list utilities documentation](list/README.md). 

## Lines

A `lines` is a valid JSON array containing only string elements.

The following is a valid `lines`:

```json
[
  "import sys",
  "",
  "def create_tag(tag_name: str, content: str) -> str:",
  "    return '<{tag_name}>{content}</{tag_name}>'.format(tag_name=tag_name, content=content)",
  "",
  "if __name__ == '__main__':",
  "  print create_tag(sys.argv[1], sys.argv[2])"
]
```

To see how `lines` utitilites works, please visit [lines utilities documentation](lines/README.md). 

<!--startTocSubtopic-->
# Sub-topics
* [Base64](base64/README.md)
  * [Decode](base64/decode.md)
  * [Encode](base64/encode.md)
* [🏝️ Env](env/README.md)
  * [Get](env/get.md)
  * [Print](env/print.md)
  * [Read](env/read.md)
* [🧩 Install](install.md)
* [📁 File](file/README.md)
  * [Copy](file/copy.md)
  * [GetLine](file/getline.md)
  * [GetLineIndex](file/getlineindex.md)
  * [InsertBefore](file/insertbefore.md)
  * [InsertAfter](file/insertafter.md)
  * [List](file/list.md)
  * [Read](file/read.md)
  * [Replace](file/replace.md)
  * [Submatch](file/submatch.md)
  * [Walk](file/walk.md)
* [🚈 Lines](lines/README.md)
  * [GetIndex](lines/getindex.md)
  * [InsertAfter](lines/insertafter.md)
  * [InsertBefore](lines/insertbefore.md)
  * [Read](lines/read.md)
  * [Replace](lines/replace.md)
  * [Submatch](lines/submatch.md)
  * [Write](lines/write.md)
* [🧺 List](list/README.md)
  * [Append](list/append.md)
  * [Contain](list/contain.md)
  * [Get](list/get.md)
  * [Join](list/join.md)
  * [Length](list/length.md)
  * [Merge](list/merge.md)
  * [RangeIndex](list/rangeindex.md)
  * [Set](list/set.md)
  * [Validate](list/validate.md)
* [🗺️ Map](map/README.md)
  * [Get](map/get.md)
  * [GetKeys](map/getkeys.md)
  * [Merge](map/merge.md)
  * [RangeKey](map/rangekey.md)
  * [Set](map/set.md)
  * [ToStringMap](map/tostringmap.md)
  * [ToVariedStringMap](map/tovariedstringmap.md)
  * [TransformKey](map/transformkey.md)
  * [Validate](map/validate.md)
* [🔢 Num](num/README.md)
  * [Add](num/add.md)
  * [Divide](num/divide.md)
  * [Multiply](num/multiply.md)
  * [Range](num/range.md)
  * [Subtract](num/subtract.md)
  * [ValidateFloat](num/validatefloat.md)
  * [ValidateInt](num/validateint.md)
* [🛣️ Path](path/README.md)
  * [GetAppName](path/getappname.md)
  * [GetEnv](path/getenv.md)
  * [GetPortConfig](path/getportconfig.md)
  * [GetRelativePath](path/getrelativepath.md)
* [🏗️ Project](project/README.md)
  * [AddTask](project/addtask.md)
  * [Include](project/include.md)
  * [SetValue](project/setvalue.md)
  * [ShowLog](project/showlog.md)
  * [SyncEnv](project/syncenv.md)
  * [SyncEnvFiles](project/syncenvfiles.md)
* [🔠 Str](str/README.md)
  * [AddPrefix](str/addprefix.md)
  * [CurrentTime](str/currenttime.md)
  * [DecodeBase64](str/decodebase64.md)
  * [DoubleQuote](str/doublequote.md)
  * [EncodeBase64](str/encodebase64.md)
  * [FullIndent](str/fullindent.md)
  * [GetIndentation](str/getindentation.md)
  * [Indent](str/indent.md)
  * [NewName](str/newname.md)
  * [NewUuid](str/newuuid.md)
  * [PadLeft](str/padleft.md)
  * [PadRight](str/padright.md)
  * [Repeat](str/repeat.md)
  * [Replace](str/replace.md)
  * [SingleQuote](str/singlequote.md)
  * [Split](str/split.md)
  * [Submatch](str/submatch.md)
  * [ToCamel](str/tocamel.md)
  * [ToKebab](str/tokebab.md)
  * [ToLower](str/tolower.md)
  * [ToPascal](str/topascal.md)
  * [ToPlural](str/toplural.md)
  * [ToSingular](str/tosingular.md)
  * [ToSnake](str/tosnake.md)
  * [ToUpper](str/toupper.md)
  * [ToUpperSnake](str/touppersnake.md)
* [🔨 Task](task/README.md)
  * [AddDependencies](task/adddependencies.md)
  * [AddParents](task/addparents.md)
  * [GetIcon](task/geticon.md)
  * [IsExist](task/isexist.md)
  * [SetConfig](task/setconfig.md)
  * [SetEnv](task/setenv.md)
  * [SyncEnv](task/syncenv.md)
* [🍠 Yaml](yaml/README.md)
  * [Print](yaml/print.md)
  * [Read](yaml/read.md)
* [🍠 Json](json/README.md)
  * [Print](json/print.md)
<!--endTocSubtopic-->