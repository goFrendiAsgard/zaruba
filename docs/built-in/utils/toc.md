<!--startTocHeader-->
[🏠](../../README.md) > [Built-in](../README.md) > [Utils](README.md)
# toc
<!--endTocHeader-->

```

Create/update documentation directory based on a TOC file.

A TOC file is markdown file representing "table of content" for your documentation.
A tag is written as <!--start[TagName] attribute="value" -->content<!--end[TagName]-->

There are several tagname available for a TOC file:
    - toc: This tag contains list of bulleted-items that will be rendered into documentation structure.
    - code: This tag has several attributes:
        - lang: code language (e.g., python, javascript, bash)
        - src: location of your source code file, relative to TOC file directory
        - cmd: Command to run your code.
      Your code tag will be rendered into a markdown.    

There are also additional tagName for your documentation files:
    - tocHeader: This tag will be filled with documentation header
    - tocSubtopic: This tag will be filled with documentation subtopics

Usage:
  zaruba toc <tocFileLocation> [flags]

Examples:

> cat README.md
# My Cool Project

This is a documentation for my cool project.

<!--startToc-->
- Getting Started
- Concepts
    - Model
    - View
    - Controller
<!--endToc-->

> zaruba toc README.md


Flags:
  -h, --help   help for toc

```

<!--startTocSubtopic-->
<!--endTocSubtopic-->