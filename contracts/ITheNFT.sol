// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface ITheNFT {

    function enableRedeem() external  ;
    function disableRedeem() external ;
    function setTokenMaker(address _addr) external ;
    
    function mint(address _to) external returns (uint256);
    function newOwner(address _newOwner) external ;
    function setCollateralController(address _addr) external;
    function increaseLevel(uint _tokenId) external returns(bool);
}
