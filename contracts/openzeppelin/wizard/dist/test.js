"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const hardhat_1 = __importDefault(require("hardhat"));
const path_1 = __importDefault(require("path"));
const sources_1 = require("./generate/sources");
const api_1 = require("./api");
(0, ava_1.default)('result compiles', async (t) => {
    const generatedSourcesPath = path_1.default.join(hardhat_1.default.config.paths.sources, 'generated');
    await (0, sources_1.writeGeneratedSources)(generatedSourcesPath, 'all');
    // We only want to check that contracts compile and we don't care about any
    // of the outputs. Setting empty outputSelection causes compilation to go a
    // lot faster and not run out of memory.
    for (const { settings } of hardhat_1.default.config.solidity.compilers) {
        settings.outputSelection = {};
    }
    await hardhat_1.default.run('compile');
    t.pass();
});
function isAccessControlRequired(opts) {
    switch (opts.kind) {
        case 'ERC20':
            return api_1.erc20.isAccessControlRequired(opts);
        case 'ERC721':
            return api_1.erc721.isAccessControlRequired(opts);
        case 'ERC1155':
            return api_1.erc1155.isAccessControlRequired(opts);
        case 'Governor':
            return api_1.governor.isAccessControlRequired(opts);
        case 'Custom':
            return api_1.custom.isAccessControlRequired(opts);
        default:
            throw new Error("No such kind");
    }
}
(0, ava_1.default)('is access control required', async (t) => {
    for (const contract of (0, sources_1.generateSources)('all')) {
        const regexOwnable = /import.*Ownable(Upgradeable)?.sol.*/gm;
        if (!contract.options.access) {
            if (isAccessControlRequired(contract.options)) {
                t.regex(contract.source, regexOwnable, JSON.stringify(contract.options));
            }
            else {
                t.notRegex(contract.source, regexOwnable, JSON.stringify(contract.options));
            }
        }
    }
});
//# sourceMappingURL=test.js.map