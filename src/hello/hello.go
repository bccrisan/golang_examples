package main

import "fmt"
import "math"
import "time"

func separate_functions() {
    fmt.Println("-------------------------------------")
}

func hello_world() {
    fmt.Println("Powered by: https://gobyexample.com/ ")
}


func values() {
    // Strings, which can be added together with `+`.
    fmt.Println("go" + "lang")

    // Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleans, with boolean operators as you'd expect.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}


func variables() {
    // `var` declares 1 or more variables.
    var a = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "short"` in this case.
    f := "short"
    fmt.Println(f)
}

func constants(){
// Go supports _constants_ of character, string, boolean,
// and numeric values.


    // `const` declares a constant value.
    const s string = "constant"


    fmt.Println(s)

    // A `const` statement can appear anywhere a `var`
    // statement can.
    const n = 500000000

    // Constant expressions perform arithmetic with
    // arbitrary precision.
    const d = 3e20 / n
    fmt.Println(d)

    // A numeric constant has no type until it's given
    // one, such as by an explicit cast.
    fmt.Println(int64(d))

    // A number can be given a type by using it in a
    // context that requires one, such as a variable
    // assignment or function call. For example, here
    // `math.Sin` expects a `float64`.
    fmt.Println(math.Sin(n))
}


func test_for_function(){
    // The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // A classic initial/condition/after `for` loop.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` without a condition will loop repeatedly
    // until you `break` out of the loop or `return` from
    // the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // You can also `continue` to the next iteration of
    // the loop.
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}


func if_else_conditions(){
    // Here's a basic example.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // You can have an `if` statement without an else.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // A statement can precede conditionals; any variables
    // declared in this statement are available in all
    // branches.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.

}


func switch_statements(){
// _Switch statements_ express conditionals across many
// branches.

// Here's a basic `switch`.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // You can use commas to separate multiple expressions
    // in the same `case` statement. We use the optional
    // `default` case in this example as well.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // `switch` without an expression is an alternate way
    // to express if/else logic. Here we also show how the
    // `case` expressions can be non-constants.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // A type `switch` compares types instead of values.  You
    // can use this to discover the type of an interface
    // value.  In this example, the variable `t` will have the
    // type corresponding to its clause.
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}


func using_arrays(){
    // In Go, an _array_ is a numbered sequence of elements of a
    // specific length.

    // Here we create an array `a` that will hold exactly
    // 5 `int`s. The type of elements and length are both
    // part of the array's type. By default an array is
    // zero-valued, which for `int`s means `0`s.
    var a [5]int
    fmt.Println("emp:", a)

    // We can set a value at an index using the
    // `array[index] = value` syntax, and get a value with
    // `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // The builtin `len` returns the length of an array.
    fmt.Println("len:", len(a))

    // Use this syntax to declare and initialize an array
    // in one line.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Array types are one-dimensional, but you can
    // compose types to build multi-dimensional data
    // structures.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}


func play_with_slices(){
    // _Slices_ are a key data type in Go, giving a more
    // powerful interface to sequences than arrays.
        
    // Unlike arrays, slices are typed only by the
    // elements they contain (not the number of elements).
    // To create an empty slice with non-zero length, use
    // the builtin `make`. Here we make a slice of
    // `string`s of length `3` (initially zero-valued).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` returns the length of the slice as expected.
    fmt.Println("len:", len(s))

    // In addition to these basic operations, slices
    // support several more that make them richer than
    // arrays. One is the builtin `append`, which
    // returns a slice containing one or more new values.
    // Note that we need to accept a return value from
    // `append` as we may get a new slice value.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Slices can also be `copy`'d. Here we create an
    // empty slice `c` of the same length as `s` and copy
    // into `c` from `s`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Slices support a "slice" operator with the syntax
    // `slice[low:high]`. For example, this gets a slice
    // of the elements `s[2]`, `s[3]`, and `s[4]`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    // This slices up to (but excluding) `s[5]`.
    l = s[:5]
    fmt.Println("sl2:", l)

    // And this slices up from (and including) `s[2]`.
    l = s[2:]
    fmt.Println("sl3:", l)

    // We can declare and initialize a variable for slice
    // in a single line as well.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Slices can be composed into multi-dimensional data
    // structures. The length of the inner slices can
    // vary, unlike with multi-dimensional arrays.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}


func using_maps(){
    // _Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
    // (sometimes called _hashes_ or _dicts_ in other languages).

    // To create an empty map, use the builtin `make`:
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

    // Set key/value pairs using typical `name[key] = val`
    // syntax.
    m["k1"] = 7
    m["k2"] = 13

    // Printing a map with e.g. `fmt.Println` will show all of
    // its key/value pairs.
    fmt.Println("map:", m)

    // Get a value for a key with `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // The builtin `len` returns the number of key/value
    // pairs when called on a map.
    fmt.Println("len:", len(m))

    // The builtin `delete` removes key/value pairs from
    // a map.
    delete(m, "k2")
    fmt.Println("map:", m)

    // The optional second return value when getting a
    // value from a map indicates if the key was present
    // in the map. This can be used to disambiguate
    // between missing keys and keys with zero values
    // like `0` or `""`. Here we didn't need the value
    // itself, so we ignored it with the _blank identifier_
    // `_`.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // You can also declare and initialize a new map in
    // the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

func play_with_range(){
    // _range_ iterates over elements in a variety of data
    // structures. Let's see how to use `range` with some
    // of the data structures we've already learned.
    
    // Here we use `range` to sum the numbers in a slice.
    // Arrays work like this too.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` on arrays and slices provides both the
    // index and value for each entry. Above we didn't
    // need the index, so we ignored it with the
    // blank identifier `_`. Sometimes we actually want
    // the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` on strings iterates over Unicode code
    // points. The first value is the starting byte index
    // of the `rune` and the second the `rune` itself.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

    // Go requires explicit returns, i.e. it won't
    // automatically return the value of the last
    // expression.
    return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
    return a + b + c
}


func functions_test(){
    // Call a function just as you'd expect, with
    // `name(args)`.
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}

// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
    return 3, 7
}

func multiple_return(){
    // Here we use the 2 different return values from the
    // call with _multiple assignment_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // If you only want a subset of the returned values,
    // use the blank identifier `_`.
    _, c := vals()
    fmt.Println(c)
}


func main() {
//    hello_world()
//    separate_functions()
//    values()
//    variables()
//    constants()
//    test_for_function()
//    if_else_conditions()
//    switch_statements()
//    using_arrays()
//    play_with_slices()
//    using_maps()
//    play_with_range()
//    functions_test()
//    multiple_return()

}
