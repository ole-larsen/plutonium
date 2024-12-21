"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateERC20Options = void 0;
const set_access_control_1 = require("../set-access-control");
const set_info_1 = require("../set-info");
const set_upgradeable_1 = require("../set-upgradeable");
const alternatives_1 = require("./alternatives");
const booleans = [true, false];
const blueprint = {
    name: ['MyToken'],
    symbol: ['MTK'],
    burnable: booleans,
    snapshots: booleans,
    pausable: booleans,
    mintable: booleans,
    permit: booleans,
    votes: booleans,
    flashmint: booleans,
    premint: ['1'],
    access: set_access_control_1.accessOptions,
    upgradeable: set_upgradeable_1.upgradeableOptions,
    info: set_info_1.infoOptions,
};
function* generateERC20Options() {
    yield* (0, alternatives_1.generateAlternatives)(blueprint);
}
exports.generateERC20Options = generateERC20Options;
//# sourceMappingURL=erc20.js.map