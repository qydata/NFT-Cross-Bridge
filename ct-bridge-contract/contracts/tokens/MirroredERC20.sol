pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MirroredERC20 is Initializable, ERC20Upgradeable, OwnableUpgradeable, PausableUpgradeable, ERC20BurnableUpgradeable {
  uint8 private __decimals;
  function initialize(string memory _name, string memory _symbol, uint8 _decimals) public initializer {
    __ERC20_init(_name,  _symbol);
    __decimals = _decimals;
    __Ownable_init();
    __Pausable_init();
    __ERC20Burnable_init();
  }

  function decimals() public view virtual override returns (uint8) {
    return __decimals;
  }

  function pause() public onlyOwner {
    _pause();
  }

  function unpause() public onlyOwner {
    _unpause();
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
  
  function _beforeTokenTransfer(address from, address to, uint256 amount)
  internal
  whenNotPaused
  override
  {
    super._beforeTokenTransfer(from, to, amount);
  }
}