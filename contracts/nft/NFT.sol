// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "../openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "../openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract NFT is ERC721URIStorage {
    uint public count;

    constructor(string memory _name, string memory _symbol) ERC721(_name, _symbol) {}

    function mint(string memory _tokenURI) external returns(uint) {
        count++;
        _safeMint(msg.sender, count);
        _setTokenURI(count, _tokenURI);
        return count;
    }
}

