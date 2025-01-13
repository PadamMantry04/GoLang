// mapVariable := make(map[keyType]valueType)
// myMap := make(map[string]int)

// What is the difference between new and make?
// Purpose:

// make(): Used for slices, maps, and channels. It allocates and initializes.
// new(): Used for any type. It only allocates memory and returns a pointer to the zero value of that type.

// Return Type:

// make(): Returns an initialized value of the type (e.g., a slice, not a pointer to a slice).
// new(): Always returns a pointer to the zero value of the type.

// Usage:

// With make(), the created slice/map/channel is ready to use without further initialization.
// With new(), you generally need further steps to initialize or set the value at the allocated address.

package main

func main() {

}
