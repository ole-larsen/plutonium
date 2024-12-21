"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateAlternatives = void 0;
require("../utils/map-values");
function* generateAlternatives(blueprint) {
    const entries = Object.entries(blueprint).map(([key, values]) => ({
        key,
        values,
        current: 0,
        limit: values.length,
    }));
    for (; !done(); advance()) {
        yield Object.fromEntries(entries.map(e => [e.key, e.values[e.current % e.limit]]));
    }
    function done() {
        const last = entries[entries.length - 1];
        return (last === null || last === void 0 ? void 0 : last.current) === (last === null || last === void 0 ? void 0 : last.limit);
    }
    function advance() {
        for (const e of entries) {
            e.current = (e.current % e.limit) + 1;
            if (e.current < e.limit) {
                break;
            }
        }
    }
}
exports.generateAlternatives = generateAlternatives;
//# sourceMappingURL=alternatives.js.map