# quiz

> [gophercises](https://github.com/gophercises/quiz)

[TOC]

## Part I

```bash
> ./quiz -h
Usage of ./quiz:
  -csv string
      a csv file in the format of "question,answer" (default "problems.csv")
```

It is very easy to implement.

## Part II

```diff
< Usage of ./quiz:
<  -csv string
<      a csv file in the format of "question,answer" (default "problems.csv")
---
> Usage of ./quiz:
>  -csv string
>      a csv file in the format of "question,answer" (default "problems.csv")
>  -limit int
>      the time limit for the quiz in seconds (default 30)
```

Use goroutines and channels to implement the timer.
