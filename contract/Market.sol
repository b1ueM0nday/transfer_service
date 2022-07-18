//SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
contract Market{

    address payable public owner;

    enum AccountType {BASE,REGULAR}

    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {}

    modifier accessLimitid(AccountType _accountType){
        if (accounts[msg.sender].accType >= _accountType) {
            _;
        }
    }
    mapping(address => Account) public accounts;
    address[] accountsArray;

    struct Account{
        uint balance;
        UserData userData;
        bool isActive;
        AccountType accType;
        uint64 itemsCount;
        string[] itemsArray;
        mapping(string => Item) Items;
    }
    struct UserData{
        string name;
        string phone;
        string email;
        uint64 birthday;
        uint64 regTime;
    }
    struct Item{
        string vendorCode;
        string name;
        string description;
        uint256 price;
        uint count;
        bool isActive;
    }
    event RegisterAccount(string name);
    event UpgradeAccount();
    event Deposit(uint amount);
    event Withdrawal(uint amount);
    event Transfer(address receiver, uint amount);

    function registerAccount(string memory _name, string memory _phone, string memory _email, uint64 _birthday, uint64 _regtime) public{
        require(!accounts[msg.sender].isActive, "Account already exists");
        emit RegisterAccount(_name);
        UserData memory userData = UserData(_name, _phone, _email, _birthday, _regtime);

        accounts[msg.sender].userData = userData;
        accounts[msg.sender].isActive = true;
        accounts[msg.sender].accType = AccountType.BASE;
    }

    function upgradeAccount() public payable{
        require(accounts[msg.sender].accType != AccountType.REGULAR, "Account already regular");
        require(accounts[msg.sender].isActive, "Account already exists");
        require(accounts[msg.sender].balance==1 ether);
        emit UpgradeAccount();
        owner.transfer(msg.value);
        accounts[msg.sender].accType = AccountType.REGULAR;
    }

    function changeName (string memory _name) public accessLimitid(AccountType.REGULAR){
        require(!accounts[msg.sender].isActive, "Account already exists");
        accounts[msg.sender].userData.name = _name;
    }

    function changePhone (string memory _phone) public accessLimitid(AccountType.REGULAR){
        require(!accounts[msg.sender].isActive, "Account already exists");
        accounts[msg.sender].userData.phone = _phone;
    }


    function changeEmail (string memory _email) public accessLimitid(AccountType.REGULAR){
        require(!accounts[msg.sender].isActive, "Account already exists");
        accounts[msg.sender].userData.email = _email;
    }

    function getAccountInfo(address accAddress) public view returns(UserData memory)
    {   require (accounts[msg.sender].accType >= AccountType.REGULAR, "Access denied, upgrade account");
        require(accounts[msg.sender].isActive, "Account does not exists");
        return accounts[accAddress].userData;
    }


    function deposit(uint amount) public payable {
        require(accounts[msg.sender].isActive, "Account does not exists");
        emit Deposit(amount);
        accounts[msg.sender].balance += amount;
    }

    function withdraw(uint amount) external{
        require (accounts[msg.sender].accType >= AccountType.REGULAR, "Access denied, upgrade account");
        require(accounts[msg.sender].isActive, "Account does not exists");
        require(accounts[msg.sender].balance >= amount, "Insufficient funds");
        emit Withdrawal(amount);
        accounts[msg.sender].balance -= amount;
    }

    function transfer(address receiver, uint amount) public {
        require(accounts[msg.sender].isActive, "Account does not exists");
        require(accounts[receiver].isActive, "Receiver does not exists");
        require(accounts[msg.sender].balance >= amount, "Insufficient funds");
        emit Transfer(receiver, amount);
        accounts[msg.sender].balance -= amount;
        accounts[receiver].balance += amount;
    }

    function addItem(string memory _vendorCode, string memory _name, string memory _description, uint256 _price, uint _count) public accessLimitid(AccountType.REGULAR) {
        require(!accounts[msg.sender].Items[_vendorCode].isActive, "Item already exists!");
        Item  memory i= Item(_vendorCode, _name, _description, _price, _count, true);
        accounts[msg.sender].Items[_vendorCode] = i;
        accounts[msg.sender].itemsArray.push(_vendorCode);
    }
    function removeItem(string memory _vendorCode) public accessLimitid(AccountType.REGULAR){
        require(!accounts[msg.sender].Items[_vendorCode].isActive, "Item already removed;");
        require(accounts[msg.sender].isActive);
        accounts[msg.sender].Items[_vendorCode].isActive = false;
    }

    function getAccountItemsList() public view returns(Item[] memory){

        Item[] memory qq;
        for (uint i = 0; i < accounts[msg.sender].itemsCount; i++){
            qq[i] = accounts[msg.sender].Items[accounts[msg.sender].itemsArray[i]];
        }
        return qq;
    }

    function updateItem(string memory _vendorCode, string memory _desc, uint256 _price, uint _count) public{
        accounts[msg.sender].Items[_vendorCode].description = _desc;
        accounts[msg.sender].Items[_vendorCode].price = _price;
        accounts[msg.sender].Items[_vendorCode].count = _count;
    }
    function buyItem(address payable seller, string memory _vendorCode, uint count) public{
        require (accounts[seller].Items[_vendorCode].price * count <= accounts[msg.sender].balance, "Insufficient funds");
        require (accounts[seller].Items[_vendorCode].count >= count, "Seller's items count is not enough");
        seller.transfer(accounts[seller].Items[_vendorCode].price * count);
        accounts[seller].Items[_vendorCode].count -= count;
    }


}
