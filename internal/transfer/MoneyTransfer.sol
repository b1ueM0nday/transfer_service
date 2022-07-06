//SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
contract MT{
    mapping(address => uint) public balances;

    event Deposit(address sender, uint amount);
    event Withdrawal(address receiver, uint amount);
    event Transfer(address sender, address receiver, uint amount);

    function deposit(uint amount) public payable {
        emit Deposit(msg.sender, amount);
        balances[msg.sender] += amount;
    }

    function withdraw(uint amount) public {
        require(balances[msg.sender] >= amount, "Insufficient funds");
        emit Withdrawal(msg.sender, amount);
        balances[msg.sender] -= amount;
    }

    function transfer(address receiver, uint amount) public {
        require(balances[msg.sender] >= amount, "Insufficient funds");
        emit Transfer(msg.sender, receiver, amount);
        balances[msg.sender] -= amount;
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

