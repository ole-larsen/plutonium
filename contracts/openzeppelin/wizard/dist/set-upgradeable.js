"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.setUpgradeable = exports.upgradeableOptions = void 0;
const set_access_control_1 = require("./set-access-control");
const define_functions_1 = require("./utils/define-functions");
exports.upgradeableOptions = [false, 'transparent', 'uups'];
function setUpgradeable(c, upgradeable, access) {
    if (upgradeable === false) {
        return;
    }
    c.upgradeable = true;
    c.addParent({
        name: 'Initializable',
        path: '@openzeppelin/contracts/proxy/utils/Initializable.sol',
    });
    switch (upgradeable) {
        case 'transparent': break;
        case 'uups': {
            (0, set_access_control_1.requireAccessControl)(c, functions._authorizeUpgrade, access, 'UPGRADER');
            c.addParent({
                name: 'UUPSUpgradeable',
                path: '@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol',
            });
            c.addOverride('UUPSUpgradeable', functions._authorizeUpgrade);
            c.setFunctionBody([], functions._authorizeUpgrade);
            break;
        }
        default: {
            const _ = upgradeable;
            throw new Error('Unknown value for `upgradeable`');
        }
    }
}
exports.setUpgradeable = setUpgradeable;
const functions = (0, define_functions_1.defineFunctions)({
    _authorizeUpgrade: {
        args: [
            { name: 'newImplementation', type: 'address' },
        ],
        kind: 'internal',
    },
});
//# sourceMappingURL=set-upgradeable.js.map