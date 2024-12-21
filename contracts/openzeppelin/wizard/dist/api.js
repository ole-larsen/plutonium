"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.custom = exports.governor = exports.erc1155 = exports.erc721 = exports.erc20 = void 0;
const erc20_1 = require("./erc20");
const erc721_1 = require("./erc721");
const erc1155_1 = require("./erc1155");
const governor_1 = require("./governor");
const custom_1 = require("./custom");
exports.erc20 = {
    print: erc20_1.printERC20,
    defaults: erc20_1.defaults,
    isAccessControlRequired: erc20_1.isAccessControlRequired
};
exports.erc721 = {
    print: erc721_1.printERC721,
    defaults: erc721_1.defaults,
    isAccessControlRequired: erc721_1.isAccessControlRequired
};
exports.erc1155 = {
    print: erc1155_1.printERC1155,
    defaults: erc1155_1.defaults,
    isAccessControlRequired: erc1155_1.isAccessControlRequired
};
exports.governor = {
    print: governor_1.printGovernor,
    defaults: governor_1.defaults,
    isAccessControlRequired: governor_1.isAccessControlRequired
};
exports.custom = {
    print: custom_1.printCustom,
    defaults: custom_1.defaults,
    isAccessControlRequired: custom_1.isAccessControlRequired
};
//# sourceMappingURL=api.js.map