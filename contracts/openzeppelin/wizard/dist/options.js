"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.withHelpers = void 0;
const path_1 = __importDefault(require("path"));
const upgradeableName = (n) => {
    if (n === 'Initializable') {
        return n;
    }
    else {
        return n.replace(/(Upgradeable)?(?=\.|$)/, 'Upgradeable');
    }
};
const upgradeableImport = (p) => {
    const { dir, ext, name } = path_1.default.parse(p);
    // Use path.posix to get forward slashes
    return path_1.default.posix.format({
        ext,
        dir: dir.replace(/^@openzeppelin\/contracts/, '@openzeppelin/contracts-upgradeable'),
        name: upgradeableName(name),
    });
};
function withHelpers(contract, opts = {}) {
    const upgradeable = contract.upgradeable;
    const transformName = (n) => upgradeable ? upgradeableName(n) : n;
    return {
        upgradeable,
        transformName,
        transformImport: p1 => {
            var _a, _b;
            const p2 = upgradeable ? upgradeableImport(p1) : p1;
            return (_b = (_a = opts.transformImport) === null || _a === void 0 ? void 0 : _a.call(opts, p2)) !== null && _b !== void 0 ? _b : p2;
        },
        transformVariable: v => v.replace(/[A-Z]\w*(?=\.|$)/, transformName),
    };
}
exports.withHelpers = withHelpers;
//# sourceMappingURL=options.js.map