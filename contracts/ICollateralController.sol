
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
interface ICollateralController {

    function deposit(uint tokenId, uint amount, address token) external payable returns(bool);
    function redeem(uint _tokenId, address _to) external  returns (bool);

}
