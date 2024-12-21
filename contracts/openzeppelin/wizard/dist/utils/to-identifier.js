"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.toIdentifier = void 0;
function toIdentifier(str, capitalize = false) {
    return str
        .normalize('NFD').replace(/[\u0300-\u036f]/g, '') // remove accents
        .replace(/^[^a-zA-Z$_]+/, '')
        .replace(/^(.)/, c => capitalize ? c.toUpperCase() : c)
        .replace(/[^\w$]+(.?)/g, (_, c) => c.toUpperCase());
}
exports.toIdentifier = toIdentifier;
//# sourceMappingURL=to-identifier.js.map