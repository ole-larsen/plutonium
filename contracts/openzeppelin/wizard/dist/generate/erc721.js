"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateERC721Options = void 0;
const set_access_control_1 = require("../set-access-control");
const set_info_1 = require("../set-info");
const set_upgradeable_1 = require("../set-upgradeable");
const alternatives_1 = require("./alternatives");
const booleans = [true, false];
const blueprint = {
    name: ['MyToken'],
    symbol: ['MTK'],
    baseUri: ['https://example.com/'],
    enumerable: booleans,
    uriStorage: booleans,
    burnable: booleans,
    pausable: booleans,
    mintable: booleans,
    incremental: booleans,
    access: set_access_control_1.accessOptions,
    upgradeable: set_upgradeable_1.upgradeableOptions,
    info: set_info_1.infoOptions,
    votes: booleans,
};
function* generateERC721Options() {
    yield* (0, alternatives_1.generateAlternatives)(blueprint);
}
exports.generateERC721Options = generateERC721Options;
//# sourceMappingURL=erc721.js.map