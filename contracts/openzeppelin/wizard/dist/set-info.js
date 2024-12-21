"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.setInfo = exports.defaults = exports.infoOptions = exports.TAG_SECURITY_CONTACT = void 0;
exports.TAG_SECURITY_CONTACT = `@custom:security-contact`;
exports.infoOptions = [{}, { securityContact: 'security@example.com', license: 'WTFPL' }];
exports.defaults = { license: 'MIT' };
function setInfo(c, info) {
    const { securityContact, license } = info;
    if (securityContact) {
        c.addNatspecTag(exports.TAG_SECURITY_CONTACT, securityContact);
    }
    if (license) {
        c.license = license;
    }
}
exports.setInfo = setInfo;
//# sourceMappingURL=set-info.js.map