"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineFunctions = void 0;
function defineFunctions(fns) {
    return Object.fromEntries(Object.entries(fns).map(([name, fn]) => [
        name,
        Object.assign({ name }, fn),
    ]));
}
exports.defineFunctions = defineFunctions;
//# sourceMappingURL=define-functions.js.map