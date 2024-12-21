"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const _1 = require(".");
const erc1155_1 = require("./erc1155");
const print_1 = require("./print");
function testERC1155(title, opts) {
    (0, ava_1.default)(title, t => {
        const c = (0, erc1155_1.buildERC1155)({
            name: 'MyToken',
            uri: 'https://gateway.pinata.cloud/ipfs/QmcP9hxrnC1T5ATPmq2saFeAM1ypFX9BnAswCdHB9JCjLA/',
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
        t.is(_1.erc1155.print(opts), (0, print_1.printContract)((0, erc1155_1.buildERC1155)({
            name: 'MyToken',
            uri: '',
            ...opts,
        })));
    });
}
testERC1155('basic', {});
testERC1155('basic + roles', {
    access: 'roles',
});
testERC1155('no updatable uri', {
    updatableUri: false,
});
testERC1155('burnable', {
    burnable: true,
});
testERC1155('pausable', {
    pausable: true,
});
testERC1155('mintable', {
    mintable: true,
});
testERC1155('mintable + roles', {
    mintable: true,
    access: 'roles',
});
testERC1155('supply tracking', {
    supply: true,
});
testERC1155('full upgradeable transparent', {
    mintable: true,
    access: 'roles',
    burnable: true,
    pausable: true,
    upgradeable: 'transparent',
});
testERC1155('full upgradeable uups', {
    mintable: true,
    access: 'roles',
    burnable: true,
    pausable: true,
    upgradeable: 'uups',
});
testAPIEquivalence('API default');
testAPIEquivalence('API basic', { name: 'CustomToken', uri: 'https://gateway.pinata.cloud/ipfs/QmcP9hxrnC1T5ATPmq2saFeAM1ypFX9BnAswCdHB9JCjLA/' });
testAPIEquivalence('API full upgradeable', {
    name: 'CustomToken',
    uri: 'https://gateway.pinata.cloud/ipfs/QmcP9hxrnC1T5ATPmq2saFeAM1ypFX9BnAswCdHB9JCjLA/',
    mintable: true,
    access: 'roles',
    burnable: true,
    pausable: true,
    upgradeable: 'uups',
});
(0, ava_1.default)('API assert defaults', async (t) => {
    t.is(_1.erc1155.print(_1.erc1155.defaults), _1.erc1155.print());
});
(0, ava_1.default)('API isAccessControlRequired', async (t) => {
    t.is(_1.erc1155.isAccessControlRequired({ updatableUri: false, mintable: true }), true);
    t.is(_1.erc1155.isAccessControlRequired({ updatableUri: false, pausable: true }), true);
    t.is(_1.erc1155.isAccessControlRequired({ updatableUri: false, upgradeable: 'uups' }), true);
    t.is(_1.erc1155.isAccessControlRequired({ updatableUri: true }), true);
    t.is(_1.erc1155.isAccessControlRequired({ updatableUri: false }), false);
    t.is(_1.erc1155.isAccessControlRequired({}), true); // updatableUri is true by default
});
//# sourceMappingURL=erc1155.test.js.map