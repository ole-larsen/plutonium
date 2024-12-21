// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/token/ERC721/IERC721.sol";
import "../openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "../openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../openzeppelin/contracts/utils/math/SafeMath.sol";

contract NFTAuction is ERC721Holder, ReentrancyGuard {
    using SafeMath for uint256;

    address private _marketPlace;
    // Parameters of the auction. Times are either absolute unix timestamps (seconds since 1970-01-01)
    // or time periods in seconds.
    uint256 private _collectionId;
    uint256 private _itemId;
    address payable private _beneficiary;
    uint256 private _startPrice;
    // The _reservePrice variable is used to set the minimum price that the auction must reach in order for the
    // NFT to be sold. If the highest bid does not meet or exceed the reserve price by the end of the auction,
    // then the NFT will not be sold and the auction will end without a winner.
    //
    // This is a common feature in auctions, particularly for high-value items such as art or real estate,
    // where the seller wants to ensure that they receive a minimum amount for their item.
    // The reserve price helps to prevent the NFT from being sold at too low of a price,
    // while still allowing for competitive bidding to drive up the price.
    uint256 private _reservePrice;
    uint256 private _auctionStartTime;
    uint256 private _auctionEndTime;
    uint256 private _highestBid;
    address payable private _highestBidder;

    // Allowed withdrawals of previous bids
    mapping(address => uint256) pendingReturns;

    // Set to true at the end, disallows any change.
    // By default initialized to `false`.
    bool private _ended;
    bool private _started;

    // Declare events
    event AuctionStarted(
        address indexed beneficiary,
        uint256 collectionId,
        uint256 itemId,
        uint256 reservePrice,
        uint256 startPrice,
        uint256 auctionStartTime,
        uint256 auctionEndTime,
        bool started);

    event BidPlaced(address indexed bidder, uint256 amount);
    event AuctionEnded(address indexed winner, uint256 amount);

    // Errors that describe failures.

    // The triple-slash comments are so-called natspec
    // comments. They will be shown when the user
    // is asked to confirm a transaction or
    // when an error is displayed.

    /// There is already a higher or equal bid.
    error BidNotHighEnough(uint highestBid);
    /// The auction has not ended yet.
    error AuctionNotYetEnded();
    /// The function auctionEnd has already been called.
    error AuctionEndAlreadyCalled();

    /// Create a simple auction with `biddingTime`
    /// seconds bidding time on behalf of the
    /// beneficiary address `beneficiaryAddress`.
    constructor(
        address marketPlaceAddress,
        uint256 nftCollectionId,
        uint256 nftItemId,
        uint256 sPrice,
        uint256 rPrice,
        uint256 startTime,
        uint256 endTime
    ) {
        _marketPlace      = payable(marketPlaceAddress);
        _collectionId     = nftCollectionId;
        _itemId           = nftItemId;
        _beneficiary      = payable(msg.sender);
        _startPrice       = sPrice;
        _reservePrice     = rPrice;
        _auctionStartTime = startTime;
        _auctionEndTime   = endTime;
        _started          = false;
        _ended            = false;
    }

    function itemId() external view returns (uint256) {
        return _itemId;
    }

    function collectionId() external view returns (uint256) {
        return _collectionId;
    }

    function beneficiary() external view returns (address payable) {
        return _beneficiary;
    }

    function startPrice() external view returns (uint256) {
        return _startPrice;
    }

    function reservePrice() external view returns (uint256) {
        return _reservePrice;
    }

    function auctionStartTime() external view returns (uint256) {
        return _auctionStartTime;
    }

    function auctionEndTime() external view returns (uint256) {
        return _auctionEndTime;
    }

    function highestBid() external view returns (uint256) {
        return _highestBid;
    }

    function highestBidder() external view returns (address payable) {
        return _highestBidder;
    }

    function isStarted() external view returns (bool) {
        return _started;
    }

    function isEnded() external view returns (bool) {
        return _ended;
    }

    function startAuction() external {
        require(msg.sender == _beneficiary || msg.sender == _marketPlace, "only the seller can start the auction");
        require(block.timestamp >= _auctionStartTime, "auction has not started yet");
        require(block.timestamp <  _auctionEndTime, "auction has already ended");
        require(!_started, "auction must not be started");

        _started = true;

        emit AuctionStarted(_beneficiary, _collectionId, _itemId, _reservePrice, _startPrice, _auctionStartTime, _auctionEndTime, _started);
    }

    /// Bid on the auction with the value sent
    /// together with this transaction.
    /// The value will only be refunded if the
    /// auction is not won.
    function placeBid() external payable nonReentrant {
        require(_started, "auction should be started");
        require(!_ended, "auction should be not ended");

        require(block.timestamp >= _auctionStartTime, "auction has not started yet");
        require(block.timestamp <  _auctionEndTime, "auction has already ended");

        if (msg.value <= _reservePrice) {
            revert BidNotHighEnough(_highestBid);
        }

        if (msg.value <= _startPrice) {
            revert BidNotHighEnough(_highestBid);
        }

        if (msg.value <= _highestBid) {
            revert BidNotHighEnough(_highestBid);
        }

        if (_highestBid != 0 && _highestBidder != address(0)) {
            // Sending back the money by simply using
            // highestBidder.send(highestBid) is a security risk
            // because it could execute an untrusted contract.
            // It is always safer to let the recipients
            // withdraw their money themselves.
            // pendingReturns[_highestBidder] += _highestBid;
            // Refund the previous highest bidder
            _highestBidder.transfer(_highestBid);
        }

        _highestBid = msg.value;
        _highestBidder = payable(msg.sender);

        emit BidPlaced(msg.sender, msg.value);
    }

    // End the auction
    function endAuction() external {
        require(_started, "auction should be started");
        require(msg.sender == _beneficiary || msg.sender == _marketPlace, "only the seller can end the auction");
        require(block.timestamp >= _auctionStartTime, "auction has not started yet");
        require(block.timestamp <  _auctionEndTime, "auction has already ended");

        // It is a good guideline to structure functions that interact
        // with other contracts (i.e. they call functions or send Ether)
        // into three phases:
        // 1. checking conditions
        // 2. performing actions (potentially changing conditions)
        // 3. interacting with other contracts
        // If these phases are mixed up, the other contract could call
        // back into the current contract and modify the state or cause
        // effects (ether payout) to be performed multiple times.
        // If functions called internally include interaction with external
        // contracts, they also have to be considered interaction with
        // external contracts.
        /// End the auction and send the highest bid to the beneficiary.

        // Send the highest bid to the seller
        _beneficiary.transfer(_highestBid);
        _started = false;
        _ended = true;
        emit AuctionEnded(_highestBidder, _highestBid);
    }
}