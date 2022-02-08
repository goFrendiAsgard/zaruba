<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# Add Subrepo
<!--endTocHeader-->


You might want to put everything in a single git repository, and allow some part of your code to have it's own public repository.

Or you probably have several existing git repository and you want to merge them into single monorepo.

To do this you need to either use:

* [git submodule](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
* [git subrepo](https://github.com/ingydotnet/git-subrepo)
* [git subtree](https://www.atlassian.com/git/tutorials/git-subtree)

`Git submodule` is a bit confusing since adding a submodule in git, doesn't add the code of the submodule to the main repository. When you add a submodule, you only add information about the submodule that is added to the main repository.

`Git subrepo` need to be installed separately, thus not always available in every computer.

`Git subtree` is a contrib script and most likely available in your git client. Thus, under the hood Zaruba choose this approach.

There are several builtin tasks you can use to manage subrepo:

* [initSubrepos](../core-tasks/initSubrepos.md)
* [addSubrepo](../core-tasks/addSubrepo.md)
* [pullSubrepos](../core-tasks/pullSubrepos.md)
* [pushSubrepos](../core-tasks/pushSubrepos.md)


# Add Subrepo

Suppose you already in a zaruba project, and you want to add [git@github.com:state-alchemists/fibonacci-clock.git](https://github.com/state-alchemists/fibonacci-clock) as a subrepo named `fibo`, then you can do:

```bash
zaruba please addSubrepo subrepoUrl="https://github.com/state-alchemists/fibonacci-clock" subrepoPrefix="fibo" 
zaruba please initSubrepos 
zaruba please pullSubrepos 
```

After performing the task, you will see `fibo` directory in your project.

# Pull from subrepos

People might contribute to your subrepos. You want the changes in your subrepo is also reflected in your zaruba project. In that case you want to pull from subrepos.

To pull from your subrepos, you can invoke:

```
zaruba please pullSubrepos
```

# Push to subrepos

Sometime you need any changes in your project to be reflected in your subrepos. In that case, you can push to subrepos.

To push to your subrepos, you can invoke:

```
zaruba please pushSubrepos
```


<!--startTocSubtopic-->

<!--endTocSubtopic-->