"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.findCover = void 0;
const sorted_by_1 = require("./sorted-by");
// Greedy approximation of minimum set cover.
function findCover(sets, getElements) {
    const sortedSets = (0, sorted_by_1.sortedBy)(sets.map(set => ({ set, elems: getElements(set) })), s => -s.elems.length);
    const seen = new Set();
    const res = [];
    for (const { set, elems } of sortedSets) {
        let included = false;
        for (const e of elems) {
            if (!included && !seen.has(e)) {
                res.push(set);
                included = true;
            }
            seen.add(e);
        }
    }
    return res;
}
exports.findCover = findCover;
//# sourceMappingURL=find-cover.js.map