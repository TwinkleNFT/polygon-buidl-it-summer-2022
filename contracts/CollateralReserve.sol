// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/utils/math/SafeMath.sol';
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract CollateralReserve is Ownable
{
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    address public nftContract;
    address public collateralController;
    uint public tokenId; 

    address[] public collaterals;
    mapping(address => bool) public isCollateral;

    constructor(address _nftContract, uint _tokenId) {
        nftContract = _nftContract;
        tokenId = _tokenId;
        collateralController = msg.sender;
    }

    /* ========== MODIFIER ========== */


    modifier onlyNFTContract() {
      require(
        nftContract== msg.sender ,
        'Only NFTContract can redeem'
      );
      _;
    }
    modifier onlyCollateralController() {
      require(
        collateralController== msg.sender ,
        'Only CollateralController can redeem'
      );
      _;
    }

    /* ========== RESTRICTED FUNCTIONS ========== */


    // add a collatereal
    function addCollateral(address _token)
        public
    {
        require(msg.sender==collateralController,"collateralController error");
        require(_token != address(0), 'invalid token');

        isCollateral[_token] = true;
        if (!_listContains(collaterals, _token)) {
          collaterals.push(_token);
        }
        // emit CollateralAdded(_token);
    }

    // function redeem(uint _total, address _recipient) external onlyNFTContract returns (bool) {
    function redeem(address _recipient) external onlyCollateralController returns (bool) {
       
        for (uint256 i = 0; i < collaterals.length; i++) {
          if (
            address(collaterals[i]) != address(0)
          ) {
            try IERC20(collaterals[i]).balanceOf(address(this)) returns (
              uint256 bal
            ) {
              if (bal>0) {
                IERC20(collaterals[i]).safeTransfer(_recipient, bal);
              }
            } catch Error(
              string memory /*reason*/
            ) {
              // value_ = 0;
            }
          }
        }
        return true;
    }


    /* ========== INTERNAL FUNCTIONS ========== */



    function _listContains(address[] storage _list, address _token) internal view returns (bool) {
        for (uint256 i = 0; i < _list.length; i++) {
          if (_list[i] == _token) {
            return true;
          }
        }
        return false;
    }
    


    /* ========== EVENTS  ========== */
    // event ManagerChanged(address indexed addr);
    // event RedeemerChanged(address indexed addr);
    
    // event BalancerV2VaultChanged(address indexed addr);
    // event UniV2RouterChanged(address indexed addr);
    // event RedemptionFeesChanged(uint fees);
    // event Redeemed(address _token, uint amt, address recipient);

    // event CollateralRemoved(address _token);
    // event CollateralAdded(address _token);

    // event CreateLiquidy(address token0, address token1, uint256 liquidity);
    // event RemoveLiquidy(address token0, address token1, uint256 liquidity);
}
