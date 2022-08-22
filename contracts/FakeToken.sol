// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract FakeToken is ERC20 {
    address _owner;
    
    constructor(
    ) ERC20('FakeToken', 'FakeToken') {
        _owner = msg.sender;
    }

    // only owner can mint
    function mint(address dst, uint256 amt) public returns (bool) {
        require(msg.sender==_owner, "only owner can mint");
        _mint(dst, amt);
        return true;
    }


}
