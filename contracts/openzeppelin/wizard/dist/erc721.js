"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildERC721 = exports.isAccessControlRequired = exports.printERC721 = exports.defaults = void 0;
const contract_1 = require("./contract");
const set_access_control_1 = require("./set-access-control");
const add_pausable_1 = require("./add-pausable");
const common_functions_1 = require("./common-functions");
const define_functions_1 = require("./utils/define-functions");
const common_options_1 = require("./common-options");
const set_upgradeable_1 = require("./set-upgradeable");
const set_info_1 = require("./set-info");
const print_1 = require("./print");
exports.defaults = {
    name: 'MyToken',
    symbol: 'MTK',
    baseUri: '',
    enumerable: false,
    uriStorage: false,
    burnable: false,
    pausable: false,
    mintable: false,
    incremental: false,
    votes: false,
    access: common_options_1.defaults.access,
    upgradeable: common_options_1.defaults.upgradeable,
    info: common_options_1.defaults.info
};
function withDefaults(opts) {
    var _a, _b, _c, _d, _e, _f, _g, _h;
    return {
        ...opts,
        ...(0, common_options_1.withCommonDefaults)(opts),
        baseUri: (_a = opts.baseUri) !== null && _a !== void 0 ? _a : exports.defaults.baseUri,
        enumerable: (_b = opts.enumerable) !== null && _b !== void 0 ? _b : exports.defaults.enumerable,
        uriStorage: (_c = opts.uriStorage) !== null && _c !== void 0 ? _c : exports.defaults.uriStorage,
        burnable: (_d = opts.burnable) !== null && _d !== void 0 ? _d : exports.defaults.burnable,
        pausable: (_e = opts.pausable) !== null && _e !== void 0 ? _e : exports.defaults.pausable,
        mintable: (_f = opts.mintable) !== null && _f !== void 0 ? _f : exports.defaults.mintable,
        incremental: (_g = opts.incremental) !== null && _g !== void 0 ? _g : exports.defaults.incremental,
        votes: (_h = opts.votes) !== null && _h !== void 0 ? _h : exports.defaults.votes,
    };
}
function printERC721(opts = exports.defaults) {
    return (0, print_1.printContract)(buildERC721(opts));
}
exports.printERC721 = printERC721;
function isAccessControlRequired(opts) {
    return opts.mintable || opts.pausable || opts.upgradeable === 'uups';
}
exports.isAccessControlRequired = isAccessControlRequired;
function buildERC721(opts) {
    const allOpts = withDefaults(opts);
    const c = new contract_1.ContractBuilder(allOpts.name);
    const { access, upgradeable, info } = allOpts;
    addBase(c, allOpts.name, allOpts.symbol);
    if (allOpts.baseUri) {
        addBaseURI(c, allOpts.baseUri);
    }
    if (allOpts.enumerable) {
        addEnumerable(c);
    }
    if (allOpts.uriStorage) {
        addURIStorage(c);
    }
    if (allOpts.pausable) {
        (0, add_pausable_1.addPausable)(c, access, [functions._beforeTokenTransfer]);
    }
    if (allOpts.burnable) {
        addBurnable(c);
    }
    if (allOpts.mintable) {
        addMintable(c, access, allOpts.incremental, allOpts.uriStorage);
    }
    if (allOpts.votes) {
        addVotes(c, allOpts.name);
    }
    (0, set_access_control_1.setAccessControl)(c, access);
    (0, set_upgradeable_1.setUpgradeable)(c, upgradeable, access);
    (0, set_info_1.setInfo)(c, info);
    return c;
}
exports.buildERC721 = buildERC721;
function addBase(c, name, symbol) {
    c.addParent({
        name: 'ERC721',
        path: '@openzeppelin/contracts/token/ERC721/ERC721.sol',
    }, [name, symbol]);
    c.addOverride('ERC721', functions._beforeTokenTransfer);
    c.addOverride('ERC721', functions._afterTokenTransfer);
    c.addOverride('ERC721', functions._burn);
    c.addOverride('ERC721', functions.tokenURI);
    c.addOverride('ERC721', common_functions_1.supportsInterface);
}
function addBaseURI(c, baseUri) {
    c.addOverride('ERC721', functions._baseURI);
    c.setFunctionBody([`return ${JSON.stringify(baseUri)};`], functions._baseURI);
}
function addEnumerable(c) {
    c.addParent({
        name: 'ERC721Enumerable',
        path: '@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol',
    });
    c.addOverride('ERC721Enumerable', functions._beforeTokenTransfer);
    c.addOverride('ERC721Enumerable', common_functions_1.supportsInterface);
}
function addURIStorage(c) {
    c.addParent({
        name: 'ERC721URIStorage',
        path: '@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol',
    });
    c.addOverride('ERC721URIStorage', functions._burn);
    c.addOverride('ERC721URIStorage', functions.tokenURI);
}
function addBurnable(c) {
    c.addParent({
        name: 'ERC721Burnable',
        path: '@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol',
    });
}
function addMintable(c, access, incremental = false, uriStorage = false) {
    const fn = getMintFunction(incremental, uriStorage);
    (0, set_access_control_1.requireAccessControl)(c, fn, access, 'MINTER');
    if (incremental) {
        c.addUsing({
            name: 'Counters',
            path: '@openzeppelin/contracts/utils/Counters.sol',
        }, 'Counters.Counter');
        c.addVariable('Counters.Counter private _tokenIdCounter;');
        c.addFunctionCode('uint256 tokenId = _tokenIdCounter.current();', fn);
        c.addFunctionCode('_tokenIdCounter.increment();', fn);
        c.addFunctionCode('_safeMint(to, tokenId);', fn);
    }
    else {
        c.addFunctionCode('_safeMint(to, tokenId);', fn);
    }
    if (uriStorage) {
        c.addFunctionCode('_setTokenURI(tokenId, uri);', fn);
    }
}
function addVotes(c, name) {
    c.addParent({
        name: 'EIP712',
        path: '@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol',
    }, [name, "1"]);
    c.addParent({
        name: 'ERC721Votes',
        path: '@openzeppelin/contracts/token/ERC721/extensions/draft-ERC721Votes.sol',
    });
    c.addOverride('ERC721Votes', functions._afterTokenTransfer);
}
const functions = (0, define_functions_1.defineFunctions)({
    _beforeTokenTransfer: {
        kind: 'internal',
        args: [
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'tokenId', type: 'uint256' },
        ],
    },
    _afterTokenTransfer: {
        kind: 'internal',
        args: [
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'tokenId', type: 'uint256' },
        ],
    },
    _burn: {
        kind: 'internal',
        args: [
            { name: 'tokenId', type: 'uint256' },
        ],
    },
    tokenURI: {
        kind: 'public',
        args: [
            { name: 'tokenId', type: 'uint256' },
        ],
        returns: ['string memory'],
        mutability: 'view',
    },
    _baseURI: {
        kind: 'internal',
        args: [],
        returns: ['string memory'],
        mutability: 'pure',
    },
});
function getMintFunction(incremental, uriStorage) {
    const fn = {
        name: 'safeMint',
        kind: 'public',
        args: [
            { name: 'to', type: 'address' },
        ],
    };
    if (!incremental) {
        fn.args.push({ name: 'tokenId', type: 'uint256' });
    }
    if (uriStorage) {
        fn.args.push({ name: 'uri', type: 'string memory' });
    }
    return fn;
}
//# sourceMappingURL=erc721.js.map