# Go-Expand
This tool generates query parameters based on given conditions (yaml file).

# How to use
expand.yaml(sample)
```
host: "localhost:8888"
path: "/sample"
query-parameters:
  - name: "name1"
    type: "list"
    null-able: true
    values:
      - "AAAAA"
      - "BBBBB"
      - "CCCCC"
  - name: "name2"
    type: "list"
    null-able: false
    values:
      - "DDDDD"
      - "EEEEE"
  - name: "name3"
    type: "range"
    null-able: false
    range-config:
      start: 0
      stop: 30
      step: 10
```

run
```
$ goexpand
http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=0
http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=10
http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=20
http://localhost:8888/sample?name1=AAAAA&name2=DDDDD&name3=30
http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=0
http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=10
http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=20
http://localhost:8888/sample?name1=AAAAA&name2=EEEEE&name3=30
http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=0
http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=10
http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=20
http://localhost:8888/sample?name1=BBBBB&name2=DDDDD&name3=30
http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=0
http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=10
http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=20
http://localhost:8888/sample?name1=BBBBB&name2=EEEEE&name3=30
http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=0
http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=10
http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=20
http://localhost:8888/sample?name1=CCCCC&name2=DDDDD&name3=30
http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=0
http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=10
http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=20
http://localhost:8888/sample?name1=CCCCC&name2=EEEEE&name3=30
```

# Todo
- [ ]  use "NullAble"
