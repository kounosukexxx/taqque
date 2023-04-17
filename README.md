# taqque

## What is taqque?

This is a task manegement tool using queues with priority concept.

## Motivation

I want to complete tasks in FIFO order, so I created taqque to manage tasks.

## Installation

```shell script
go install github.com/kounosukexxx/taqque@latest
```

## How to use

### list tasks

```shell script
taqque
```

### push a task

```shell script
taqque push {title}
```

### pop a task

```shell script
taqque pop
```

### push a task with priority

You can't set negative priority.

Tasks are listed in descending order of priority.

By default, a task is pushed with 1 priority.

```shell script
taqque push {title} {priority}
```

### pop a task with priority

You can also pop a task specifing priority.

By default, a task with 1 priority is popped.

```shell script
taqque pop {priority}
```

## Demonstration

```shell script
taqque push assinmentA
+-------+----------+------------+
| INDEX | PRIORITY |   TITLE    |
+-------+----------+------------+
|     0 |     1.00 | assinmentA |
+-------+----------+------------+

taqque push assinmentB
+-------+----------+------------+
| INDEX | PRIORITY |   TITLE    |
+-------+----------+------------+
|     0 |     1.00 | assinmentA |
|     1 |     1.00 | assinmentB |
+-------+----------+------------+

taqque pop
+-------+----------+------------+
| INDEX | PRIORITY |   TITLE    |
+-------+----------+------------+
|     0 |     1.00 | assinmentB |
+-------+----------+------------+

taqque push high_priority_task 2
+-------+----------+--------------------+
| INDEX | PRIORITY |       TITLE        |
+-------+----------+--------------------+
|     0 |     2.00 | high_priority_task |
|     1 |     1.00 | assinmentB         |
+-------+----------+--------------------+

taqque pop
+-------+----------+--------------------+
| INDEX | PRIORITY |       TITLE        |
+-------+----------+--------------------+
|     0 |     2.00 | high_priority_task |
+-------+----------+--------------------+

 taqque pop 2
+-------+----------+-------+
| INDEX | PRIORITY | TITLE |
+-------+----------+-------+
+-------+----------+-------+
```

## Future improvement plan

- undo previous task
- make other queues (We can use only one single queue now)
- push and pop a task specifing index
