pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

interface IERC20Mirrored is IERC20Upgradeable {
  function mint(address account, uint256 amount) external;

  function burn(address account, uint256 value) external;
}
