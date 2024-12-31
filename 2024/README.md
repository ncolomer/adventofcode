# Year 2024

https://adventofcode.com/2024

```shell
# run test
go get
get test -v
```

```shell
# create new day files
(export FROM_DAY=06 TO_DAY=10 \
&& touch testdata/day${TO_DAY}_{sample,puzzle}.txt \
&& cp -r day${FROM_DAY} day${TO_DAY} \
&& cp day${FROM_DAY}_test.go day${TO_DAY}_test.go \
&& sed -i '' -E -e "s/([Dd]ay[ ]?)${FROM_DAY}/\1${TO_DAY}/g" day${TO_DAY}_test.go day${TO_DAY}/*)
```
