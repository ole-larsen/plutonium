// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "../openzeppelin/contracts/token/ERC721/IERC721.sol";
import "../openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../openzeppelin/contracts/utils/math/SafeMath.sol";

import "../collection/NFTCollection.sol";
// import "../auction/NFTAuction.sol";

contract Marketplace is ReentrancyGuard {
    using SafeMath for uint256;

    address payable private immutable _owner;
    string  private _name;
    uint256 private _fee;
    uint256 private _collectionCounter;

    mapping(address => uint256) private _userFunds;

    struct Collection {
        string name;
        string symbol;
        string description;
        NFTCollection nftCollection;
        uint256 fee;
        address payable owner;
        address payable creator;
        bool isApproved;
        bool isLocked;
        uint256 maxItems;
    }

    mapping(string => uint256)     private _collectionsNames;
    mapping(uint256 => Collection) private _collections;
    mapping(address => bool)       private _collectionAddressExists;

    event CreateCollection(
        uint256 id,
        string name,
        string symbol,
        uint256 fee,
        uint256 maxItems,
        address indexed collection,
        address indexed creator,
        address indexed owner
    );

    event EditCollection(
        uint256 id,
        string name,
        string symbol,
        uint256 fee,
        uint256 maxItems,
        address indexed collection,
        address indexed creator,
        address indexed owner,
        bool isApproved,
        bool isLocked
    );

    constructor(string memory marketName, uint256 marketFee) {
        _name = marketName;
        _fee = marketFee;
        _owner = payable(msg.sender);
    }

    function getName() public view returns (string memory) {
        return _name;
    }

    function getFee() public view returns (uint256) {
        return _fee;
    }

    function getOwner() public view returns (address) {
        return _owner;
    }

    function getUserFunds(address userAddress) public view returns (uint256) {
        return _userFunds[userAddress];
    }

    function setName(string memory marketName) public {
        require(msg.sender == _owner);
        require(bytes(marketName).length != 0);
        _name = marketName;
    }

    function setFee(uint256 marketFee) public {
        require(msg.sender == _owner);
        require(marketFee > 0);
        _fee = marketFee;
    }

    function createCollection(
        string memory name,
        string memory symbol,
        string memory description,
        uint256 fee,
        uint256 maxItemsInCollection,
        address nftCollection,
        address payable owner
    ) external nonReentrant {
        require(bytes(name).length != 0);
        require(bytes(symbol).length != 0);
        require(_collectionsNames[name] == 0);
        require(!_collectionAddressExists[nftCollection]);
        require(maxItemsInCollection > 0);
        _collectionCounter++;

        Collection storage newCollection = _collections[_collectionCounter];
        
        newCollection.name = name;
        newCollection.symbol = symbol;
        newCollection.description = description;
        newCollection.nftCollection = NFTCollection(nftCollection);
        newCollection.fee = fee;
        newCollection.owner = payable(owner);
        newCollection.creator = payable(msg.sender);
        newCollection.isApproved = true;
        newCollection.isLocked = false;
        newCollection.maxItems = maxItemsInCollection;

         _collectionsNames[name] = _collectionCounter;
        _collectionAddressExists[nftCollection] = true;

        emit CreateCollection(
            _collectionCounter,
            name,
            symbol,
            fee,
            maxItemsInCollection,
            nftCollection,
            msg.sender,
            owner
        );
    }

    function editCollection(
        uint256 id,
        string memory name,
        string memory symbol,
        string memory description,
        uint256 fee,
        uint256 maxItemsInCollection,
        address nftCollection,
        address payable owner,
        bool isApproved,
        bool isLocked
    ) external nonReentrant {
        require(msg.sender == owner);
        Collection storage collection = _collections[id];
        require(bytes(collection.name).length != 0);
        require(maxItemsInCollection >= collection.maxItems);

        if (bytes(collection.name).length != bytes(name).length || keccak256(bytes(collection.name)) != keccak256(bytes(name))) {
            require(_collectionsNames[name] > 0);
            collection.name = name;
        }
        if (bytes(collection.symbol).length != bytes(symbol).length || keccak256(bytes(collection.symbol)) != keccak256(bytes(symbol))) {
            collection.symbol = symbol;
        }
        if (bytes(collection.description).length != bytes(description).length || keccak256(bytes(collection.description)) != keccak256(bytes(description))) {
            collection.description = description;
        }
        if (collection.fee != fee) {
            collection.fee = fee;
        }
        if (collection.isApproved != isApproved) {
            collection.isApproved = isApproved;
        }
        if (collection.isLocked != isLocked) {
            collection.isLocked = isLocked;
        }
        if (collection.maxItems != maxItemsInCollection) {
            collection.maxItems = maxItemsInCollection;
        }
        emit EditCollection(id, name, symbol, fee, maxItemsInCollection, nftCollection, msg.sender, owner, isApproved, isLocked);
    }

    function getCollectionCounter() public view returns (uint256) {
        return _collectionCounter;
    }

    function getCollection(uint256 collectionId) public view returns (Collection memory) {
        require(collectionId > 0 && collectionId <= _collectionCounter);
        return _collections[collectionId];
    }

    function getCollectionIdByName(string memory name) public view returns (uint256) {
        require(bytes(name).length != 0);
        return _collectionsNames[name];
    }
    
    mapping(uint256 => mapping (uint256 => Collectible)) private _collectibles;
    mapping(uint256 => uint256) private _collectionCollectibleCount;

    struct Collectible {
        uint256           id;
        uint256           collectionId;
        uint256[]         tokenIds;
        address payable[] owners;
        address payable   creator;
        bool              isAuction;
        uint256           price;
        bool[]            fulfilled;
        bool              isLocked;
    }

    event CreateCollectible(uint256 id, uint256 collectionId, uint256[] tokenIds, address payable[] owners, address indexed creator, 
        bool isAuction, uint256 price);

    modifier validCollectibleParams(uint256[] memory _tokenIds, uint256 _collectionId) {
        require(_collectionId > 0 && _collectionId <= _collectionCounter);
        require(_tokenIds.length > 0);
        _;
    }

    function createCollectible(uint256[] memory _tokenIds, uint256 _collectionId, bool _isAuction, uint256 _price) external nonReentrant validCollectibleParams(_tokenIds, _collectionId){
        require(msg.sender != address(0));
        require(address(this) != address(0));

        Collection storage collection = _collections[_collectionId];

        require(collection.maxItems >= _tokenIds.length);
        require(collection.isApproved);
        require(!collection.isLocked);

        bool[] memory _fulfilled = new bool[](_tokenIds.length);
        address payable[] memory _owners = new address payable[](_tokenIds.length);

        for (uint i = 0; i < _tokenIds.length; i++) {
            collection.nftCollection.approve(address(this), _tokenIds[i]);
            collection.nftCollection.transferFrom(msg.sender, address(this), _tokenIds[i]);
            require(collection.nftCollection.ownerOf(_tokenIds[i]) == address(this));
            _fulfilled[i] = false;
            _owners[i] = payable(msg.sender);
        }
        _collectionCollectibleCount[_collectionId]++;

        _collectibles[_collectionId][_collectionCollectibleCount[_collectionId]] = Collectible(
            _collectionCollectibleCount[_collectionId],
            _collectionId,
            _tokenIds,
            _owners,
            payable(msg.sender),
            _isAuction,
            _price,
            _fulfilled,
            false
        );

        emit CreateCollectible(
            _collectionCollectibleCount[_collectionId],
            _collectionId,
            _tokenIds,
            _owners,
            msg.sender,
            _isAuction,
            _price
        );
    }

    function getCollectibleCount(uint256 _collectionId) public view returns (uint) {
        require(_collectionId > 0 && _collectionId <= _collectionCounter);
        return _collectionCollectibleCount[_collectionId];
    }

    function getCollectible(uint256 _collectionId, uint256 _collectibleId) public view returns (Collectible memory) {
        require(_collectionId > 0 && _collectionId <= _collectionCounter);
        require(_collectibleId > 0 && _collectibleId <= _collectionCollectibleCount[_collectionId]);
        return _collectibles[_collectionId][_collectibleId];
    }

    
    modifier validateCollectibleParams(uint256 _collectionId, uint256 _collectibleId) {
        require(_collectionId > 0 && _collectionId <= _collectionCounter);
        require(_collectibleId > 0 && _collectibleId <= _collectionCollectibleCount[_collectionId]);
        _;
    }

    event Buy(
        uint256 id,
        uint256 collectionId,
        uint256[] tokenIds,
        uint256 price,
        address indexed buyer,
        uint256 quantity
    );

    function checkFreeTokens(bool[] memory _fulfilled) private pure returns (uint) {
        uint result = 0;
        for (uint i = 0; i < _fulfilled.length; i++) {
            if (_fulfilled[i] == false) {
                result++;
            }
        }
        return result;
    }

    function buy(uint256 _collectionId, uint256 _collectibleId, uint256 _quantity) external payable nonReentrant validateCollectibleParams(_collectionId, _collectibleId) {
        Collection  storage collection = _collections[_collectionId];
        Collectible storage collectible = _collectibles[_collectionId][_collectibleId];
        require(_quantity > 0 && _quantity <= collectible.tokenIds.length);
        require(collectible.id == _collectibleId);
        require(!collectible.isAuction);
        require(!collectible.isLocked);
        uint256 percent = (collectible.price.mul(collection.fee).div(100)).div(1e18);
        uint256 totalPrice = collectible.price.add(percent);
        require(msg.value == totalPrice.mul(_quantity));
        require(checkFreeTokens(collectible.fulfilled) >= _quantity);
        uint256 quantity = _quantity;
        for (uint i = 0; i < collectible.tokenIds.length; i++) {
            if (!collectible.fulfilled[i] && collectible.owners[i] != msg.sender) {
                if (quantity > 0) {
                    _userFunds[collectible.owners[i]] += totalPrice.sub(percent);
                    collectible.fulfilled[i] = true;
                    collection.nftCollection.transferFrom(address(this), msg.sender, collectible.tokenIds[i]);
                    collectible.owners[i] = payable(msg.sender);
                    quantity--;
                }
            }
        }
        if (quantity != _quantity) {
            _owner.transfer(percent.mul(_quantity));
        } else {
            revert("Trying to buy from itself");
        }
        emit Buy(
            collectible.id,
            collectible.collectionId,
            collectible.tokenIds,
            collectible.price,
            msg.sender,
            _quantity
        );
    }
    
    event Sell(
        uint256 id,
        uint256 collectionId,
        uint256[] tokenIds,
        address indexed creator,
        address indexed seller,
        uint256 quantity
    );
    
    function sell(uint256 _collectionId, uint256 _collectibleId, uint256 _quantity) external payable nonReentrant validateCollectibleParams(_collectionId, _collectibleId) {
        Collection  storage collection = _collections[_collectionId];
        Collectible storage collectible = _collectibles[_collectionId][_collectibleId];
        require(collectible.id == _collectibleId);
        require(_quantity > 0 && _quantity <= collectible.tokenIds.length);
        require(!collectible.isAuction);
        require(!collectible.isLocked);
        require(collectible.fulfilled.length - checkFreeTokens(collectible.fulfilled) >= _quantity);
        uint256 quantity = _quantity;
        for (uint i = 0; i < collectible.tokenIds.length; i++) {
            if (collectible.fulfilled[i] && collectible.owners[i] == msg.sender) {
                if (quantity > 0) {
                    collectible.fulfilled[i] = false;
                    collection.nftCollection.transferFrom(msg.sender, address(this), collectible.tokenIds[i]);
                    quantity--;
                }
            }
        }
        if (quantity == _quantity) {
            revert("Can sell only tokens that belongs to owner");
        } 
        emit Sell(
            collectible.id,
            collectible.collectionId,
            collectible.tokenIds,
            collectible.creator,
            msg.sender,
            _quantity
        );
    }

    event EditCollectible(uint256 collectionId, uint256 collectibleId, bool isLocked, bool isAuction, address owner);

    function editCollectible(uint256 _collectionId, uint256 _collectibleId, uint256 _price, bool _isLocked, bool _isAuction) external validateCollectibleParams(_collectionId, _collectibleId) {
        Collectible storage collectible = _collectibles[_collectionId][_collectibleId];
        require(collectible.id == _collectibleId);
        require(!collectible.isAuction);
        require(msg.sender == _owner);
        for (uint i = 0; i < collectible.tokenIds.length; i++) {
            require(!collectible.fulfilled[i]);
        }
        require(_price > 0);
        if (_price != collectible.price) {
            collectible.price = _price;
        }
        if (_isLocked != collectible.isLocked) {
            collectible.isLocked = _isLocked;
        }

        if (_isAuction != collectible.isAuction) {
            collectible.isAuction = _isAuction;
        }
        emit EditCollectible(_collectionId, _collectibleId, _isLocked, _isAuction, msg.sender);
    }

    event ClaimFunds(address user, uint256 amount);

    function claimFunds() external payable nonReentrant {
        require(_userFunds[msg.sender] > 0);
        payable(msg.sender).transfer(_userFunds[msg.sender]);
        emit ClaimFunds(msg.sender, _userFunds[msg.sender]);
        _userFunds[msg.sender] = 0;
    }

    function startAuction(uint256 _collectionId, uint256 _collectibleId) external validateCollectibleParams(_collectionId, _collectibleId) {
        Collectible storage collectible = _collectibles[_collectionId][_collectibleId];
        require(collectible.id == _collectibleId, "collectible should exist");
        require(collectible.isAuction, "can start only auction collectible");
        require(msg.sender == _owner, "only market owner can start auction");
        collectible.isAuction = true;
        // require(collectible.nftAuction.itemId() == _collectibleId, "item id must be same");
        // require(collectible.nftAuction.collectionId() == collectible.collectionId, "collection id must be same");
        // require(msg.sender == collectible.owner, "only the collectible owner can start the auction");

        // collectible.nftAuction.startAuction();
    }

    function endAuction(uint256 _collectionId, uint256 _collectibleId) external validateCollectibleParams(_collectionId, _collectibleId) {
        Collectible storage collectible = _collectibles[_collectionId][_collectibleId];
        require(collectible.id == _collectibleId, "collectible should exist");
        require(collectible.isAuction, "can stop only auction collectible");
        require(msg.sender == _owner, "only market owner can stop auction");
        collectible.isAuction = false;
        // require(collectible.nftAuction.tokenId() == collectible.tokenId, "token id must be same");
        // require(collectible.nftAuction.collectionId() == collectible.collectionId, "collection id must be same");
        // require(msg.sender == collectible.owner, "only the collectible owner can stop the auction");
        // collectible.nftAuction.endAuction();
        // transfer just token to highest bidder
        // collection.nftCollection.transferFrom(address(this), collectible.nftAuction.highestBidder(), collectible.tokenId);
    }
    // Fallback: reverts if Ether is sent to this smart-contract by mistake
    fallback () external {
        revert();
    }
}