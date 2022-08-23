pragma solidity >=0.7.0 <0.9.0;

contract Todo {
    address public owner;
    Task[] tasks;

    struct Task{
        string content;
        bool status;
    }

    constructor(){
        owner = msg.sender;
    }

    // Much like RAM, Memory in Solidity is a temporary place to store data whereas Storage holds data between function calls
    function add(string memory _content) public {
        require(owner == msg.sender);
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view returns (Task memory){
        require(owner == msg.sender);
        return tasks[_id];
    }

    function list() public view returns(Task[] memory){
        require(owner == msg.sender);
        return tasks;
    }

    function update(uint _id, string memory _content) public {
        require(owner == msg.sender);
        tasks[_id].content = _content;
    }

    function remove(uint _id) public {
        require(owner == msg.sender);
        delete tasks[_id];
    }
}