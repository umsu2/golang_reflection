# golang_reflection

Motivation:

During serialization of certain lib or data, for eg. csv,sql, etc the data is given back in nested arrays
instead of iterating through each array obj and trying to cast them to particular type I would want to have a function that have the following interface


`Serialize(data [][]interface{}, ...i interface{})`  
Or  
`SqlDecode(sqlStmt, params, ...i interface{})`

where i are list of ptrs to hold the deserialized data

so what if i's are list of structs of which the data could deserialize into?


basically the func need to pass ptr to a list of structs of which data will be populated

this example basically shows how to use the reflection lib in go to grab the type info from the ptr ,and create a
temp list and temp struct then populate it through some deserialization.

some steps.
1. get the interface which is the ptr that points to array of some struct
2. get the type of the slice which ptr is pointing to
3. get the type of the slice element
4. create new using the type from 3
5. create slice using the type from 2
6. do manipulation to 4 by passing the interface of 4 to another func
7. append the dereferenced of 4 to 5
8. set the original slice to new slice.