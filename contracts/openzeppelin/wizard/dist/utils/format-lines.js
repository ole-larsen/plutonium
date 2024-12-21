"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.spaceBetween = exports.formatLines = void 0;
require("array.prototype.flatmap/auto");
const whitespace = Symbol('whitespace');
function formatLines(...lines) {
    return [...indentEach(0, lines)].join('\n') + '\n';
}
exports.formatLines = formatLines;
function* indentEach(indent, lines) {
    for (const line of lines) {
        if (line === whitespace) {
            yield '';
        }
        else if (Array.isArray(line)) {
            yield* indentEach(indent + 1, line);
        }
        else {
            yield '    '.repeat(indent) + line;
        }
    }
}
function spaceBetween(...lines) {
    return lines
        .filter(l => l.length > 0)
        .flatMap(l => [whitespace, ...l])
        .slice(1);
}
exports.spaceBetween = spaceBetween;
//# sourceMappingURL=format-lines.js.map