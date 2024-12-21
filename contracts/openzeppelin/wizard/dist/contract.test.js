"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ava_1 = __importDefault(require("ava"));
const contract_1 = require("./contract");
const print_1 = require("./print");
const set_info_1 = require("./set-info");
(0, ava_1.default)('contract basics', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with a parent', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const Bar = { name: 'Bar', path: './Bar.sol' };
    Foo.addParent(Bar);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with two parents', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const Bar = { name: 'Bar', path: './Bar.sol' };
    const Quux = { name: 'Quux', path: './Quux.sol' };
    Foo.addParent(Bar);
    Foo.addParent(Quux);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with a parent with parameters', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const Bar = { name: 'Bar', path: './Bar.sol' };
    Foo.addParent(Bar, ["param1", "param2"]);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with two parents only one with parameters', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const Bar = { name: 'Bar', path: './Bar.sol' };
    const Quux = { name: 'Quux', path: './Quux.sol' };
    Foo.addParent(Bar, ["param1", "param2"]);
    Foo.addParent(Quux);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with one override', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const _beforeTokenTransfer = {
        name: '_beforeTokenTransfer',
        kind: 'internal',
        args: [
            { name: 'from', type: 'address' },
            { name: 'to', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
    };
    Foo.addOverride('ERC20', _beforeTokenTransfer);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with two overrides', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addOverride('ERC20', _beforeTokenTransfer);
    Foo.addOverride('ERC20Snapshot', _beforeTokenTransfer);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with two different overrides', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addOverride('ERC20', _beforeTokenTransfer);
    Foo.addOverride('OtherParent', _beforeTokenTransfer);
    Foo.addOverride('ERC20', _otherFunction);
    Foo.addOverride('OtherParent', _otherFunction);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with a modifier', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addModifier('whenNotPaused', _otherFunction);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with a modifier and override', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addModifier('whenNotPaused', _otherFunction);
    Foo.addOverride('ERC20', _otherFunction);
    Foo.addOverride('OtherParent', _otherFunction);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with constructor code', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addConstructorCode('_mint(msg.sender, 10 ether);');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with constructor code and a parent', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    const Bar = { name: 'Bar', path: './Bar.sol' };
    Foo.addParent(Bar, ["param1", "param2"]);
    Foo.addConstructorCode('_mint(msg.sender, 10 ether);');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with function code', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addFunctionCode('_mint(msg.sender, 10 ether);', _otherFunction);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with overriden function with code', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addOverride('Bar', _otherFunction);
    Foo.addFunctionCode('_mint(msg.sender, 10 ether);', _otherFunction);
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with one variable', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addVariable('uint value = 42;');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with two variables', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addVariable('uint value = 42;');
    Foo.addVariable('string name = "john";');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('name with special characters', t => {
    const Foo = new contract_1.ContractBuilder('foo bar baz');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('using for statement', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addUsing({
        name: 'Counters',
        path: './Counters.sol',
    }, 'Counters.Counter');
    t.snapshot((0, print_1.printContract)(Foo));
});
(0, ava_1.default)('contract with info', t => {
    const Foo = new contract_1.ContractBuilder('Foo');
    Foo.addNatspecTag(set_info_1.TAG_SECURITY_CONTACT, 'security@example.com');
    t.snapshot((0, print_1.printContract)(Foo));
});
const _beforeTokenTransfer = {
    name: '_beforeTokenTransfer',
    kind: 'internal',
    args: [
        { name: 'from', type: 'address' },
        { name: 'to', type: 'address' },
        { name: 'amount', type: 'uint256' },
    ],
};
const _otherFunction = {
    name: '_otherFunction',
    kind: 'internal',
    args: [],
};
//# sourceMappingURL=contract.test.js.map