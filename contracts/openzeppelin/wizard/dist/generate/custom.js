"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateCustomOptions = void 0;
const set_access_control_1 = require("../set-access-control");
const set_info_1 = require("../set-info");
const set_upgradeable_1 = require("../set-upgradeable");
const alternatives_1 = require("./alternatives");
const booleans = [true, false];
const blueprint = {
    name: ['MyContract'],
    pausable: booleans,
    access: set_access_control_1.accessOptions,
    upgradeable: set_upgradeable_1.upgradeableOptions,
    info: set_info_1.infoOptions,
};
function* generateCustomOptions() {
    yield* (0, alternatives_1.generateAlternatives)(blueprint);
}
exports.generateCustomOptions = generateCustomOptions;
//# sourceMappingURL=custom.js.map