pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/introspection/ERC165CheckerUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../tokens/MirroredERC20.sol";
import "../interfaces/IERC20Mirrored.sol";

contract ERC20SwapAgent is
Initializable,
OwnableUpgradeable
{
  using ERC165CheckerUpgradeable for address;

  // -- ERC20 INTERFACE ID --
  bytes4 private constant _ERC_20_INTERFACE_ID = 0xd9b67a26;

  // -- Error Constants --
  // ERC20SwapAgent::registerSwapPair:: token does not conform ERC20 standard 
  string private constant _ERR20_REGISTER_NOT_20 = "A20.20.1.1";
  // ERC20SwapAgent::registerSwapPair:: token is already registered
  string private constant _ERR20_REGISTER_TOKEN_EXISTS = "A20.20.1.2";
  // ERC20SwapAgent::registerSwapPair:: empty token name
  string private constant _ERR20_REGISTER_EMPTY_TOKEN_NAME = "A20.20.1.3";
  // ERC20SwapAgent::registerSwapPair:: empty token symbol
  string private constant _ERR20_REGISTER_EMPTY_TOKEN_SYMBOL = "A20.20.1.4";

  // ERC20SwapAgent::createSwapPair:: mirrored token is already deployed
  string private constant _ERR20_CREATE_MIRRORED_EXISTS = "A20.20.2.1";

  // ERC20SwapAgent::swap:: token has no swap pair
  string private constant _ERR20_SWAP_NO_PAIR = "A20.20.3.1";


  // ERC20SwapAgent::fill:: tx hash was already filled
  string private constant _ERR20_FILL_ALREADY_FILLED = "A20.20.4.1";
  // ERC20SwapAgent::fill:: token has no swap pair
  string private constant _ERR20_FILL_NO_PAIR = "A20.20.4.2";

  
  
  // -- Storage Variables --
  mapping(uint256 => mapping(address => bool)) public registeredToken;
  mapping(bytes32 => bool) public filledSwap;
  mapping(uint256 => mapping(address => address)) public swapMappingIncoming;
  mapping(uint256 => mapping(address => address)) public swapMappingOutgoing;

  /* Events from outgoing messages */
  event SwapPairRegister(
    address indexed sponsor,
    address indexed tokenAddress,
    uint256 toChainId,
    uint256 feeAmount
  );

  event SwapStarted(
    address indexed tokenAddr,
    address indexed sender,
    address indexed recipient,
    uint256 dstChainId,
    uint256 amount,
    uint256 feeAmount
  );

  event BackwardSwapStarted(
    address indexed mirroredTokenAddr,
    address indexed sender,
    address indexed recipient,
    uint256 dstChainId,
    uint256 amount,
    uint256 feeAmount
  );

  /* Events from incoming messages */
  event SwapPairCreated(
    bytes32 indexed registerTxHash,
    address indexed fromTokenAddr,
    address indexed mirroredTokenAddr,
    uint256 fromChainId
  );

  event SwapFilled(
    bytes32 indexed swapTxHash,
    address indexed fromTokenAddr,
    address indexed recipient,
    address mirroredTokenAddr,
    uint256 fromChainId,
    uint256 amount
  );

  event BackwardSwapFilled(
    bytes32 indexed swapTxHash,
    address indexed tokenAddr,
    address indexed recipient,
    uint256 fromChainId,
    uint256 amount
  );

  function initialize() public initializer {
    __Ownable_init();
  }

  function createSwapPair(
    bytes32 registerTxHash,
    address fromTokenAddr,
    uint256 fromChainId,
    string memory _name, 
    string memory _symbol,
    uint8 _decimals
  ) public onlyOwner {
    require(
      swapMappingIncoming[fromChainId][fromTokenAddr] == address(0x0),
      _ERR20_CREATE_MIRRORED_EXISTS
    );

    MirroredERC20 mirrored = new MirroredERC20();
    mirrored.initialize(_name, _symbol, _decimals);

    swapMappingIncoming[fromChainId][fromTokenAddr] = address(mirrored);
    swapMappingOutgoing[fromChainId][address(mirrored)] = fromTokenAddr;

    emit SwapPairCreated(
      registerTxHash,
      fromTokenAddr,
      address(mirrored),
      fromChainId
    );
  }

  function registerSwapPair(address tokenAddr, uint256 chainId)
  external
  payable
  {
    // 没有有效的方式判断是否是erc20,通过在页面来判断
//    require(
//      tokenAddr.supportsInterface(_ERC_20_INTERFACE_ID),
//      _ERR20_REGISTER_NOT_20
//    );

    require(
      !registeredToken[chainId][tokenAddr],
      _ERR20_REGISTER_TOKEN_EXISTS
    );

    registeredToken[chainId][tokenAddr] = true;

    if (msg.value != 0) {
      payable(owner()).transfer(msg.value);
    }

    emit SwapPairRegister(
      msg.sender,
      tokenAddr,
      chainId,
      msg.value
    );
  }

  function swap(
    address tokenAddr,
    address recipient,
    uint256 amount,
    uint256 dstChainId
  ) external payable {
    if (msg.value != 0) {
      payable(owner()).transfer(msg.value);
    }

    // try forward swap
    if (registeredToken[dstChainId][tokenAddr]) {
      IERC20 token = IERC20(tokenAddr);
      token.transferFrom(msg.sender, address(this), amount);

      emit SwapStarted(
        tokenAddr,
        msg.sender,
        recipient,
        dstChainId,
        amount,
        msg.value
      );

      return;
    }

    // try backward swap
    address dstTokenAddr = swapMappingOutgoing[dstChainId][tokenAddr];
    if (dstTokenAddr != address(0x0)) {
      IERC20Mirrored mirroredToken = IERC20Mirrored(tokenAddr);
      mirroredToken.transferFrom(msg.sender, address(this), amount);
      mirroredToken.burn(address(this), amount);

      emit BackwardSwapStarted(
        tokenAddr,
        msg.sender,
        recipient,
        dstChainId,
        amount,
        msg.value
      );

      return;
    }

    revert(_ERR20_SWAP_NO_PAIR);
  }

  function fill(
    bytes32 swapTxHash,
    address fromTokenAddr,
    address recipient,
    uint256 fromChainId,
    uint256 amount
  ) public onlyOwner {
    require(!filledSwap[swapTxHash], _ERR20_FILL_ALREADY_FILLED);
    filledSwap[swapTxHash] = true;

    // fill forward swap, it means our core server will find the related mirrored token
    // and assign the value to fromTokenAddr
    address mirroredTokenAddr = swapMappingIncoming[fromChainId][
    fromTokenAddr
    ];
    if (mirroredTokenAddr != address(0x0)) {
      IERC20Mirrored mirroredToken = IERC20Mirrored(mirroredTokenAddr);
      mirroredToken.mint(recipient, amount);

      emit SwapFilled(
        swapTxHash,
        fromTokenAddr,
        recipient,
        mirroredTokenAddr,
        fromChainId,
        amount
      );

      return;
    }

    // fill backward swap, it means that this token is the one users have been sent before
    // our server will find this token from the given mirrored token in the BackwardSwapStarted event
    // and assign the value to fromTokenAddr
    if (registeredToken[fromChainId][fromTokenAddr]) {
      IERC20 token = IERC20(fromTokenAddr);
      token.transferFrom(address(this), recipient, amount);

      emit BackwardSwapFilled(
        swapTxHash,
        fromTokenAddr,
        recipient,
        fromChainId,
        amount
      );

      return;
    }

    revert(_ERR20_FILL_NO_PAIR);
  }
}
