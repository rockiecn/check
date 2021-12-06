// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "./library/Recover.sol";
import "./library/SafeMath.sol";


struct Check {
    uint256 value;          // value of the check, payvalue shoud not exceed value
    address tokenAddr;      // token address, point out which token to pay
    uint64  nonce;          // nonce of the check, check's nonce should not smaller than it.
    address fromAddr;       // buyer of this check, should be check's signer
	address toAddr;         // receiver of check's money, point out who to pay
	address opAddr;         // operator of this cheuqe, shuould be contract's owner
	address ctrAddr;        // contract address, should be this contract
    bytes   checkSig;       // signer of this signature should be operator.
}

struct Paycheck {
	Check check;
    uint256 payValue;       // money to pay, should not exceed value
    bytes paycheckSig;      // signer of this signature should be user.
}

struct BatchCheck  {
    address opAddr;      // operator address
	address toAddr;      // 存储节点号
	address ctrAddr;     // 合约地址
	address tokenAddr;
	uint256 batchValue;  // 聚合后的支票面额
	uint64 minNonce;     // 聚合的nonce最小值
	uint64 maxNonce;     // 聚合的nonce最大值
	bytes batchSig;      // signature of operator
}


contract Cash  {
    using SafeMath for uint256;
    
    event Pos(uint256);
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
        //require(paycheck.check.CtrAddr == address(this), "contract address error");
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
            paycheck.check.ctrAddr
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
        payable(paycheck.check.toAddr).transfer(paycheck.payValue); //pay money to storage
        emit Paid(paycheck.check.toAddr, paycheck.payValue);
        
        // update nonce after paid
        nodeNonce[paycheck.check.toAddr] = paycheck.check.nonce + 1;

        return true;
    }
    

    function withdrawBatch(BatchCheck memory bc) public payable returns(bool) {
        
        require(bc.minNonce >= nodeNonce[bc.toAddr], "batch check nonce too old");
        emit Pos(1);
        require(bc.ctrAddr == address(this), "contract address error");
        emit Pos(2);
        require(bc.toAddr == msg.sender, "caller shuould be check.toAddr");
        emit Pos(3);
        require(bc.opAddr == owner, "operator should be owner of this contract");
        emit Pos(4);
        

        // verify check's signer
        bytes memory bcPack = 
        abi.encodePacked(
            bc.opAddr,
            bc.toAddr,
            bc.ctrAddr, 
            bc.tokenAddr,
            bc.batchValue,
            bc.minNonce,
            bc.maxNonce
    		);
        bytes32 bcHash = keccak256(bcPack);
        address bcSigner = Recover.recover(bcHash, bc.batchSig);
        
        require(bc.opAddr == bcSigner, "illegal bc sig");
        emit Pos(5);
    	
        // pay
        payable(bc.toAddr).transfer(bc.batchValue); // pay money to storage
        emit Paid(bc.toAddr, bc.batchValue);
        
        // update nonce after paid
        nodeNonce[bc.toAddr] = bc.maxNonce + 1;

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
    
    // set nonce for test
    function setNonce(address node, uint64 nonce) public {
        nodeNonce[node]=nonce;
    }
    
    // get owner of the contract
    function getOwner() public view returns(address) {
        return owner;
    }
}