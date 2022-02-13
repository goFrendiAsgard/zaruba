<!--startTocHeader-->
[🏠](../../../README.md) > [🧠 Core Concepts](../../README.md) > [🏗️ Project](../README.md) > [Task](README.md)
# Define task dependencies
<!--endTocHeader-->


Some tasks might require several pre-requisites.


For example, a typescript developer needs to install npm packages and compile their typescript before running/testing their application.

To __start__ a typescript application, you need to do:

```bash
# prepare app
npm install
tsc
# start app
npm start
```

while to __test__ the application, you need to do:

```bash
# prepate app
npm install
tsc
# test app
npm test
```

You can make a zaruba script to execute those actions:

```yaml
tasks:

  prepareApp:
    extend: zrbRunShellScript
    configs:
      start: npm install && tsc

  startApp:
    extend: zrbStartApp
    dependencies:
      - prepareApp
    configs:
      ports: 3000
      start: npm start

  testApp:
    extend: zrbRunShellScript
    dependencies:
      - prepareApp
    configs:
      start: npm test
```

The scripts has some advantages:

* You can update `prepareApp` without touching `startApp` and `testApp`.
* In case of you have many dependencies, Zaruba will run the dependencies in parallel.

Let's modify the script a little bit to run `startApp` and `testApp` in parallel:


```yaml
tasks:

  prepareApp:
    extend: zrbRunShellScript
    configs:
      start: npm install && tsc

  startApp:
    extend: zrbStartApp
    dependencies:
      - prepareApp
    configs:
      ports: 3000
      start: npm start

  testApp:
    extend: zrbRunShellScript
    dependencies:
      - prepareApp
    configs:
      start: npm test

  startAndTestApp:
    dependencies:
      - startApp
      - testApp
```

Cool. Now, whenever you run `zaruba please startAndTestApp`, things will be executed in this order:

![](images/task-dependencies.png)

<!--startTocSubTopic-->
<!--endTocSubTopic-->
