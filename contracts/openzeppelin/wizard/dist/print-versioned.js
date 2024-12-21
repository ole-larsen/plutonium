"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.printContractVersioned = void 0;
const package_json_1 = require("@openzeppelin/contracts/package.json");
const print_1 = require("./print");
function printContractVersioned(contract) {
    return (0, print_1.printContract)(contract, {
        transformImport: p => p.replace(/^@openzeppelin\/contracts(-upgradeable)?/, `$&@${package_json_1.version}`),
    });
}
exports.printContractVersioned = printContractVersioned;
//# sourceMappingURL=print-versioned.js.map