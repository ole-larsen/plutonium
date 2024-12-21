import { Contract } from './contract';
import { CommonOptions } from './common-options';
export interface ERC721Options extends CommonOptions {
    name: string;
    symbol: string;
    baseUri?: string;
    enumerable?: boolean;
    uriStorage?: boolean;
    burnable?: boolean;
    pausable?: boolean;
    mintable?: boolean;
    incremental?: boolean;
    votes?: boolean;
}
export declare const defaults: Required<ERC721Options>;
export declare function printERC721(opts?: ERC721Options): string;
export declare function isAccessControlRequired(opts: Partial<ERC721Options>): boolean;
export declare function buildERC721(opts: ERC721Options): Contract;
//# sourceMappingURL=erc721.d.ts.map