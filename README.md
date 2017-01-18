This is a test case for go install -pkgdir <dir> failures on go 1.7.3

To reproduce this:
1. git clone git@github.com:stevenh/golang-pkgdir-test.git
2. cd golang-pkgdir-test/cmd/test
3. go install -pkgdir pkgs
