pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MockERC20 is ERC20, Ownable {
  uint8 private _decimals;
  constructor(string memory name, string memory symbol, uint8 decimals) ERC20(name, symbol) {
    _decimals = decimals;
  }

  function decimals() public view virtual override returns (uint8) {
    return _decimals;
  }

  function mint(address account, uint256 amount)
  public
  onlyOwner
  {
    _mint(account, amount);
  }

  function burn(address account, uint256 amount)
  public
  onlyOwner
  {
    _burn(account, amount);
  }

}
