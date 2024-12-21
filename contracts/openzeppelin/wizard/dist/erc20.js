"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.premintPattern = exports.buildERC20 = exports.isAccessControlRequired = exports.printERC20 = exports.defaults = void 0;
const contract_1 = require("./contract");
const set_access_control_1 = require("./set-access-control");
const add_pausable_1 = require("./add-pausable");
const define_functions_1 = require("./utils/define-functions");
const common_options_1 = require("./common-options");
const set_upgradeable_1 = require("./set-upgradeable");
const set_info_1 = require("./set-info");
const print_1 = require("./print");
exports.defaults = {
    name: 'MyToken',
    symbol: 'MTK',
    burnable: false,
    snapshots: false,
    pausable: false,
    premint: '0',
    mintable: false,
    permit: false,
    votes: false,
    flashmint: false,
    access: common_options_1.defaults.access,
    upgradeable: common_options_1.defaults.upgradeable,
    info: common_options_1.defaults.info,
};
function withDefaults(opts) {
    var _a, _b, _c, _d, _e, _f, _g;
    return {
        ...opts,
        ...(0, common_options_1.withCommonDefaults)(opts),
        burnable: (_a = opts.burnable) !== null && _a !== void 0 ? _a : exports.defaults.burnable,
        snapshots: (_b = opts.snapshots) !== null && _b !== void 0 ? _b : exports.defaults.snapshots,
        pausable: (_c = opts.pausable) !== null && _c !== void 0 ? _c : exports.defaults.pausable,
        premint: opts.premint || exports.defaults.premint,
        mintable: (_d = opts.mintable) !== null && _d !== void 0 ? _d : exports.defaults.mintable,
        permit: (_e = opts.permit) !== null && _e !== void 0 ? _e : exports.defaults.permit,
        votes: (_f = opts.votes) !== null && _f !== void 0 ? _f : exports.defaults.votes,
        flashmint: (_g = opts.flashmint) !== null && _g !== void 0 ? _g : exports.defaults.flashmint,
    };
}
function printERC20(opts = exports.defaults) {
    return (0, print_1.printContract)(buildERC20(opts));
}
exports.printERC20 = printERC20;
function isAccessControlRequired(opts) {
    return opts.mintable || opts.pausable || opts.snapshots || opts.upgradeable === 'uups';
}
exports.isAccessControlRequired = isAccessControlRequired;
function buildERC20(opts) {
    const allOpts = withDefaults(opts);
    const c = new contract_1.ContractBuilder(allOpts.name);
    const { access, upgradeable, info } = allOpts;
    addBase(c, allOpts.name, allOpts.symbol);
    if (allOpts.burnable) {
        addBurnable(c);
    }
    if (allOpts.snapshots) {
        addSnapshot(c, access);
    }
    if (allOpts.pausable) {
        (0, add_pausable_1.addPausable)(c, access, [functions._beforeTokenTransfer]);
    }
    if (allOpts.premint) {
        addPremint(c, allOpts.premint);
    }
    if (allOpts.mintable) {
        addMintable(c, access);
    }
    // Note: Votes requires Permit
    if (allOpts.permit || allOpts.votes) {
        addPermit(c, allOpts.name);
    }
    if (allOpts.votes) {
        addVotes(c);
    }
    if (allOpts.flashmint) {
        addFlashMint(c);
    }
    (0, set_access_control_1.setAccessControl)(c, access);
    (0, set_upgradeable_1.setUpgradeable)(c, upgradeable, access);
    (0, set_info_1.setInfo)(c, info);
    return c;
}
exports.buildERC20 = buildERC20;
function addBase(c, name, symbol) {
    c.addParent({
        name: 'ERC20',
        path: '@openzeppelin/contracts/token/ERC20/ERC20.sol',
    }, [name, symbol]);
    c.addOverride('ERC20', functions._beforeTokenTransfer);
    c.addOverride('ERC20', functions._afterTokenTransfer);
    c.addOverride('ERC20', functions._mint);
    c.addOverride('ERC20', functions._burn);
}
function addBurnable(c) {
    c.addParent({
        name: 'ERC20Burnable',
        path: '@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol',
    });
}
function addSnapshot(c, access) {
    c.addParent({
        name: 'ERC20Snapshot',
        path: '@openzeppelin/contracts/token/ERC20/extensions/ERC20Snapshot.sol',
    });
    c.addOverride('ERC20Snapshot', functions._beforeTokenTransfer);
    (0, set_access_control_1.requireAccessControl)(c, functions.snapshot, access, 'SNAPSHOT');
    c.addFunctionCode('_snapshot();', functions.snapshot);
}
exports.premintPattern = /^(\d*)(?:\.(\d+))?(?:e(\d+))?$/;
function addPremint(c, amount) {
    var _a, _b, _c, _d, _e;
    const m = amount.match(exports.premintPattern);
    if (m) {
        const integer = (_b = (_a = m[1]) === null || _a === void 0 ? void 0 : _a.replace(/^0+/, '')) !== null && _b !== void 0 ? _b : '';
        const decimals = (_d = (_c = m[2]) === null || _c === void 0 ? void 0 : _c.replace(/0+$/, '')) !== null && _d !== void 0 ? _d : '';
        const exponent = Number((_e = m[3]) !== null && _e !== void 0 ? _e : 0);
        if (Number(integer + decimals) > 0) {
            const decimalPlace = decimals.length - exponent;
            const zeroes = new Array(Math.max(0, -decimalPlace)).fill('0').join('');
            const units = integer + decimals + zeroes;
            const exp = decimalPlace <= 0 ? 'decimals()' : `(decimals() - ${decimalPlace})`;
            c.addConstructorCode(`_mint(msg.sender, ${units} * 10 ** ${exp});`);
        }
    }
}
function addMintable(c, access) {
    (0, set_access_control_1.requireAccessControl)(c, functions.mint, access, 'MINTER');
    c.addFunctionCode('_mint(to, amount);', functions.mint);
}
function addPermit(c, name) {
    c.addParent({
        name: 'ERC20Permit',
        path: '@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol',
    }, [name]);
}
function addVotes(c) {
    if (!c.parents.some(p => p.contract.name === 'ERC20Permit')) {
        throw new Error('Missing ERC20Permit requirement for ERC20Votes');
    }
    c.addParent({
        name: 'ERC20Votes',
        path: '@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol',
    });
    c.addOverride('ERC20Votes', functions._mint);
    c.addOverride('ERC20Votes', functions._burn);
    c.addOverride('ERC20Votes', functions._afterTokenTransfer);
}
function addFlashMint(c) {
    c.addParent({
        name: 'ERC20FlashMint',
        path: '@openzeppelin/contracts/token/ERC20/extensions/ERC20FlashMint.sol',
    });
}
const functions = (0, define_functions_1.defineFunctions)({
    _beforeTokenTransfer: {
        kind: 'internal',
        args: [
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    },
    _afterTokenTransfer: {
        kind: 'internal',
        args: [
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    },
    _burn: {
        kind: 'internal',
        args: [
            { name: 'account', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    },
    _mint: {
        kind: 'internal',
        args: [
            { name: 'to', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    },
    mint: {
        kind: 'public',
        args: [
            { name: 'to', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    },
    pause: {
        kind: 'public',
        args: [],
    },
    unpause: {
        kind: 'public',
        args: [],
    },
    snapshot: {
        kind: 'public',
        args: [],
    },
});
//# sourceMappingURL=erc20.js.map