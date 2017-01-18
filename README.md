This is a test case for go install -pkgdir <dir> failures on go 1.7.3

To reproduce this simply do the following:

```shell
git clone git@github.com:stevenh/golang-pkgdir-test.git
cd golang-pkgdir-test/cmd/test
go install -pkgdir pkgs
```
