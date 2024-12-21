import { CustomOptions } from './custom';
import { ERC20Options } from './erc20';
import { ERC721Options } from './erc721';
import { ERC1155Options } from './erc1155';
import { GovernorOptions } from './governor';
export interface KindedOptions {
    ERC20: {
        kind: 'ERC20';
    } & ERC20Options;
    ERC721: {
        kind: 'ERC721';
    } & ERC721Options;
    ERC1155: {
        kind: 'ERC1155';
    } & ERC1155Options;
    Governor: {
        kind: 'Governor';
    } & GovernorOptions;
    Custom: {
        kind: 'Custom';
    } & CustomOptions;
}
export declare type GenericOptions = KindedOptions[keyof KindedOptions];
export declare function buildGeneric(opts: GenericOptions): import("./contract").Contract;
//# sourceMappingURL=build-generic.d.ts.map