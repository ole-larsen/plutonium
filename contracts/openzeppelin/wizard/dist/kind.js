"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.sanitizeKind = void 0;
function sanitizeKind(kind) {
    if (typeof kind === 'string') {
        const sanitized = kind.replace(/^(ERC|.)/i, c => c.toUpperCase());
        if (isKind(sanitized)) {
            return sanitized;
        }
    }
    return 'ERC20';
}
exports.sanitizeKind = sanitizeKind;
function isKind(value) {
    switch (value) {
        case 'ERC20':
        case 'ERC1155':
        case 'ERC721':
        case 'Governor':
        case 'Custom':
            return true;
        default: {
            // Static assert that we've checked all kinds.
            const _ = value;
            return false;
        }
    }
}
//# sourceMappingURL=kind.js.map