"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs_1 = require("fs");
const path_1 = __importDefault(require("path"));
const hardhat_1 = __importDefault(require("hardhat"));
const utils_1 = require("solidity-ast/utils");
const rimraf_1 = __importDefault(require("rimraf"));
const util_1 = require("util");
const package_json_1 = require("@openzeppelin/contracts/package.json");
const rimraf = (0, util_1.promisify)(rimraf_1.default);
const sources_1 = require("../generate/sources");
const map_values_1 = require("../utils/map-values");
const transitive_closure_1 = require("../utils/transitive-closure");
async function main() {
    var _a;
    const generatedSourcesPath = path_1.default.join(hardhat_1.default.config.paths.sources, 'generated');
    await rimraf(generatedSourcesPath);
    await (0, sources_1.writeGeneratedSources)(generatedSourcesPath, 'minimal-cover');
    await hardhat_1.default.run('compile');
    const dependencies = {};
    const sources = {};
    for (const buildInfoPath of await hardhat_1.default.artifacts.getBuildInfoPaths()) {
        const buildInfo = JSON.parse(await fs_1.promises.readFile(buildInfoPath, 'utf8'));
        for (const [sourceFile, { ast }] of Object.entries(buildInfo.output.sources)) {
            if (sourceFile.startsWith('@openzeppelin/contracts')) {
                const sourceDependencies = ((_a = dependencies[sourceFile]) !== null && _a !== void 0 ? _a : (dependencies[sourceFile] = new Set()));
                for (const imp of (0, utils_1.findAll)('ImportDirective', ast)) {
                    sourceDependencies.add(imp.absolutePath);
                }
            }
        }
        for (const [sourceFile, { content }] of Object.entries(buildInfo.input.sources)) {
            if (sourceFile.startsWith('@openzeppelin/contracts')) {
                sources[sourceFile] = content;
            }
        }
    }
    const contracts = {
        version: package_json_1.version,
        sources,
        dependencies: (0, map_values_1.mapValues)((0, transitive_closure_1.transitiveClosure)(dependencies), d => Array.from(d)),
    };
    await fs_1.promises.writeFile('openzeppelin-contracts.json', JSON.stringify(contracts, null, 2));
}
main().catch(e => {
    console.error(e);
    process.exit(1);
});
//# sourceMappingURL=prepare.js.map