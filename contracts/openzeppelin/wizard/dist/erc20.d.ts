import { Contract } from './contract';
import { CommonOptions } from './common-options';
export interface ERC20Options extends CommonOptions {
    name: string;
    symbol: string;
    burnable?: boolean;
    snapshots?: boolean;
    pausable?: boolean;
    premint?: string;
    mintable?: boolean;
    permit?: boolean;
    votes?: boolean;
    flashmint?: boolean;
}
export declare const defaults: Required<ERC20Options>;
export declare function printERC20(opts?: ERC20Options): string;
export declare function isAccessControlRequired(opts: Partial<ERC20Options>): boolean;
export declare function buildERC20(opts: ERC20Options): Contract;
export declare const premintPattern: RegExp;
//# sourceMappingURL=erc20.d.ts.map