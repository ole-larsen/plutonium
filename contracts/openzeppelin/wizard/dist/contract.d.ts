export interface Contract {
    name: string;
    license: string;
    parents: Parent[];
    using: Using[];
    natspecTags: NatspecTag[];
    imports: string[];
    functions: ContractFunction[];
    constructorCode: string[];
    constructorArgs: FunctionArgument[];
    variables: string[];
    upgradeable: boolean;
}
export declare type Value = string | number | {
    lit: string;
} | {
    note: string;
    value: Value;
};
export interface Parent {
    contract: ParentContract;
    params: Value[];
}
export interface ParentContract {
    name: string;
    path: string;
}
export interface Using {
    library: ParentContract;
    usingFor: string;
}
export interface BaseFunction {
    name: string;
    args: FunctionArgument[];
    returns?: string[];
    kind: FunctionKind;
    mutability?: FunctionMutability;
}
export interface ContractFunction extends BaseFunction {
    override: Set<string>;
    modifiers: string[];
    code: string[];
    mutability: FunctionMutability;
    final: boolean;
}
export declare type FunctionKind = 'internal' | 'public';
export declare type FunctionMutability = typeof mutabilityRank[number];
declare const mutabilityRank: readonly ["pure", "view", "nonpayable", "payable"];
export interface FunctionArgument {
    type: string;
    name: string;
}
export interface NatspecTag {
    key: string;
    value: string;
}
export declare class ContractBuilder implements Contract {
    readonly name: string;
    license: string;
    upgradeable: boolean;
    readonly using: Using[];
    readonly natspecTags: NatspecTag[];
    readonly constructorArgs: FunctionArgument[];
    readonly constructorCode: string[];
    readonly variableSet: Set<string>;
    private parentMap;
    private functionMap;
    constructor(name: string);
    get parents(): Parent[];
    get imports(): string[];
    get functions(): ContractFunction[];
    get variables(): string[];
    addParent(contract: ParentContract, params?: Value[]): boolean;
    addUsing(library: ParentContract, usingFor: string): void;
    addOverride(parent: string, baseFn: BaseFunction, mutability?: FunctionMutability): void;
    addModifier(modifier: string, baseFn: BaseFunction): void;
    addNatspecTag(key: string, value: string): void;
    private addFunction;
    addConstructorArgument(arg: FunctionArgument): void;
    addConstructorCode(code: string): void;
    addFunctionCode(code: string, baseFn: BaseFunction, mutability?: FunctionMutability): void;
    setFunctionBody(code: string[], baseFn: BaseFunction, mutability?: FunctionMutability): void;
    addVariable(code: string): boolean;
}
export {};
//# sourceMappingURL=contract.d.ts.map