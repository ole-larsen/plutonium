"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addPausable = void 0;
const set_access_control_1 = require("./set-access-control");
const define_functions_1 = require("./utils/define-functions");
function addPausable(c, access, pausableFns) {
    c.addParent({
        name: 'Pausable',
        path: '@openzeppelin/contracts/security/Pausable.sol',
    });
    for (const fn of pausableFns) {
        c.addModifier('whenNotPaused', fn);
    }
    (0, set_access_control_1.requireAccessControl)(c, functions.pause, access, 'PAUSER');
    c.addFunctionCode('_pause();', functions.pause);
    (0, set_access_control_1.requireAccessControl)(c, functions.unpause, access, 'PAUSER');
    c.addFunctionCode('_unpause();', functions.unpause);
}
exports.addPausable = addPausable;
const functions = (0, define_functions_1.defineFunctions)({
    pause: {
        kind: 'public',
        args: [],
    },
    unpause: {
        kind: 'public',
        args: [],
    },
});
//# sourceMappingURL=add-pausable.js.map