"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildGeneric = void 0;
const custom_1 = require("./custom");
const erc20_1 = require("./erc20");
const erc721_1 = require("./erc721");
const erc1155_1 = require("./erc1155");
const governor_1 = require("./governor");
function buildGeneric(opts) {
    switch (opts.kind) {
        case 'ERC20':
            return (0, erc20_1.buildERC20)(opts);
        case 'ERC721':
            return (0, erc721_1.buildERC721)(opts);
        case 'ERC1155':
            return (0, erc1155_1.buildERC1155)(opts);
        case 'Governor':
            return (0, governor_1.buildGovernor)(opts);
        case 'Custom':
            return (0, custom_1.buildCustom)(opts);
        default:
            const _ = opts;
            throw new Error('Unknown ERC');
    }
}
exports.buildGeneric = buildGeneric;
//# sourceMappingURL=build-generic.js.map