"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.writeGeneratedSources = exports.generateSources = exports.generateOptions = void 0;
const fs_1 = require("fs");
const path_1 = __importDefault(require("path"));
const crypto_1 = __importDefault(require("crypto"));
const erc20_1 = require("./erc20");
const erc721_1 = require("./erc721");
const erc1155_1 = require("./erc1155");
const governor_1 = require("./governor");
const custom_1 = require("./custom");
const build_generic_1 = require("../build-generic");
const print_1 = require("../print");
const error_1 = require("../error");
const find_cover_1 = require("../utils/find-cover");
function* generateOptions() {
    for (const kindOpts of (0, erc20_1.generateERC20Options)()) {
        yield { kind: 'ERC20', ...kindOpts };
    }
    for (const kindOpts of (0, erc721_1.generateERC721Options)()) {
        yield { kind: 'ERC721', ...kindOpts };
    }
    for (const kindOpts of (0, erc1155_1.generateERC1155Options)()) {
        yield { kind: 'ERC1155', ...kindOpts };
    }
    for (const kindOpts of (0, governor_1.generateGovernorOptions)()) {
        yield { kind: 'Governor', ...kindOpts };
    }
    for (const kindOpts of (0, custom_1.generateCustomOptions)()) {
        yield { kind: 'Custom', ...kindOpts };
    }
}
exports.generateOptions = generateOptions;
function generateContractSubset(subset) {
    const contracts = [];
    for (const options of generateOptions()) {
        const id = crypto_1.default
            .createHash('sha1')
            .update(JSON.stringify(options))
            .digest()
            .toString('hex');
        try {
            const contract = (0, build_generic_1.buildGeneric)(options);
            contracts.push({ id, options, contract });
        }
        catch (e) {
            if (e instanceof error_1.OptionsError) {
                continue;
            }
            else {
                throw e;
            }
        }
    }
    if (subset === 'all') {
        return contracts;
    }
    else {
        const getParents = (c) => c.contract.parents.map(p => p.contract.path);
        return [
            ...(0, find_cover_1.findCover)(contracts.filter(c => c.options.upgradeable), getParents),
            ...(0, find_cover_1.findCover)(contracts.filter(c => !c.options.upgradeable), getParents),
        ];
    }
}
function* generateSources(subset) {
    for (const c of generateContractSubset(subset)) {
        const source = (0, print_1.printContract)(c.contract);
        yield { ...c, source };
    }
}
exports.generateSources = generateSources;
async function writeGeneratedSources(dir, subset) {
    await fs_1.promises.mkdir(dir, { recursive: true });
    for (const { id, source } of generateSources(subset)) {
        await fs_1.promises.writeFile(path_1.default.format({ dir, name: id, ext: '.sol' }), source);
    }
}
exports.writeGeneratedSources = writeGeneratedSources;
//# sourceMappingURL=sources.js.map