```console
[ngocthieu@NT-majaro 11-generate-bin-abi-file-ethereum]$ solc --bin --abi contract/todo.sol -o build
```
 
result:
```console
Warning: SPDX license identifier not provided in source file. Before publishing, consider adding a comment containing "SPDX-License-Identifier: <SPDX-License>" to each source file. Use "SPDX-License-Identifier: UNLICENSED" for non-open-source code. Please see https://spdx.org for more information.
--> contract/todo.sol
```

- Create folder gen
  
  ```console
  abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go
  ```
  <img src="/11-generate-bin-abi-file-ethereum/err%20abigen.png">

  If the above command is not working, use this instead:

  ```console
  abigen --bin=/home/ngocthieu/Documents/go-ethereum-idirdev-implementation/11-generate-bin-abi-file-ethereum/build/Todo.bin --abi=/home/ngocthieu/Documents/go-ethereum-idirdev-implementation/11-generate-bin-abi-file-ethereum/build/Todo.abi --pkg=todo --out=/home/ngocthieu/Documents/go-ethereum-idirdev-implementation/11-generate-bin-abi-file-ethereum/gen/todo.go
  ```

  <img src="/11-generate-bin-abi-file-ethereum/abigen.png">
   
  Result is go file is generated in folder gen

  <img src="/11-generate-bin-abi-file-ethereum/abigen result.png">
