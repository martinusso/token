# token 

Creates a token hash and verifies that a token matches a key and a cost.

## How to use

### Hash()

Creates a token hash using a strong one-way hashing algorithm (bcrypt).

#### Parameters

- key (string) - a key
- cost (integer) - which denotes the algorithmic cost that should be used

#### Example

```
import "github.com/martinusso/token"

token := token.Hash("123", 10)
fmt.Println(token)
```

The above example will optput something similar to:

```
JDJhJDEwJDY1L3QwVFIyeFBMOWlNZWxxbWtzRU9HWUJ5MXZoajg2NExOU0FNTVZUQjEzR3dndjExTlYy
```


### Verify()

Verifies that a token hash matches a key and cost.

#### Parameters

- key (string) - a key
- cost (integer) - which denotes the algorithmic cost that should be used
- hash (string) - a token hash created by Hash method

#### Example

```
hash := "JDJhJDEwJDY1L3QwVFIyeFBMOWlNZWxxbWtzRU9HWUJ5MXZoajg2NExOU0FNTVZUQjEzR3dndjExTlYy"
fmt.Println(Verify("123", 10, hash))
```

The above example will output `true`.
