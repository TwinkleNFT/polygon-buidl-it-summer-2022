
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface CollateralReserve {
  function addCollateral(address _token) external ;
  // function redeem(uint _total, address _recipient) external  returns (bool);
  function redeem(address _recipient) external  returns (bool);
}