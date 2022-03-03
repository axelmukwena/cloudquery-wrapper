# Cloudquery wrapper

A Go based wrapper to automate Cloudquery executions.


## Technical notes

---

### Embedding Go into Ruby
- https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf
- Parsing strings: https://github.com/ffi/ffi/issues/585#issuecomment-546549387
- https://gist.github.com/schweigert/385cd8e2267140674b6c4818d8f0c373
- https://guides.rubygems.org/specification-reference/#files

---

### How to initialize cloudquery in this project

```
$ cd config/initializers/cloudquery
$ go build -o cloudquery.so -buildmode=c-shared cloudquery.go

# Test the file
$ ruby cloudquery.rb  
```

---

### ffi Error
```
block in ffi_lib': Could not open library './cloudquery.so': dlopen(./cloudquery.so, 5): image not found (LoadError)
```
- https://stackoverflow.com/a/58833833/8050183

---

### How to parse json in Go
- https://www.sohamkamani.com/golang/json/

---

### Specifying AWS credentials
- https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials

---
