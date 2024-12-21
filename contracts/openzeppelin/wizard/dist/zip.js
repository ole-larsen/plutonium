"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.zipContract = void 0;
const jszip_1 = __importDefault(require("jszip"));
const print_1 = require("./print");
const transitive_closure_1 = require("./utils/transitive-closure");
const openzeppelin_contracts_1 = __importDefault(require("../openzeppelin-contracts"));
const options_1 = require("./options");
const readme = (variant) => `\
# OpenZeppelin Contracts

The files in this directory were sourced unmodified from OpenZeppelin Contracts v${openzeppelin_contracts_1.default.version}.

They are not meant to be edited.

The originals can be found on [GitHub] and [npm].

[GitHub]: https://github.com/OpenZeppelin/openzeppelin-contracts${variant}/tree/v${openzeppelin_contracts_1.default.version}
[npm]: https://www.npmjs.com/package/@openzeppelin/contracts${variant}/v/${openzeppelin_contracts_1.default.version}

Generated with OpenZeppelin Contracts Wizard (https://zpl.in/wizard).
`;
function zipContract(c) {
    const { transformImport } = (0, options_1.withHelpers)(c);
    const contractsVariant = c.upgradeable ? '-upgradeable' : '';
    const fileName = c.name + '.sol';
    const dependencies = {
        [fileName]: c.parents.map(p => transformImport(p.contract.path)),
        ...openzeppelin_contracts_1.default.dependencies,
    };
    const allImports = (0, transitive_closure_1.reachable)(dependencies, fileName);
    const zip = new jszip_1.default();
    zip.file(fileName, (0, print_1.printContract)(c, { transformImport: p => './' + p }));
    zip.file(`@openzeppelin/contracts${contractsVariant}/README.md`, readme(contractsVariant));
    for (const importPath of allImports) {
        const source = openzeppelin_contracts_1.default.sources[importPath];
        if (source === undefined) {
            throw new Error(`Source for ${importPath} not found`);
        }
        zip.file(importPath, source);
    }
    return zip;
}
exports.zipContract = zipContract;
//# sourceMappingURL=zip.js.map