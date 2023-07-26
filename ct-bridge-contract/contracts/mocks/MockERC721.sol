pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MockERC721 is ERC721, ERC721URIStorage, Ownable {
  string public baseURI;

  constructor(string memory name, string memory symbol) ERC721(name, symbol) {}

  function setTokenURI(uint256 tokenId, string calldata _tokenURI) public onlyOwner {
    _setTokenURI(tokenId, _tokenURI);
  }

  function safeMint(address to, uint256 tokenId) public onlyOwner {
    _safeMint(to, tokenId);
  }

  function setBaseURI(string memory baseURI_) external {
    baseURI = baseURI_;
  }

  // The following functions are overrides required by Solidity.

  function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
    super._burn(tokenId);
  }

  function _baseURI() internal view override returns (string memory) {
    return baseURI;
  }

  function tokenURI(uint256 tokenId)
  public
  view
  override(ERC721, ERC721URIStorage)
  returns (string memory)
  {
    return super.tokenURI(tokenId);
  }
}
