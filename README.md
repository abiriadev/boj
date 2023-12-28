# BOJ

Solutions of BOJ([Baekjoon Online Judge](https://www.acmicpc.net/)) in various languages.

## Structure

```
solved/
├── xxxx/
│  ├── lang1/
│  │  └── ...
│  └── lang2/
│     └── ...
└── xxxx/
   ├── lang1/
   │  └── ...
   └── lang2/
      └── ...
unsolved/
├── xxxx/
│  ├── lang1/
│  │  └── ...
│  └── lang2/
│     └── ...
└── xxxx/
   ├── lang1/
   │  └── ...
   └── lang2/
      └── ...
```

There are two top level directories:

* `solved/`: Solved problems.
* `unsolved/`: Unsolved problems.

If I solved problem A with language X, but failed to solve the same one with language Y, the former will go under `solved/`, and later will go under `unsolved/`.

Every problem in BOJ has its own unique integer number, so the directory containg solutions of the problem should have the problem number as its name.

Under the solution directory, each attempted language has its own sub directory.
