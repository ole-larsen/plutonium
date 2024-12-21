"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const zip_1 = require("./zip");
const erc20_1 = require("./erc20");
const sources_1 = require("./generate/sources");
const build_generic_1 = require("./build-generic");
(0, ava_1.default)('erc20 basic', t => {
    const c = (0, erc20_1.buildERC20)({ name: 'MyToken', symbol: 'MTK' });
    const zip = (0, zip_1.zipContract)(c);
    const files = Object.values(zip.files).map(f => f.name).sort();
    t.deepEqual(files, [
        '@openzeppelin/',
        '@openzeppelin/contracts/',
        '@openzeppelin/contracts/README.md',
        '@openzeppelin/contracts/token/',
        '@openzeppelin/contracts/token/ERC20/',
        '@openzeppelin/contracts/token/ERC20/ERC20.sol',
        '@openzeppelin/contracts/token/ERC20/IERC20.sol',
        '@openzeppelin/contracts/token/ERC20/extensions/',
        '@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol',
        '@openzeppelin/contracts/utils/',
        '@openzeppelin/contracts/utils/Context.sol',
        'MyToken.sol',
    ]);
});
(0, ava_1.default)('can zip all combinations', t => {
    for (const { options } of (0, sources_1.generateSources)('all')) {
        const c = (0, build_generic_1.buildGeneric)(options);
        (0, zip_1.zipContract)(c);
    }
    t.pass();
});
//# sourceMappingURL=zip.test.js.map