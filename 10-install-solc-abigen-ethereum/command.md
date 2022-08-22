1. Install solidity compiler (Solc):
   Reference: https://archlinux.org/packages/community/x86_64/solidity/
2. https://www.youtube.com/watch?v=epAA4J5UKc4&list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_&index=10
   - Install Protoc Arch Linux:
    ```console
    curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-osx-universal_binary.zip

    unzip -d protoc protoc-21.5-osx-universal_binary.zip

    cd protoc

    ls
    ```
    - Installing Go 
  
    - Make sure your gopath is correct
  
    ```console
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
    - Install abigen
    ```console
    go get -u github.com/ethereum/go-ethereum
    ```
    <img src="/10-install-solc-abigen-ethereum/get%20go-ethereum.png" />

    - After get go-ethereum
    ```console
    cd $GOPATH/src/github.com/ethereum/go-ethereum/
    make
    make devtools
    ```

    - Check if abigen is installed
    ```console
        abigen --help
    ```
    <img src="/10-install-solc-abigen-ethereum/makedevtools.png" />