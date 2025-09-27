package main

import "fmt"

func main() {
    // 1. Array Declaration and Initialization
    
    // Declare an array of 5 integers (zero-initialized)
    var numbers [5]int
    fmt.Println("Empty array:", numbers) // [0 0 0 0 0]
    
    // Declare and initialize with values
    var fruits [3]string = [3]string{"apple", "banana", "orange"}
    fmt.Println("Fruits array:", fruits)
    
    // Short declaration with automatic size inference
    colors := [...]string{"red", "green", "blue", "yellow"}
    fmt.Println("Colors array:", colors)
    fmt.Println("Length of colors:", len(colors))
    
    // Initialize specific indices
    var scores [5]int = [5]int{0: 100, 2: 85, 4: 92}
    fmt.Println("Scores array:", scores) // [100 0 85 0 92]
    
    // 2. Accessing and Modifying Array Elements
    
    // Access elements by index

    fmt.Println("First fruit:", fruits[0])
    fmt.Println("Last color:", colors[len(colors)-1])
    
    // Modify elements
    numbers[0] = 10
    numbers[1] = 20
    numbers[4] = 50
    fmt.Println("Modified numbers:", numbers)
    
    // 3. Iterating Through Arrays
    
    // Using traditional for loop
    fmt.Println("\nUsing index loop:")
    for i := 0; i < len(fruits); i++ {
        fmt.Printf("Index %d:", i, fruits[i])
    }
    
    // Using range (gets index and value)
    fmt.Println("\nUsing range (index and value):")
    for index, value := range colors {
        fmt.Printf("Index %d: %s\n", index, value)
    }
    
    // Using range (value only)
    fmt.Println("\nUsing range (value only):")
    for _, value := range fruits {
        fmt.Printf("Fruit: %s\n", value)
    }
    
    // 4. Multi-dimensional Arrays
    
    // 2D array (matrix)
    var matrix [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("\n2D Array (Matrix):")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    // 5. Array Comparison
    
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 4}
    
    fmt.Println("\nArray Comparisons:")
    fmt.Println("arr1 == arr2:", arr1 == arr2) // true
    fmt.Println("arr1 == arr3:", arr1 == arr3) // false
    
    // 6. Passing Arrays to Functions
    
    fmt.Println("\nFunction Examples:")
    printArray(numbers)
    sum := sumArray(arr1)
    fmt.Println("Sum of arr1:", sum)
    
    // Modify array through function (pass by reference using pointer)
    fmt.Println("Before modification:", numbers)
    modifyArray(&numbers)
    fmt.Println("After modification:", numbers)
    
    // 7. Array vs Slice (Important Difference)
    
    // Array - fixed size
    var fixedArray [5]int = [5]int{1, 2, 3, 4, 5}
    
    // Slice - dynamic size (more commonly used)
    var dynamicSlice []int = []int{1, 2, 3, 4, 5}
    dynamicSlice = append(dynamicSlice, 6, 7) // Can grow
    
    fmt.Println("\nArray vs Slice:")
    fmt.Printf("Fixed Array: %v (length: %d)\n", fixedArray, len(fixedArray))
    fmt.Printf("Dynamic Slice: %v (length: %d)\n", dynamicSlice, len(dynamicSlice))
}

// Function that takes array by value (creates a copy)
func printArray(arr [5]int) {
    fmt.Println("Array in function:", arr)
}

// Function that calculates sum of array elements
func sumArray(arr [3]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}

// Function that modifies array by reference (using pointer)
func modifyArray(arr *[5]int) {
    arr[0] = 999
    arr[1] = 888
}
