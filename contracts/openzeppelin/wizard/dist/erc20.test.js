"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const _1 = require(".");
const erc20_1 = require("./erc20");
const print_1 = require("./print");
function testERC20(title, opts) {
    (0, ava_1.default)(title, t => {
        const c = (0, erc20_1.buildERC20)({
            name: 'MyToken',
            symbol: 'MTK',
            ...opts,
        });
        t.snapshot((0, print_1.printContract)(c));
    });
}
/**
 * Tests external API for equivalence with internal API
 */
function testAPIEquivalence(title, opts) {
    (0, ava_1.default)(title, t => {
        t.is(_1.erc20.print(opts), (0, print_1.printContract)((0, erc20_1.buildERC20)({
            name: 'MyToken',
            symbol: 'MTK',
            ...opts,
        })));
    });
}
testERC20('basic erc20', {});
testERC20('erc20 with snapshots', {
    snapshots: true,
});
testERC20('erc20 burnable', {
    burnable: true,
});
testERC20('erc20 burnable with snapshots', {
    burnable: true,
    snapshots: true,
});
testERC20('erc20 pausable', {
    pausable: true,
    access: 'ownable',
});
testERC20('erc20 pausable with roles', {
    pausable: true,
    access: 'roles',
});
testERC20('erc20 burnable pausable', {
    burnable: true,
    pausable: true,
});
testERC20('erc20 burnable pausable with snapshots', {
    burnable: true,
    pausable: true,
    snapshots: true,
});
testERC20('erc20 preminted', {
    premint: '1000',
});
testERC20('erc20 premint of 0', {
    premint: '0',
});
testERC20('erc20 mintable', {
    mintable: true,
    access: 'ownable',
});
testERC20('erc20 mintable with roles', {
    mintable: true,
    access: 'roles',
});
testERC20('erc20 permit', {
    permit: true,
});
testERC20('erc20 votes', {
    votes: true,
});
testERC20('erc20 flashmint', {
    flashmint: true,
});
testERC20('erc20 full upgradeable transparent', {
    premint: '2000',
    access: 'roles',
    burnable: true,
    mintable: true,
    pausable: true,
    snapshots: true,
    permit: true,
    votes: true,
    flashmint: true,
    upgradeable: 'transparent',
});
testERC20('erc20 full upgradeable uups', {
    premint: '2000',
    access: 'roles',
    burnable: true,
    mintable: true,
    pausable: true,
    snapshots: true,
    permit: true,
    votes: true,
    flashmint: true,
    upgradeable: 'uups',
});
testAPIEquivalence('erc20 API default');
testAPIEquivalence('erc20 API basic', { name: 'CustomToken', symbol: 'CTK' });
testAPIEquivalence('erc20 API full upgradeable', {
    name: 'CustomToken',
    symbol: 'CTK',
    premint: '2000',
    access: 'roles',
    burnable: true,
    mintable: true,
    pausable: true,
    snapshots: true,
    permit: true,
    votes: true,
    flashmint: true,
    upgradeable: 'uups',
});
(0, ava_1.default)('erc20 API assert defaults', async (t) => {
    t.is(_1.erc20.print(_1.erc20.defaults), _1.erc20.print());
});
(0, ava_1.default)('erc20 API isAccessControlRequired', async (t) => {
    t.is(_1.erc20.isAccessControlRequired({ mintable: true }), true);
    t.is(_1.erc20.isAccessControlRequired({ pausable: true }), true);
    t.is(_1.erc20.isAccessControlRequired({ snapshots: true }), true);
    t.is(_1.erc20.isAccessControlRequired({ upgradeable: 'uups' }), true);
    t.is(_1.erc20.isAccessControlRequired({ upgradeable: 'transparent' }), false);
});
//# sourceMappingURL=erc20.test.js.map