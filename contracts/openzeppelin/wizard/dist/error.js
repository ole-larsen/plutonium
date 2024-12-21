"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.OptionsError = void 0;
class OptionsError extends Error {
    constructor(messages) {
        super("Invalid options for Governor");
        this.messages = messages;
    }
}
exports.OptionsError = OptionsError;
//# sourceMappingURL=error.js.map