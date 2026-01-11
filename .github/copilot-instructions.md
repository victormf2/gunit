# Code Style Instructions

Variable declarations:

- Use short variable declarations (:=) for local variables whenever possible.

Maps and slices:

- When iterating over slices, prefer using range loops (for i := range slice) instead of traditional for loops with index counters.
- When creating slices or maps, use the literal syntax (e.g., []Type{} or map[KeyType]ValueType{}) instead of the make function, unless you need to specify an initial capacity.
