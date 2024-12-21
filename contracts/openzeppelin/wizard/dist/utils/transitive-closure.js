"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reachable = exports.transitiveClosure = void 0;
function transitiveClosure(obj) {
    const closure = {};
    for (const key in obj) {
        closure[key] = reachable(obj, key);
    }
    return closure;
}
exports.transitiveClosure = transitiveClosure;
function reachable(obj, key) {
    var _a;
    let prevSize = 0;
    const res = new Set(obj[key]);
    while (prevSize < res.size) {
        prevSize = res.size;
        for (const k of res) {
            for (const v of (_a = obj[k]) !== null && _a !== void 0 ? _a : []) {
                res.add(v);
            }
        }
    }
    return res;
}
exports.reachable = reachable;
//# sourceMappingURL=transitive-closure.js.map