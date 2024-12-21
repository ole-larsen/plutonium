"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mapValues = void 0;
function mapValues(obj, fn) {
    const res = {};
    for (const key in obj) {
        res[key] = fn(obj[key]);
    }
    return res;
}
exports.mapValues = mapValues;
//# sourceMappingURL=map-values.js.map