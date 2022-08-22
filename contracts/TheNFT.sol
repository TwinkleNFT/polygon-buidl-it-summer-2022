// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

import "./TheStorage.sol";
import "./ITheNFT.sol";
import "./ITokenMaker.sol";
import './ICollateralController.sol';

contract TheNFT is ERC721, ERC721Enumerable, ERC721Burnable, Ownable, ITheNFT {
    using Counters for Counters.Counter;
    bool public redemptionEnabled = false;
    address public collateralController; 
    Counters.Counter private _tokenIdCounter;
    constructor() ERC721("TheNFT", "TheNFT") {}

    event NFTCreated(uint indexed mechaId);

    mapping (uint256 => TheStorage) public items;

    function tokenURI(uint _tokenId) view public override returns (string  memory) {
        require(exists(_tokenId), "not yet minted");
        TheStorage storage item = items[_tokenId];
        string memory output = ITokenMaker(token_maker).createJson(_tokenId, item);
        return output;
    }


    function increaseLevel(uint _tokenId) public onlyOwner returns (bool) {
        require(exists(_tokenId), "not yet minted");
        TheStorage storage item = items[_tokenId];
        item.level = item.level + 1; 
        return true;
    }

    function mint(
        address _to
    )
        external onlyOwner
        returns (uint256)
    {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(_to, tokenId);
        TheStorage memory newItem;
        newItem.level = 0;
        newItem.trait_face = ITokenMaker(token_maker).randomFaceNumber(tokenId);
        newItem.trait_body = ITokenMaker(token_maker).randomBodyNumber(tokenId);
        
        items[tokenId] = newItem;
        emit NFTCreated(tokenId);
        return tokenId;
    }

    function newOwner(address _newOwner) external onlyOwner {
        super.transferOwnership(_newOwner);
    }


    address public token_maker;
    function setTokenMaker(address _addr) external onlyOwner{
        token_maker = _addr;
    }

    
    function enableRedeem() public onlyOwner {
        redemptionEnabled = true;
    }
    function disableRedeem() public onlyOwner {
        redemptionEnabled = false;
    }

    
    // setCollateralController
    function setCollateralController(address _addr) external onlyOwner {
        collateralController = _addr;
    }

    // burn 
    function burn(uint256 _tokenId)
        public virtual override    
    {
        require(ownerOf(_tokenId)==msg.sender, "only owner can burn");
        require(redemptionEnabled, "redeem disabled");
        _burn(_tokenId);
        if (collateralController!=address(0)) {
            
            ICollateralController(collateralController).redeem(_tokenId, msg.sender);
            emit Burned(_tokenId, msg.sender);
        }
    }

    function exists(uint256 tokenId) public view  returns (bool) 
    {
        return _exists(tokenId);
    }
    
    // The following functions are overrides required by Solidity.

    function _beforeTokenTransfer(address from, address to, uint256 tokenId)
        internal
        override(ERC721, ERC721Enumerable)
    {
        super._beforeTokenTransfer(from, to, tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
    // 
    event Burned(uint256 _tokenId, address _owner);
}