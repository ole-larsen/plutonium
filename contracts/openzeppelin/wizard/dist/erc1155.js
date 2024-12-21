"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildERC1155 = exports.isAccessControlRequired = exports.printERC1155 = exports.defaults = void 0;
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
    uri: '',
    burnable: false,
    pausable: false,
    mintable: false,
    supply: false,
    updatableUri: true,
    access: false,
    upgradeable: common_options_1.defaults.upgradeable,
    info: common_options_1.defaults.info
};
function withDefaults(opts) {
    var _a, _b, _c, _d, _e;
    return {
        ...opts,
        ...(0, common_options_1.withCommonDefaults)(opts),
        burnable: (_a = opts.burnable) !== null && _a !== void 0 ? _a : exports.defaults.burnable,
        pausable: (_b = opts.pausable) !== null && _b !== void 0 ? _b : exports.defaults.pausable,
        mintable: (_c = opts.mintable) !== null && _c !== void 0 ? _c : exports.defaults.mintable,
        supply: (_d = opts.supply) !== null && _d !== void 0 ? _d : exports.defaults.supply,
        updatableUri: (_e = opts.updatableUri) !== null && _e !== void 0 ? _e : exports.defaults.updatableUri,
    };
}
function printERC1155(opts = exports.defaults) {
    return (0, print_1.printContract)(buildERC1155(opts));
}
exports.printERC1155 = printERC1155;
function isAccessControlRequired(opts) {
    return opts.mintable || opts.pausable || opts.updatableUri !== false || opts.upgradeable === 'uups';
}
exports.isAccessControlRequired = isAccessControlRequired;
function buildERC1155(opts) {
    const allOpts = withDefaults(opts);
    const c = new contract_1.ContractBuilder(allOpts.name);
    const { access, upgradeable, info } = allOpts;
    addBase(c, allOpts.uri);
    if (allOpts.updatableUri) {
        addSetUri(c, access);
    }
    if (allOpts.pausable) {
        (0, add_pausable_1.addPausable)(c, access, [functions._beforeTokenTransfer]);
    }
    if (allOpts.burnable) {
        addBurnable(c);
    }
    if (allOpts.mintable) {
        addMintable(c, access);
    }
    if (allOpts.supply) {
        addSupply(c);
    }
    (0, set_access_control_1.setAccessControl)(c, access);
    (0, set_upgradeable_1.setUpgradeable)(c, upgradeable, access);
    (0, set_info_1.setInfo)(c, info);
    return c;
}
exports.buildERC1155 = buildERC1155;
function addBase(c, uri) {
    c.addParent({
        name: 'ERC1155',
        path: '@openzeppelin/contracts/token/ERC1155/ERC1155.sol',
    }, [uri]);
    c.addOverride('ERC1155', functions._beforeTokenTransfer);
    c.addOverride('ERC1155', common_functions_1.supportsInterface);
}
function addBurnable(c) {
    c.addParent({
        name: 'ERC1155Burnable',
        path: '@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Burnable.sol',
    });
}
function addMintable(c, access) {
    (0, set_access_control_1.requireAccessControl)(c, functions.mint, access, 'MINTER');
    (0, set_access_control_1.requireAccessControl)(c, functions.mintBatch, access, 'MINTER');
    c.addFunctionCode('_mint(account, id, amount, data);', functions.mint);
    c.addFunctionCode('_mintBatch(to, ids, amounts, data);', functions.mintBatch);
}
function addSetUri(c, access) {
    (0, set_access_control_1.requireAccessControl)(c, functions.setURI, access, 'URI_SETTER');
    c.addFunctionCode('_setURI(newuri);', functions.setURI);
}
function addSupply(c) {
    c.addParent({
        name: 'ERC1155Supply',
        path: '@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol',
    });
    c.addOverride('ERC1155Supply', functions._beforeTokenTransfer);
}
const functions = (0, define_functions_1.defineFunctions)({
    _beforeTokenTransfer: {
        kind: 'internal',
        args: [
            { name: 'operator', type: 'address' },
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'ids', type: 'uint256[] memory' },
            { name: 'amounts', type: 'uint256[] memory' },
            { name: 'data', type: 'bytes memory' },
        ],
    },
    setURI: {
        kind: 'public',
        args: [
            { name: 'newuri', type: 'string memory' },
        ],
    },
    mint: {
        kind: 'public',
        args: [
            { name: 'account', type: 'address' },
            { name: 'id', type: 'uint256' },
            { name: 'amount', type: 'uint256' },
            { name: 'data', type: 'bytes memory' },
        ],
    },
    mintBatch: {
        kind: 'public',
        args: [
            { name: 'to', type: 'address' },
            { name: 'ids', type: 'uint256[] memory' },
            { name: 'amounts', type: 'uint256[] memory' },
            { name: 'data', type: 'bytes memory' },
        ],
    },
});
//# sourceMappingURL=erc1155.js.map