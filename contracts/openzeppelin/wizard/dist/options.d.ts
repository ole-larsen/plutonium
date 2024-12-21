import type { Contract } from './contract';
export interface Options {
    transformImport?: (path: string) => string;
}
export interface Helpers extends Required<Options> {
    upgradeable: boolean;
    transformName: (name: string) => string;
    transformVariable: (code: string) => string;
}
export declare function withHelpers(contract: Contract, opts?: Options): Helpers;
//# sourceMappingURL=options.d.ts.map