"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateGovernorOptions = void 0;
const governor_1 = require("../governor");
const set_access_control_1 = require("../set-access-control");
const set_info_1 = require("../set-info");
const set_upgradeable_1 = require("../set-upgradeable");
const alternatives_1 = require("./alternatives");
const booleans = [true, false];
const blueprint = {
    name: ['MyGovernor'],
    delay: ['1 week'],
    period: ['1 week'],
    blockTime: [governor_1.defaults.blockTime],
    proposalThreshold: ['0', '1000'],
    decimals: [18],
    quorumMode: ['percent', 'absolute'],
    quorumPercent: [4],
    quorumAbsolute: ['1000'],
    votes: governor_1.votesOptions,
    timelock: governor_1.timelockOptions,
    bravo: booleans,
    settings: booleans,
    upgradeable: set_upgradeable_1.upgradeableOptions,
    access: set_access_control_1.accessOptions,
    info: set_info_1.infoOptions,
};
function* generateGovernorOptions() {
    yield* (0, alternatives_1.generateAlternatives)(blueprint);
}
exports.generateGovernorOptions = generateGovernorOptions;
//# sourceMappingURL=governor.js.map