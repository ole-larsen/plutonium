"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildCustom = exports.isAccessControlRequired = exports.printCustom = exports.defaults = void 0;
const contract_1 = require("./contract");
const common_options_1 = require("./common-options");
const set_upgradeable_1 = require("./set-upgradeable");
const set_info_1 = require("./set-info");
const set_access_control_1 = require("./set-access-control");
const add_pausable_1 = require("./add-pausable");
const print_1 = require("./print");
exports.defaults = {
    name: 'MyContract',
    pausable: false,
    access: common_options_1.defaults.access,
    upgradeable: common_options_1.defaults.upgradeable,
    info: common_options_1.defaults.info,
};
function withDefaults(opts) {
    var _a;
    return {
        ...opts,
        ...(0, common_options_1.withCommonDefaults)(opts),
        pausable: (_a = opts.pausable) !== null && _a !== void 0 ? _a : exports.defaults.pausable,
    };
}
function printCustom(opts = exports.defaults) {
    return (0, print_1.printContract)(buildCustom(opts));
}
exports.printCustom = printCustom;
function isAccessControlRequired(opts) {
    return opts.pausable || opts.upgradeable === 'uups';
}
exports.isAccessControlRequired = isAccessControlRequired;
function buildCustom(opts) {
    const allOpts = withDefaults(opts);
    const c = new contract_1.ContractBuilder(allOpts.name);
    const { access, upgradeable, info } = allOpts;
    if (allOpts.pausable) {
        (0, add_pausable_1.addPausable)(c, access, []);
    }
    (0, set_access_control_1.setAccessControl)(c, access);
    (0, set_upgradeable_1.setUpgradeable)(c, upgradeable, access);
    (0, set_info_1.setInfo)(c, info);
    return c;
}
exports.buildCustom = buildCustom;
//# sourceMappingURL=custom.js.map