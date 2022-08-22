// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/utils/math/SafeMath.sol';

import "./ITheNFT.sol";
import './CollateralReserve.sol';
import './IUniswapV2Router.sol';
import './IUniswapV2Factory.sol';

interface IWETH is IERC20 {
    function deposit() external payable;
    function withdraw(uint256 amount) external;
}

contract CollateralController is Ownable {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    address public routerAddress;
    address public nftContract;
    address public thenft;
    mapping(address => bool) public isCollateral;
    address[] public collateralTokens;
    mapping (uint => address) public collaterals;
    
    uint256 private constant LIMIT_SWAP_TIME = 10 minutes;

    constructor (address _nftContract, address _thenft) {
        nftContract = _nftContract;
        thenft = _thenft;
    }
    // struct Parameters {
    //     string tokenId;
    // }

    function deposit(uint tokenId, uint amount, address token) public payable returns(bool){
        require(amount>0, "amount <= 0 ");
        if (collaterals[tokenId]==address(0)) {
            // create new collateral 
            address cr = address(new CollateralReserve{salt: keccak256(abi.encode(address(this), msg.sender))}(
                nftContract, tokenId
            ));
            collaterals[tokenId] = cr;

        } 
        require(collaterals[tokenId]!=address(0), "no collateral address yet...");

        address WETH = 0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889; // mumbai wmatich

        if (token==address(0)) {
            require(msg.value==amount, "diff amount");
            CollateralReserve(collaterals[tokenId]).addCollateral(WETH);
            IWETH(WETH).deposit{value:msg.value}();
            IERC20(WETH).safeTransfer(collaterals[tokenId], msg.value);
        } else {
            CollateralReserve(collaterals[tokenId]).addCollateral(token);
            IERC20(token).safeTransferFrom(msg.sender, collaterals[tokenId], amount);
        }
        return true;
    }



    function redeem(uint _tokenId, address _to) public  returns (bool){
        require(nftContract==msg.sender,"nftContract !=msg.sender");
        
        return CollateralReserve(collaterals[_tokenId]).redeem(_to);
    }



    function zapDeposit(uint _tokenId, address _token, uint _amount) public {
        require(isCollateral[_token], 'Not accepted.');

        IERC20(_token).safeTransferFrom(msg.sender, address(this), _amount);
        _rebalance(_token,thenft,_amount.div(2),0,routerAddress);
        
        uint amt1 = IERC20(_token).balanceOf(address(this));
        uint amt2 = IERC20(thenft).balanceOf(address(this));

        _addLiquidity(_token,thenft,amt1,amt2,0,0,collaterals[_tokenId]);

        address lptoken = IUniswapV2Factory(IUniswapV2Router(routerAddress).factory()).getPair(_token, thenft);
        CollateralReserve(collaterals[_tokenId]).addCollateral(_token);
        CollateralReserve(collaterals[_tokenId]).addCollateral(thenft);
        CollateralReserve(collaterals[_tokenId]).addCollateral(lptoken);


        amt1 = IERC20(_token).balanceOf(address(this));
        amt2 = IERC20(thenft).balanceOf(address(this));
        if (amt1>0) {
            IERC20(_token).safeTransfer(collaterals[_tokenId], amt1);
        }
        if (amt2>0) {
            IERC20(thenft).safeTransfer(collaterals[_tokenId], amt2);
        }
    }


    /* ========== ADMIN FUNCTIONS ========== */

    function setRouterAddress(address _addr) public onlyOwner {
        require(_addr != address(0), 'invalid ');
        routerAddress = _addr;
        emit UpdateRouter(_addr);

    }

    function addBondingCollateral(address _token) public onlyOwner {
        require(_token != address(0), 'invalid token');

        isCollateral[_token] = true;
        if (!_listContains(collateralTokens, _token)) {
          collateralTokens.push(_token);
        }
        emit CollateralAdded(_token);

    }

    // Remove a BondingCollateral
    function deleteBondingCollateral(address _token) public onlyOwner {
        require(_token != address(0), 'invalid token');
        // Delete from the mapping
        // delete collateralCalc[_token];
        delete isCollateral[_token];
        // delete acceptDeposit[_token];
        // delete acceptWithdraw[_token];

        // 'Delete' from the array by setting the address to 0x0
        for (uint256 i = 0; i < collateralTokens.length; i++) {
        if (collateralTokens[i] == _token) {
            // coffin_pools_array[i] = address(0);
            // This will leave a null in the array and keep the indices the same
            delete collateralTokens[i];
            break;
        }
        }
        emit CollateralRemoved(_token);
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
    

  function _rebalance(
    address token0,
    address token1,
    uint256 _amount,
    uint256 _min_output_amount,
    address _routerAddress
  ) internal {

    IERC20(token0).approve(address(_routerAddress), 0);
    IERC20(token0).approve(address(_routerAddress), _amount);

    address[] memory router_path = new address[](2);
    router_path[0] = token0;
    router_path[1] = token1;

    uint256[] memory _received_amounts = IUniswapV2Router(_routerAddress)
      .swapExactTokensForTokens(
        _amount,
        _min_output_amount,
        router_path,
        address(this),
        block.timestamp + LIMIT_SWAP_TIME
      );

    require(
      _received_amounts[_received_amounts.length - 1] >= _min_output_amount,
      'Slippage limit reached'
    );
  }


    function _addLiquidity(
      address token0,
      address token1,
      uint256 amtToken0,
      uint256 amtToken1,
      uint256 minToken0,
      uint256 minToken1,
      address dest
    )
      internal
      
      returns (
        uint256,
        uint256,
        uint256
      )
    {
        require(address(routerAddress)!=address(0), "need to setup uni router");
        require(amtToken0 != 0 && amtToken1 != 0, "amounts can't be 0");
        
        IERC20(token0).approve(address(routerAddress), 0);
        IERC20(token0).approve(address(routerAddress), amtToken0);

        IERC20(token1).approve(address(routerAddress), 0);
        IERC20(token1).approve(address(routerAddress), amtToken1);

        uint256 resultAmtToken0;
        uint256 resultAmtToken1;
        uint256 liquidity;

        (resultAmtToken0, resultAmtToken1, liquidity) = IUniswapV2Router(routerAddress)
          .addLiquidity(
            token0,
            token1,
            amtToken0,
            amtToken1,
            minToken0,
            minToken1,
            dest, // address(this),
            block.timestamp + LIMIT_SWAP_TIME
          );
        emit CreateLiquidy(token0, token1, liquidity);
        return (resultAmtToken0, resultAmtToken1, liquidity);
    }

    event CreateLiquidy(address token0, address token1, uint256 liquidity);

  event CollateralRemoved(address _token);
  event CollateralAdded(address _token);
  event UpdateRouter(address _addr);
}