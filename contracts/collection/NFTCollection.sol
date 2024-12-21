// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "../openzeppelin/contracts/utils/math/SafeMath.sol";
import "../auction/NFTAuction.sol";

contract NFTCollection is ERC721, ERC721Enumerable {
  using SafeMath for uint256;
  string[]   public tokenURIs;
  mapping(string => bool) _tokenURIExists;
  mapping(uint => string) _tokenIdToTokenURI;

  constructor(string memory _name, string memory _symbol) ERC721(_name, _symbol){}

  function _beforeTokenTransfer(address from, address to, uint256 tokenId, uint256 batchSize) internal override(ERC721, ERC721Enumerable) {
    super._beforeTokenTransfer(from, to, tokenId, batchSize);
  }

  function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721Enumerable) returns (bool) {
    return super.supportsInterface(interfaceId);
  }

  function tokenURI(uint256 tokenId) public override view returns (string memory) {
    require(_exists(tokenId), 'ERC721Metadata: URI query for nonexistent token');
    return _tokenIdToTokenURI[tokenId];
  }

  function safeMint(string memory _tokenURI) public {
    require(!_tokenURIExists[_tokenURI], 'The token URI should be unique');
    tokenURIs.push(_tokenURI);
    uint _id = totalSupply().add(1);
    _tokenIdToTokenURI[_id] = _tokenURI;
    _safeMint(msg.sender, _id);
    _tokenURIExists[_tokenURI] = true;
  }

  /**
    * @dev Mints Multiple NFT
    */
  function mint(string memory _tokenURI, uint256 _quantity) public payable {
    require(!_tokenURIExists[_tokenURI], 'The token URI should be unique');
    require(_quantity > 0, "quantity cannot be 0");

    for (uint i = 0; i < _quantity; i++) {
      uint _id = totalSupply().add(1);
      _tokenIdToTokenURI[_id] = _tokenURI;
      _safeMint(msg.sender, _id);
    }

    tokenURIs.push(_tokenURI);
    _tokenURIExists[_tokenURI] = true;
  }
}