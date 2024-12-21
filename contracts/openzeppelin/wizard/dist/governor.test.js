"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const _1 = require(".");
const governor_1 = require("./governor");
const print_1 = require("./print");
function testGovernor(title, opts) {
    (0, ava_1.default)(title, t => {
        const c = (0, governor_1.buildGovernor)({
            name: 'MyGovernor',
            delay: '1 week',
            period: '1 week',
            settings: false,
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
        t.is(_1.governor.print(opts), (0, print_1.printContract)((0, governor_1.buildGovernor)({
            name: 'MyGovernor',
            delay: '1 block',
            period: '1 week',
            ...opts,
        })));
    });
}
testGovernor('governor with proposal threshold', {
    proposalThreshold: '1',
});
testGovernor('governor with custom block time', {
    blockTime: 6,
});
testGovernor('governor with custom decimals', {
    decimals: 6,
    proposalThreshold: '1',
    quorumMode: 'absolute',
    quorumAbsolute: '1',
});
testGovernor('governor with 0 decimals', {
    decimals: 0,
    proposalThreshold: '1',
    quorumMode: 'absolute',
    quorumAbsolute: '1',
});
testGovernor('governor with settings', {
    settings: true,
    proposalThreshold: '1',
});
testGovernor('governor with bravo', {
    bravo: true,
});
testGovernor('governor with erc20votes', {
    votes: 'erc20votes',
});
testGovernor('governor with erc721votes', {
    votes: 'erc721votes',
});
testGovernor('governor with erc721votes omit decimals', {
    votes: 'erc721votes',
    decimals: 6,
    proposalThreshold: '1',
    quorumMode: 'absolute',
    quorumAbsolute: '5',
});
testGovernor('governor with erc721votes settings omit decimals', {
    votes: 'erc721votes',
    decimals: 6,
    proposalThreshold: '10',
    settings: true,
});
testGovernor('governor with comp', {
    votes: 'comp',
    quorumMode: 'absolute',
    quorumAbsolute: '1',
});
testGovernor('governor with percent quorum', {
    quorumMode: 'percent',
    quorumPercent: 6,
});
testGovernor('governor with fractional percent quorum', {
    quorumMode: 'percent',
    quorumPercent: 0.5,
});
testGovernor('governor with openzeppelin timelock', {
    timelock: 'openzeppelin',
});
testGovernor('governor with compound timelock', {
    timelock: 'compound',
});
testAPIEquivalence('API default');
testAPIEquivalence('API basic', { name: 'CustomGovernor', delay: '2 weeks', period: '2 week' });
testAPIEquivalence('API basic upgradeable', { name: 'CustomGovernor', delay: '2 weeks', period: '2 week', upgradeable: 'uups' });
(0, ava_1.default)('API assert defaults', async (t) => {
    t.is(_1.governor.print(_1.governor.defaults), _1.governor.print());
});
(0, ava_1.default)('API isAccessControlRequired', async (t) => {
    t.is(_1.governor.isAccessControlRequired({ upgradeable: 'uups' }), true);
    t.is(_1.governor.isAccessControlRequired({}), false);
});
//# sourceMappingURL=governor.test.js.map