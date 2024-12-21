"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.durationToBlocks = exports.durationPattern = void 0;
const durationUnits = ['block', 'second', 'minute', 'hour', 'day', 'week', 'month', 'year'];
exports.durationPattern = new RegExp(`^(\\d+(?:\\.\\d+)?) +(${durationUnits.join('|')})s?$`);
const second = 1;
const minute = 60 * second;
const hour = 60 * minute;
const day = 24 * hour;
const week = 7 * day;
const month = 30 * day;
const year = 365 * day;
const secondsForUnit = { second, minute, hour, day, week, month, year };
function durationToBlocks(duration, blockTime) {
    const match = duration.trim().match(exports.durationPattern);
    if (!match) {
        throw new Error('Bad duration format');
    }
    const value = parseFloat(match[1]);
    const unit = match[2];
    if (unit === 'block') {
        if (!Number.isInteger(value)) {
            throw new Error('Invalid number of blocks');
        }
        return value;
    }
    const durationSeconds = value * secondsForUnit[unit];
    return Math.round(durationSeconds / blockTime);
}
exports.durationToBlocks = durationToBlocks;
//# sourceMappingURL=duration.js.map