package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func main() {
	// ### Exercise 1: Concurrent Map Access

	// **Objective**: Write a Go program that creates a map with string keys and integer values. The program should launch multiple goroutines. Each goroutine should access the map to read and update its elements in a thread-safe manner.

	// **Steps**:
	// 1. Define a map with string keys and integer values.
	myMap := make(map[string]int)
	// 2. Implement a mutex to ensure safe access to the map.

	var mu sync.Mutex
	
	numGoroutines := 5

	done := make(chan bool)

	rand.Seed(time.Now().UnixNano())

	// 3. Create multiple goroutines. Each goroutine should:
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			//    - Read a random key from the map.
			key := generateRandomKey()

			//    - Increment its value.
			mu.Lock()
			myMap[key]++
			value := myMap[key]
			mu.Unlock()

			//    - Write the value back to the map.
			fmt.Printf("Goroutine %d: Key '%s' incremented to %d\n", goroutineID, key, value)

			done <- true
		}(i)
	}

	// 4. Ensure that all goroutines complete their execution.
	for i :=0; i < numGoroutines; i++ {
		<-done
	}
	
	// 5. Print the final state of the map.
	fmt.Println("Final Map:", myMap)
	
}

func generateRandomKey() string {
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	return keys[rand.Intn(len(keys))]
}