// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import "@openzeppelin/contracts/access/Ownable.sol";

import "./ITheNFT.sol";

contract TwinkleController is Ownable {
    address public nftContract;
    constructor (address _nftContract) {
        nftContract = _nftContract;
    }
    function enableRedeem() external onlyOwner{
        ITheNFT(nftContract).enableRedeem();
    }
    function disableRedeem() external  onlyOwner{
        ITheNFT(nftContract).disableRedeem();
    }
    function mint(address _to) external onlyOwner returns (uint256){
        uint tokenId = ITheNFT(nftContract).mint(_to);
        return tokenId;
    }
    
    function newOwner(address _newOwner) public  onlyOwner {
        ITheNFT(nftContract).newOwner(_newOwner);
    }
    function setCollateralController(address _addr) public  onlyOwner {
        ITheNFT(nftContract).setCollateralController(_addr);
    }
    function setTokenMaker(address _addr) public  onlyOwner {
        ITheNFT(nftContract).setTokenMaker(_addr);
    }

    function increaseLevel(uint _tokenId) public  onlyOwner returns(bool){
        return ITheNFT(nftContract).increaseLevel(_tokenId);
    }
}