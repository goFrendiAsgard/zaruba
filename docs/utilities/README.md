<!--startTocHeader-->
[🏠](../README.md)
# 🔧 Utilities
<!--endTocHeader-->

Zaruba provide several utilities to deal with environment, files, tasks, and JSON/YAML.

<!--startCode-->
```bash
zaruba
```

````
,,                      
MMM"""AMV                               *MM              db      
M'   AMV                                 MM             ;MM:     
'   AMV    ,6"Yb.  '7Mb,od8 '7MM  '7MM   MM,dMMb.      ,V^MM.    
   AMV    8)   MM    MM' "'   MM    MM   MM    'Mb    ,M  'MM    
  AMV   ,  ,pm9MM    MM       MM    MM   MM     M8    AbmmmqMA   
 AMV   ,M 8M   MM    MM       MM    MM   MM.   ,M9   A'     VML  
AMVmmmmMM 'Moo9^Yo..JMML.     'Mbod"YML. P^YbmdP'  .AMA.   .AMMA.
--.. .- .-. ..- -... .-    .--. .-.. . .- ... .    ... - .- .-. - 
                                    Task runner and CLI utilities
v0.9.0-alpha-2-88-gcf5f855d

Usage:
  zaruba [command]

Available Commands:
  advertisement Advertisement utilities
  completion    generate the autocompletion script for the specified shell
  env           Env utilities
  file          File utilities
  generate      Make something based on template
  help          Help about any command
  install       Install external tools
  lines         Lines manipulation utilities
  list          List manipulation utilities
  map           Map manipulation utilities
  num           Number manipulation utilities
  path          Path manipulation utilities
  please        Run Task(s)
  project       Project manipulation utilities
  serve         Serve static files in location at a specified port
  str           String manipulation utilities
  task          Task manipulation utilities
  version       Show current version
  yaml          YAML utilities

Flags:
  -h, --help   help for zaruba

Use "zaruba [command] --help" for more information about a command.
````
<!--endCode-->


# JSON Data Structures

There are several complex data structures that are not supported by `bash` out of the box.

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

<!--startTocSubTopic-->
# Sub-topics
* [Env](env/README.md)
  * [Get](env/get.md)
  * [Print](env/print.md)
  * [Read](env/read.md)
  * [Write](env/write.md)
* [Install](install.md)
* [File](file/README.md)
  * [Copy](file/copy.md)
  * [List](file/list.md)
  * [Walk](file/walk.md)
* [Lines](lines/README.md)
  * [Fill](lines/fill.md)
  * [GetIndex](lines/get-index.md)
  * [InsertAfter](lines/insert-after.md)
  * [InsertBefore](lines/insert-before.md)
  * [Read](lines/read.md)
  * [Replace](lines/replace.md)
  * [Submatch](lines/submatch.md)
  * [Write](lines/write.md)
* [List](list/README.md)
  * [Append](list/append.md)
  * [Contain](list/contain.md)
  * [Get](list/get.md)
  * [Join](list/join.md)
  * [Length](list/length.md)
  * [Merge](list/merge.md)
  * [RangeIndex](list/range-index.md)
  * [Set](list/set.md)
  * [Validate](list/validate.md)
* [Map](map/README.md)
  * [Get](map/get.md)
  * [GetKeys](map/get-keys.md)
  * [Merge](map/merge.md)
  * [RangeKey](map/range-key.md)
  * [Set](map/set.md)
  * [ToStringMap](map/to-string-map.md)
  * [ToVariedStringMap](map/to-varied-string-map.md)
  * [TransformKey](map/transform-key.md)
  * [Validate](map/validate.md)
* [Num](num/README.md)
  * [Range](num/range.md)
  * [ValidateFloat](num/validate-float.md)
  * [ValidateInt](num/validate-int.md)
* [Path](path/README.md)
  * [GetAppName](path/get-app-name.md)
  * [GetEnv](path/get-env.md)
  * [GetPortConfig](path/get-port-config.md)
  * [GetRelativePath](path/get-relative-path.md)
* [Project](project/README.md)
  * [AddTaskIfNotExist](project/add-task-if-not-exist.md)
  * [Include](project/include.md)
  * [SetValue](project/set-value.md)
  * [ShowLog](project/show-log.md)
  * [SyncEnv](project/sync-env.md)
  * [SyncEnvFiles](project/sync-env-files.md)
* [Str](str/README.md)
  * [AddPrefix](str/add-prefix.md)
  * [DoubleQuote](str/double-quote.md)
  * [FullIndent](str/full-indent.md)
  * [GetIndentation](str/get-indentation.md)
  * [Indent](str/indent.md)
  * [NewName](str/new-name.md)
  * [NewUuid](str/new-uuid.md)
  * [PadLeft](str/pad-left.md)
  * [PadRight](str/pad-right.md)
  * [Repeat](str/repeat.md)
  * [Replace](str/replace.md)
  * [SingleQuote](str/single-quote.md)
  * [Split](str/split.md)
  * [Submatch](str/submatch.md)
  * [ToCamel](str/to-camel.md)
  * [ToKebab](str/to-kebab.md)
  * [ToLower](str/to-lower.md)
  * [ToPascal](str/to-pascal.md)
  * [ToPlural](str/to-plural.md)
  * [ToSingular](str/to-singular.md)
  * [ToSnake](str/to-snake.md)
  * [ToUpper](str/to-upper.md)
  * [ToUpperSnake](str/to-upper-snake.md)
* [Task](task/README.md)
  * [AddDependency](task/add-dependency.md)
  * [AddParent](task/add-parent.md)
  * [GetIcon](task/get-icon.md)
  * [IsExist](task/is-exist.md)
  * [SetConfig](task/set-config.md)
  * [SetEnv](task/set-env.md)
  * [SyncEnv](task/sync-env.md)
* [Yaml](yaml/README.md)
  * [Print](yaml/print.md)
  * [Read](yaml/read.md)
  * [Write](yaml/write.md)
<!--endTocSubTopic-->