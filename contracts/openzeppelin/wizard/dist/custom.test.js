"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const _1 = require(".");
const custom_1 = require("./custom");
const print_1 = require("./print");
function testCustom(title, opts) {
    (0, ava_1.default)(title, t => {
        const c = (0, custom_1.buildCustom)({
            name: 'MyContract',
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
        t.is(_1.custom.print(opts), (0, print_1.printContract)((0, custom_1.buildCustom)({
            name: 'MyContract',
            ...opts,
        })));
    });
}
testCustom('custom', {});
testCustom('pausable', {
    pausable: true,
});
testCustom('upgradeable transparent', {
    upgradeable: 'transparent',
});
testCustom('upgradeable uups', {
    upgradeable: 'uups',
});
testCustom('access control disabled', {
    access: false,
});
testCustom('access control ownable', {
    access: 'ownable',
});
testCustom('access control roles', {
    access: 'roles',
});
testCustom('upgradeable uups with access control disabled', {
    // API should override access to true since it is required for UUPS
    access: false,
    upgradeable: 'uups',
});
testAPIEquivalence('custom API default');
testAPIEquivalence('custom API basic', { name: 'CustomContract' });
testAPIEquivalence('custom API full upgradeable', {
    name: 'CustomContract',
    access: 'roles',
    pausable: true,
    upgradeable: 'uups',
});
(0, ava_1.default)('custom API assert defaults', async (t) => {
    t.is(_1.custom.print(_1.custom.defaults), _1.custom.print());
});
(0, ava_1.default)('API isAccessControlRequired', async (t) => {
    t.is(_1.custom.isAccessControlRequired({ pausable: true }), true);
    t.is(_1.custom.isAccessControlRequired({ upgradeable: 'uups' }), true);
    t.is(_1.custom.isAccessControlRequired({ upgradeable: 'transparent' }), false);
});
//# sourceMappingURL=custom.test.js.map