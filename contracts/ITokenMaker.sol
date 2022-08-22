
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import './TheStorage.sol';

interface ITokenMaker {
    function createJson(uint256 _tokenId, TheStorage memory _item) external pure returns(string memory);

    function randomFaceNumber(uint tokenId) external view returns (uint8) ;
    function randomBodyNumber(uint tokenId) external view returns (uint8) ;

}
