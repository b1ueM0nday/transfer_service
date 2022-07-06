//SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
contract BalanceOperations{
    address payable public owner;

    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {}

    mapping(address => uint) public balances;

    event Deposit(uint amount);
    event Withdrawal(uint amount);
    event Transfer(address receiver, uint amount);

    function deposit(uint amount) public payable {
        emit Deposit(amount);
        balances[owner] += amount;
    }

    function withdraw(uint amount) external{
        require(msg.sender == owner, "caller is not owner");
        require(balances[owner] >= amount, "Insufficient funds");
        emit Withdrawal(amount);
        balances[owner] -= amount;
    }

    function transfer(address receiver, uint amount) public {
        require(msg.sender == owner, "caller is not owner");
        require(balances[owner] >= amount, "Insufficient funds");
        emit Transfer(receiver, amount);
        balances[owner] -= amount;
        balances[receiver] += amount;
    }

    // In a Batch

 /*   function transfer(address[] memory receivers, uint amount) public {
        require(balances[msg.sender] >= receivers.length * amount, "Insufficient funds");
        for (uint i=0; i<receivers.length; i++) {
            emit Transfer(msg.sender, receivers[i], amount);
            balances[msg.sender] -= amount;
            balances[receivers[i]] += amount;
        }
    }*/
}

