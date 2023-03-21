**Link to issue** : [devfile/api#1067](https://github.com/devfile/api/issues/1067)

**Notes**
- Related to https://github.com/redhat-developer/odo/issues/5694#issuecomment-1465778398

**How to reproduce the issue**

Using Go >= 1.19, the tests should not pass:

```bash  
go test -v ./...
```
