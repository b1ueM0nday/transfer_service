pragma solidity >=0.4.22 <0.9.0;
contract BalanceOperations{
    address payable public owner;
    enum WalletType{BASE, REGULAR, PREMIUM, VIP}
    WalletType wallet_type;
    WalletType constant default_type = wallet_type.BASE;

    function upgradeType(uint8 levels) public {
        choice = FreshJuiceSize.LARGE;
    }
    function degradeType(uint8 levels) public {
        choice = FreshJuiceSize.LARGE;
    }
    function getChoice() public view returns (FreshJuiceSize) {
        return choice;
    }
    function getDefaultChoice() public pure returns (uint) {
        return uint(defaultChoice);
    }
    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {}

    mapping(address => Wallet) public balances;
    struct Wallet {
        uint _balance;
        string _name;
        WalletType _wallet_type;
    }
    event Deposit(uint amount);
    event Withdrawal(uint amount);
    event Transfer(address receiver, uint amount);
    function registerWallet(address _userAddress, string memory _username, uint _age) public {
        _user[_userAddress] = User(_username, _age);
    }

    function getWalletData(address _address) view public returns(User memory) {
        return _user[_address];
    }
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

