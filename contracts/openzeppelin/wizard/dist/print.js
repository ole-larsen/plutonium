"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.printValue = exports.printContract = void 0;
require("array.prototype.flatmap/auto");
const options_1 = require("./options");
const format_lines_1 = require("./utils/format-lines");
const map_values_1 = require("./utils/map-values");
const SOLIDITY_VERSION = '0.8.4';
function printContract(contract, opts) {
    const helpers = (0, options_1.withHelpers)(contract, opts);
    const fns = (0, map_values_1.mapValues)(sortedFunctions(contract), fns => fns.map(fn => printFunction(fn, helpers)));
    const hasOverrides = fns.override.some(l => l.length > 0);
    return (0, format_lines_1.formatLines)(...(0, format_lines_1.spaceBetween)([
        `// SPDX-License-Identifier: ${contract.license}`,
        `pragma solidity ^${SOLIDITY_VERSION};`,
    ], contract.imports.map(p => `import "${helpers.transformImport(p)}";`), [
        ...printNatspecTags(contract.natspecTags),
        [`contract ${contract.name}`, ...printInheritance(contract, helpers), '{'].join(' '),
        (0, format_lines_1.spaceBetween)(printUsingFor(contract, helpers), contract.variables.map(helpers.transformVariable), printConstructor(contract, helpers), ...fns.code, ...fns.modifiers, hasOverrides ? [`// The following functions are overrides required by Solidity.`] : [], ...fns.override),
        `}`,
    ]));
}
exports.printContract = printContract;
function printInheritance(contract, { transformName }) {
    if (contract.parents.length > 0) {
        return ['is ' + contract.parents.map(p => transformName(p.contract.name)).join(', ')];
    }
    else {
        return [];
    }
}
function printUsingFor(contract, { transformName }) {
    return contract.using.map(u => `using ${transformName(u.library.name)} for ${transformName(u.usingFor)};`);
}
function printConstructor(contract, helpers) {
    const hasParentParams = contract.parents.some(p => p.params.length > 0);
    const hasConstructorCode = contract.constructorCode.length > 0;
    const parentsWithInitializers = contract.parents.filter(hasInitializer);
    if (hasParentParams || hasConstructorCode || (helpers.upgradeable && parentsWithInitializers.length > 0)) {
        const parents = parentsWithInitializers
            .flatMap(p => printParentConstructor(p, helpers));
        const modifiers = helpers.upgradeable ? ['initializer public'] : parents;
        const args = contract.constructorArgs.map(a => printArgument(a, helpers));
        const body = helpers.upgradeable
            ? (0, format_lines_1.spaceBetween)(parents.map(p => p + ';'), contract.constructorCode)
            : contract.constructorCode;
        const head = helpers.upgradeable ? 'function initialize' : 'constructor';
        const constructor = printFunction2(head, args, modifiers, body);
        if (!helpers.upgradeable) {
            return constructor;
        }
        else {
            return (0, format_lines_1.spaceBetween)(DISABLE_INITIALIZERS, constructor);
        }
    }
    else if (!helpers.upgradeable) {
        return [];
    }
    else {
        return DISABLE_INITIALIZERS;
    }
}
const DISABLE_INITIALIZERS = [
    '/// @custom:oz-upgrades-unsafe-allow constructor',
    'constructor() {',
    [
        '_disableInitializers();'
    ],
    '}'
];
function hasInitializer(parent) {
    // CAUTION
    // This list is validated by compilation of SafetyCheck.sol.
    // Always keep this list and that file in sync.
    return !['Initializable'].includes(parent.contract.name);
}
// Functions with code first, then those with modifiers, then the rest
function sortedFunctions(contract) {
    const fns = { code: [], modifiers: [], override: [] };
    for (const fn of contract.functions) {
        if (fn.code.length > 0) {
            fns.code.push(fn);
        }
        else if (fn.modifiers.length > 0) {
            fns.modifiers.push(fn);
        }
        else {
            fns.override.push(fn);
        }
    }
    return fns;
}
function printParentConstructor({ contract, params }, helpers) {
    const fn = helpers.upgradeable ? `__${contract.name}_init` : contract.name;
    if (helpers.upgradeable || params.length > 0) {
        return [
            fn + '(' + params.map(printValue).join(', ') + ')',
        ];
    }
    else {
        return [];
    }
}
function printValue(value) {
    if (typeof value === 'object') {
        if ('lit' in value) {
            return value.lit;
        }
        else if ('note' in value) {
            return `${printValue(value.value)} /* ${value.note} */`;
        }
        else {
            throw Error('Unknown value type');
        }
    }
    else if (typeof value === 'number') {
        if (Number.isSafeInteger(value)) {
            return value.toFixed(0);
        }
        else {
            throw new Error(`Number not representable (${value})`);
        }
    }
    else {
        return JSON.stringify(value);
    }
}
exports.printValue = printValue;
function printFunction(fn, helpers) {
    var _a, _b;
    const { transformName } = helpers;
    if (fn.override.size <= 1 && fn.modifiers.length === 0 && fn.code.length === 0 && !fn.final) {
        return [];
    }
    const modifiers = [fn.kind, ...fn.modifiers];
    if (fn.mutability !== 'nonpayable') {
        modifiers.splice(1, 0, fn.mutability);
    }
    if (fn.override.size === 1) {
        modifiers.push(`override`);
    }
    else if (fn.override.size > 1) {
        modifiers.push(`override(${[...fn.override].map(transformName).join(', ')})`);
    }
    if ((_a = fn.returns) === null || _a === void 0 ? void 0 : _a.length) {
        modifiers.push(`returns (${fn.returns.join(', ')})`);
    }
    const code = [...fn.code];
    if (fn.override.size > 0 && !fn.final) {
        const superCall = `super.${fn.name}(${fn.args.map(a => a.name).join(', ')});`;
        code.push(((_b = fn.returns) === null || _b === void 0 ? void 0 : _b.length) ? 'return ' + superCall : superCall);
    }
    if (modifiers.length + fn.code.length > 1) {
        return printFunction2('function ' + fn.name, fn.args.map(a => printArgument(a, helpers)), modifiers, code);
    }
    else {
        return [];
    }
}
// generic for functions and constructors
// kindedName = 'function foo' or 'constructor'
function printFunction2(kindedName, args, modifiers, code) {
    const fn = [];
    const headingLength = [kindedName, ...args, ...modifiers]
        .map(s => s.length)
        .reduce((a, b) => a + b);
    const braces = code.length > 0 ? '{' : '{}';
    if (headingLength <= 72) {
        fn.push([`${kindedName}(${args.join(', ')})`, ...modifiers, braces].join(' '));
    }
    else {
        fn.push(`${kindedName}(${args.join(', ')})`, modifiers, braces);
    }
    if (code.length > 0) {
        fn.push(code, '}');
    }
    return fn;
}
function printArgument(arg, { transformName }) {
    const type = /^[A-Z]/.test(arg.type) ? transformName(arg.type) : arg.type;
    return [type, arg.name].join(' ');
}
function printNatspecTags(tags) {
    return tags.map(({ key, value }) => `/// ${key} ${value}`);
}
//# sourceMappingURL=print.js.map