# DI

## 1.1 - Initial setup

* we have a straightforward codebase with global variables to ease object sharing
* unit tests have to pollute global scope to insert custom behavior

hint : https://github.com/nathanleclaire/blog/blob/master/content/post/interfaces-and-composition-for-effective-unit-testing-in-golang.markdown


`01-initial`


## 1.2. Pass shared objects as parameters

* we move objects from global scope and pass it to handler as parameter
* unit tests can now create their own isolated instances of objects


`02-attributes`


## 1.3. Handler as a struct

* HandlerFunc's can become messy and repetitive

`03-struct`

## 1.4. Objects as interface

* DB and log still have to be a specific type and have to be configured differently to allow us to test it
* they make it hard to do actual unit testing (testing only your unit of code)
* interfaces allow us to abstract that

`04-interfaces`

## 2. Context

## 2.1. About it

* carries a deadline, a cancellation signal, and other values across API boundaries
* in general:
 * informs a goroutine it is not needed anymore (and it should handle it as it pleases)
 * allows a process to broadcast its cancellation to other linked processes
* Context interface
```go
Deadline() (deadline time.Time, ok bool)
Done() context struct{}
Err() error
Value(key interface{}) interface{}
```
* Additional helper functions
 * `Background() Context` returns an empty context
 * `TODO()` returns an empty context but also implies that the correct context still has to be assigned

5.2. WithCancel

* allows one process to send cancellation signal to another
* to create it `ctx, cancelFn := context.WithContext(context.Background())`
* `<- ctx.Done()` blocks until `cancelFn` is called

## 2.2.1 naive

`05-cancel`

## 2.2.2 cautious

`06-cautious`

## 2.3. WithTimeout/Deadline

* Allows us to limit execution of goroutines / actions
* `context.Timeout` accepts `time.Duration`
* `context.Deadline` accepts `time.Time`
* both also return `cancel` function, which should always be called to ensure goroutines get cancelled

`07-timeout`

## 2.4. WithValue

* useful to cary around request specific values
* should only be used for smaller / scalar values
* can make code error prone as compiler can't catch possible issues
* you should create your own type for keys

`08-withValue`

## 3. Databases

* A database is an organized collection of data
* As a process application can hold / store data in memory
* When designed well, simple in-memory database can be easily replaced with a more powerful setup

```golang
package bucket

type (
    BucketReader interface{
        Get(int) (string, error)
        GetAll() ([]string, error)
    }

    BucketWriter interface{
        Append(string...) error
        Remove(int) error
    }

    BucketReadWriter interface{
        BucketReader
        BucketWriter
    }

    BucketStorage interface{
        func Create(string) (BucketReadWriter, error)
        func Remove(string) error
    }
)
```

```go
package kvMemory

import (
    "errors"
)

// ErrNotFound is returned when an unexisting key is requested
var ErrNotFound = errors.New("Key not found")

// KVStorage describes basic functions required for KV storage
type KVStorage interface{
    Add(string, string) error
    Get(string) (string, error)
    Remove(string) error
}

// KV hold all the data in memory
type KV struct{
    data map[string]string
}

// New returns an initialized instance of KV
func New() *KV {
    return &KV{map[string]string{}}
}

// Add writes a given value under the given key
func (kv *KV) Add(k, v string) error {
    kv.data[k] = v
    return nil
}

// Get returns value linked to given key. Returns ErrNotFound when key does not
// exist.
func (kv *KV) Get(k string) (string, error) {
    if v, ok := kv.data[k]; ok {
        return v, nil
    }

    return "", ErrNotFound
}

// Remove deletes the given key from the map. Returns ErrNotFound if key does
// not exist.
func (kv *KV) Remove(k string) error {
    if _, ok := kv.data[k]; !ok {
        return ErrNotFound
    }

    delete(kv.data, k)
    return nil
}
```

```go
package main

import "kvStorage"

func main() {
    kv := kvStorage.New()
    kv.Add("first", "value")
    kv.Add("second", "value")

    v, _ := kv.Get("second")
    fmt.Printf("Second value = %s", v)
}
```

## 3.1 Race conditions

