"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.sortedBy = void 0;
function sortedBy(xs, fn) {
    return Array.from(xs).sort((a, b) => fn(a) - fn(b));
}
exports.sortedBy = sortedBy;
//# sourceMappingURL=sorted-by.js.map