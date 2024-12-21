"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.withCommonDefaults = exports.defaults = void 0;
const set_info_1 = require("./set-info");
exports.defaults = {
    access: false,
    upgradeable: false,
    info: set_info_1.defaults,
};
function withCommonDefaults(opts) {
    var _a, _b, _c;
    return {
        access: (_a = opts.access) !== null && _a !== void 0 ? _a : false,
        upgradeable: (_b = opts.upgradeable) !== null && _b !== void 0 ? _b : false,
        info: (_c = opts.info) !== null && _c !== void 0 ? _c : {},
    };
}
exports.withCommonDefaults = withCommonDefaults;
//# sourceMappingURL=common-options.js.map