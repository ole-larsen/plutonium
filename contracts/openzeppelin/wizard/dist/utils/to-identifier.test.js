"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const to_identifier_1 = require("./to-identifier");
(0, ava_1.default)('unmodified', t => {
    t.is((0, to_identifier_1.toIdentifier)('abc'), 'abc');
});
(0, ava_1.default)('remove leading specials', t => {
    t.is((0, to_identifier_1.toIdentifier)('--abc'), 'abc');
});
(0, ava_1.default)('remove specials and upcase next char', t => {
    t.is((0, to_identifier_1.toIdentifier)('abc-def'), 'abcDef');
    t.is((0, to_identifier_1.toIdentifier)('abc--def'), 'abcDef');
});
(0, ava_1.default)('capitalize', t => {
    t.is((0, to_identifier_1.toIdentifier)('abc', true), 'Abc');
});
//# sourceMappingURL=to-identifier.test.js.map