"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.requireAccessControl = exports.setAccessControl = exports.accessOptions = void 0;
const common_functions_1 = require("./common-functions");
exports.accessOptions = [false, 'ownable', 'roles'];
/**
 * Sets access control for the contract by adding inheritance.
 */
function setAccessControl(c, access) {
    switch (access) {
        case 'ownable': {
            c.addParent(parents.Ownable);
            break;
        }
        case 'roles': {
            if (c.addParent(parents.AccessControl)) {
                c.addConstructorCode('_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);');
            }
            c.addOverride(parents.AccessControl.name, common_functions_1.supportsInterface);
            break;
        }
    }
}
exports.setAccessControl = setAccessControl;
/**
 * Enables access control for the contract and restricts the given function with access control.
 */
function requireAccessControl(c, fn, access, role) {
    if (access === false) {
        access = 'ownable';
    }
    setAccessControl(c, access);
    switch (access) {
        case 'ownable': {
            c.addModifier('onlyOwner', fn);
            break;
        }
        case 'roles': {
            const roleId = role + '_ROLE';
            if (c.addVariable(`bytes32 public constant ${roleId} = keccak256("${roleId}");`)) {
                c.addConstructorCode(`_grantRole(${roleId}, msg.sender);`);
            }
            c.addModifier(`onlyRole(${roleId})`, fn);
            break;
        }
    }
}
exports.requireAccessControl = requireAccessControl;
const parents = {
    Ownable: {
        name: 'Ownable',
        path: '@openzeppelin/contracts/access/Ownable.sol',
    },
    AccessControl: {
        name: 'AccessControl',
        path: '@openzeppelin/contracts/access/AccessControl.sol',
    },
};
//# sourceMappingURL=set-access-control.js.map