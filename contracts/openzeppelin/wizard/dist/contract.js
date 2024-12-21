"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ContractBuilder = void 0;
const to_identifier_1 = require("./utils/to-identifier");
// Order is important
const mutabilityRank = ['pure', 'view', 'nonpayable', 'payable'];
function maxMutability(a, b) {
    return mutabilityRank[Math.max(mutabilityRank.indexOf(a), mutabilityRank.indexOf(b))];
}
class ContractBuilder {
    constructor(name) {
        this.license = 'MIT';
        this.upgradeable = false;
        this.using = [];
        this.natspecTags = [];
        this.constructorArgs = [];
        this.constructorCode = [];
        this.variableSet = new Set();
        this.parentMap = new Map();
        this.functionMap = new Map();
        this.name = (0, to_identifier_1.toIdentifier)(name, true);
    }
    get parents() {
        return [...this.parentMap.values()].sort((a, b) => {
            if (a.contract.name === 'Initializable') {
                return -1;
            }
            else if (b.contract.name === 'Initializable') {
                return 1;
            }
            else {
                return 0;
            }
        });
    }
    get imports() {
        return [
            ...[...this.parentMap.values()].map(p => p.contract.path),
            ...this.using.map(u => u.library.path),
        ];
    }
    get functions() {
        return [...this.functionMap.values()];
    }
    get variables() {
        return [...this.variableSet];
    }
    addParent(contract, params = []) {
        const present = this.parentMap.has(contract.name);
        this.parentMap.set(contract.name, { contract, params });
        return !present;
    }
    addUsing(library, usingFor) {
        this.using.push({ library, usingFor });
    }
    addOverride(parent, baseFn, mutability) {
        const fn = this.addFunction(baseFn);
        fn.override.add(parent);
        if (mutability) {
            fn.mutability = maxMutability(fn.mutability, mutability);
        }
    }
    addModifier(modifier, baseFn) {
        const fn = this.addFunction(baseFn);
        fn.modifiers.push(modifier);
    }
    addNatspecTag(key, value) {
        if (!/^(@custom:)?[a-z][a-z\-]*$/.exec(key))
            throw new Error(`Invalid natspec key: ${key}`);
        this.natspecTags.push({ key, value });
    }
    addFunction(baseFn) {
        const signature = [baseFn.name, '(', ...baseFn.args.map(a => a.name), ')'].join('');
        const got = this.functionMap.get(signature);
        if (got !== undefined) {
            return got;
        }
        else {
            const fn = {
                override: new Set(),
                modifiers: [],
                code: [],
                mutability: 'nonpayable',
                final: false,
                ...baseFn,
            };
            this.functionMap.set(signature, fn);
            return fn;
        }
    }
    addConstructorArgument(arg) {
        this.constructorArgs.push(arg);
    }
    addConstructorCode(code) {
        this.constructorCode.push(code);
    }
    addFunctionCode(code, baseFn, mutability) {
        const fn = this.addFunction(baseFn);
        if (fn.final) {
            throw new Error(`Function ${baseFn.name} is already finalized`);
        }
        fn.code.push(code);
        if (mutability) {
            fn.mutability = maxMutability(fn.mutability, mutability);
        }
    }
    setFunctionBody(code, baseFn, mutability) {
        const fn = this.addFunction(baseFn);
        if (fn.code.length > 0) {
            throw new Error(`Function ${baseFn.name} has additional code`);
        }
        fn.code.push(...code);
        fn.final = true;
        if (mutability) {
            fn.mutability = mutability;
        }
    }
    addVariable(code) {
        const present = this.variableSet.has(code);
        this.variableSet.add(code);
        return !present;
    }
}
exports.ContractBuilder = ContractBuilder;
//# sourceMappingURL=contract.js.map