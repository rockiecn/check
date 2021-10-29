// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;


import "./library/Recover.sol";
import "./library/SafeMath.sol";


struct Check {
    uint256 value;         // value of the check, payvalue shoud not exceed value
    address tokenAddr;      // token address, point out which token to pay
    uint64  nonce;          // nonce of the check, check's nonce should not smaller than it.
    address fromAddr;       // buyer of this check, should be check's signer
	address toAddr;         // receiver of check's money, point out who to pay
	address opAddr;         // operator of this cheuqe, shuould be contract's owner
	address contractAddr;   // contract address, should be this contract
    bytes   checkSig;         // signer of this signature should be operator.
}

struct Paycheck {
	Check check;
    uint256 payValue;       // money to pay, should not exceed value
    bytes paycheckSig;      // signer of this signature should be user.
}


contract Cash  {
    using SafeMath for uint256;
    
    event Received(address, uint256);
    event Paid(address, uint256);
    
    address owner;
    mapping(address => uint64) public nodeNonce;
    

    // constructor
    constructor() payable {
        owner = msg.sender;
    }
    
    // receiver
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    // called by storage
    function withdraw(Paycheck memory paycheck) public payable returns(bool) {
        
        require(paycheck.check.nonce >= nodeNonce[paycheck.check.toAddr], "check.nonce too old");
        require(paycheck.payValue <= paycheck.check.value, "payvalue should not exceed value of check.");
        //require(paycheck.check.contractAddr == address(this), "contract address error");
        require(paycheck.check.toAddr == msg.sender, "caller shuould be check.toAddr");
        require(paycheck.check.opAddr == owner, "operator should be owner of this contract");
        

        // verify check's signer
        bytes memory checkPack = 
        abi.encodePacked(
            paycheck.check.value,
            paycheck.check.tokenAddr,
            paycheck.check.nonce,
            paycheck.check.fromAddr,
            paycheck.check.toAddr,
            paycheck.check.opAddr,
            paycheck.check.contractAddr
    		);
        bytes32 checkHash = keccak256(checkPack);
        address checkSigner = Recover.recover(checkHash,paycheck.check.checkSig);
        
        require(paycheck.check.opAddr == checkSigner, "illegal check sig");
    	
        // verify paycheck's signer
        bytes memory paycheckPack = 
        abi.encodePacked(
            checkPack,
            paycheck.payValue
        );
        bytes32 paycheckHash = keccak256(paycheckPack);
        address paycheckSigner = Recover.recover(paycheckHash,paycheck.paycheckSig);
        require(paycheck.check.fromAddr == paycheckSigner, "illegal paycheck sig");
        
        // pay
        payable(paycheck.check.toAddr).transfer(paycheck.payValue); //pay value to storage
        emit Paid(paycheck.check.toAddr, paycheck.payValue);
        
        // update nonce after paid
        //nodeNonce[paycheck.check.toAddr] = paycheck.check.nonce;

        return true;
    }

    // deposit some money to contract
    function deposit() public payable{
    }
  
    // get balance of contract
    function getBalance() public view returns(uint256) {
        return address(this).balance;
    }
    
    // get nonce of a specified node
    function getNonce(address node) public view returns(uint64) {
        return nodeNonce[node];
    }
    
    // get owner of the contract
    function getOwner() public view returns(address) {
        return owner;
    }
}